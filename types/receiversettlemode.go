package types

/*
<type name="receiver-settle-mode" class="restricted" source="ubyte">
    <choice name="first" value="0"/>
    <choice name="second" value="1"/>
</type>
*/
type ReceiverSettleMode struct {
	*UByte
	choiceValue byte
}

func (r *ReceiverSettleMode) SetFirst() {
	r.choiceValue = 0
}

func (r *ReceiverSettleMode) SetSecond() {
	r.choiceValue = 1
}
