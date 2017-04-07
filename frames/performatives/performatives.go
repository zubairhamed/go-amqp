package performatives

import (
	"net"
	"encoding/binary"
	"github.com/zubairhamed/go-amqp/frames"
	"log"
)

type Performative interface {
	Encode() ([]byte, error)
}

func SendPerformative(c net.Conn, p Performative) (int, error) {
	b, err := p.Encode()
	if err != nil {
		panic(err.Error())
	}

	var frameSize uint32 = 8 + uint32(len(b))
	var frameSizeBytes = make([]byte, 4)
	binary.BigEndian.PutUint32(frameSizeBytes, frameSize)

	frameContent := frames.EncodeFrame(b)

	log.Println("SEND PERFORMTIVE", frameContent)

	c.Write(frameContent)

	return 0, nil
}

