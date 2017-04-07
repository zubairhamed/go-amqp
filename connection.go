package amqp

import (
	"net"
	"bufio"
	"log"
	. "github.com/zubairhamed/go-amqp/frames"
	. "github.com/zubairhamed/go-amqp/frames/performatives"
	"github.com/zubairhamed/go-amqp/types"
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

	readBuf := make([]byte, 1500)

	// Handshake
	SendHandshake(conn)
	_, err = bufio.NewReader(conn).Read(readBuf)
	log.Println("From Handshake", readBuf)

	err = HandleHandshake(readBuf)
	if err != nil {
		panic(err.Error())
	}

	// Send Open Performative
	log.Println("sending open peformative")
	openPerformative := NewOpenPerformative()
	openPerformative.ContainerId = types.NewString("MyContainer")

	_, err = SendPerformative(conn, openPerformative)
	if err != nil {
		panic(err.Error())
	}
	log.Println("gpt back reply open peformative")

	_, err = bufio.NewReader(conn).Read(readBuf)
	openPerformative, err = DecodeOpenPerformative(readBuf)
	if err != nil {
		log.Panic(err)
		return
	}
	// Read Incoming Open Performative

	log.Println("Protocol Negotiation OK")
	log.Println("ChannelMax", openPerformative.ChannelMax)
	log.Println("ContainerId", openPerformative.ContainerId)
	log.Println("DesiredCapabilities", openPerformative.DesiredCapabilities)
	log.Println("HostName", openPerformative.Hostname)
	log.Println("Idle Timeout", openPerformative.IdleTimeout)
	log.Println("Incoming Locales", openPerformative.IncomingLocales)
	log.Println("MaxFrameSize", openPerformative.MaxFrameSize)
	log.Println("Offered Capabilities", openPerformative.OfferedCapabilities)
	log.Println("Outgoing Locales", openPerformative.OutgoingLocales)
	log.Println("Properties", openPerformative.Properties)

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

	// log.Println("",
	log.Println("Properties", beginPerformative.Properties)
	log.Println("HandleMax", beginPerformative.HandleMax)
	log.Println("OfferedCapabilities", beginPerformative.OfferedCapabilities)
	log.Println("OutgoingWindow", beginPerformative.OutgoingWindow)
	log.Println("IncomingWindow", beginPerformative.IncomingWindow)
	log.Println("NextOutgoingId", beginPerformative.NextOutgoingId)
	log.Println("RemoteChannel", beginPerformative.RemoteChannel)


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
