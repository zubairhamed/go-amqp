package types

import (
	"encoding/binary"
	"errors"
)

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
		valueLength = uint(v[1])
		fieldLength = valueLength + 2
		strValue = string(v[2:valueLength])

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
		return []byte { byte(TYPE_NULL) }, 1, nil
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

		byteVal := make([]byte, 2)
		binary.BigEndian.PutUint16(byteVal, uint16(vlen))
		b = append(b, byteVal...)
	}

	b = append(b, []byte(v)...)

	return b, uint(len(b)), nil
}


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
