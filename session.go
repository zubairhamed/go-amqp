package amqp

func NewSession(conn *Connection) *Session {
	return &Session{
		connection: conn,
	}
}

type Session struct {
	connection *Connection
}

func (s *Session) CreateSender(queue string) (sender *Sender, err error){
	err = s.initConnection()
	if err != nil {
		return
	}

	sender = &Sender{
		session: s,
	}
	return
}

func (s *Session) CreateReceiver(queue string) (receiver *Receiver, err error) {
	err = s.initConnection()
	if err != nil {
		return
	}

	receiver = &Receiver{
		session: s,
	}
	return
}

func (s *Session) Close() {

}

func (s *Session) initConnection() (err error) {
	if !s.connection.connected {
		return s.connection.doConnect()
	}

	return
}
