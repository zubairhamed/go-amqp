package types

import "encoding/binary"

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
	v := s.value
	b := []byte{}

	switch {
	case v == 0:
		b = append(b, byte(TYPE_ULONG_0))

	case v > 255:
		b = append(b, byte(TYPE_ULONG))

		var vb []byte
		binary.BigEndian.PutUint64(vb, v)
		b = append(b, vb...)

	case v < 256:
		b = append(b, []byte{ byte(TYPE_ULONG_SMALL),  byte(v) }...)
	}

	return b, uint(len(b)), nil
}
