package amqp

import (
	"net"
	"bufio"
	"log"
	. "github.com/zubairhamed/go-amqp/frames"
	. "github.com/zubairhamed/go-amqp/frames/performatives"
)

func NewConnection(url string) *Connection {
	return &Connection{
		url: url,
		connected: false,
	}
}

type Connection struct {
	netConn net.Conn
	url string
	connected bool
}

func (c *Connection) doConnect() (err error) {
	// Connect
	conn, err := net.Dial("tcp", c.url)
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

	beginPerformative := NewBeginPerformative()

	_, err = SendPerformative(conn, beginPerformative)
	err = beginPerformative.Decode(readBuf)
	if err != nil {
		log.Panic(err)
		return
	}


	// >> Begin
	// << Begin

	// >> Attach
	// << Attach

	// << Flow

	c.connected = true

	return
}

func (c *Connection) Close() {

}
