package primitives

type ListType struct {
}

type ListEmptyType struct {
	ListType
}

func (t ListEmptyType) Validate() error {
	return nil
}

func (t ListEmptyType) Encode() []byte {
	return []byte{}
}

func (t Array8Type) GetTypeFormatCode() byte {
	return byte{}
}

type List8Type struct {
	ListType
}

func (t List8Type) Validate() error {
	return nil
}

func (t List8Type) Encode() []byte {
	return []byte{}
}

func (t Array8Type) GetTypeFormatCode() byte {
	return byte{}
}

type List32Type struct {
	ListType
}

func (t List32Type) Validate() error {
	return nil
}

func (t List32Type) Encode() []byte {
	return []byte{}
}

func (t Array8Type) GetTypeFormatCode() byte {
	return byte{}
}
