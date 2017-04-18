package types

func NewFloat(v float32) *Float {
	return &Float{
		value: v,
	}
}

type Float struct {
	BaseAMQPType
	value float32
}
