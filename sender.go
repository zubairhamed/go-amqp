package amqp

import "log"

type Sender struct {
	session *Session
}

func (s *Sender) Send(msg *Message) {
	log.Println("Sender:Send")
	// >> transfer
	// << disposition
}

func (r *Sender) Close() {
	log.Println("Sender:Close")
}
