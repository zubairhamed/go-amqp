package performatives

import . "github.com/zubairhamed/go-amqp/types"

func NewEndPerformative() *PerformativeEnd {
	return &PerformativeEnd{}
}

/*
<type name="end" class="composite" source="list" provides="frame">
    <descriptor name="amqp:end:list" code="0x00000000:0x00000017"/>
    <field name="error" type="error"/>
</type>
*/

type PerformativeEnd struct {
	BasePerformative
	Error *Error
}
