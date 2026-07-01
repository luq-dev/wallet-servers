package routing

import (
	"encoding/json"
	"fmt"
	"net/http"

	"user/services/auth"
)

type TransactionMessage map[string]string

func TransactionHandler(w http.ResponseWriter, r *http.Request) {
	// Transaction {to: Account.email, from: Account.email, amount: float64} *amount in sender's currency *security checks like sender's [valid or not] client to be added
	// notify reciever on new transaction added

	_, err := auth.GetToken(r.Header)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	// need security checks here

	var t TransactionMessage

	if dec_err := json.NewDecoder(r.Body).Decode(&t); dec_err != nil {
		http.Error(w, "Error", http.StatusBadRequest)
		return
	}

	// verify involved parties
	
	go worker(t)

	fmt.Fprintf(w, "Transaction Processing...")
}

func worker(t TransactionMessage) {

}
