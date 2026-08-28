package message

import (
	"context"
	"database/sql"
	"encoding/json"
)

// DB transaction model
type Transaction struct {
	Id string
	From string
	To string
	Currency string
	Amount int64
	Destination_bank string
	details map[string]string
}

func(t *Transaction) Details() string{
	reciept, _ := json.Marshal(t.details)
	return string(reciept)
}

type TransactionDAO struct {
	DB *sql.DB
}

func(dao *TransactionDAO) RecordTransaction(ctx context.Context, t Transaction) string{
	
	dao.DB.ExecContext(ctx, "INSERT INTO transactions(transaction_id, from_account_number, to_account_number, currency, amount, destination_bank, details, transaction_status) VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id", t.Id, t.From, t.To, t.Currency, t.Amount, t.Destination_bank, t.Details(), "COMPLETE")

	return ""
}


