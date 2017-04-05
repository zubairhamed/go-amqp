package types

import (
	"errors"
)

type Symbol struct {
	BaseAMQPType
	value string
}

func DecodeSymbolField(v []byte) (val *Symbol, fieldLength uint, err error) {
	ctor := Type(v[0])

	if ctor == TYPE_NULL {
		val = &Symbol{
			BaseAMQPType: BaseAMQPType{
				encoding: ctor,
			},
		}
		fieldLength = 1
		return
	}

	if ctor != TYPE_SYMBOL_8 && ctor != TYPE_SYMBOL_32 {
		err = errors.New("Malformed error. Expecting symbol field")
		return
	}

	return
}

func DecodeSymbolArrayField(v []byte) (val []*Symbol, fieldLength uint, err error) {
	ctor := Type(v[0])

	if ctor == TYPE_NULL {
		val = []*Symbol{}
		fieldLength = 1
		return
	}
	return
}

