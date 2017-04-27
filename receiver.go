package amqp

func NewReceiver(name string, ch chan *Event) *Receiver {
	return &Receiver{
		Client: Client{
			name: name,
			ch:   ch,
			role: ROLE_RECEIVER,
		},
	}
}

type Receiver struct {
	Client
	//	session *Session
}
