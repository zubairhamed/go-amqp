package types

func NewIetfLanguageTag(v string) *IetfLanguageTag {
	return &IetfLanguageTag{
		Symbol: NewSymbol(v),
	}
}

type IetfLanguageTag struct {
	*Symbol
}

func DecodeIetfLanguageTagField(v []byte) (val *IetfLanguageTag, fieldLength uint, err error) {
	iVal, fieldLength, err := DecodeSymbolField(v)

	val = &IetfLanguageTag{
		Symbol: iVal,
	}

	return
}

// Array Type
type IetfLanguageTagArray struct {
	*SymbolArray
}

func (s *IetfLanguageTagArray) Encode() ([]byte, uint, error) {
	if s == nil {
		return []byte{byte(TYPE_NULL)}, 1, nil
	}
	return EncodeIetfLanguageTagArrayField(s)
}

func (b *IetfLanguageTagArray) Stringify() string {
	return "String: IetfLanguageTagArray"
}

func NewIetfLanguageTagArray() *IetfLanguageTagArray {
	return &IetfLanguageTagArray{
		&SymbolArray{
			value: []*Symbol{},
		},
	}
}

func EncodeIetfLanguageTagArrayField(v *IetfLanguageTagArray) (enc []byte, l uint, err error) {
	if v.Length() == 0 {
		return []byte{byte(TYPE_NULL)}, 1, nil
	}

	return nil, 0, nil
}

func DecodeIetfLanguageTagArrayField(v []byte) (val *IetfLanguageTagArray, fieldLength uint, err error) {
	iVal, fieldLength, err := DecodeSymbolArrayField(v)

	val = NewIetfLanguageTagArray()

	for _, v2 := range iVal.value {
		val.Append(v2)
	}
	return
}
