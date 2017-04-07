package performatives

import (
	"encoding/binary"
	"github.com/zubairhamed/go-amqp/frames"
	"github.com/zubairhamed/go-amqp/types"
	"net"
)

type Performative interface {
	types.AMQPType
}

func SendPerformative(c net.Conn, p Performative) (int, error) {
	b, _, err := p.Encode()
	if err != nil {
		panic(err.Error())
	}

	var frameSize uint32 = 8 + uint32(len(b))
	var frameSizeBytes = make([]byte, 4)
	binary.BigEndian.PutUint32(frameSizeBytes, frameSize)

	frameContent := frames.EncodeFrame(b)

	c.Write(frameContent)

	return 0, nil
}
