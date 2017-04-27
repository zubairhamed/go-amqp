package performatives

import (
	"errors"
	. "github.com/zubairhamed/go-amqp/types"
)

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
	State     []byte
	Batchable *Boolean
}

func (b *PerformativeDisposition) GetType() Type {
	return TYPE_PERFORMATIVE_DISPOSITION
}

func (p *PerformativeDisposition) Encode() (enc []byte, l uint, err error) {
	return
}

func (b *PerformativeDisposition) Stringify() string {
	return "Stringify: Performative Disposition"
}

func DecodeDispositionPerformative(b []byte) (op *PerformativeDisposition, err error) {
	op = NewDispositionPerformative()
	frameData, listCount, err := HandleBasePerformative(b, TYPE_PERFORMATIVE_DISPOSITION)
	if err != nil {
		return
	}

	remainingBytes := frameData
	var fieldSize uint

	if listCount > 0 {
		// role
		op.Role, fieldSize, err = DecodeRoleField(remainingBytes)

		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 1 {
		// first
		op.First, fieldSize, err = DecodeDeliveryNumberField(remainingBytes)

		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 2 {
		// last
		op.Last, fieldSize, err = DecodeDeliveryNumberField(remainingBytes)

		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 3 {
		// settled
		op.Settled, fieldSize, err = DecodeBooleanField(remainingBytes)

		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 4 {
		// state
		op.State, fieldSize, err = DecodeAnyTypeField(remainingBytes)
		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 5 {
		// batchable
		op.Batchable, fieldSize, err = DecodeBooleanField(remainingBytes)
		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if len(remainingBytes) > 0 {
		err = errors.New("Disposition Performative: There should not be any bytes left")
	}

	return
}
