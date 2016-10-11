package primitives

type UUIDType struct {
}

func (t UUIDType) Validate() error {
	return nil
}

func (t UUIDType) Encode() []byte {
	return []byte{}
}

func (t UUIDType) GetTypeFormatCode() byte {
	return 0
}
