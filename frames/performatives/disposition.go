package performatives

import . "github.com/zubairhamed/go-amqp/types"

func NewDispositionPerformative() *PerformativeDisposition {
	return &PerformativeDisposition{}
}

/*
<type name="disposition" class="composite" source="list" provides="frame">
    <descriptor name="amqp:disposition:list" code="0x00000000:0x00000015"/>
    <field name="role" type="role" mandatory="true"/>
    <field name="first" type="delivery-number" mandatory="true"/>
    <field name="last" type="delivery-number"/>
    <field name="settled" type="boolean" default="false"/>
    <field name="state" type="*" requires="delivery-state"/>
    <field name="batchable" type="boolean" default="false"/>
</type>
*/

type PerformativeDisposition struct {
	BasePerformative
	Role      *Role
	First     *DeliveryNumber
	Last      *DeliveryNumber
	Settled   *Boolean
	State     interface{}
	Batchable *Boolean
}
