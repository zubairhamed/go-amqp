package primitives

type UShortType struct {
}

func (t UShortType) Validate() error {
	return nil
}

func (t UShortType) Encode() []byte {
	return []byte{}
}

type ShortType struct {
}

func (t ShortType) Validate() error {
	return nil
}

func (t ShortType) Encode() []byte {
	return []byte{}
}
