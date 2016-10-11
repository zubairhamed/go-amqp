package amqp

import (
	"bufio"
	"github.com/wendal/errors"
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
	log.Println("Incoming:", readBuf[:len], "of length", len)

	err = c.validateProtocolHeader(protocolType, readBuf)
	if err != nil {
		panic(err)
	}

	var out = []byte{0x00, 0x53, 0x10, 0xC0, 0x20, 0x0A}
	out = append(out, []byte{0xA1, 0x0B}...)
	out = append(out, []byte("MyContainer")...)
	out = append(out, []byte{0xA1, 0x09}...)
	out = append(out, []byte("localhost")...)
	out = append(out, []byte{0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40}...)

	conn.Write(out)
	len, err = bufio.NewReader(conn).Read(readBuf)
	log.Println("Incoming #2:", readBuf[:len], "of length", len)
	log.Println("str:", string(readBuf))

	//CreatePerformativeFrame(t TypeFormatCode, containerId, hostname string) (b []byte) {
	//
	//}

	//max-frame-size uint
	//null
	//
	//channel-max ushort
	//null
	//
	//idle-time-out milliseconds
	//null
	//
	//outgoing-locales
	//null
	//
	//incoming-locales
	//null
	//
	//offered-capabilities
	//null
	//
	//desired-capabilities
	//null
	//
	//properties
	//null

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

func (c *Container) validateProtocolHeader(t ProtocolType, buf []byte) (err error) {
	var amqpLiteral [4]byte
	copy(amqpLiteral[:], buf[:4])

	amqpProtocolType := byte(buf[4])
	amqpProtocolMajor := buf[5]
	amqpProtocolMinor := buf[6]
	amqpProtocolRev := buf[7]

	if amqpLiteral != [4]byte{65, 77, 81, 80} {
		err = errors.New("Invalid Header, not AMQP")
	}

	if amqpProtocolType != byte(t) {
		err = errors.New("Mismatched Protocol Type")
	}

	if amqpProtocolMajor != 1 {
		err = errors.New("Mismatched Protocol Version Major")
	}

	if amqpProtocolMinor != 0 {
		err = errors.New("Mismatched Protocol Version Minor")
	}

	if amqpProtocolRev != 0 {
		err = errors.New("Mismatched Protocol Version Revision")
	}
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
