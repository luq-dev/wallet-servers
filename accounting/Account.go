package accounting

import "time"

type Account interface {
	updateBalance()
	Send(to Account, amount float32, desc string)
	Receive(from Account, amount float32, desc string)
	recordTransaction(from, to Account, amount float32, desc string) bool
}

type LocalAccount struct {
	number  int
	Name    string
	book    []Transaction
	balance float32
}

func (acc *LocalAccount) updateBalance() {

	var total float32

	for i := 0; i < len(acc.book); i++ {
		total += acc.book[i].amount
	}

	acc.balance = total
}

func (acc *LocalAccount) Send(to Account, amount float32, desc string) {
	if acc.balance >= amount {
		to.Receive(acc, amount, desc)
		acc.recordTransaction(acc, to, -amount, desc)
	}
}

func (acc *LocalAccount) Receive(from Account, amount float32, desc string) {
	acc.recordTransaction(from, acc, amount, desc)
}

// Main Event
func (acc *LocalAccount) recordTransaction(from, to Account, amount float32, desc string) bool {
	/*
	* Checks
	* record
	*/
	acc.book = append(acc.book, Transaction{time: time.Now(), from: from, to: to, amount: amount, description: desc})
	acc.updateBalance()
	return true
}
