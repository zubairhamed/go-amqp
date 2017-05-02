package types

import (
	"encoding/binary"
	"errors"
)

func NewString(v string) *String {
	return &String{
		value: v,
	}
}

type String struct {
	BaseAMQPType
	value string
}

func (s *String) Value() string {
	return s.value
}

func (s *String) Encode() ([]byte, uint, error) {
	return EncodeStringField(s)
}

func (b *String) Stringify() string {
	return b.Value()
}

func (b *String) GetType() Type {
	v := b.value
	vlen := len(v)

	if vlen == 0 || vlen < 256 {
		return TYPE_STRING_8_UTF8
	} else if vlen > 255 && vlen < 4294967295 {
		return TYPE_STRING_32_UTF8
	}
	return Type(0x00)
}

func DecodeStringField(v []byte) (val *String, fieldLength uint, err error) {
	ctor := Type(v[0])

	if ctor == TYPE_NULL {
		val = &String{}
		fieldLength = 1
		return
	}

	if ctor != TYPE_STRING_8_UTF8 && ctor != TYPE_STRING_32_UTF8 {
		err = errors.New("Malformed error. Expecting string field")
		return
	}

	var valueLength uint
	var strValue string
	if ctor == TYPE_STRING_8_UTF8 {
		offset := uint(2)
		valueLength = uint(v[1])
		fieldLength = valueLength + offset
		strValue = string(v[offset : valueLength+offset])

	} else if ctor == TYPE_STRING_8_UTF8 {
		valueLength = uint(binary.BigEndian.Uint32(v[1:4]))
		fieldLength = valueLength + 5
		strValue = string(v[5:valueLength])
	}

	val = &String{
		value: string(strValue),
	}

	return
}

func EncodeStringField(s *String) ([]byte, uint, error) {
	if s == nil {
		return NullValue()
	}

	v := s.value
	b := []byte{}

	vlen := len(v)

	switch {
	case vlen == 0 || vlen < 256:
		b = append(b, byte(TYPE_STRING_8_UTF8))
		b = append(b, byte(vlen))

	case vlen > 255 && vlen < 4294967295:
		b = append(b, byte(TYPE_STRING_32_UTF8))

		byteVal := make([]byte, TYPE_SIZE_4)
		binary.BigEndian.PutUint32(byteVal, uint32(vlen))
		b = append(b, byteVal...)
	}
	b = append(b, []byte(v)...)

	return b, uint(len(b)), nil
}
