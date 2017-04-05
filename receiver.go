package amqp

type Receiver struct {
	session *Session
}

func (r *Receiver) Receive() *Message {
	return nil
}

func (r *Receiver) Close() {

}
