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
	"sync"
	"github.com/zubairhamed/go-amqp/util"
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
	beginWg     sync.WaitGroup
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
	LogIn("HANDSHAKE", c.name)
	err = HandleHandshake(readBuf)
	if err != nil {
		panic(err.Error())
	}

	// Send Open Performative
	openPerformative := NewOpenPerformative()
	openPerformative.ContainerId = NewString("ContainerID-" + c.name)

	// TODO (?)
	openPerformative.MaxFrameSize = NewUInt(16384)

	LogOut("OPEN", c.name)
	_, err = c.SendPerformative(openPerformative)
	if err != nil {
		panic(err.Error())
	}

	// dispatch loop
	c.beginWg.Add(1)
	go c.handleMessages(conn, fn)
	c.beginWg.Wait()

	return
}

func (c *Connection) handleMessages(conn net.Conn, fn func(b []byte)) {

	for {
		buf := []byte{}
		tmp := make([]byte, 2014)
		bytesRead, err := conn.Read(tmp)
		log.Println("-------- Reading of Bytes -------")
		log.Println("BYTES READ: ", util.ToHex(tmp))
		log.Println("BYTES READ HEX: ", util.ToHex(tmp[:bytesRead]))
		log.Println("BYTES READ STRING: ", string(tmp[:bytesRead]))
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}

		// Extract frame data and append to buf
		// Extract Frame Data
		_, b, err := c.extractFrameData(tmp)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}

		buf = append(buf, b...)

		for len(buf) > 0 {
			// Get frame
			l, fb, err := c.extractPerformative(buf)
			log.Println("read in frame length", l)
			if err != nil {
				log.Println("An error occcured dispatching frame", err.Error())
			}

			buf = buf[l:]
			log.Println("New Buffer is", util.ToHex(buf))

			// Dispatch frame
			log.Println("Dispatching", util.ToHex(fb))
			log.Println("Dispatching (string)", string(fb))
			fn(fb)
		}
	}
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

func (c *Connection) extractPerformative(b []byte) (n int, fr []byte, err error) {
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