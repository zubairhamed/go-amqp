package amqp

type ProtocolType byte

const (
	ProtocolTypeNone = ProtocolType(0)
	ProtocolTypeSsl  = ProtocolType(2)
	ProtocolTypeSasl = ProtocolType(3)
)

const (
	EVENT_MSG_OPEN        = EventType(0)
	EVENT_MSG_BEGIN       = EventType(1)
	EVENT_MSG_ATTACH      = EventType(2)
	EVENT_MSG_FLOW        = EventType(3)
	EVENT_MSG_TRANSFER    = EventType(4)
	EVENT_MSG_DISPOSITION = EventType(5)
	EVENT_MSG_DETACH      = EventType(6)
	EVENT_MSG_END         = EventType(7)
	EVENT_MSG_CLOSE       = EventType(8)

	EVENT_ERROR = EventType(255)
)

type EventType byte

type RoleType byte

const (
	ROLE_SENDER   = 0
	ROLE_RECEIVER = 1
)
