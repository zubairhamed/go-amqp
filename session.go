package amqp

//import (
//	"log"
//)

//func NewSession(conn *Connection) (s *Session, err error) {
//	s = &Session{
//		connection: conn,
//	}
//	// err = s.initConnection()
//
//	return
//}
//
//type Session struct {
//	connection *Connection
//}

//func (s *Session) CreateSender(queue string) (sender *Sender, err error) {
//	log.Println("Session:CreateSender")
//
//	attach := performatives.NewAttachPerformative()
//
//	attach.Name = types.NewString("sender")
//	attach.Handle = types.NewHandle(0)
//	attach.Role = types.NewRole(true)
//	attach.Target = types.NewFields(map[string]types.AMQPType{
//		"Address": types.NewString("my_queue"),
//	})
//	attach.InitialDeliveryCount = types.NewSequenceNumber(0)
//
//	s.connection.SendPerformative(attach)
//	readBuf, err := ReadFromConnection(s.connection.netConn)
//	perf, err := performatives.DecodeAttachPerformative(readBuf)
//
//	DescribeType(perf)
//
//	if err != nil {
//		log.Panic(err)
//		return
//	}
//
//	DescribeType(perf)
//	if err != nil {
//		return
//	}
//
//	sender = &Sender{
//		session: s,
//	}
//	return
//}
//
//func (s *Session) CreateReceiver(queue string) (receiver *Receiver, err error) {
//	log.Println("Session:CreateReceiver")
//
//	attach := performatives.NewAttachPerformative()
//
//	attach.Name = types.NewString("receiver")
//	attach.Handle = types.NewHandle(1)
//	attach.Role = types.NewRole(false)
//	attach.Source = types.NewFields(map[string]types.AMQPType{
//		"Address": types.NewString("my_queue"),
//	})
//
//	s.connection.SendPerformative(attach)
//	readBuf, err := ReadFromConnection(s.connection.netConn)
//	perf, err := performatives.DecodeAttachPerformative(readBuf)
//	if err != nil {
//		log.Panic(err)
//		return
//	}
//
//	log.Println("OK CreateReceiver")
//
//	DescribeType(perf)
//
//	receiver = &Receiver{
//		session: s,
//	}
//	return
//}

//func (s *Session) Close() {
//	log.Println("Session:Close")
//}
