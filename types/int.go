package types

func NewInt(v int32) *Int {
	return &Int {
		value: v,
	}
}

type Int struct {
	BaseAMQPType
	value int32
}
