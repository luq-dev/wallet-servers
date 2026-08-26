package routing

import (
	"encoding/json"
	"net/http"

	"finance/models"
	"storage/database"
	"user/services/auth"
	"user/services/dao"

	"github.com/golang-jwt/jwt/v5"
)

var cardDao = dao.NewCardDAO(database.DB)
var accountDao = dao.NewAccountDAO(database.DB)

func addAccount(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()

	t, err := auth.GetHeaderToken(req.Header)
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

	card, card_err := cardDao.NewCard(ctx, accNumber)

	if card_err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Error while creating a card"})
		return
	}
	resp_data := map[any]any{
		"account_number": accNumber,
		"card":           card,
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
	ctx := req.Context()
	t, err := auth.GetHeaderToken(req.Header)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	mapClaims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid Claims"})
		return
	}

	email, ok := mapClaims["email"].(string)
	
	if !ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid uid"})
		return
	}

	accounts, accounts_err := accountDao.GetUserAccounts(ctx, email)
	if accounts_err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": accounts_err.Error()})
		return
	}

	w.Header().Set("Content Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accounts)
}
