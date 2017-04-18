package types

func NewBinary(b []byte) *Binary {
	return &Binary{
		value: b,
	}
}

// A sequence of octets
type Binary struct {
	BaseAMQPType
	value []byte
}
