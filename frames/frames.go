package frames

import (
	"bytes"
	"encoding/binary"
	"errors"
	. "github.com/zubairhamed/go-amqp/types"
	"log"
	"net"
	"github.com/zubairhamed/go-amqp/util"
)

const MINIMUM_FRAME_SIZE = 8

type FrameType byte

const (
	FrameTypeAmqp = FrameType(0x00)
	FrameTypeSasl = FrameType(0x01)
)

type Frame struct {
	header    *FrameHeader
	extHeader *ExtendedHeader
	body      *FrameBody
}

func NewFrameHeader(d uint8, f FrameType, c uint16, s uint32) *FrameHeader {
	fh := &FrameHeader{}

	fh.Channel = c
	fh.DataOffset = d
	fh.FrameType = f
	fh.Size = s

	return fh
}

type FrameHeader struct {
	DataOffset uint8
	FrameType  FrameType
	Channel    uint16
	Size       uint32
}

func (h *FrameHeader) GetSize() (sz uint32, err error) {
	return
}

type ExtendedHeader struct {
}

type FrameBody struct {
}

func UnmarshalProtocolHeader(b []byte) (error, *ProtocolHeader) {
	if len(b) != 8 {
		return errors.New("Invalid header size. Expecting 8 bytes"), nil
	}

	var amqpLiteral [4]byte
	copy(amqpLiteral[:], b[:4])
	amqpProtocolType := byte(b[4])
	amqpProtocolMajor := b[5]
	amqpProtocolMinor := b[6]
	amqpProtocolRev := b[7]

	if amqpLiteral != [4]byte{65, 77, 81, 80} {
		return errors.New("Invalid Header, not AMQP"), nil
	}

	if amqpProtocolMajor != 1 {
		return errors.New("Mismatched Protocol Version Major"), nil
	}

	if amqpProtocolMinor != 0 {
		return errors.New("Mismatched Protocol Version Minor"), nil
	}

	if amqpProtocolRev != 0 {
		return errors.New("Mismatched Protocol Version Revision"), nil
	}

	h := &ProtocolHeader{}
	h.ProtocolType = amqpProtocolType
	h.ProtocolMajor = amqpProtocolMajor
	h.ProtocolMinor = amqpProtocolMinor
	h.ProtocolRevision = amqpProtocolRev

	return nil, h
}

type ProtocolHeader struct {
	ProtocolType     byte
	ProtocolMajor    byte
	ProtocolMinor    byte
	ProtocolRevision byte
}

type AMQPFrame struct {
}

func UnmarshalFrameHeader(b []byte) (f *FrameHeader, err error) {
	bLen := len(b)
	if bLen < MINIMUM_FRAME_SIZE {
		err = errors.New("Malformed frame. Invalid frame size")
		return
	}

	sz := binary.BigEndian.Uint32(b[:4])

	doff := b[4]
	if doff < 2 {
		err = errors.New("Malformed frame. Data offset less than 2")
		return
	}

	var ft FrameType
	ftByte := b[5]
	if ftByte == 0 {
		ft = FrameTypeAmqp
	} else {
		ft = FrameTypeSasl
	}

	f = NewFrameHeader(doff, ft, 0, sz)

	return
}

func EncodeFrame(b []byte) (fb []byte) {
	var frameSize uint32 = 8 + uint32(len(b))
	var frameSizeBytes = make([]byte, 4)
	binary.BigEndian.PutUint32(frameSizeBytes, frameSize)

	frameContent := []byte{}
	frameContent = append(frameContent, frameSizeBytes...)
	frameContent = append(frameContent, []byte{0x02, 0x00, 0x00, 0x00}...)
	frameContent = append(frameContent, b...)

	return frameContent
}

var HANDSHAKE_MSG = []byte{0x41, 0x4d, 0x51, 0x50, 0x00, 0x01, 0x00, 0x00}

func SendHandshake(c net.Conn) (int, error) {
	c.Write(HANDSHAKE_MSG)

	return 0, nil
}

func HandleHandshake(b []byte) error {
	log.Println(util.ToHex(b))
	if !bytes.Equal(b[0:8], HANDSHAKE_MSG) {
		log.Println(b)
		return errors.New("Invalid handshake message")
	}
	return nil
}

func ValidateFrame(b []byte) (err error) {
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

	return
}

func ExtractPerformative(b []byte) (n int, fr []byte, err error) {
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
