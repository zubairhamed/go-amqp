package types

import (
	"errors"
)

type Symbol struct {
	BaseAMQPType
	value string
}

func (s *Symbol) Encode() ([]byte, uint, error) {
	if s == nil {
		return []byte { byte(TYPE_NULL) }, 1, nil
	}
	return EncodeSymbolField(s)
}

func EncodeSymbolField(s *Symbol) ([]byte, uint, error) {
	return nil, 0, nil
}

func DecodeSymbolField(v []byte) (val *Symbol, fieldLength uint, err error) {
	ctor := Type(v[0])

	if ctor == TYPE_NULL {
		val = &Symbol{}
		fieldLength = 1
		return
	}

	if ctor != TYPE_SYMBOL_8 && ctor != TYPE_SYMBOL_32 {
		err = errors.New("Malformed error. Expecting symbol field")
		return
	}

	return
}

func EncodeSymbolArrayField(v []*Symbol) (enc []byte, l uint, err error) {
	return nil, 0, nil
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

