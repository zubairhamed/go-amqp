package types

import "encoding/binary"

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
		return []byte { byte(TYPE_NULL) }, 1, nil
	}
	return EncodeULongField(s)
}

func EncodeULongField(s *ULong) ([]byte, uint, error) {
	if s == nil {
		return []byte { byte(TYPE_ULONG_0) }, 1, nil
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
		b = append(b, []byte{ byte(TYPE_ULONG_SMALL),  byte(v) }...)
	}

	return b, uint(len(b)), nil
}
