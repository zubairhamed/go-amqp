package performatives

import (
	"errors"
	. "github.com/zubairhamed/go-amqp/types"
)

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

func (b *PerformativeDetach) GetType() Type {
	return TYPE_PERFORMATIVE_DETACH
}

func (p *PerformativeDetach) Encode() (enc []byte, l uint, err error) {
	return
}

func (b *PerformativeDetach) Stringify() string {
	return "Stringify: Performative Detatch"
}

func DecodeDetachPerformative(b []byte) (op *PerformativeDetach, err error) {
	op = NewDetatchPerformative()
	frameData, listCount, err := HandleBasePerformative(b, TYPE_PERFORMATIVE_DETACH)
	if err != nil {
		return
	}

	remainingBytes := frameData
	var fieldSize uint

	if listCount > 0 {
		// handle
		op.Handle, fieldSize, err = DecodeHandleField(remainingBytes)
		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 1 {
		// closed
		op.Closed, fieldSize, err = DecodeBooleanField(remainingBytes)
		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 2 {
		// error
		op.Error, fieldSize, err = DecodeErrorField(remainingBytes)
		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if len(remainingBytes) > 0 {
		err = errors.New("Detach Performative: There should not be any bytes left")
	}

	return
}
