package primitives

type MapType struct {
}

func (t MapType) Validate() error {
	return nil
}

func (t MapType) Encode() []byte {
	return []byte{}
}

type Map8Type struct {
}

func (t Map8Type) Validate() error {
	return nil
}

func (t Map8Type) Encode() []byte {
	return []byte{}
}

func (t Map8Type) GetTypeFormatCode() byte {
	return 0
}

type Map32Type struct {
}

func (t Map32Type) Validate() error {
	return nil
}

func (t Map32Type) Encode() []byte {
	return []byte{}
}

func (t Map32Type) GetTypeFormatCode() byte {
	return 0
}

type MapKeyValue struct {
}

func (t MapKeyValue) Validate() error {
	return nil
}

func (t MapKeyValue) Encode() []byte {
	return []byte{}
}

func (t MapKeyValue) GetTypeFormatCode() byte {
	return 0
}
