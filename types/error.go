package types

import (
	"errors"
)

func NewError(condition string, description string, info map[string]AMQPType) *Error {
	return &Error{
		Condition:   NewSymbol(condition),
		Description: NewString(description),
		Info:        NewFields(info),
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

func DecodeErrorField(v []byte) (val *Error, fieldLength uint, err error) {
	if Type(v[0]) != TYPE_CONSTRUCTOR {
		err = errors.New("Malformed or unexpected frame. Expecting constructor.")
		return
	}

	if Type(v[1]) != TYPE_ULONG_SMALL {
		err = errors.New("Malformed or unexpected frame. Expecting small ulong type")
		return
	}

	if Type(v[2]) != TYPE_ERROR {
		err = errors.New("Malformed or unexpected frame. Expecting Error Type")
		return
	}

	if Type(v[3]) != TYPE_LIST_8 {
		err = errors.New("Malformed or unexpected frame. Expecting list 8")
		return
	}

	listBytes := int(v[4])
	listCount := int(v[5])

	if listCount > 3 {
		err = errors.New("Invalid list count. Expecting 3 or less")
		return
	}

	frameData := v[6:]
	if len(frameData)+1 != listBytes {
		err = errors.New("Malformed or unexpected frame. list size not equal or expected")
		return
	}

	remainingBytes := frameData

	var fieldSize uint
	val = &Error{}
	if listCount > 0 {
		val.Condition, fieldSize, err = DecodeSymbolField(remainingBytes)
		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 1 {
		val.Description, fieldSize, err = DecodeStringField(remainingBytes)
		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 2 {
		val.Info, fieldSize, err = DecodeFieldsField(remainingBytes)
		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if len(remainingBytes) > 0 {
		err = errors.New("There should not be any bytes left")
	}

	fieldLength = uint(listBytes + 5)

	return
}
