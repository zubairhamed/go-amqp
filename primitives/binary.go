package primitives

type Binary8Type struct {
}

func (t Binary8Type) Validate() error {
	return nil
}

func (t Binary8Type) Encode() []byte {
	return []byte{}
}

func (t Binary8Type) GetTypeFormatCode() byte {
	return 0
}

type Binary32Type struct {
}

func (t Binary32Type) Validate() error {
	return nil
}

func (t Binary32Type) Encode() []byte {
	return []byte{}
}

func (t Binary32Type) GetTypeFormatCode() byte {
	return 0
}
