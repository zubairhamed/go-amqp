package types

import (
	"errors"
)

type Timestamp struct {
	BaseAMQPType
	value uint64
}

func DecodeTimestampField(v []byte) (val *Timestamp, fieldLength int, err error) {
	ctor := Type(v[0])
	if ctor != TYPE_TIMESTAMP {
		err = errors.New("Malformed error. Expecting timestamp field")
		return
	}
	return
}
