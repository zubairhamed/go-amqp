package types

type Decimal32 struct {
	BaseAMQPType
	value float32
}

type Decimal64 struct {
	BaseAMQPType
	value float64
}

type Decimal128 struct {
	BaseAMQPType
}
