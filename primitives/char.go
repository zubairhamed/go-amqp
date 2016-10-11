package primitives

type CharType struct {
}

func (t CharType) Validate() error {
	return nil
}

func (t CharType) Encode() []byte {
	return []byte{}
}

func (t Array8Type) GetTypeFormatCode() byte {
	return byte{}
}
