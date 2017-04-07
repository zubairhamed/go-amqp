package amqp

import (
	"bufio"
	. "github.com/zubairhamed/go-amqp/frames"
	. "github.com/zubairhamed/go-amqp/frames/performatives"
	"github.com/zubairhamed/go-amqp/types"
	"log"
	"net"
)

func NewConnection(url string) *Connection {
	return &Connection{
		url:       url,
		connected: false,
	}
}

type Connection struct {
	netConn   net.Conn
	url       string
	connected bool
}

func (c *Connection) doConnect() (err error) {
	// Connect
	conn, err := net.Dial("tcp", c.url)
	if err != nil {
		panic(err)
	}

	readBuf := make([]byte, 1500)

	// Handshake
	SendHandshake(conn)
	_, err = bufio.NewReader(conn).Read(readBuf)
	err = HandleHandshake(readBuf)
	if err != nil {
		panic(err.Error())
	}

	// Send Open Performative
	openPerformative := NewOpenPerformative()
	openPerformative.ContainerId = types.NewString("MyContainer")

	_, err = SendPerformative(conn, openPerformative)
	if err != nil {
		panic(err.Error())
	}

	_, err = bufio.NewReader(conn).Read(readBuf)
	openPerformative, err = DecodeOpenPerformative(readBuf)
	if err != nil {
		log.Panic(err)
		return
	}

	DescribeType(openPerformative)

	// Read Incoming Open Performative
	beginPerformative := NewBeginPerformative()
	beginPerformative.NextOutgoingId = types.NewUInt(4294967293)
	beginPerformative.IncomingWindow = types.NewUInt(2048)
	beginPerformative.OutgoingWindow = types.NewUInt(2048)
	beginPerformative.HandleMax = types.NewUInt(7)

	_, err = SendPerformative(conn, beginPerformative)
	_, err = bufio.NewReader(conn).Read(readBuf)

	beginPerformative, err = DecodeBeginPerformative(readBuf)
	if err != nil {
		log.Panic(err)
		return
	}

	DescribeType(beginPerformative)

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
