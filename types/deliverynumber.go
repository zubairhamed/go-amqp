package types

func NewDeliveryNumber(v uint32) *DeliveryNumber {
	return &DeliveryNumber{
		SequenceNumber: NewSequenceNumber(v),
	}
}

type DeliveryNumber struct {
	*SequenceNumber
}
