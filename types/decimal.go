package types

func NewDecimal32(value float32) *Decimal32 {
	return &Decimal32{
		value: value,
	}
}

// 32-bit decimal number (IEEE 754-2008 decimal32)
type Decimal32 struct {
	BaseAMQPType
	value float32
}

func NewDecimal64(value float64) *Decimal64 {
	return &Decimal64{
		value: value,
	}
}

// 64-bit decimal number (IEEE 754-2008 decimal64)
type Decimal64 struct {
	BaseAMQPType
	value float64
}

func NewDecimal128() *Decimal128 {
	return &Decimal128{}
}

// 128-bit decimal number (IEEE 754-2008 decimal128)
type Decimal128 struct {
	BaseAMQPType
}
