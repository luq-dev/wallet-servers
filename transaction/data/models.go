package data

import "time"

type Transaction struct {
	From int64
	To int64
	Currency string
	Amount float64
	description string
	status string
	ExchangeRate float64
	Date time.Time
}

 