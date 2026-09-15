package iso8583

import (
	"fmt"
)

type ISO8583Message struct {
	MTI    []byte
	Bitmap []byte
	Data   []byte
}

func NewMessage() *ISO8583Message {
	return &ISO8583Message{}
}

func (msg *ISO8583Message) Encode() []byte {

	var buff []byte

	buff = append(buff, msg.MTI...)
	buff = append(buff, msg.Bitmap...)
	buff = append(buff, msg.Data...)

	return buff
}

func Decode(src []byte, dst *ISO8583Message) error {

	var msg ISO8583Message

	if src == nil {
		return fmt.Errorf("Invalid Message")
	}

	if len(src) < 12 {
		return fmt.Errorf("Invalid Message length")
	}

	msg.MTI = src[:4]

	var bitmap []byte

	var data []byte

	if int((src)[5]) >= 128 {
		copy((src)[5:12], bitmap)
		if len(src) > 12 {
			copy((src)[13:], data)
		}
	} else {
		copy((src)[5:8], bitmap)
		if len(src) > 8 {
			copy((src)[9:], data)
		}
	}

	msg.Data = data

	dst = &msg

	return nil
}
