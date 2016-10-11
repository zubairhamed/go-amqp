package amqp

import "github.com/zubairhamed/go-amqp/primitives"

func NewOpenPerformative(containerId, hostname string) *PerformativeOpen {
	return &PerformativeOpen{
		ContainerId: containerId,
		Hostname:    hostname,
	}
}

type PerformativeOpen struct {
	ContainerId         string
	Hostname            string
	MaxFrameSize        primitives.TypeFormatCode
	ChannelMax          primitives.TypeFormatCode
	IdleTimeout         primitives.TypeFormatCode
	OutgoingLocales     []string
	IncomingLocales     []string
	OfferedCapabilities []string
	DesiredCapabilities []string
	Properties          primitives.TypeFormatCode
}

func (p *PerformativeOpen) Encode() []byte {
	out := []byte{}

	return out
}
