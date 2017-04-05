package amqp

type Sender struct {
	session *Session
}

func (s *Sender) Send(msg *Message) {
	// >> transfer
	// << disposition
}

func (r *Sender) Close() {

}
