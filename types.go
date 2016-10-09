package amqp

type TypeFormatCode byte

const (
	TypeDescriptor = TypeFormatCode(0x00)

	TypeFixedNull       = TypeFormatCode(0x40)
	TypeFixedBool       = TypeFormatCode(0x56)
	TypeFixedTrue       = TypeFormatCode(0x41)
	TypeFixedFalse      = TypeFormatCode(0x42)
	TypeFixedUByte      = TypeFormatCode(0x50)
	TypeFixedUShort     = TypeFormatCode(0x60)
	TypeFixedUInt       = TypeFormatCode(0x70)
	TypeFixedSmallUInt  = TypeFormatCode(0x52)
	TypeFixedUInt0      = TypeFormatCode(0x43)
	TypeFixedULong      = TypeFormatCode(0x80)
	TypeFixedSmallULong = TypeFormatCode(0x53)
	TypeFixedULong0     = TypeFormatCode(0x44)
	TypeFixedByte       = TypeFormatCode(0x51)
	TypeFixedShort      = TypeFormatCode(0x61)
	TypeFixedInt        = TypeFormatCode(0x71)
	TypeFixedSmallInt   = TypeFormatCode(0x54)
	TypeFixedLong       = TypeFormatCode(0x81)
	TypeFixedSmallLong  = TypeFormatCode(0x55)
	TypeFixedFloat      = TypeFormatCode(0x72)
	TypeFixedDouble     = TypeFormatCode(0x82)
	TypeFixedDecimal32  = TypeFormatCode(0x74)
	TypeFixedDecimal64  = TypeFormatCode(0x84)
	TypeFixedDecimal128 = TypeFormatCode(0x94)
	TypeFixedChar       = TypeFormatCode(0x73)
	TypeFixedTimestamp  = TypeFormatCode(0x83)
	TypeFixedUuid       = TypeFormatCode(0x98)
	TypeFixedList0      = TypeFormatCode(0x45)

	TypeVariableBinaryVBin8  = TypeFormatCode(0xA0)
	TypeVariableBinaryVBin32 = TypeFormatCode(0xB0)
	TypeVariableStringStr8   = TypeFormatCode(0xA1)
	TypeVariableStringStr32  = TypeFormatCode(0xB1)
	TypeVariableSymbol8      = TypeFormatCode(0xA3)
	TypeVariableSymbol32     = TypeFormatCode(0xB3)

	TypeCompoundList8  = TypeFormatCode(0xC0)
	TypeCompoundList32 = TypeFormatCode(0xD0)
	TypeCompoundMap8   = TypeFormatCode(0xC1)
	TypeCompoundMap32  = TypeFormatCode(0xD1)

	TypeArray8  = TypeFormatCode(0xE0)
	TypeArray32 = TypeFormatCode(0xF0)

	TypePerformativeOpen        = TypeFormatCode(0x10)
	TypePerformativeBegin       = TypeFormatCode(0x11)
	TypePerformativeAttach      = TypeFormatCode(0x12)
	TypePerformativeFlow        = TypeFormatCode(0x13)
	TypePerformativeTransfer    = TypeFormatCode(0x14)
	TypePerformativeDisposition = TypeFormatCode(0x15)
	TypePerformativeDetach      = TypeFormatCode(0x16)
	TypePerformativeEnd         = TypeFormatCode(0x17)
	TypePerformativeClose       = TypeFormatCode(0x18)
)

type Data interface {
	Encode() []byte
	GetTypeFormatCode() TypeFormatCode
}

type BaseData struct {
}
