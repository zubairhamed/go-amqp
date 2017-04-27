package performatives

import (
	"errors"
	. "github.com/zubairhamed/go-amqp/types"
)

func NewFlowPerformative() *PerformativeFlow {
	return &PerformativeFlow{}
}

/*
<type name="flow" class="composite" source="list" provides="frame">
    <descriptor name="amqp:flow:list" code="0x00000000:0x00000013"/>
    <field name="next-incoming-id" type="transfer-number"/>
    <field name="incoming-window" type="uint" mandatory="true"/>
    <field name="next-outgoing-id" type="transfer-number" mandatory="true"/>
    <field name="outgoing-window" type="uint" mandatory="true"/>
    <field name="handle" type="handle"/>
    <field name="delivery-count" type="sequence-no"/>
    <field name="link-credit" type="uint"/>
    <field name="available" type="uint"/>
    <field name="drain" type="boolean" default="false"/>
    <field name="echo" type="boolean" default="false"/>
    <field name="properties" type="fields"/>
</type>
*/
type PerformativeFlow struct {
	BasePerformative
	NextIncomingId *TransferNumber
	IncomingWindow *UInt
	NextOutgoingId *TransferNumber
	OutgoingWindow *UInt
	Handle         *Handle
	DeliveryCount  *SequenceNumber
	LinkCredit     *UInt
	Available      *UInt
	Drain          *Boolean
	Echo           *Boolean
	Properties     *Fields
}

func (b *PerformativeFlow) GetType() Type {
	return TYPE_PERFORMATIVE_FLOW
}

func (p *PerformativeFlow) Encode() (enc []byte, l uint, err error) {
	return
}

func (b *PerformativeFlow) Stringify() string {
	return "Stringify: Performative Flow"
}

func DecodeFlowPerformative(b []byte) (op *PerformativeFlow, err error) {
	op = NewFlowPerformative()
	frameData, listCount, err := HandleBasePerformative(b, TYPE_PERFORMATIVE_FLOW)
	if err != nil {
		return
	}

	remainingBytes := frameData
	var fieldSize uint

	if listCount > 0 {
		// next-incoming-id
		op.NextIncomingId, fieldSize, err = DecodeTransferNumberField(remainingBytes)

		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 1 {
		// incoming-window
		op.IncomingWindow, fieldSize, err = DecodeUIntField(remainingBytes)

		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 2 {
		// next-outgoing-id
		op.NextIncomingId, fieldSize, err = DecodeTransferNumberField(remainingBytes)
		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 3 {
		// outgoing-window
		op.OutgoingWindow, fieldSize, err = DecodeUIntField(remainingBytes)
		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 4 {
		// handle
		op.Handle, fieldSize, err = DecodeHandleField(remainingBytes)
		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 5 {
		// delivery-count
		op.DeliveryCount, fieldSize, err = DecodeSequenceNumber(remainingBytes)
		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 6 {
		// link-credit
		op.LinkCredit, fieldSize, err = DecodeUIntField(remainingBytes)
		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 7 {
		// available
		op.Available, fieldSize, err = DecodeUIntField(remainingBytes)
		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 8 {
		// drain
		op.Drain, fieldSize, err = DecodeBooleanField(remainingBytes)
		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 9 {
		// echo
		op.Echo, fieldSize, err = DecodeBooleanField(remainingBytes)
		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 10 {
		// properties
		op.Properties, fieldSize, err = DecodeFieldsField(remainingBytes)
		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if len(remainingBytes) > 0 {
		err = errors.New("Flow Performative: There should not be any bytes left")
	}

	return
}
