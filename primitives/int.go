package primitives

type UIntType struct {
}

func (t UIntType) Validate() error {
	return nil
}

func (t UIntType) Encode() []byte {
	return []byte{}
}

func (t Array8Type) GetTypeFormatCode() byte {
	return byte{}
}


type SmallUIntType struct {
}

func (t SmallUIntType) Validate() error {
	return nil
}

func (t SmallUIntType) Encode() []byte {
	return []byte{}
}

func (t Array8Type) GetTypeFormatCode() byte {
	return byte{}
}


type UIntZeroType struct {
}

func (t UIntZeroType) Validate() error {
	return nil
}

func (t UIntZeroType) Encode() []byte {
	return []byte{}
}

func (t Array8Type) GetTypeFormatCode() byte {
	return byte{}
}

type IntType struct {
}

func (t IntType) Validate() error {
	return nil
}

func (t IntType) Encode() []byte {
	return []byte{}
}

func (t Array8Type) GetTypeFormatCode() byte {
	return byte{}
}

type SmallIntType struct {
}

func (t SmallIntType) Validate() error {
	return nil
}

func (t SmallIntType) Encode() []byte {
	return []byte{}
}

func (t Array8Type) GetTypeFormatCode() byte {
	return byte{}
}

