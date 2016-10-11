package primitives

type NullType struct {
}

func (t NullType) Validate() error {
	return nil
}

func (t NullType) Encode() []byte {
	return []byte{}
}

func (t NullType) GetTypeFormatCode() byte {
	return 0
}
