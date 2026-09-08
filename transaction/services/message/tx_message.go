package message

type TransactionMessage interface {
	Decode([]byte) error
	Encode() byte
}
