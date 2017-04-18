package types

func NewChar(v string) *Char {
	return &Char {
		value: v,
	}
}

// a single unicode character
type Char struct {
	BaseAMQPType
	value string
}
