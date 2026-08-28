package routing

import (
	"encoding/json"
	"net/http"

	"user/services/auth"
	"transaction/services/message"
)


func RegisterRoutes() {
	http.HandleFunc("POST /t", TransactionHandler)
}


func TransactionHandler(w http.ResponseWriter, r *http.Request) {
	var msg message.ISO8583

	_, err := auth.GetHeaderToken(r.Header)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if dec_err := json.NewDecoder(r.Body).Decode(&msg); dec_err != nil {
		http.Error(w, "Error", http.StatusBadRequest)
		return
	}

	// process transaction message
}
