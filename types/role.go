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

func (r *Role) IsSender() bool {
	return !r.value
}

func (r *Role) IsReceiver() bool {
	return r.value
}

func NewRole(sender bool) *Role {
	return &Role{
		Boolean: &Boolean{
			value: sender,
		},
	}
}

func DecodeRoleField(v []byte) (val *Role, fieldLength uint, err error) {
	iVal, fieldLength, err := DecodeBooleanField(v)

	val = &Role{
		Boolean: iVal,
	}

	return
}
