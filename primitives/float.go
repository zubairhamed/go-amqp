package primitives

type FloatType struct {
}

func (t FloatType) Validate() error {
	return nil
}

func (t FloatType) Encode() []byte {
	return []byte{}
}

func (t Array8Type) GetTypeFormatCode() byte {
	return byte{}
}

