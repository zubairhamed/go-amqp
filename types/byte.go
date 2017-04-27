package types

func NewByte(v byte) *Byte {
	return &Byte{
		value: v,
	}
}

//  integer in the range  (27) to 27 - 1 inclusive
type Byte struct {
	BaseAMQPType
	value byte
}

func DecodeAnyTypeField(v []byte) (val []byte, fieldLength uint, err error) {
	return
}
