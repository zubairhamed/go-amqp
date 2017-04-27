package types

func NewHandle(v uint32) *Handle {
	return &Handle{
		UInt: NewUInt(v),
	}
}

/*
<type name="handle" class="restricted" source="uint"/>
*/
type Handle struct {
	*UInt
}

func DecodeHandleField(v []byte) (val *Handle, fieldLength uint, err error) {
	iVal, fieldLength, err := DecodeUIntField(v)

	val = &Handle{
		UInt: iVal,
	}

	return
}
