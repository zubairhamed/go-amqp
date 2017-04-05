package types

import (
	"encoding/binary"
	"errors"
)

func DecodeUShortField(v []byte) (val *UShort, fieldLength uint, err error) {
	ctor := Type(v[0])
	if ctor != TYPE_USHORT {
		err = errors.New("Malformed error. Expecting ushort field")
		return
	}

	fieldLength = 3
	val = &UShort{
		value: binary.BigEndian.Uint16(v[1:3]),
	}
	return
}

func EncodeUShortField(s *UShort) (b []byte, l uint, err error) {
	if s == nil {
		return []byte { byte(TYPE_USHORT), 0, 0 }, 3, nil
	}

	v := s.value
	b = []byte{ byte(TYPE_USHORT) }

	byteVal := make([]byte, TYPE_SIZE_2)
	binary.BigEndian.PutUint16(byteVal, v)
	b = append(b, byteVal...)

	return b, uint(len(b)), nil
}

func NewUShort(v uint16) *UShort {
	return &UShort{
		value: v,
	}
}

type UShort struct {
	BaseAMQPType
	value uint16
}

func (s *UShort) Encode() ([]byte, uint, error) {
	if s == nil {
		return []byte { byte(TYPE_NULL) }, 1, nil
	}
	return EncodeUShortField(s)
}

func (b *UShort) GetType() Type {
	return TYPE_USHORT
}

