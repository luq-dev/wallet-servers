package dao

import (
	"context"
	"database/sql"
)

type AccountDTO struct {
	Name string `json:"name"`
	Number string `json:"number"`
	UserEmail string `json:"user_email`
	Type string `json:"type"`
	Currency string `json:"currency"`
}

type AccountDAO struct {
	DB *sql.DB
}

func NewAccountDAO(db *sql.DB) *AccountDAO {
	return &AccountDAO{
		DB: db,
	}
}

func(a *AccountDAO) GetUserAccounts(ctx context.Context, userEmail string) ([]AccountDTO, error) {
	tx, txBegin_err := a.DB.BeginTx(ctx, nil)
	defer tx.Rollback()

	var accounts []AccountDTO

	if txBegin_err != nil {
		return []AccountDTO{}, txBegin_err
	}

	rows, query_err := tx.QueryContext(ctx, "SELECT a.account_name, a.account_number, a.user_email, a.currency, t.type_name FROM accounts a JOIN account_types t ON t.id = t.account_type_id WHERE a.user_email = $1", userEmail)

	if query_err != nil {
		return []AccountDTO{}, nil
	}

	defer rows.Close()
	
	for rows.Next() {
		var tempAccount AccountDTO

		if err := rows.Scan(&tempAccount.Name, &tempAccount.Number, &tempAccount.UserEmail, &tempAccount.Type, &tempAccount.Currency); err != nil {
			return []AccountDTO{}, err
		}

		if err := rows.Err(); err != nil {
			return []AccountDTO{}, err
		}

		accounts = append(accounts, tempAccount)
	}

	tx.Commit()

	return accounts, nil
}
