package types

import (
	"encoding/binary"
	"errors"
	"fmt"
)

func NewULong(v uint64) *ULong {
	return &ULong{
		value: v,
	}
}

type ULong struct {
	BaseAMQPType
	value uint64
}

func (s *ULong) Encode() ([]byte, uint, error) {
	if s == nil {
		return []byte{byte(TYPE_NULL)}, 1, nil
	}
	return EncodeULongField(s)
}

func (b *ULong) Stringify() string {
	return fmt.Sprint(b)
}

func (s *ULong) Value() uint64 {
	return s.value
}

func EncodeULongField(s *ULong) ([]byte, uint, error) {
	if s == nil {
		return []byte{byte(TYPE_ULONG_0)}, 1, nil
	}

	v := s.value
	b := []byte{}

	switch {
	case v == 0:
		b = append(b, byte(TYPE_ULONG_0))

	case v > 255:
		b = append(b, byte(TYPE_ULONG))

		byteVal := make([]byte, TYPE_SIZE_8)
		binary.BigEndian.PutUint64(byteVal, v)
		b = append(b, byteVal...)

	case v < 256:
		b = append(b, []byte{byte(TYPE_ULONG_SMALL), byte(v)}...)
	}

	return b, uint(len(b)), nil
}

func DecodeULongField(v []byte) (val *ULong, fieldLength uint, err error) {
	ctor := Type(v[0])
	if ctor == TYPE_NULL {
		fieldLength = 1
		return
	}

	if ctor != TYPE_ULONG && ctor != TYPE_ULONG_0 && ctor != TYPE_ULONG_SMALL {
		err = errors.New("Malformed error. Expecting ulong field")
		return
	}

	var fieldValue uint64

	switch {
	case ctor == TYPE_ULONG_0:
		fieldLength = 1
		fieldValue = 0
		break

	case ctor == TYPE_ULONG:
		fieldLength = 9
		fieldValue = binary.BigEndian.Uint64(v[1:9])
		break

	case ctor == TYPE_ULONG_SMALL:
		fieldLength = 2
		fieldValue = uint64(v[0])
		break
	}

	val = &ULong{
		value: fieldValue,
	}
	return
}
