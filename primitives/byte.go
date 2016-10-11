package primitives

type UByteType struct {
}

func (t UByteType) Validate() error {
	return nil
}

func (t UByteType) Encode() []byte {
	return []byte{}
}

func (t Array8Type) GetTypeFormatCode() byte {
	return byte{}
}

type ByteType struct {
}

func (t ByteType) Validate() error {
	return nil
}

func (t ByteType) Encode() []byte {
	return []byte{}
}

func (t Array8Type) GetTypeFormatCode() byte {
	return byte{}
}
