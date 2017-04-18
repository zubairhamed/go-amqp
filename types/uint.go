package types

import (
	"encoding/binary"
	"errors"
	"fmt"
)

func NewUInt(v uint32) *UInt {
	return &UInt{
		value: v,
	}
}

type UInt struct {
	*BaseAMQPType
	value uint32
}

func (b *UInt) Stringify() string {
	return fmt.Sprint(b.Value())
}

func (s *UInt) Encode() ([]byte, uint, error) {
	if s == nil {
		return []byte{byte(TYPE_NULL)}, 1, nil
	}
	return EncodeUIntField(s)
}

func (s *UInt) Value() uint32 {
	return s.value
}

func EncodeUIntField(s *UInt) ([]byte, uint, error) {
	if s == nil {
		return []byte{byte(TYPE_UINT_0)}, 1, nil
	}

	v := s.value
	b := []byte{}

	switch {
	case v == 0:
		b = append(b, byte(TYPE_UINT_0))

	case v > 255:
		b = append(b, byte(TYPE_UINT))

		byteVal := make([]byte, TYPE_SIZE_4)
		binary.BigEndian.PutUint32(byteVal, v)
		b = append(b, byteVal...)

	case v < 256:
		b = append(b, []byte{byte(TYPE_UINT_SMALL), byte(v)}...)
	}

	return b, uint(len(b)), nil
}

func DecodeUIntField(v []byte) (val *UInt, fieldLength uint, err error) {
	ctor := Type(v[0])
	if ctor == TYPE_NULL {
		fieldLength = 1
		return
	}

	if ctor != TYPE_UINT && ctor != TYPE_UINT_0 && ctor != TYPE_UINT_SMALL {
		err = errors.New("Malformed error. Expecting uint field")
		return
	}

	var fieldValue uint32

	switch {
	case ctor == TYPE_UINT_0:
		fieldLength = 1
		fieldValue = 0
		break

	case ctor == TYPE_UINT:
		fieldLength = 5
		fieldValue = binary.BigEndian.Uint32(v[1:5])
		break

	case ctor == TYPE_UINT_SMALL:
		fieldLength = 2
		fieldValue = binary.BigEndian.Uint32(v[1:2])
		break
	}

	val = &UInt{
		value: fieldValue,
	}
	return
}
