package services

import (
	. "user/database"
)

// 0000000000	base16
// Number


func GetLastAccountNumber() (string,error){
	var acc string
	err := DB.QueryRow("SELECT account_number FROM accounts WHERE created_at=max(created_at)").Scan(&acc)
	return acc, err
}

func GetNextAccountNumber() (string, error) {
	last, err := GetLastAccountNumber()
	if err == nil {
		return string(last)[0:8] + string(int(string(last)[9])+1), err
	}
	return "", err
}