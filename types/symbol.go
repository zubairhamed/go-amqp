package types

import (
	"encoding/binary"
	"errors"
)

func NewSymbol(v string) *Symbol {
	return &Symbol{
		value: v,
	}
}

type Symbol struct {
	BaseAMQPType
	value string
}

func (s *Symbol) Encode() ([]byte, uint, error) {
	if s == nil {
		return []byte{byte(TYPE_NULL)}, 1, nil
	}
	return EncodeSymbolField(s)
}

func (b *Symbol) Stringify() string {
	return b.value
}

func EncodeSymbolField(s *Symbol) ([]byte, uint, error) {
	if s == nil {
		return []byte{byte(TYPE_NULL)}, 1, nil
	}

	v := s.value
	b := []byte{}

	vlen := len(v)

	switch {
	case vlen == 0 || vlen < 256:
		b = append(b, byte(TYPE_SYMBOL_8))
		b = append(b, byte(vlen))

	case vlen > 255 && vlen < 4294967295:
		b = append(b, byte(TYPE_SYMBOL_32))

		byteVal := make([]byte, TYPE_SIZE_4)
		binary.BigEndian.PutUint32(byteVal, uint32(vlen))
		b = append(b, byteVal...)
	}

	b = append(b, []byte(v)...)

	return b, uint(len(b)), nil
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

	offset := uint(2)
	valueLength := uint(v[1])
	fieldLength = valueLength + offset

	strVal := string(v[2:fieldLength])

	val = NewSymbol(strVal)

	return
}

// Array Type
type SymbolArray struct {
	BaseAMQPType
	value []*Symbol
}

func (s *SymbolArray) Encode() ([]byte, uint, error) {
	if s == nil {
		return NullValue()
	}
	return EncodeSymbolArrayField(s)
}

func (b *SymbolArray) Stringify() string {
	return "String: SymbolArray"
}

func (s *SymbolArray) Append(symbol *Symbol) {
	s.value = append(s.value, symbol)
}

func (s *SymbolArray) Length() int {
	return len(s.value)
}

func EncodeSymbolArrayField(v *SymbolArray) (enc []byte, l uint, err error) {
	if v.Length() == 0 {
		return NullValue()
	}

	return nil, 0, nil
}

func DecodeSymbolArrayField(v []byte) (val *SymbolArray, fieldLength uint, err error) {
	ctor := Type(v[0])

	if ctor == TYPE_NULL {
		val = &SymbolArray{}
		fieldLength = 1
		return
	}

	return
}
