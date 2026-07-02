package dao

import (
	"database/sql"
	. "finance/models"
)

type AccountDAO struct {
	db *sql.DB
}

func (dao *AccountDAO) CreateAccount(acc Account) {

}

func (dao *AccountDAO) CheckIfExists(accNumber string) bool{
	var n int64;
	dao.db.QueryRow("SELECT 1 FROM accounts WHERE account_number=$1", accNumber).Scan(&n)
	if n == 1 {
		return true
	}
	return false
}
