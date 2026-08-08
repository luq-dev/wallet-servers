package routing

import(
	"bytes"
	"net/http"
	"encoding/json"

	"user/services/auth"
	"finance/models"
	"storage/database"

	"github.com/golang-jwt/jwt/v5"
)

// incomplete
func addAccount(w http.ResponseWriter, req *http.Request) {

	t, err := auth.GetToken(req.Header)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid Token"})
		return
	}

	mapClaims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid Token Claims"})
		return
	}

	uid, ok := mapClaims["uid"].(float64)
	if !ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid Token Claims"})
		return
	}

	var acc models.Account
	dec_err := json.NewDecoder(req.Body).Decode(&acc)

	if dec_err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid Request Format"})
		return
	}
	accNumber := ""
	if acc.Name != "" && acc.Type != "" {
		err := database.DB.QueryRow("INSERT INTO accounts(user_id, account_name, account_type) VALUES ($1, $2, $3) RETURNING account_number", int64(uid), acc.Name, acc.Type).Scan(&accNumber)

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to add account"})
			return
		}
	}

	// assign a card
	r, req_err := http.Post("localhost:8083", "Content-Type: application/json", bytes.NewReader([]byte(accNumber)))
	if req_err != nil {

	}

	var card map[string]string

	defer r.Body.Close()

	decode_err := json.NewDecoder(r.Body).Decode(&card)

	if decode_err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Error while creating a card"})
		return
	}
	resp_data := map[any]any {
		"account_number": accNumber,
		"card": card,
	}
	rdata, enc_err := json.Marshal(resp_data)

	if enc_err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Error while creating a card"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(rdata)
}

func getUserAccounts(w http.ResponseWriter, req *http.Request) {
	t, err := auth.GetToken(req.Header)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error":err.Error()})
		return
	}

	mapClaims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error":"Invalid Claims"})
		return
	}

	uid, ok := mapClaims["uid"].(float64)
	if !ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error":"Invalid uid"})
		return
	}

	rows, err := database.DB.Query("SELECT account_id, account_name, account_type FROM user_account_details WHERE user_id = $1", int64(uid))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	var accounts []models.Account

	for rows.Next() {
		var r1 int64
		var r2 string
		var r3 string
		rows.Scan(&r1, &r2, &r3)
		accounts = append(accounts, models.Account{ID: r1, Type: r2, Name: r3})
	}
	rows.Close()

	w.Header().Set("Content Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accounts)
}
