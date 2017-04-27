package types

func NewReceiverSettleMode(v byte) *ReceiverSettleMode {
	r := &ReceiverSettleMode{}
	if v == 0 {
		r.SetFirst()
	} else if v == 1 {
		r.SetSecond()
	}
	return r
}

/*
<type name="receiver-settle-mode" class="restricted" source="ubyte">
    <choice name="first" value="0"/>
    <choice name="second" value="1"/>
</type>
*/
type ReceiverSettleMode struct {
	*UByte
}

func (r *ReceiverSettleMode) SetFirst() {
	r.UByte.value = 0
}

func (r *ReceiverSettleMode) SetSecond() {
	r.UByte.value = 1
}

func (s *ReceiverSettleMode) Encode() ([]byte, uint, error) {
	if s == nil {
		return NullValue()
	}
	return EncodeReceiverSettleModeField(s)
}

func (b *ReceiverSettleMode) Stringify() string {
	if b.UByte.value == 0 {
		return "first"
	} else if b.UByte.value == 1 {
		return "second"
	}

	return "?"
}

func DecodeReceiverSettleModeField(v []byte) (val *ReceiverSettleMode, fieldLength uint, err error) {
	ctor := Type(v[0])
	if ctor == TYPE_NULL {
		fieldLength = 1
		return
	}

	switch {
	case ctor == TYPE_NULL:
		fieldLength = 1
		return

	case ctor == TYPE_BOOLEAN_TRUE:
		val = NewReceiverSettleMode(0)
		fieldLength = 1
		return

	case ctor == TYPE_BOOLEAN_FALSE:
		val = NewReceiverSettleMode(1)
		fieldLength = 1
		return
	}

	return
}

func EncodeReceiverSettleModeField(s *ReceiverSettleMode) ([]byte, uint, error) {
	return EncodeUByteField(s.UByte)
}
