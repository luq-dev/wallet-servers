package models

type Transaction struct {
	Id              string `json:"id"`
	From            string `json:"from"`
	To              string `json:"to"`
	Amount          int64  `json:"amount"` // cents
	Currency        string `json:"currency"`
	DestinationBank string `json:"bank"`
	Details         string `json:"details"`
}

type Account struct {
	ID   int64  `json:"account_id"`
	Type string `json:"account_type"`
	Name string `json:"account_name"`
}

