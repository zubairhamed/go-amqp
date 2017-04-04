package amqp

import "errors"

type Frame struct {
	header    *FrameHeader
	extHeader *ExtendedHeader
	body      *FrameBody
}

func NewFrameHeader(d uint8, f FrameType, c uint16, s uint32) *FrameHeader {
	fh := &FrameHeader{}

	fh.channel = c
	fh.dataOffset = d
	fh.frameType = f
	fh.size = s

	return fh
}

type FrameHeader struct {
	dataOffset uint8
	frameType  FrameType
	channel    uint16
	size       uint32
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
