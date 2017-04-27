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

func (s *Binary) Encode() ([]byte, uint, error) {
	if s == nil {
		return NullValue()
	}
	return EncodeBinaryField(s)
}

func EncodeBinaryField(s *Binary) ([]byte, uint, error) {
	b := s.value

	if len(b) == 0 {
		return NullValue()
	}

	return b, uint(len(b)), nil
}
