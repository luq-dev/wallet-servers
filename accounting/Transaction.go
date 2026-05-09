package accounting

import "time"

type Transaction struct {
	id          int
	time        time.Time
	to          Account
	from        Account
	amount      float32
	description string // reciept, JSON
}

/*
	Transaction = "
		{
			from:
			to:
			amount:
			desc: ""
		}
	"
*/

type TransactionMessage string
