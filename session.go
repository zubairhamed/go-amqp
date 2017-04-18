package amqp

import (
	"log"
	"github.com/zubairhamed/go-amqp/frames/performatives"
)

func NewSession(conn *Connection) (s *Session, err error) {
	s = &Session{
		connection: conn,
	}
	err = s.initConnection()

	return
}

type Session struct {
	connection *Connection
}

func (s *Session) CreateSender(queue string) (sender *Sender, err error) {
	log.Println("Session:CreateSender")

	attach := performatives.NewAttachPerformative()

	s.connection.SendPerformative(attach)

	// >> Attach sender
	// << Attach sender

	if err != nil {
		return
	}

	sender = &Sender{
		session: s,
	}
	return
}

func (s *Session) CreateReceiver(queue string) (receiver *Receiver, err error) {
	log.Println("Session:CreateReceiver")
	// >> attach receiver
	// << attach receiver

	if err != nil {
		return
	}

	receiver = &Receiver{
		session: s,
	}
	return
}

func (s *Session) Close() {
	log.Println("Session:Close")
}

func (s *Session) initConnection() (err error) {
	if !s.connection.connected {
		return s.connection.doConnect()
	}

	return
}
