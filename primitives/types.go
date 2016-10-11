package primitives

type TypeFormatCode byte

const (
	TypeDescriptor = TypeFormatCode(0x00)

	TypeNull       = TypeFormatCode(0x40)
	TypeBool       = TypeFormatCode(0x56)
	TypeTrue       = TypeFormatCode(0x41)
	TypeFalse      = TypeFormatCode(0x42)
	TypeUByte      = TypeFormatCode(0x50)
	TypeUShort     = TypeFormatCode(0x60)
	TypeUInt       = TypeFormatCode(0x70)
	TypeSmallUInt  = TypeFormatCode(0x52)
	TypeUInt0      = TypeFormatCode(0x43)
	TypeULong      = TypeFormatCode(0x80)
	TypeSmallULong = TypeFormatCode(0x53)
	TypeULong0     = TypeFormatCode(0x44)
	TypeByte       = TypeFormatCode(0x51)
	TypeShort      = TypeFormatCode(0x61)
	TypeInt        = TypeFormatCode(0x71)
	TypeSmallInt   = TypeFormatCode(0x54)
	TypeLong       = TypeFormatCode(0x81)
	TypeSmallLong  = TypeFormatCode(0x55)
	TypeFloat      = TypeFormatCode(0x72)
	TypeDouble     = TypeFormatCode(0x82)
	TypeDecimal32  = TypeFormatCode(0x74)
	TypeDecimal64  = TypeFormatCode(0x84)
	TypeDecimal128 = TypeFormatCode(0x94)
	TypeChar       = TypeFormatCode(0x73)
	TypeTimestamp  = TypeFormatCode(0x83)
	TypeUuid       = TypeFormatCode(0x98)
	TypeList0      = TypeFormatCode(0x45)

	TypeBinary8  = TypeFormatCode(0xA0)
	TypeBinary32 = TypeFormatCode(0xB0)
	TypeString8  = TypeFormatCode(0xA1)
	TypeString32 = TypeFormatCode(0xB1)
	TypeSymbol8  = TypeFormatCode(0xA3)
	TypeSymbol32 = TypeFormatCode(0xB3)

	TypeList8  = TypeFormatCode(0xC0)
	TypeList32 = TypeFormatCode(0xD0)
	TypeMap8   = TypeFormatCode(0xC1)
	TypeMap32  = TypeFormatCode(0xD1)

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
