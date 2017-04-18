package types

func NewDeliveryTag(b []byte) *DeliveryTag {
	return &DeliveryTag{
		Binary: NewBinary(b),
	}
}

/*
<type name="delivery-tag" class="restricted" source="binary"/>
*/
type DeliveryTag struct {
	*Binary
}
