package amqp

import "github.com/zubairhamed/go-amqp/primitives"

func NewPerformativeOpen(containerId, hostname string) *PerformativeOpen {
	return &PerformativeOpen{
		ContainerId: containerId,
		Hostname:    hostname,
	}
}

type PerformativeOpen struct {
	ContainerId         primitives.StringType
	Hostname            primitives.StringType
	MaxFrameSize        primitives.UIntType
	ChannelMax          primitives.UShortType
	IdleTimeOut         primitives.UIntType
	OutgoingLocales     primitives.ListType
	IncomingLocales     primitives.ListType
	OfferedCapabilities primitives.ListType
	DesiredCapabilities primitives.ListType
	Properties          primitives.MapType
}

func (p *PerformativeOpen) Encode() []byte {
	out := []byte{}

	return out
}
