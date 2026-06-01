package workers

import "transaction/data"

type TransactionHandler struct {
	Transactions []data.Transaction
}

// process transactions and return

func (h *TransactionHandler) AddTransactionJob(t data.Transaction) {

}
