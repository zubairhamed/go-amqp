package performatives

import . "github.com/zubairhamed/go-amqp/types"

func NewClosePerformative() *PerformativeClose {
	return &PerformativeClose{}
}

/*
<type name="close" class="composite" source="list" provides="frame">
    <descriptor name="amqp:close:list" code="0x00000000:0x00000018"/>
    <field name="error" type="error"/>
</type>
*/
type PerformativeClose struct {
	BasePerformative
	Error *Error
}
