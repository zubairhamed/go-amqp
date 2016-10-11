package primitives

type DoubleType struct {
}

func (t DoubleType) Validate() error {
	return nil
}

func (t DoubleType) Encode() []byte {
	return []byte{}
}

func (t DoubleType) GetTypeFormatCode() byte {
	return 0
}
