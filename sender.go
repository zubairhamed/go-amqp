package amqp

type Sender struct {
}

func (s *Sender) Send(msg *Message) {
	// >> transfer
	// << disposition
}
