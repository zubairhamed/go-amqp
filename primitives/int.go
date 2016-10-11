package primitives

type UIntType struct {
}

func (t UIntType) Validate() error {
	return nil
}

func (t UIntType) Encode() []byte {
	return []byte{}
}

func (t UIntType) GetTypeFormatCode() byte {
	return 0
}

type SmallUIntType struct {
}

func (t SmallUIntType) Validate() error {
	return nil
}

func (t SmallUIntType) Encode() []byte {
	return []byte{}
}

func (t SmallUIntType) GetTypeFormatCode() byte {
	return 0
}

type UIntZeroType struct {
}

func (t UIntZeroType) Validate() error {
	return nil
}

func (t UIntZeroType) Encode() []byte {
	return []byte{}
}

func (t UIntZeroType) GetTypeFormatCode() byte {
	return 0
}

type IntType struct {
}

func (t IntType) Validate() error {
	return nil
}

func (t IntType) Encode() []byte {
	return []byte{}
}

func (t IntType) GetTypeFormatCode() byte {
	return 0
}

type SmallIntType struct {
}

func (t SmallIntType) Validate() error {
	return nil
}

func (t SmallIntType) Encode() []byte {
	return []byte{}
}

func (t SmallIntType) GetTypeFormatCode() byte {
	return 0
}
