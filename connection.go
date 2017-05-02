package amqp

import (
	"encoding/binary"
	"errors"
	"fmt"
	. "github.com/zubairhamed/go-amqp/frames"
	. "github.com/zubairhamed/go-amqp/frames/performatives"
	. "github.com/zubairhamed/go-amqp/types"
	"log"
	"net"
)

func NewConnectInfo(url, nodeAddress string) *ConnectInfo {
	return &ConnectInfo{
		url:         url,
		nodeAddress: nodeAddress,
	}
}

type ConnectInfo struct {
	nodeAddress string
	url         string
}

func NewConnection(c *ConnectInfo) *Connection {
	return &Connection{
		connectInfo: c,
		connected:   false,
	}
}

type Connection struct {
	connectInfo *ConnectInfo
	netConn     net.Conn
	connected   bool
	name        string
}

func (c *Connection) doConnect(fn func(b []byte), connName string) (err error) {

	c.name = connName

	// Connect
	conn, err := net.Dial("tcp", c.connectInfo.url)
	if err != nil {
		panic(err)
	}

	c.netConn = conn

	var readBuf []byte

	// Handshake
	LogOut("HANDSHAKE", c.name)
	SendHandshake(conn)
	readBuf, err = ReadFromConnection(conn)
	err = HandleHandshake(readBuf)
	if err != nil {
		panic(err.Error())
	}

	// Send Open Performative
	openPerformative := NewOpenPerformative()
	openPerformative.ContainerId = NewString("ContainerID-" + c.name)

	LogOut("OPEN", c.name)
	_, err = c.SendPerformative(openPerformative)
	if err != nil {
		panic(err.Error())
	}

	// dispatch loop
	go c.handleMessages(conn, fn)

	return
}

func (c *Connection) handleMessages(conn net.Conn, fn func(b []byte)) {

	for {
		buf := []byte{}
		tmp := make([]byte, 2014)
		bytesRead, err := conn.Read(tmp)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}

		buf = append(buf, tmp[:bytesRead]...)

		_, buf, err = c.extractFrameData(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}
		for len(buf) > 0 {
			// Get frame`
			l, fb, err := c.extractFrame(buf)
			if err != nil {
				log.Println("An error occcured dispatching frame", err.Error())
			}

			buf = buf[l:]

			// Dispatch frame
			fn(fb)
		}
	}
}

func (c *Connection) extractFrame(b []byte) (n int, fr []byte, err error) {
	if Type(b[0]) != TYPE_CONSTRUCTOR {
		err = errors.New("Malformed or unexpected frame. Expecting constructor.")
		return
	}

	if Type(b[1]) != TYPE_ULONG_SMALL {
		err = errors.New("Malformed or unexpected frame. Expecting small ulong type")
		return
	}

	perf := Type(b[2])
	if perf != TYPE_PERFORMATIVE_ATTACH &&
		perf != TYPE_PERFORMATIVE_END &&
		perf != TYPE_PERFORMATIVE_OPEN &&
		perf != TYPE_PERFORMATIVE_BEGIN &&
		perf != TYPE_PERFORMATIVE_DISPOSITION &&
		perf != TYPE_PERFORMATIVE_FLOW &&
		perf != TYPE_PERFORMATIVE_TRANSFER &&
		perf != TYPE_PERFORMATIVE_CLOSE &&
		perf != TYPE_PERFORMATIVE_DETACH {

		err = errors.New("Malformed or unexpected frame. Expecting a Performative")
	}

	if Type(b[3]) != TYPE_LIST_8 {
		err = errors.New("Malformed or unexpected frame. Expecting list 8")
		return
	}

	n = int(b[4]) + 5
	fr = b[:n]

	return
}

func (c *Connection) extractFrameData(b []byte) (n int, fr []byte, err error) {
	f, err := UnmarshalFrameHeader(b)
	if err != nil {
		return
	}

	doff := f.DataOffset
	if uint32(len(b)) < f.Size {
		err = errors.New("Malformed frame. Invalid size")
		return
	}

	fr = b[doff*4 : f.Size]
	n = len(fr)

	return
}

func (c *Connection) Close() {
	log.Println("Connection:Close")
}

func (c *Connection) Write(b []byte) (int, error) {
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
