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

