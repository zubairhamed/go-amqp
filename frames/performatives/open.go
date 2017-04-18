package performatives

import (
	"errors"
	"github.com/zubairhamed/go-amqp/frames"
	. "github.com/zubairhamed/go-amqp/types"
)

func NewOpenPerformative() *PerformativeOpen {
	return &PerformativeOpen{}
}

/*
<type name="open" class="composite" source="list" provides="frame">
    <descriptor name="amqp:open:list" code="0x00000000:0x00000010"/>
    <field name="container-id" type="string" mandatory="true"/>
    <field name="hostname" type="string"/>
    <field name="max-frame-size" type="uint" default="4294967295"/>
    <field name="channel-max" type="ushort" default="65535"/>
    <field name="idle-time-out" type="milliseconds"/>
    <field name="outgoing-locales" type="ietf-language-tag" multiple="true"/>
    <field name="incoming-locales" type="ietf-language-tag" multiple="true"/>
    <field name="offered-capabilities" type="symbol" multiple="true"/>
    <field name="desired-capabilities" type="symbol" multiple="true"/>
    <field name="properties" type="fields"/>
</type>
*/
type PerformativeOpen struct {
	BasePerformative
	ContainerId         *String
	Hostname            *String
	MaxFrameSize        *UInt
	ChannelMax          *UShort
	IdleTimeout         *Milliseconds
	OutgoingLocales     []*IetfLanguageTag
	IncomingLocales     []*IetfLanguageTag
	OfferedCapabilities []*Symbol
	DesiredCapabilities []*Symbol
	Properties          *Fields
}

func (b *PerformativeOpen) Stringify() string {
	return "Stringify: Performative Open"
}

func (p *PerformativeOpen) Encode() (enc []byte, l uint, err error) {
	var bodyFieldBytes []byte = []byte{}
	var bodyFieldLength uint = 0
	var encField []byte
	var fieldLen uint

	encField, fieldLen, err = EncodeField(p.ContainerId)
	if err != nil {
		return
	}
	bodyFieldLength += fieldLen
	bodyFieldBytes = append(bodyFieldBytes, encField...)

	encField, fieldLen, err = EncodeField(p.Hostname)
	if err != nil {
		return
	}
	bodyFieldLength += fieldLen
	bodyFieldBytes = append(bodyFieldBytes, encField...)

	encField, fieldLen, err = EncodeField(p.MaxFrameSize)
	if err != nil {
		return
	}
	bodyFieldLength += fieldLen
	bodyFieldBytes = append(bodyFieldBytes, encField...)

	encField, fieldLen, err = EncodeField(p.ChannelMax)
	if err != nil {
		return
	}
	bodyFieldLength += fieldLen
	bodyFieldBytes = append(bodyFieldBytes, encField...)

	encField, fieldLen, err = EncodeField(p.IdleTimeout)
	if err != nil {
		return
	}
	bodyFieldLength += fieldLen
	bodyFieldBytes = append(bodyFieldBytes, encField...)

	encField, fieldLen, err = EncodeIetfLanguageTagArrayField(p.OutgoingLocales)
	if err != nil {
		return
	}
	bodyFieldLength += fieldLen
	bodyFieldBytes = append(bodyFieldBytes, encField...)

	encField, fieldLen, err = EncodeIetfLanguageTagArrayField(p.IncomingLocales)
	if err != nil {
		return
	}
	bodyFieldLength += fieldLen
	bodyFieldBytes = append(bodyFieldBytes, encField...)

	encField, fieldLen, err = EncodeSymbolArrayField(p.OfferedCapabilities)
	if err != nil {
		return
	}
	bodyFieldLength += fieldLen
	bodyFieldBytes = append(bodyFieldBytes, encField...)

	encField, fieldLen, err = EncodeSymbolArrayField(p.DesiredCapabilities)
	if err != nil {
		return
	}
	bodyFieldLength += fieldLen
	bodyFieldBytes = append(bodyFieldBytes, encField...)

	encField, fieldLen, err = EncodeField(p.Properties)
	if err != nil {
		return
	}
	bodyFieldLength += fieldLen
	bodyFieldBytes = append(bodyFieldBytes, encField...)

	performativeBytes := []byte{
		0x00, // Constructor
		0x53, // ulong small
		0x10, // performative open
		0xC0, // list
		byte(bodyFieldLength), // body bytes size
		0x0A, // field count
	}

	performativeBytes = append(performativeBytes, bodyFieldBytes...)

	return performativeBytes, uint(len(performativeBytes)), nil
}

func (b *PerformativeOpen) GetType() Type {
	return TYPE_PERFORMATIVE_OPEN
}

func DecodeOpenPerformative(b []byte) (op *PerformativeOpen, err error) {
	op = NewOpenPerformative()

	f, err := frames.UnmarshalFrameHeader(b)
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

	if Type(frameBytes[2]) != TYPE_PERFORMATIVE_OPEN {
		err = errors.New("Malformed or unexpected frame. Expecting Open Performative.")
		return
	}

	if Type(frameBytes[3]) != TYPE_LIST_8 {
		err = errors.New("Malformed or unexpected frame. Expecting list 8")
		return
	}

	listBytes := int(frameBytes[4])
	listCount := frameBytes[5]
	if listCount > 10 {
		err = errors.New("Invalid list count. Expecting 10 or less.")
		return
	}

	frameData := frameBytes[6:]

	if len(frameData)+1 != listBytes {
		err = errors.New("Malformed or unexpected frame. list size not equal or expected")
		return
	}
	remainingBytes := frameData

	var fieldSize uint

	if listCount > 0 {
		// container-id
		op.ContainerId, fieldSize, err = DecodeStringField(remainingBytes)
		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 1 {
		// hostname
		op.Hostname, fieldSize, err = DecodeStringField(remainingBytes)
		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 2 {
		// max-frame-size
		op.MaxFrameSize, fieldSize, err = DecodeUIntField(remainingBytes)
		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 3 {
		// channel-max
		op.ChannelMax, fieldSize, err = DecodeUShortField(remainingBytes)
		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 4 {
		// idle-time-out
		op.IdleTimeout, fieldSize, err = DecodeMillisecondsField(remainingBytes)
		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 5 {
		// outgoing-locales
		op.OutgoingLocales, fieldSize, err = DecodeIetfLanguageTagArrayField(remainingBytes)
		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 6 {
		// incoming-locales
		op.IncomingLocales, fieldSize, err = DecodeIetfLanguageTagArrayField(remainingBytes)
		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 7 {
		// offered-capabilities
		op.OfferedCapabilities, fieldSize, err = DecodeSymbolArrayField(remainingBytes)
		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 8 {
		// desired-capabiliites
		op.DesiredCapabilities, fieldSize, err = DecodeSymbolArrayField(remainingBytes)
		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 9 {
		// properties
		op.Properties, fieldSize, err = DecodeFieldsField(remainingBytes)
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
