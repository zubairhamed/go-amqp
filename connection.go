package amqp

import (
	. "github.com/zubairhamed/go-amqp/frames"
	. "github.com/zubairhamed/go-amqp/frames/performatives"
	"github.com/zubairhamed/go-amqp/types"
	"log"
	"net"
	"encoding/binary"
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

	c.netConn = conn

	var readBuf []byte

	// Handshake
	SendHandshake(conn)
	readBuf, err = ReadFromConnection(conn)
	err = HandleHandshake(readBuf)
	if err != nil {
		panic(err.Error())
	}

	// Send Open Performative
	openPerformative := NewOpenPerformative()
	openPerformative.ContainerId = types.NewString("MyContainer")


	_, err = c.SendPerformative(openPerformative)
	if err != nil {
		panic(err.Error())
	}

	readBuf, err = ReadFromConnection(conn)
	openPerformative, err = DecodeOpenPerformative(readBuf)
	if err != nil {
		log.Panic(err)
		return
	}

	DescribeType(openPerformative)

	// Read Incoming Open Performative
	beginPerformative := NewBeginPerformative()
	beginPerformative.NextOutgoingId = types.NewTransferNumber(4294967293)
	beginPerformative.IncomingWindow = types.NewUInt(2048)
	beginPerformative.OutgoingWindow = types.NewUInt(2048)
	beginPerformative.HandleMax = types.NewHandle(7)

	_, err = c.SendPerformative(beginPerformative)
	readBuf, err = ReadFromConnection(conn)
	beginPerformative, err = DecodeBeginPerformative(readBuf)
	if err != nil {
		log.Panic(err)
		return
	}

	DescribeType(beginPerformative)

	c.connected = true

	return
}

func (c *Connection) Close() {
	log.Println("Connection:Close")
}


func (c *Connection) Write(b []byte) (int, error){
	return c.netConn.Write(b)
}

func (c *Connection) SendPerformative(p Performative) (int, error) {
	b, _, err := p.Encode()
	if err != nil {
		panic(err.Error())
	}

	var frameSize uint32 = 8 + uint32(len(b))
	var frameSizeBytes = make([]byte, 4)
	binary.BigEndian.PutUint32(frameSizeBytes, frameSize)

	frameContent := EncodeFrame(b)

	return c.Write(frameContent)
}