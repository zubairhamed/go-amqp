package primitives

type StringType struct {
}

func (t *StringType) Validate() error {
	return nil
}

func (t *StringType) Encode() []byte {
	return []byte{}
}

type String8Type struct {
	StringType
}

func (t *String8Type) Validate() error {
	return nil
}

func (t *String8Type) Encode() []byte {
	return []byte{}
}

type String32Type struct {
	StringType
}

func (t *String32Type) Validate() error {
	return nil
}

func (t *String32Type) Encode() []byte {
	return []byte{}
}
