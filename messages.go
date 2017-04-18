package amqp

import "log"

func CreateProtocolHeader(t ProtocolType) []byte {
	return []byte{
		65, 77, 81, 80, byte(t), 1, 0, 0,
	}
}

func NewMessage() *Message {
	log.Println("NewMessage")
	return &Message{}
}

type Message struct {
}
