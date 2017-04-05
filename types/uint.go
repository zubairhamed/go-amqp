package types

import (
	"encoding/binary"
	"errors"
)

type UInt struct {
	BaseAMQPType
	value uint32
}

func DecodeUIntField(v []byte) (val *UInt, fieldLength uint, err error) {
	ctor := Type(v[0])
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
		BaseAMQPType: BaseAMQPType{
			encoding: ctor,
		},
	}
	return
}
