package main

import (
	. "github.com/zubairhamed/go-amqp"
	// "log"
)

func main() {
	conn := NewConnectInfo(":5672", "my_queue")
	var err error

	//receiverEvents := make(chan *Event)
	//r := NewReceiver("receiver", receiverEvents)
	//err = r.Dial(conn)
	//if err != nil {
	//	panic(err.Error())
	//}

	senderEvents := make(chan *Event)
	s := NewSender("sender", senderEvents)
	err = s.Dial(conn)
	if err != nil {
		panic(err.Error())
	}

	// 10 times, create a message and send to broker
	//outMsg := NewMessage()
	//s.Send(outMsg)

	// Close Sender and Receiver after 10 times

	//for {
	//	select {
	//	case evt, open := <-receiverEvents:
	//		if open {
	//			switch evt.Type {
	//			case EVENT_MSG_TRANSFER:
	//				log.Println("Got message from sender..", evt.Performative)
	//			}
	//		}
	//	}
	//}
}
