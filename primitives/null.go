package primitives

type NullType struct {
}

func (t NullType) Validate() error {
	return nil
}

func (t NullType) Encode() []byte {
	return []byte{}
}
