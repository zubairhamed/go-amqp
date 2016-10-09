package amqp

func NewPerformativeOpen(containerId, hostname string) *PerformativeOpen {
	return &PerformativeOpen{
		containerId: containerId,
		hostname: hostname,
	}
}

type PerformativeOpen struct {
	containerId string
	hostname    string
}

func (p *PerformativeOpen) Encode() []byte {
	out := []byte {}

	return out
}


