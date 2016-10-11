package primitives

type BooleanType struct {
}

func (t BooleanType) Validate() error {
	return nil
}

func (t BooleanType) Encode() []byte {
	return []byte{}
}

func (t Array8Type) GetTypeFormatCode() byte {
	return byte{}
}

type TrueType struct {
}

func (t TrueType) Validate() error {
	return nil
}

func (t TrueType) Encode() []byte {
	return []byte{}
}

func (t Array8Type) GetTypeFormatCode() byte {
	return byte{}
}

type FalseType struct {
}

func (t FalseType) Validate() error {
	return nil
}

func (t FalseType) Encode() []byte {
	return []byte{}
}

func (t Array8Type) GetTypeFormatCode() byte {
	return byte{}
}
