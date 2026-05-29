package accounting

import "time"

type Transaction struct {
	id          int
	time        time.Time  // to be taken from the reciever's client or sender's client local time
	to          Account
	from        Account
	amount      float32
	description string // reciept, JSON
}
/*
	TransactionMessage = "
		{
			from:
			to:
			amount:
			desc: ""
		}
	"
*/

type TransactionMessage struct {
	from   Account
	to     Account
	amount float32
	desc   string
}

func TransactionProcessor(msg TransactionMessage){
	// check for fraud (sources n shit)
	// check if msg.to exists
	// add msg.amount to msg.to account
}
