package types

func NewSequenceNumber(v uint32) *SequenceNumber {
	return &SequenceNumber{
		UInt: NewUInt(v),
	}
}

type SequenceNumber struct {
	*UInt
}

func DecodeSequenceNumber(v []byte) (val *SequenceNumber, fieldLength uint, err error) {
	v2, fieldLength, err := DecodeUIntField(v)

	val = &SequenceNumber{
		UInt: v2,
	}
	return
}
