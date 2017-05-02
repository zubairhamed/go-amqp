package types

func NewSenderSettleMode(v byte) *SenderSettleMode {
	r := &SenderSettleMode{}
	if v == 0 {
		r.SetUnsettled()
	} else if v == 1 {
		r.SetSettled()
	} else if v == 2 {
		r.SetMixed()
	}
	return r
}

/*
Settlement policy for a Sender.

<type name="sender-settle-mode" class="restricted" source="ubyte">
    <choice name="unsettled" value="0"/>
    <choice name="settled" value="1"/>
    <choice name="mixed" value="2"/>
</type>

Valid Values
0	The Sender will send all deliveries initially unsettled to the Receiver.
1	The Sender will send all deliveries settled to the Receiver.
2	The Sender may send a mixture of settled and unsettled deliveries to the Receiver.
*/
type SenderSettleMode struct {
	*UByte
}

func (s *SenderSettleMode) SetUnsettled() {
	s.UByte.value = 0
}

func (s *SenderSettleMode) SetSettled() {
	s.UByte.value = 1
}

func (s *SenderSettleMode) SetMixed() {
	s.UByte.value = 2
}

func (b *SenderSettleMode) Stringify() string {
	if b.UByte.value == 0 {
		return "unsettled"
	} else if b.UByte.value == 1 {
		return "settled"
	} else if b.UByte.value == 2 {
		return "mixed"
	}
	return "?"
}

func DecodeSenderSettleModeField(v []byte) (val *SenderSettleMode, fieldLength uint, err error) {
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
		val = NewSenderSettleMode(0)
		fieldLength = 1
		return

	case ctor == TYPE_BOOLEAN_FALSE:
		val = NewSenderSettleMode(1)
		fieldLength = 1
		return
	}

	return
}
