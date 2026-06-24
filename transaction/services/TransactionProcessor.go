package services

import (
	"database/sql"
	. "finance/models"
	. "message/models"
)

type TransactionProcessor struct {
	db *sql.DB
}

func NewTrasactionProcessor(db *sql.DB) *TransactionProcessor {
	return &TransactionProcessor{db: db}
}

func (p *TransactionProcessor) FetchAndProcessTransactions() error {
	var buff Transaction
	var msgs []Message
	c := make(chan Message);

	res, err := p.db.Query(
		`SELECT transaction_id, from_account_number, to_account_number, currency, amount, destination_bank 
			FROM transactions WHERE transaction_state = 'PENDING'`)

	if err != nil {
		return err
	}
	for t_count := 0 ;; {
		for res.Next() {
			res.Scan(&buff.Id, &buff.From, &buff.To, &buff.Currency, &buff.Amount, &buff.DestinationBank)
			go p.process(buff, c) // pass a copy
	
			t_count++
			if(t_count>10){	// 10 for controlling db pool usage / speed or sth
				msgs = append(msgs, <- c)

				t_count = len(msgs)
				if len(msgs) > 10 {
					clear(msgs)
				}
				break
			}
		}
	}
}

func (p *TransactionProcessor) process(t Transaction, msg_chn chan Message) { // should be fast (very fast)
	if t.DestinationBank == "" { // here
		
	}
	
	// check for fraud
	// Send for process externally
	// Process internally
	// notify success (waiting)
	// give tracking code (server notifies on complete)
}
