package amqp

import (
	"encoding/binary"
	"errors"
	"log"
)

func GetField(v []byte) (val []byte, fieldLength uint, err error) {
	ctor := val[0]


	var t AMQPType

	switch {

	}

	fieldLength = uint(v[1])

	if uint(len(v)) < fieldLength {
		err = errors.New("Malformed field. Not enough bytes in field declared by field length")
		return
	}
	val = v[2:fieldLength]
	fieldLength = fieldLength + 2

	return
}

func GetUIntField(v []byte) (val *UInt, fieldLength uint, err error) {
	ctor := Type(v[0])
	if ctor != TYPE_UINT && ctor != TYPE_UINT_0 && ctor != TYPE_UINT_SMALL {
		err = errors.New("Malformed error. Expecting uint field")
		return
	}

	var fieldValue uint32

	switch {
	case ctor == TYPE_UINT_0:
		fieldLength = 1
		fieldValue = 0
		break

	case ctor == TYPE_UINT:
		fieldLength = 5
		fieldValue = binary.BigEndian.Uint32(v[1:5])
		break

	case ctor == TYPE_UINT_SMALL:
		fieldLength = 2
		fieldValue = binary.BigEndian.Uint32(v[1:2])
		break
	}

	val = &UInt{
		value: fieldValue,
		BaseAMQPType: BaseAMQPType{
			encoding: ctor,
		},
	}
	return
}

type Type byte

const (
	TYPE_CONSTRUCTOR = Type(0x00)

	TYPE_PERFORMATIVE_OPEN = Type(0x10)

	TYPE_NULL        = Type(0x40)

	TYPE_BOOLEAN       = Type(0x56)
	TYPE_BOOLEAN_TRUE  = Type(0x41)
	TYPE_BOOLEAN_FALSE = Type(0x42)

	TYPE_UBYTE = Type(0x50)

	TYPE_USHORT = Type(0x60)

	TYPE_UINT       = Type(0x70)
	TYPE_UINT_SMALL = Type(0x52)
	TYPE_UINT_0     = Type(0x43)

	TYPE_ULONG = Type(0x80)

	TYPE_ULONG_SMALL = Type(0x53)
	TYPE_ULONG_0     = Type(0x44)

	TYPE_BYTE = Type(0x51)

	TYPE_SHORT = Type(0x61)

	TYPE_INT       = Type(0x71)
	TYPE_INT_SMALL = Type(0x54)

	TYPE_LONG       = Type(0x81)
	TYPE_LONG_SMALL = Type(0x55)

	TYPE_FLOAT = Type(0x72)

	TYPE_DOUBLE = Type(0x82)

	TYPE_DECIMAL32  = Type(0x74)
	TYPE_DECIMAl64  = Type(0x84)
	TYPE_DECIMAL128 = Type(0x94)
	TYPE_CHAR       = Type(0x73)
	TYPE_TIMESTAMP  = Type(0x83)
	TYPE_UUID       = Type(0x98)

	TYPE_BINARY_VBIN8  = Type(0xa0)
	TYPE_BINARY_VBIN32 = Type(0xb0)

	TYPE_STRING_8_UTF8  = Type(0xa1)
	TYPE_STRING_32_UTF8 = Type(0xb1)

	TYPE_SYMBOL_8  = Type(0xa3)
	TYPE_SYMBOL_32 = Type(0xb3)

	TYPE_LIST_0  = Type(0x45)
	TYPE_LIST_8  = Type(0xc0)
	TYPE_LIST_32 = Type(0xd0)

	TYPE_MAP_8  = Type(0xc1)
	TYPE_MAP_32 = Type(0xd1)

	TYPE_ARRAY_8  = Type(0xe0)
	TYPE_ARRAY_32 = Type(0xf0)
)

type TypeEncoding byte

const (
	TYPEENCODING_NONE       = TypeEncoding(0)
	TYPEENCODING_TRUE       = TypeEncoding(1)
	TYPEENCODING_FALSE      = TypeEncoding(2)
	TYPEENCODING_SMALLUINT  = TypeEncoding(3)
	TYPEENCODING_UINT0      = TypeEncoding(4)
	TYPEENCODING_SMALLULONG = TypeEncoding(5)
	TYPEENCODING_ULONG0     = TypeEncoding(6)
	TYPEENCODING_SMALLINT   = TypeEncoding(7)
	TYPEENCODING_SMALLLONG  = TypeEncoding(8)
	TYPEENCODING_IEEE_754   = TypeEncoding(9)
	TYPEENCODING_UTF32      = TypeEncoding(10)
	TYPEENCODING_MS64       = TypeEncoding(11)
	TYPEENCODING_VBIN8      = TypeEncoding(12)
	TYPEENCODING_VBIN32     = TypeEncoding(13)
	TYPEENCODING_STR8_UTF8  = TypeEncoding(14)
	TYPEENCODING_STR32_UTF8 = TypeEncoding(15)
	TYPEENCODING_SYM8       = TypeEncoding(16)
	TYPEENCODING_SYM32      = TypeEncoding(17)
	TYPEENCODING_LIST0      = TypeEncoding(18)
	TYPEENCODING_LIST8      = TypeEncoding(19)
	TYPEENCODING_LIST32     = TypeEncoding(20)
	TYPEENCODING_MAP8       = TypeEncoding(21)
	TYPEENCODING_MAP32      = TypeEncoding(22)
	TYPEENCODING_ARRAY8     = TypeEncoding(23)
	TYPEENCODING_ARRAY32    = TypeEncoding(24)
)

type AMQPType interface {
	GetEncoding() TypeEncoding
}

type BaseAMQPType struct {
	encoding Type
}

func (b *BaseAMQPType) GetEncoding() TypeEncoding {
	return TYPEENCODING_NONE
}

type Null struct {
	BaseAMQPType
}

type Boolean struct {
	BaseAMQPType
	value bool
}

type UByte struct {
	BaseAMQPType
}

type UInt struct {
	BaseAMQPType
	value uint32
}

type ULong struct {
	BaseAMQPType
	value uint64
}

type Byte struct {
	BaseAMQPType
	value byte
}

type Int struct {
	BaseAMQPType
	value int32
}

type Long struct {
	BaseAMQPType
	value int64
}

type Float struct {
	BaseAMQPType
	value float32
}

type Double struct {
	BaseAMQPType
	value float64
}

type Decimal32 struct {
	BaseAMQPType
	value float32
}

type Decimal64 struct {
	BaseAMQPType
	value float64
}

type Decimal128 struct {
	BaseAMQPType
}

type Char struct {
	BaseAMQPType
	value string
}

type Timestamp struct {
	BaseAMQPType
	value uint64
}

func GetTimestampField(v []byte) (val *Timestamp, fieldLength int, err error) {
	ctor := Type(v[0])
	if ctor != TYPE_TIMESTAMP {
		err = errors.New("Malformed error. Expecting timestamp field")
		return
	}
	return
}

type UUID struct {
	BaseAMQPType
	value string
}

type Binary struct {
	BaseAMQPType
	value []byte
}

func GetStringField(v []byte) (val *String, fieldLength uint, err error) {
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

	rawVal, fieldLength, err := GetField(v)

	val = &String{
		value: string(rawVal),
		BaseAMQPType: BaseAMQPType{
			encoding: ctor,
		},
	}
	val = NewString(string(rawVal))

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

type Symbol struct {
	BaseAMQPType
	value string
}

func GetSymbolField(v []byte) (val *Symbol, fieldLength int, err error) {
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

func GetSymbolArrayField(v []byte) (val []*Symbol, fieldLength uint, err error) {
	ctor := Type(v[0])

	if ctor == TYPE_NULL {
		val = []*Symbol{}
		fieldLength = 1
		return
	}
	return
}

type List struct {
	BaseAMQPType
}

type Map struct {
	BaseAMQPType
}

/*

 */
func GetMapField(v []byte) (val *Map, fieldLength uint, err error) {
	ctor := Type(v[0])
	if ctor != TYPE_MAP_8 && ctor != TYPE_MAP_32 {
		err = errors.New("Malformed error. Expecting map field")
		return
	}

	var fieldCount uint
	if ctor == TYPE_MAP_8 {
		fieldLength = uint(v[1]) + 2
		fieldCount = uint(v[2])
	} else if ctor == TYPE_MAP_32 {
		fieldLength = uint(binary.BigEndian.Uint32(v[1:4])) + 5
		fieldCount = uint(v[5])
	}

	log.Println("Fields", fieldCount)





	/*
	// properties
	c1      map8
	51      81 bytes,

	08      // 8 values

	a3 07   //07
	70 72 6f 64 75 63 74    // product
	a1 04
	71 70 69 64             // qpid

	a3 07   // 07
	76 65 72 73 69 6f 6e
	a1 05
	36 2e 31 2e 32

	a3 0a   // 10
	71 70 69 64 2e 62 75 69 6c 64
	a1 07
	31 37 38 37 30 36 37

	a3 12   // 18
	71 70 69 64 2e 69 6e 73 74 61 6e 63 65 5f 6e 61 6d 65
	a1 06
	42 72 6f 6b 65 72

	*/

	return
}

type Array struct {
	BaseAMQPType
}

type Short struct {
	BaseAMQPType
	value int16
}

type UShort struct {
	BaseAMQPType
	value uint16
}

func GetUShortField(v []byte) (val *UShort, fieldLength uint, err error) {
	ctor := Type(v[0])
	if ctor != TYPE_USHORT {
		err = errors.New("Malformed error. Expecting ushort field")
		return
	}

	fieldLength = 3
	val = &UShort{
		value: binary.BigEndian.Uint16(v[1:3]),
	}

	return
	// 0x60
}
