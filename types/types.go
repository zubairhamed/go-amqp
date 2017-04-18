package types

import (
	"errors"
	"log"
)

type Type byte

const (
	TYPE_SIZE_1 = byte(1)
	TYPE_SIZE_2 = byte(2)
	TYPE_SIZE_4 = byte(4)
	TYPE_SIZE_8 = byte(8)
)

const (
	TYPE_CONSTRUCTOR              = Type(0x00)
	TYPE_PERFORMATIVE_OPEN        = Type(0x10)
	TYPE_PERFORMATIVE_BEGIN       = Type(0x11)
	TYPE_PERFORMATIVE_ATTACH      = Type(0x12)
	TYPE_PERFORMATIVE_FLOW        = Type(0x13)
	TYPE_PERFORMATIVE_TRANSFER    = Type(0x14)
	TYPE_PERFORMATIVE_DISPOSITION = Type(0x15)
	TYPE_PERFORMATIVE_DETACH      = Type(0x16)
	TYPE_PERFORMATIVE_END         = Type(0x17)
	TYPE_PERFORMATIVE_CLOSE       = Type(0x18)
	TYPE_NULL                     = Type(0x40)
	TYPE_BOOLEAN_TRUE             = Type(0x41)
	TYPE_BOOLEAN_FALSE            = Type(0x42)
	TYPE_UINT_0                   = Type(0x43)
	TYPE_ULONG_0                  = Type(0x44)
	TYPE_LIST_0                   = Type(0x45)
	TYPE_UBYTE                    = Type(0x50)
	TYPE_BYTE                     = Type(0x51)
	TYPE_UINT_SMALL               = Type(0x52)
	TYPE_ULONG_SMALL              = Type(0x53)
	TYPE_INT_SMALL                = Type(0x54)
	TYPE_LONG_SMALL               = Type(0x55)
	TYPE_BOOLEAN                  = Type(0x56)
	TYPE_USHORT                   = Type(0x60)
	TYPE_SHORT                    = Type(0x61)
	TYPE_UINT                     = Type(0x70)
	TYPE_INT                      = Type(0x71)
	TYPE_FLOAT                    = Type(0x72)
	TYPE_CHAR                     = Type(0x73)
	TYPE_DECIMAL32                = Type(0x74)
	TYPE_ULONG                    = Type(0x80)
	TYPE_LONG                     = Type(0x81)
	TYPE_DOUBLE                   = Type(0x82)
	TYPE_TIMESTAMP                = Type(0x83)
	TYPE_DECIMAl64                = Type(0x84)
	TYPE_DECIMAL128               = Type(0x94)
	TYPE_UUID                     = Type(0x98)
	TYPE_BINARY_VBIN8             = Type(0xA0)
	TYPE_BINARY_VBIN32            = Type(0xB0)
	TYPE_STRING_8_UTF8            = Type(0xA1)
	TYPE_STRING_32_UTF8           = Type(0xB1)
	TYPE_SYMBOL_8                 = Type(0xA3)
	TYPE_SYMBOL_32                = Type(0xB3)
	TYPE_LIST_8                   = Type(0xC0)
	TYPE_LIST_32                  = Type(0xD0)
	TYPE_MAP_8                    = Type(0xC1)
	TYPE_MAP_32                   = Type(0xD1)
	TYPE_ARRAY_8                  = Type(0xE0)
	TYPE_ARRAY_32                 = Type(0xF0)
)

func EncodeField(t AMQPType) (b []byte, l uint, err error) {
	log.Println("Encoding", t)
	if t == nil {
		err = errors.New("nil field found")
		return
	}
	return t.Encode()
}

func DecodeField(v []byte) (AMQPType, uint, error) {
	ctor := Type(v[0])

	switch {
	case ctor == TYPE_SYMBOL_8 || ctor == TYPE_SYMBOL_32:
		return DecodeSymbolField(v)

	case ctor == TYPE_STRING_8_UTF8 || ctor == TYPE_STRING_32_UTF8:
		return DecodeStringField(v)
	}

	return nil, 0, errors.New("Unknown field found")
}

type AMQPType interface {
	GetType() Type
	Encode() ([]byte, uint, error)
	Stringify() string
}

type BaseAMQPType struct {
	encoding Type
}

func (b *BaseAMQPType) GetType() Type {
	return b.encoding
}
