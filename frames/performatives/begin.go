package performatives

import (
	"errors"
	. "github.com/zubairhamed/go-amqp/types"
)

func NewBeginPerformative() *PerformativeBegin {
	return &PerformativeBegin{}
}

/*
<type name="begin" class="composite" source="list" provides="frame">
    <descriptor name="amqp:begin:list" code="0x00000000:0x00000011"/>
    <field name="remote-channel" type="ushort"/>
    <field name="next-outgoing-id" type="transfer-number" mandatory="true"/>
    <field name="incoming-window" type="uint" mandatory="true"/>
    <field name="outgoing-window" type="uint" mandatory="true"/>
    <field name="handle-max" type="handle" default="4294967295"/>
    <field name="offered-capabilities" type="symbol" multiple="true"/>
    <field name="desired-capabilities" type="symbol" multiple="true"/>
    <field name="properties" type="fields"/>
</type>
*/
type PerformativeBegin struct {
	BasePerformative
	RemoteChannel       *UShort
	NextOutgoingId      *TransferNumber
	IncomingWindow      *UInt
	OutgoingWindow      *UInt
	HandleMax           *Handle
	OfferedCapabilities *SymbolArray
	DesiredCapabilities *SymbolArray
	Properties          *Fields
}

func (p *PerformativeBegin) Encode() (enc []byte, l uint, err error) {
	bodyFieldBytes, bodyFieldLength, err := EncodeFields(
		p.RemoteChannel,
		p.NextOutgoingId,
		p.IncomingWindow,
		p.OutgoingWindow,
		p.HandleMax,
		p.OfferedCapabilities,
		p.DesiredCapabilities,
		p.Properties,
	)

	performativeBytes := []byte{
		0x00,
		0x53,
		byte(TYPE_PERFORMATIVE_BEGIN),
		0xC0,
		byte(bodyFieldLength),
		0x08,
	}

	performativeBytes = append(performativeBytes, bodyFieldBytes...)

	return performativeBytes, uint(len(performativeBytes)), nil
}

func (b *PerformativeBegin) Stringify() string {
	return "Stringify: Performative Begin"
}

func (b *PerformativeBegin) GetType() Type {
	return TYPE_PERFORMATIVE_BEGIN
}

func DecodeBeginPerformative(b []byte) (op *PerformativeBegin, err error) {
	op = NewBeginPerformative()

	frameData, listCount, err := HandleBasePerformative(b, TYPE_PERFORMATIVE_BEGIN)
	if err != nil {
		return
	}

	remainingBytes := frameData

	var fieldSize uint
	if listCount > 0 {
		// remote-channel
		op.RemoteChannel, fieldSize, err = DecodeUShortField(remainingBytes)
		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 1 {
		// next-outgoing-id
		op.NextOutgoingId, fieldSize, err = DecodeTransferNumberField(remainingBytes)
		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 2 {
		// incoming-window
		op.IncomingWindow, fieldSize, err = DecodeUIntField(remainingBytes)
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
		// handle-max
		op.HandleMax, fieldSize, err = DecodeHandleField(remainingBytes)
		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 5 {
		// offered-capabilities
		op.OfferedCapabilities, fieldSize, err = DecodeSymbolArrayField(remainingBytes)
		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 6 {
		// desired-capabilities
		op.DesiredCapabilities, fieldSize, err = DecodeSymbolArrayField(remainingBytes)
		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 7 {
		// properties
		op.Properties, fieldSize, err = DecodeFieldsField(remainingBytes)
		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if len(remainingBytes) > 0 {
		err = errors.New("Begin Performative: There should not be any bytes left")
	}
	return
}
