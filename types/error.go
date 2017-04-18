package types


func NewError(condition string, description string, info map[string]AMQPType) *Error {
	return &Error{
		Condition: NewSymbol(condition),
		Description: NewString(description),
		Info: NewFields(info),
	}
}

/*
<type name="error" class="composite" source="list">
    <descriptor name="amqp:error:list" code="0x00000000:0x0000001d"/>
    <field name="condition" type="symbol" requires="error-condition" mandatory="true"/>
    <field name="description" type="string"/>
    <field name="info" type="fields"/>
</type>
*/
type Error struct {
	*List
	Condition   *Symbol
	Description *String
	Info        *Fields
}
