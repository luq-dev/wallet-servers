package accounting

import (
	"time"
)

type Account struct {
	number  int
	Name    string
	book    []Transaction
	balance float32
}


func (acc *Account) updateBalance() {

	var total float32

	for i := 0; i < len(acc.book); i++ {
		total += acc.book[i].amount
	}

	acc.balance = total
}

func (acc *Account) Send(amount float32, to Account, desc string) {
	if acc.balance >= amount {
		to.Receive(amount, *acc, desc)
		acc.recordTransaction(*acc, to, -amount, desc)
	}
}

func (acc *Account) Receive(amount float32, from Account, desc string) {
	acc.recordTransaction(from, *acc, amount, desc)
}

// Main Event
func (acc *Account) recordTransaction(from, to Account, amount float32, desc string) bool {
	acc.book = append(acc.book, Transaction{time: time.Now(), from: from, to: to, amount: amount, description: desc})
	return true
}
