package message

import (
	"fmt"
	"math"
)

type ISO8583 struct {
	MTI string
	Bitmap string
	DE map[int]any
}

func (msg *ISO8583) SetMTI(mti string) {
	msg.MTI = mti
}

func (msg *ISO8583) SetDE(n int, value any){
	msg.DE[n] = value
}

func (msg *ISO8583) Build() string {
	
	var bitmap int64
	var DEs []any
	var msgString string

	for k, _ := range msg.DE {
		bitmap += int64(math.Pow(2, float64(64-k)))
	}

	for i := 1; i<=64; i++ {
		if _, exists := msg.DE[i]; exists {
			DEs = append(DEs, msg.DE[i])
		}
	}

	msgString = fmt.Sprintf("%s\n", msg.MTI)
	msgString = fmt.Sprintf("%s%x\n", msgString, bitmap)

	for n := range DEs{
		msgString = fmt.Sprintf("%s%s\n", msgString,DEs[n])
	}

	return msgString
}