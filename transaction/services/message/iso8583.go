package message

import (
	"fmt"
	"math"
)

type ISO8583 struct {
	MTI    []byte
	Bitmap int64
	DE     map[int][]byte
}

func (msg *ISO8583) Encode() []byte {

	var bitmap int64
	var DEs [][]byte
	var msgString string

	var buff []byte

	for k, _ := range msg.DE {
		bitmap += int64(math.Pow(2, float64(64-k)))
	}

	for i := 1; i <= 64; i++ {
		if _, exists := msg.DE[i]; exists {
			DEs = append(DEs, msg.DE[i])
		}
	}

	buff = append(buff, msg.MTI...)
	buff = append(buff)

	msgString = fmt.Sprintf("%d\n", int(msg.MTI))
	msgString = fmt.Sprintf("%s%x\n", msgString, bitmap)

	for n := range DEs {
		msgString = fmt.Sprintf("%s%s\n", msgString, DEs[n])
	}

	return []byte(msgString)
}

func NewMessage() *ISO8583 {
	return &ISO8583{}
}

func (msg *ISO8583) SetMTI(mti []byte) error {

	if len(mti) != 4 {
		return fmt.Errorf("Invalid MTI length")
	}

	msg.MTI = mti
	return nil
}

func (msg *ISO8583) SetDE(n int, value []byte) {
	msg.DE[n] = value
}
