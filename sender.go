package amqp

func NewSender(name string, ch chan *Event) *Sender {
	return &Sender{
		Client: Client{
			name: name,
			ch:   ch,
			role: ROLE_SENDER,
		},
	}
}

type Sender struct {
	Client
	//	session *Session
}

func (s *Sender) Send(msg *Message) {
	//	log.Println("Sender:Send")
	//
	//	transfer := performatives.NewTransferPerformative()
	//
	//	s.session.connection.SendPerformative(transfer)
	//	log.Println("Receiving from send..")
	//	readBuf, err := ReadFromConnection(s.session.connection.netConn)
	//	if err != nil {
	//		log.Panic(err)
	//		return
	//	}
	//
	//	log.Println("Return from Send", readBuf)

	// >> transfer
	// << disposition
}
