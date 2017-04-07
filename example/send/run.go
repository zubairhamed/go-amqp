package main

import (
	"github.com/zubairhamed/go-amqp"
	"log"
)

func main() {

	url := ":5672"
	queue := "my_queue"

	conn := amqp.NewConnection(url)
	session := amqp.NewSession(conn)

	sender, err := session.CreateSender(queue)
	if err != nil {
		panic(err.Error())
	}

	receiver, err := session.CreateReceiver(queue)
	if err != nil {
		panic(err.Error())
	}

	outMsg := amqp.NewMessage()
	sender.Send(outMsg)

	inMsg := receiver.Receive()
	log.Println(inMsg)
	receiver.Close()
	sender.Close()
	session.Close()
	conn.Close()

	<-make(chan struct{})
}
