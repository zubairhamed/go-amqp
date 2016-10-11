package primitives

type UByteType struct {
}

func (t UByteType) Validate() error {
	return nil
}

func (t UByteType) Encode() []byte {
	return []byte{}
}

func (t UByteType) GetTypeFormatCode() byte {
	return 0
}

type ByteType struct {
}

func (t ByteType) Validate() error {
	return nil
}

func (t ByteType) Encode() []byte {
	return []byte{}
}

func (t ByteType) GetTypeFormatCode() byte {
	return 0
}
