package primitives

type StringType struct {
}

func (t StringType) Validate() error {
	return nil
}

func (t StringType) Encode() []byte {
	return []byte{}
}

func (t StringType) GetTypeFormatCode() byte {
	return 0
}

type String8Type struct {
	StringType
}

func (t String8Type) Validate() error {
	return nil
}

func (t String8Type) Encode() []byte {
	return []byte{}
}

func (t String8Type) GetTypeFormatCode() byte {
	return 0
}

type String32Type struct {
	StringType
}

func (t String32Type) Validate() error {
	return nil
}

func (t String32Type) Encode() []byte {
	return []byte{}
}

func (t String32Type) GetTypeFormatCode() byte {
	return 0
}
