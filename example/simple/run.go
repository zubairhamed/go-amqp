package main

import (
	"github.com/apex/log"
	"github.com/zubairhamed/go-amqp"
)

func main() {
	log.SetLevel(log.DebugLevel)

	url := ":5672"
	// queue := "my_queue"

	conn := amqp.NewConnection(url)
	// session, err := amqp.NewSession(conn)
	amqp.NewSession(conn)
	//if err != nil {
	//	panic(err.Error())
	//}

	//sender, err := session.CreateSender(queue)
	//if err != nil {
	//	panic(err.Error())
	//}
	//
	//receiver, err := session.CreateReceiver(queue)
	//if err != nil {
	//	panic(err.Error())
	//}
	//
	//outMsg := amqp.NewMessage()
	//sender.Send(outMsg)
	//
	//inMsg := receiver.Receive()
	//
	//log.Debug(fmt.Sprint(inMsg))
	//
	//
	//receiver.Close()
	//sender.Close()
	//session.Close()
	//conn.Close()

	<-make(chan struct{})
}
