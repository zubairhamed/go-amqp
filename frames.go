package amqp

type Frame struct {
	header    *FrameHeader
	extHeader *ExtendedHeader
	body      *FrameBody
}

type FrameHeader struct {
	dataOffset uint8
	frameType  FrameType
	channel    uint16
}

func (h *FrameHeader) GetSize() (sz uint32, err error) {
	return
}

type ExtendedHeader struct {
}

type FrameBody struct {
}
