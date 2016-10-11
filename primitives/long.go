package primitives

type ULongType struct {
}

func (t ULongType) Validate() error {
	return nil
}

func (t ULongType) Encode() []byte {
	return []byte{}
}

func (t Array8Type) GetTypeFormatCode() byte {
	return byte{}
}

type SmallULongType struct {
}

func (t SmallULongType) Validate() error {
	return nil
}

func (t SmallULongType) Encode() []byte {
	return []byte{}
}

func (t Array8Type) GetTypeFormatCode() byte {
	return byte{}
}

type ULongZeroType struct {
}

func (t ULongZeroType) Validate() error {
	return nil
}

func (t ULongZeroType) Encode() []byte {
	return []byte{}
}

func (t Array8Type) GetTypeFormatCode() byte {
	return byte{}
}

type LongType struct {
}

func (t LongType) Validate() error {
	return nil
}

func (t LongType) Encode() []byte {
	return []byte{}
}

func (t Array8Type) GetTypeFormatCode() byte {
	return byte{}
}

type SmallLongType struct {
}

func (t SmallLongType) Validate() error {
	return nil
}

func (t SmallLongType) Encode() []byte {
	return []byte{}
}

func (t Array8Type) GetTypeFormatCode() byte {
	return byte{}
}