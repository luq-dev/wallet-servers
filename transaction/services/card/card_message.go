package message

type CardMessage struct {
	Type    string `json:"msgtype"`
	Data    string `json:"data"`
	Reciept string `json:"reciept"`
}

type TransactionDTO struct {
	To       string `json:"to"`
	From     string `json:"from"`
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
	Refund   bool   `json:"refund"`
	Auto     bool   `json:"auto"`
}
