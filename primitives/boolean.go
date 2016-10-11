package primitives

type BooleanType struct {
}

func (t BooleanType) Validate() error {
	return nil
}

func (t BooleanType) Encode() []byte {
	return []byte{}
}

func (t BooleanType) GetTypeFormatCode() byte {
	return 0
}

type TrueType struct {
}

func (t TrueType) Validate() error {
	return nil
}

func (t TrueType) Encode() []byte {
	return []byte{}
}

func (t TrueType) GetTypeFormatCode() byte {
	return 0
}

type FalseType struct {
}

func (t FalseType) Validate() error {
	return nil
}

func (t FalseType) Encode() []byte {
	return []byte{}
}

func (t FalseType) GetTypeFormatCode() byte {
	return 0
}
