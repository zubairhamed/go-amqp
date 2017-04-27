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

func (s *DeliveryTag) Encode() ([]byte, uint, error) {
	if s == nil {
		return NullValue()
	}
	return EncodeDeliveryTagField(s)
}

func (b *DeliveryTag) Stringify() string {
	return string(b.value)
}

func EncodeDeliveryTagField(s *DeliveryTag) ([]byte, uint, error) {
	return EncodeBinaryField(s.Binary)
}

func DecodeDeliveryTagField(v []byte) (val *DeliveryTag, fieldLength uint, err error) {
	return
}
