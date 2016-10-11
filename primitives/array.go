package primitives

type ArrayType struct {
}

type Array8Type struct {
	ArrayType
}

func (t Array8Type) Validate() error {
	return nil
}

func (t Array8Type) Encode() []byte {
	return []byte{}
}

func (t Array8Type) GetTypeFormatCode() byte {
	return byte{}
}

type Array32Type struct {
	ArrayType
}

func (t Array32Type) Validate() error {
	return nil
}

func (t Array32Type) Encode() []byte {
	return []byte{}
}

func (t Array32Type) GetTypeFormatCode() byte {
	return byte{}
}
