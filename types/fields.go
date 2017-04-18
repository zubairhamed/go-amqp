package types


func NewFields(v map[string]AMQPType) *Fields {
	return &Fields{
		Map: NewMap(v),
	}
}

type Fields struct {
	*Map
}

func (s *Fields) Encode() ([]byte, uint, error) {
	if s == nil {
		return []byte{byte(TYPE_NULL)}, 1, nil
	}

	return s.Map.Encode()
}


func DecodeFieldsField(v []byte) (val *Fields, fieldLength uint, err error) {
	iVal, fieldLength, err := DecodeMapField(v)

	val = &Fields{
		Map: iVal,
	}

	return
}