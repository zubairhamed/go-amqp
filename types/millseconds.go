package types

func NewMilliseconds(v uint32) *Milliseconds {
	return &Milliseconds{
		UInt: NewUInt(v),
	}
}

/*
<type name="milliseconds" class="restricted" source="uint"/>
*/
type Milliseconds struct {
	*UInt
}

func (s *Milliseconds) Encode() ([]byte, uint, error) {
	if s == nil {
		return []byte{byte(TYPE_NULL)}, 1, nil
	}
	return s.UInt.Encode()
}

func DecodeMillisecondsField(v []byte) (val *Milliseconds, fieldLength uint, err error) {
	iVal, fieldLength, err := DecodeUIntField(v)

	val = &Milliseconds{
		UInt: iVal,
	}

	return

}
