package primitives

type TimestampType struct {
}

func (t TimestampType) Validate() error {
	return nil
}

func (t TimestampType) Encode() []byte {
	return []byte{}
}
