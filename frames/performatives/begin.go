package performatives

import (
	"errors"
	. "github.com/zubairhamed/go-amqp/types"
	. "github.com/zubairhamed/go-amqp/frames"
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
	RemoteChannel       *UShort
	NextOutgoingId      *UInt
	IncomingWindow      *UInt
	OutgoingWindow      *UInt
	HandleMax           *UInt
	OfferedCapabilities []*Symbol
	DesiredCapabilities []*Symbol
	Properties          *Map
}

func (p *PerformativeBegin) Encode() (enc []byte, err error) {
	var bodyFieldBytes []byte = []byte{}
	var bodyFieldLength uint = 0
	var encField []byte
	var fieldLen uint

	encField, fieldLen, err = EncodeField(p.RemoteChannel)
	if err != nil {
		return
	}
	bodyFieldLength += fieldLen
	bodyFieldBytes = append(bodyFieldBytes, encField... )

	encField, fieldLen, err = EncodeField(p.NextOutgoingId)
	if err != nil {
		return
	}
	bodyFieldLength += fieldLen
	bodyFieldBytes = append(bodyFieldBytes, encField... )

	encField, fieldLen, err = EncodeField(p.IncomingWindow)
	if err != nil {
		return
	}
	bodyFieldLength += fieldLen
	bodyFieldBytes = append(bodyFieldBytes, encField... )

	encField, fieldLen, err = EncodeField(p.OutgoingWindow)
	if err != nil {
		return
	}
	bodyFieldLength += fieldLen
	bodyFieldBytes = append(bodyFieldBytes, encField... )

	encField, fieldLen, err = EncodeField(p.HandleMax)
	if err != nil {
		return
	}
	bodyFieldLength += fieldLen
	bodyFieldBytes = append(bodyFieldBytes, encField... )

	encField, fieldLen, err = EncodeSymbolArrayField(p.OfferedCapabilities)
	if err != nil {
		return
	}
	bodyFieldLength += fieldLen
	bodyFieldBytes = append(bodyFieldBytes, encField... )

	encField, fieldLen, err = EncodeSymbolArrayField(p.DesiredCapabilities)
	if err != nil {
		return
	}
	bodyFieldLength += fieldLen
	bodyFieldBytes = append(bodyFieldBytes, encField... )

	encField, fieldLen, err = EncodeField(p.Properties)
	if err != nil {
		return
	}
	bodyFieldLength += fieldLen
	bodyFieldBytes = append(bodyFieldBytes, encField... )

	performativeBytes := []byte{
		0x00,
		0x53,
		byte(TYPE_PERFORMATIVE_BEGIN),
		0xC0,
		byte(bodyFieldLength),
		0x08,
	}

	performativeBytes = append(performativeBytes, bodyFieldBytes...)

	return performativeBytes, nil
}

func DecodeBeginPerformative(b []byte) (op *PerformativeBegin, err error) {
	op = NewBeginPerformative()

	f, err := UnmarshalFrameHeader(b)
	if err != nil {
		return
	}

	doff := f.DataOffset
	if uint32(len(b)) < f.Size {
		err = errors.New("Malformed frame. Invalid size")
		return
	}

	frameBytes := b[doff*4 : f.Size]

	if Type(frameBytes[0]) != TYPE_CONSTRUCTOR {
		err = errors.New("Malformed or unexpected frame. Expecting constructor.")
		return
	}

	if Type(frameBytes[1]) != TYPE_ULONG_SMALL {
		err = errors.New("Malformed or unexpected frame. Expecting small ulong type")
		return
	}

	if Type(frameBytes[2]) != TYPE_PERFORMATIVE_BEGIN {
		err = errors.New("Malformed or unexpected frame. Expecting Begin Performative.")
		return
	}

	if Type(frameBytes[3]) != TYPE_LIST_8 {
		err = errors.New("Malformed or unexpected frame. Expecting list 8")
		return
	}

	listBytes := int(frameBytes[4])
	listCount := frameBytes[5]
	if listCount > 8 {
		err = errors.New("Invalid list count. Expecting 8 or less.")
		return
	}

	frameData := frameBytes[6:]

	if len(frameData)+1 != listBytes {
		err = errors.New("Malformed or unexpected frame. list size not equal or expected")
		return
	}
	remainingBytes := frameData

	var fieldSize uint

	if listCount  > 0 {
		// remote-channel
		op.RemoteChannel, fieldSize, err = DecodeUShortField(remainingBytes)
		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 1 {
		// next-outgoing-id
		op.NextOutgoingId, fieldSize, err = DecodeUIntField(remainingBytes)
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
		op.HandleMax, fieldSize, err = DecodeUIntField(remainingBytes)
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
		op.Properties, fieldSize, err = DecodeMapField(remainingBytes)
		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if len(remainingBytes) > 0 {
		err = errors.New("There should not be any bytes left")
	}
	return
}


