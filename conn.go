package amqp

import "net"

type Connection struct {
	netConn net.Conn
}

func (c *Connection) CreateSender(address string) *Sender {
	return &Sender{}
}

func SendPerformative(c net.Conn, p *PerformativeOpen) (int, error) {

	b, err := p.Encode()
	if err != nil {
		panic(err.Error())
	}

	c.Write(b)

	return 0, nil
}
