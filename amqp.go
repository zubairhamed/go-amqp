package amqp

type ProtocolType byte

const (
	ProtocolTypeNone = ProtocolType(0)
	ProtocolTypeSsl  = ProtocolType(2)
	ProtocolTypeSasl = ProtocolType(3)
)

