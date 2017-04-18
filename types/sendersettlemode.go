package types

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
	choiceValue byte
}

func (s *SenderSettleMode) SetUnsettled() {
	s.choiceValue = 0
}

func (s *SenderSettleMode) SetSettled() {
	s.choiceValue = 1
}

func (s *SenderSettleMode) SetMixed() {
	s.choiceValue = 2
}
