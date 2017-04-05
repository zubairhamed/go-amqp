package frames

import (
	"errors"
	"encoding/binary"
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

func UnmarshalAMQPFrame(b []byte) (error, *AMQPFrame) {

	return nil, nil
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