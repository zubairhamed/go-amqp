package types

func NewIetfLanguageTag(v string) *IetfLanguageTag {
	return &IetfLanguageTag {
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

func EncodeIetfLanguageTagArrayField(v []*IetfLanguageTag) (enc []byte, l uint, err error) {
	if len(v) == 0 {
		return []byte{byte(TYPE_NULL)}, 1, nil
	}

	return nil, 0, nil
}


func DecodeIetfLanguageTagArrayField(v []byte) (val []*IetfLanguageTag, fieldLength uint, err error) {
	iVal, fieldLength, err := DecodeSymbolArrayField(v)

	val = []*IetfLanguageTag{}
	for _, v2 := range iVal {
		val2 := &IetfLanguageTag{
			Symbol: v2,
		}

		val = append(val, val2)
	}
	return
}

