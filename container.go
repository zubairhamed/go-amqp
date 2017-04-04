package amqp

import (
	"bufio"
	"log"
	"net"
)

func NewContainer(network string) *Container {
	c := &Container{
		network: network,
	}

	return c
}

type Container struct {
	network string
}

func (c *Container) doConnect() (conn net.Conn) {
	// Connect
	conn, err := net.Dial("tcp", c.network)
	if err != nil {
		panic(err)
	}

	// Handshake
	protocolType := ProtocolTypeNone
	conn.Write(CreateProtocolHeader(protocolType))

	readBuf := make([]byte, 1500)
	len, err := bufio.NewReader(conn).Read(readBuf)

	err, _ = UnmarshalProtocolHeader(readBuf[:len])
	if err != nil {
		panic(err.Error())
	}

	// Send Open Performative
	openPerformative := NewOpenPerformative("MyContainer", "localhost")
	_, err = SendPerformative(conn, openPerformative)
	if err != nil {
		panic(err.Error())
	}

	log.Println("Receiving Open Peformative")
	_, err = bufio.NewReader(conn).Read(readBuf)

	err = openPerformative.Decode(readBuf)
	if err != nil {
		log.Panic(err)
		return
	}
	// Read Incoming Open Performative

	log.Println("Protocol Negotiation OK")

	// >> Open
	// << Open
	// >> Begin
	// << Begin
	// >> Attach
	// << Attach
	// << Flow

	return
}

func (c *Container) Connect() *Connection {
	conn := c.doConnect()

	return &Connection{
		netConn: conn,
	}
}

func (c *Container) Send(msg *Message) {
	log.Println("Send", msg)
}

func (c *Container) Shutdown() {

}
