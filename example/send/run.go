package main

import "github.com/zubairhamed/go-amqp"

func main() {

	container := amqp.NewContainer(":5672")

	conn := container.Connect()
	sender := conn.CreateSender("myQueue")

	msg := amqp.NewMessage()
	sender.Send(msg)

	<-make(chan struct{})
}
