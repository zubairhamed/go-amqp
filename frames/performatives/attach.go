package performatives

import . "github.com/zubairhamed/go-amqp/types"

func NewAttachPerformative() *PerformativeAttach {
	return &PerformativeAttach{}
}

/*
<type name="attach" class="composite" source="list" provides="frame">
    <descriptor name="amqp:attach:list" code="0x00000000:0x00000012"/>
    <field name="name" type="string" mandatory="true"/>
    <field name="handle" type="handle" mandatory="true"/>
    <field name="role" type="role" mandatory="true"/>
    <field name="snd-settle-mode" type="sender-settle-mode" default="mixed"/>
    <field name="rcv-settle-mode" type="receiver-settle-mode" default="first"/>
    <field name="source" type="*" requires="source"/>
    <field name="target" type="*" requires="target"/>
    <field name="unsettled" type="map"/>
    <field name="incomplete-unsettled" type="boolean" default="false"/>
    <field name="initial-delivery-count" type="sequence-no"/>
    <field name="max-message-size" type="ulong"/>
    <field name="offered-capabilities" type="symbol" multiple="true"/>
    <field name="desired-capabilities" type="symbol" multiple="true"/>
    <field name="properties" type="fields"/>
</type>
*/
type PerformativeAttach struct {
	BasePerformative
	Name                 *String
	Handle               *Handle
	Role                 *Role
	SendSettleMode       *SenderSettleMode
	ReceiverSetleMode    *ReceiverSettleMode
	Source               *String
	Target               *String
	Unsettled            *Map
	IncompleteUnsettled  *Boolean
	InitialDeliveryCount *SequenceNumber
	MaxMessageSize       *ULong
	OfferedCapabilities  []*Symbol
	DesiredCapabilities  []*Symbol
	Properties           *Fields
}

func (p *PerformativeAttach) Encode() (enc []byte, l uint, err error) {
	return
}

func (b *PerformativeAttach) Stringify() string {
	return "Stringify: Performative Attach"
}
