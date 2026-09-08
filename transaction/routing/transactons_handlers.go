package routing

import (
	"encoding/json"
	"net/http"

	"storage/database"
	"transaction/services/message"

	"user/services/auth"
)

func RegisterRoutes() {
	http.HandleFunc("POST /t", TransactionHandler)
}

var transactionDAO = message.NewTransactionDAO(database.DB)

func TransactionHandler(w http.ResponseWriter, req *http.Request) {
	var req_msg []byte
	ctx := req.Context()

	_, err := auth.GetHeaderToken(req.Header)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if dec_err := json.NewDecoder(req.Body).Decode(&msg); dec_err != nil {
		http.Error(w, "Error", http.StatusBadRequest)
		return
	}

	var tx_msg message.Transaction

	if _err := json.Unmarshal(req_msg, &tx_msg); _err != nil {
		iso_msg, parse_err := message.ParseBytesToISO8583(req_msg)
		if parse_err != nil {
			http.Error(w, parse_err.Error(), http.StatusBadRequest)
			return
		}

	} else {
		transactionDAO.RecordTransaction(ctx, tx_msg)
	}

	// process transaction message
}
