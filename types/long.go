package types

func NewLong(v int64) *Long {
	return &Long{
		value: v,
	}
}

type Long struct {
	BaseAMQPType
	value int64
}
