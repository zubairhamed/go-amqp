package types

import (
	"encoding/binary"
	"errors"
)

func DecodeStringField(v []byte) (val *String, fieldLength uint, err error) {

	ctor := Type(v[0])

	if ctor == TYPE_NULL {
		val = &String{
			BaseAMQPType: BaseAMQPType{
				encoding: TYPE_NULL,
			},
		}
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
		BaseAMQPType: BaseAMQPType{
			encoding: ctor,
		},
	}

	return
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