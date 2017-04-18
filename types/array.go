package types

// A sequence of values of a single type
type Array struct {
	BaseAMQPType
}

func (b *Array) Stringify() string {
	return ""
}
