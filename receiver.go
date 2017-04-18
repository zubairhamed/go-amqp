package amqp

import "log"

type Receiver struct {
	session *Session
}

func (r *Receiver) Receive() *Message {
	log.Println("Receiver:Receive")

	return nil
}

func (r *Receiver) Close() {
	log.Println("Receiver:Close")
}
