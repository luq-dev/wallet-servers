package routing

import (
	"encoding/json"
	"fmt"
	"net/http"
	"finance/models"

	database "user/database"
	"user/services/auth"
)



func TransactionHandler(w http.ResponseWriter, r *http.Request){
	// Transaction {to: Account.email, from: Account.email, amount: float64} *amount in sender's currency *security checks like sender's [valid or not] client to be added
	// notify reciever on new transaction added

	_, err := auth.GetToken(r.Header)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	// need security checks here
	
	var t models.Transaction // and herservicee

	if dec_err := json.NewDecoder(r.Body).Decode(&t); dec_err != nil {
		http.Error(w, "Error", http.StatusBadRequest)
		return
	}

	_, stmt_err := database.DB.Exec(`
	INSERT INTO transactions(
		from_account_number, 
		to_account_number, 
		currency, 
		amount, 
		destination_bank, 
		description, 
		transaction_state
		) VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		t.From, t.To, t.Currency, t.Amount, t.DestinationBank, t.Details, "PENDING");

	if stmt_err != nil {
		http.Error(w, stmt_err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Transaction Processing...")
}