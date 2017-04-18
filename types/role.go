package types

/*
<type name="role" class="restricted" source="boolean">
    <choice name="sender" value="false"/>
    <choice name="receiver" value="true"/>
</type>
*/
type Role struct {
	*Boolean
}
