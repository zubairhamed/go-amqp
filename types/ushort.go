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
	// 0x60
}

type UShort struct {
	BaseAMQPType
	value uint16
}


