package types

func NewDouble(v float64) *Double {
	return &Double{
		value: v,
	}
}

type Double struct {
	BaseAMQPType
	value float64
}
