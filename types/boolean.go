package types

func NewBoolean(b bool) *Boolean {
	return &Boolean{
		value: b,
	}
}

// Represents a true or false value
type Boolean struct {
	BaseAMQPType
	value bool
}
