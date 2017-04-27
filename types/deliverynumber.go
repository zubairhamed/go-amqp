package types

func NewDeliveryNumber(v uint32) *DeliveryNumber {
	return &DeliveryNumber{
		SequenceNumber: NewSequenceNumber(v),
	}
}

type DeliveryNumber struct {
	*SequenceNumber
}

func DecodeDeliveryNumberField(v []byte) (val *DeliveryNumber, fieldLength uint, err error) {
	return
}
