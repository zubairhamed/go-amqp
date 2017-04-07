package types

import (
	"encoding/binary"
	"errors"
)

type Map struct {
	BaseAMQPType
	values map[string]AMQPType
}

func (m *Map) Put(k string, v AMQPType) {
	m.values[k] = v
}

func (s *Map) Encode() ([]byte, uint, error) {
	return EncodeMapField(s)
}

func EncodeMapField(s *Map) ([]byte, uint, error) {
	if s == nil {
		return []byte { byte(TYPE_NULL) }, 1, nil
	}

	return nil, 0, nil
}

func DecodeMapField(v []byte) (val *Map, fieldLength uint, err error) {
	ctor := Type(v[0])
	if ctor == TYPE_NULL {
		fieldLength = 1
		return
	}

	if ctor != TYPE_MAP_8 && ctor != TYPE_MAP_32 {
		err = errors.New("Malformed error. Expecting map field")
		return
	}

	var valueLength uint
	var valueIdx uint
	if ctor == TYPE_MAP_8 {
		valueLength = uint(v[1])
		valueIdx = 2
	} else if ctor == TYPE_MAP_32 {
		valueLength = uint(binary.BigEndian.Uint32(v[1:4]))
		valueIdx = 5
	}

	fieldCount := int(v[valueIdx])
	fieldLength = valueLength + valueIdx

	remainingBytes := v[valueIdx+1:]
	val = &Map{
		values: make(map[string]AMQPType),
	}
	for i := 0; i < fieldCount/2; i++ {
		k, v, keyValueLength, e := DecodeKeyValueField(remainingBytes)

		if e != nil {
			err = e
			return
		}

		val.Put(k, v)

		remainingBytes = remainingBytes[keyValueLength:]
	}
	return
}

type KeyValue struct {
	BaseAMQPType
}

func EncodeKeyValueField(s *Symbol) ([]byte, uint, error) {
	return nil, 0, nil
}

func DecodeKeyValueField(v []byte) (key string, val AMQPType, fieldLength uint, err error) {
	ctor := Type(v[0])

	var readBytes uint
	var keyValueLength uint
	var valueIdx uint
	if ctor == TYPE_SYMBOL_8 {
		keyValueLength = uint(v[1])
		valueIdx = 2
	} else if ctor == TYPE_SYMBOL_8 {
		keyValueLength = uint(binary.BigEndian.Uint32(v[1:4]))
		valueIdx = 5
	} else {
		err = errors.New("Expecting first field as symbol.")
		return
	}

	key = string(v[valueIdx : keyValueLength+valueIdx])
	readBytes = keyValueLength + valueIdx
	val, valueFieldLength, err := DecodeField(v[readBytes:])
	if err != nil {
		return
	}

	fieldLength = readBytes + valueFieldLength
	return
}
