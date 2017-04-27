package types

type UByte struct {
	BaseAMQPType
	value byte
}

func (s *UByte) Encode() ([]byte, uint, error) {
	return EncodeUByteField(s)
}

func EncodeUByteField(s *UByte) ([]byte, uint, error) {
	b := s.value

	return []byte{byte(TYPE_UBYTE), b}, 2, nil
}
