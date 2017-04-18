package performatives

import . "github.com/zubairhamed/go-amqp/types"

func NewDetatchPerformative() *PerformativeDetach {
	return &PerformativeDetach{}
}

/*
<type name="detach" class="composite" source="list" provides="frame">
    <descriptor name="amqp:detach:list" code="0x00000000:0x00000016"/>
    <field name="handle" type="handle" mandatory="true"/>
    <field name="closed" type="boolean" default="false"/>
    <field name="error" type="error"/>
</type>
*/
type PerformativeDetach struct {
	BasePerformative
	Handle *Handle
	Closed *Boolean
	Error  *Error
}
