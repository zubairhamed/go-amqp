package amqp

import "net"

type Connection struct {
	netConn net.Conn
}

func (c *Connection) CreateSender(address string) *Sender {
	return &Sender{}
}
