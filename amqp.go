package amqp

type ProtocolType byte

const (
	ProtocolTypeNone = ProtocolType(0)
	ProtocolTypeSsl  = ProtocolType(2)
	ProtocolTypeSasl = ProtocolType(3)
)

type FrameType byte

const (
	FrameTypeAmqp = FrameType(0x00)
	FrameTypeSasl = FrameType(0x01)
)

const MINIMUM_FRAME_SIZE = 8
