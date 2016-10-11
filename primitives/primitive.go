package primitives

type Primitive interface {
	Validate() error
	Encode() []byte
	GetTypeFormatCode() byte
}

type BasePrimitive struct {
}

func NewPrimitive(t TypeFormatCode) (p Primitive) {
	switch t {
	case TypeArray8:
		p = Array8Type{}

	case TypeArray32:
		p = Array32Type{}

	case TypeBinary8:
		p = Binary8Type{}

	case TypeBinary32:
		p = Binary32Type{}

	case TypeBool:
		p = BooleanType{}

	case TypeByte:
		p = ByteType{}

	case TypeChar:
		p = CharType{}

	case TypeDecimal32:
		p = Decimal32Type{}

	case TypeDecimal64:
		p = Decimal64Type{}

	case TypeDecimal128:
		p = Decimal128Type{}

	case TypeDouble:
		p = DoubleType{}

	case TypeFalse:
		p = FalseType{}

	case TypeFloat:
		p = FloatType{}

	case TypeInt:
		p = IntType{}

	case TypeList0:
		p = ListEmptyType{}

	case TypeList8:
		p = List8Type{}

	case TypeList32:
		p = List32Type{}

	case TypeLong:
		p = LongType{}

	case TypeMap8:
		p = Map8Type{}

	case TypeMap32:
		p = Map32Type{}

	case TypeNull:
		p = NullType{}

	case TypeShort:
		p = ShortType{}

	case TypeSmallInt:
		p = SmallIntType{}

	case TypeSmallLong:
		p = SmallLongType{}

	case TypeSmallUInt:
		p = SmallUIntType{}

	case TypeSmallULong:
		p = SmallULongType{}

	case TypeString8:
		p = String8Type{}

	case TypeString32:
		p = String32Type{}

	case TypeSymbol8:
		p = Symbol8Type{}

	case TypeSymbol32:
		p = Symbol32Type{}

	case TypeTimestamp:
		p = TimestampType{}

	case TypeTrue:
		p = TrueType{}

	case TypeUByte:
		p = UByteType{}

	case TypeUInt:
		p = UIntType{}

	case TypeUInt0:
		p = UIntZeroType{}

	case TypeULong:
		p = ULongType{}

	case TypeULong0:
		p = ULongZeroType{}

	case TypeUShort:
		p = UShortType{}

	case TypeUuid:
		p = UUIDType{}

	default:
		break
	}
	return
}
