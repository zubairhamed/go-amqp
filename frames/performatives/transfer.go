package performatives

import . "github.com/zubairhamed/go-amqp/types"

func NewTransferPerformative() *PerformativeTransfer {
	return &PerformativeTransfer{}
}

/*
<type name="transfer" class="composite" source="list" provides="frame">
    <descriptor name="amqp:transfer:list" code="0x00000000:0x00000014"/>
    <field name="handle" type="handle" mandatory="true"/>
    <field name="delivery-id" type="delivery-number"/>
    <field name="delivery-tag" type="delivery-tag"/>
    <field name="message-format" type="message-format"/>
    <field name="settled" type="boolean"/>
    <field name="more" type="boolean" default="false"/>
    <field name="rcv-settle-mode" type="receiver-settle-mode"/>
    <field name="state" type="*" requires="delivery-state"/>
    <field name="resume" type="boolean" default="false"/>
    <field name="aborted" type="boolean" default="false"/>
    <field name="batchable" type="boolean" default="false"/>
</type>
*/

type PerformativeTransfer struct {
	BasePerformative
	Handle             *Handle
	DeliveryId         *DeliveryNumber
	DeliveryTag        *DeliveryTag
	MessageFormat      *MessageFormat
	Settled            *Boolean
	More               *Boolean
	ReceiverSettleMode *ReceiverSettleMode
	State              interface{}
	Resume             *Boolean
	Aborted            *Boolean
	Batchable          *Boolean
}
