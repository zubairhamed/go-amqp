package types

func NewAmqpValue(v AMQPType) *AMQPValue {
	return &AMQPValue{
		value: v,
	}
}

type AMQPValue struct {
	value AMQPType
}

func (b *AMQPValue) GetType() Type {
	return TYPE_AMQP_VALUE
}

func (v *AMQPValue) Stringify() string {
	return ""
}

func (v *AMQPValue) Encode() ([]byte, uint, error) {
	if v == nil {
		return []byte{byte(TYPE_NULL)}, 1, nil
	}
	return EncodeAmqpValueField(v)
}

func EncodeAmqpValueField(s *AMQPValue) (b []byte, l uint, err error) {
	v := s.value
	b = []byte{0x00, 0x53, 0x77}

	vb, _, ve := v.Encode()
	if ve != nil {
		err = ve
		return
	}

	b = append(b, vb...)
	l = uint(len(b))

	return
}
