package primitives

type SymbolType struct {
}

func (t SymbolType) Validate() error {
	return nil
}

func (t SymbolType) Encode() []byte {
	return []byte{}
}

type Symbol8Type struct {
}

func (t Symbol8Type) Validate() error {
	return nil
}

func (t Symbol8Type) Encode() []byte {
	return []byte{}
}

type Symbol32Type struct {
}

func (t Symbol32Type) Validate() error {
	return nil
}

func (t Symbol32Type) Encode() []byte {
	return []byte{}
}