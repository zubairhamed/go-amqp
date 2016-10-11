package primitives

type Decimal32Type struct {
}

func (t Decimal32Type) Validate() error {
	return nil
}

func (t Decimal32Type) Encode() []byte {
	return []byte{}
}

func (t Decimal32Type) GetTypeFormatCode() byte {
	return 0
}

type Decimal64Type struct {
}

func (t Decimal64Type) Validate() error {
	return nil
}

func (t Decimal64Type) Encode() []byte {
	return []byte{}
}

func (t Decimal64Type) GetTypeFormatCode() byte {
	return 0
}

type Decimal128Type struct {
}

func (t Decimal128Type) Validate() error {
	return nil
}

func (t Decimal128Type) Encode() []byte {
	return []byte{}
}

func (t Decimal128Type) GetTypeFormatCode() byte {
	return 0
}
