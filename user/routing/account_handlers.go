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
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	mapClaims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		http.Error(w, "Invalid Claims", http.StatusUnauthorized)
		return
	}

	uid, ok := mapClaims["uid"].(float64)
	if !ok {
		http.Error(w, "Invalid user ID in token", http.StatusUnauthorized)
		return
	}

	var acc models.Account
	dec_err := json.NewDecoder(req.Body).Decode(&acc)

	if dec_err != nil {
		http.Error(w, dec_err.Error(), http.StatusBadRequest)
		return
	}
	accNumber := ""
	if acc.Name != "" && acc.Type != "" {

		tx, _ := database.DB.Begin()
		_, err := database.DB.Exec("INSERT INTO accounts(user_id, account_name, account_type) VALUES ($1, $2, $3)", int64(uid), acc.Name, acc.Type)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tx.Commit()
	}

	accNo_err := database.DB.QueryRow("SELECT accounts.account_number FROM accounts JOIN users ON users.email = accounts.user_id WHERE users.id = $1", uid).Scan(&accNumber)
	
	if accNo_err != nil {
		http.Error(w, accNo_err.Error(), http.StatusNotFound)
	}

	// assign a card
	r, req_err := http.Post("localhost:8083", "Content-Type: application/json", bytes.NewReader([]byte(accNumber)))
	if req_err != nil {

	}

	var card map[string]string

	defer r.Body.Close()

	decode_err := json.NewDecoder(r.Body).Decode(&card)

	if decode_err != nil {
		http.Error(w, decode_err.Error(), http.StatusExpectationFailed)
	}
	resp_data := map[any]any {
		"account_number": accNumber,
		"card": card,
	}
	rdata, marsh_err := json.Marshal(resp_data)

	if marsh_err != nil {
		http.Error(w, marsh_err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(rdata)
}

func getUserAccounts(w http.ResponseWriter, req *http.Request) {
	t, err := auth.GetToken(req.Header)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	mapClaims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		http.Error(w, "Invalid Map Claims", http.StatusUnauthorized)
		return
	}

	uid, ok := mapClaims["uid"].(float64)
	if !ok {
		http.Error(w, "Invalid Map claims", http.StatusUnauthorized)
		return
	}

	rows, err := database.DB.Query("SELECT account_id, account_name, account_type FROM user_account_details WHERE user_id = $1", int64(uid))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
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
	json.NewEncoder(w).Encode(accounts)
}
