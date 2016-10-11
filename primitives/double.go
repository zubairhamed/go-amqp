package primitives

type DoubleType struct {
}

func (t DoubleType) Validate() error {
	return nil
}

func (t DoubleType) Encode() []byte {
	return []byte{}
}

func (t Array8Type) GetTypeFormatCode() byte {
	return byte{}
}
