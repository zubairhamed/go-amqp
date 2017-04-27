package types

import (
	"log"
)

func NewBoolean(b bool) *Boolean {
	return &Boolean{
		value: b,
	}
}

// Represents a true or false value
type Boolean struct {
	BaseAMQPType
	value bool
}

func (s *Boolean) Encode() ([]byte, uint, error) {
	if s == nil {
		return []byte{byte(TYPE_NULL)}, 1, nil
	}
	return EncodeBooleanField(s)
}

func (b *Boolean) Stringify() string {
	if b.value {
		return "true"
	}
	return "false"
}

func EncodeBooleanField(s *Boolean) ([]byte, uint, error) {
	if s == nil {
		return []byte{byte(TYPE_UINT_0)}, 1, nil
	}

	v := s.value
	b := []byte{}

	if v {
		b = append(b, byte(TYPE_BOOLEAN_TRUE))
	} else {
		b = append(b, byte(TYPE_BOOLEAN_FALSE))
	}

	return b, uint(len(b)), nil
}

func DecodeBooleanField(v []byte) (val *Boolean, fieldLength uint, err error) {
	ctor := Type(v[0])

	switch ctor {
	case TYPE_NULL:
		fieldLength = 1

	case TYPE_BOOLEAN:
		fieldLength = 2
		boolVal := v[1]

		if boolVal == 0x00 {
			val = NewBoolean(false)
		} else if boolVal == 0x01 {
			val = NewBoolean(true)
		} else {
			log.Println("ERROR: Invalid BOolean Value found")
		}

	case TYPE_BOOLEAN_FALSE:
		val = NewBoolean(false)
		fieldLength = 1

	case TYPE_BOOLEAN_TRUE:
		val = NewBoolean(true)
		fieldLength = 1
	}

	return
}
