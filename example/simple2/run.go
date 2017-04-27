package main

import (
	. "github.com/zubairhamed/go-amqp"
)

func main() {
	conn := NewConnection(":5672", "my_queue")
	var err error

	senderEvents := make(chan *Event)
	s := NewSender("sender", senderEvents)
	err = s.Dial(conn)
	if err != nil {
		panic(err.Error())
	}

	receiverEvents := make(chan *Event)
	r := NewReceiver("receiver", receiverEvents)
	err = r.Dial(conn)
	if err != nil {
		panic(err.Error())
	}

	outMsg := NewMessage()
	s.Send(outMsg)

	for {
		select {
		case evt, open := <-senderEvents:
			if open {
				switch evt.Type {
				case EVENT_MSG_OPEN:
				case EVENT_MSG_ATTACH:
				case EVENT_MSG_BEGIN:
				case EVENT_MSG_CLOSE:
				case EVENT_MSG_DETACH:
				case EVENT_MSG_DISPOSITION:
				case EVENT_MSG_END:
				case EVENT_MSG_FLOW:
				case EVENT_MSG_TRANSFER:
				}
			}

		case evt, open := <-receiverEvents:
			if open {
				switch evt.Type {
				case EVENT_MSG_OPEN:
				case EVENT_MSG_ATTACH:
				case EVENT_MSG_BEGIN:
				case EVENT_MSG_CLOSE:
				case EVENT_MSG_DETACH:
				case EVENT_MSG_DISPOSITION:
				case EVENT_MSG_END:
				case EVENT_MSG_FLOW:
				case EVENT_MSG_TRANSFER:
				}
			}
		}
	}
}
