package primitives

type ULongType struct {
}

func (t ULongType) Validate() error {
	return nil
}

func (t ULongType) Encode() []byte {
	return []byte{}
}

func (t ULongType) GetTypeFormatCode() byte {
	return 0
}

type SmallULongType struct {
}

func (t SmallULongType) Validate() error {
	return nil
}

func (t SmallULongType) Encode() []byte {
	return []byte{}
}

func (t SmallULongType) GetTypeFormatCode() byte {
	return 0
}

type ULongZeroType struct {
}

func (t ULongZeroType) Validate() error {
	return nil
}

func (t ULongZeroType) Encode() []byte {
	return []byte{}
}

func (t ULongZeroType) GetTypeFormatCode() byte {
	return 0
}

type LongType struct {
}

func (t LongType) Validate() error {
	return nil
}

func (t LongType) Encode() []byte {
	return []byte{}
}

func (t LongType) GetTypeFormatCode() byte {
	return 0
}

type SmallLongType struct {
}

func (t SmallLongType) Validate() error {
	return nil
}

func (t SmallLongType) Encode() []byte {
	return []byte{}
}

func (t SmallLongType) GetTypeFormatCode() byte {
	return 0
}
