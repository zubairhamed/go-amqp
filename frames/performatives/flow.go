package performatives

import . "github.com/zubairhamed/go-amqp/types"

func NewFlowPerformative() *PerformativeFlow {
	return &PerformativeFlow{}
}

/*
<type name="flow" class="composite" source="list" provides="frame">
    <descriptor name="amqp:flow:list" code="0x00000000:0x00000013"/>
    <field name="next-incoming-id" type="transfer-number"/>
    <field name="incoming-window" type="uint" mandatory="true"/>
    <field name="next-outgoing-id" type="transfer-number" mandatory="true"/>
    <field name="outgoing-window" type="uint" mandatory="true"/>
    <field name="handle" type="handle"/>
    <field name="delivery-count" type="sequence-no"/>
    <field name="link-credit" type="uint"/>
    <field name="available" type="uint"/>
    <field name="drain" type="boolean" default="false"/>
    <field name="echo" type="boolean" default="false"/>
    <field name="properties" type="fields"/>
</type>
*/
type PerformativeFlow struct {
	BasePerformative
	NextIncomingId *TransferNumber
	IncomingWindow *UInt
	NextOutgoingId *TransferNumber
	OutgoingWindow *UInt
	Handle         *Handle
	DeliveryCount  *SequenceNumber
	LinkCredit     *UInt
	Available      *UInt
	Drain          *Boolean
	Echo           *Boolean
	Properties     *Fields
}
