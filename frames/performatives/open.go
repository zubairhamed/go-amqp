package performatives

import (
	"errors"
	. "github.com/zubairhamed/go-amqp/types"
	"log"
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
	OutgoingLocales     *IetfLanguageTagArray
	IncomingLocales     *IetfLanguageTagArray
	OfferedCapabilities *SymbolArray
	DesiredCapabilities *SymbolArray
	Properties          *Fields
}

func (b *PerformativeOpen) Stringify() string {
	return "Stringify: Performative Open"
}

func (p *PerformativeOpen) Encode() (enc []byte, l uint, err error) {
	bodyFieldBytes, bodyFieldLength, err := EncodeFields(
		p.ContainerId,
		p.Hostname,
		p.MaxFrameSize,
		p.ChannelMax,
		p.IdleTimeout,
		p.OutgoingLocales,     // EncodeIetfLanguageTagArrayField
		p.IncomingLocales,     // EncodeIetfLanguageTagArrayField
		p.OfferedCapabilities, // EncodeSymbolArrayField
		p.DesiredCapabilities, // EncodeSymbolArrayField
		p.Properties,
	)

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

	frameData, listCount, err := HandleBasePerformative(b, TYPE_PERFORMATIVE_OPEN)
	if err != nil {
		return
	}

	remainingBytes := frameData

	var fieldSize uint

	if listCount > 0 {
		// container-id
		op.ContainerId, fieldSize, err = DecodeStringField(remainingBytes)
		if err != nil {
			log.Println("Open Performative: Error occured decoding container-id Field")
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 1 {
		// hostname
		op.Hostname, fieldSize, err = DecodeStringField(remainingBytes)
		if err != nil {
			log.Println("Open Performative: Error occured decoding hostname Field")
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 2 {
		// max-frame-size
		op.MaxFrameSize, fieldSize, err = DecodeUIntField(remainingBytes)
		if err != nil {
			log.Println("Open Performative: Error occured decoding max-frame-size Field")
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 3 {
		// channel-max
		op.ChannelMax, fieldSize, err = DecodeUShortField(remainingBytes)
		if err != nil {
			log.Println("Open Performative: Error occured decoding channel-max Field")
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 4 {
		// idle-time-out
		op.IdleTimeout, fieldSize, err = DecodeMillisecondsField(remainingBytes)
		if err != nil {
			log.Println("Open Performative: Error occured decoding idle-timee-out Field")
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 5 {
		// outgoing-locales
		op.OutgoingLocales, fieldSize, err = DecodeIetfLanguageTagArrayField(remainingBytes)
		if err != nil {
			log.Println("Open Performative: Error occured decoding outgoing-locales Field")
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 6 {
		// incoming-locales
		op.IncomingLocales, fieldSize, err = DecodeIetfLanguageTagArrayField(remainingBytes)
		if err != nil {
			log.Println("Open Performative: Error occured decoding incoming-locales Field")
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 7 {
		// offered-capabilities
		op.OfferedCapabilities, fieldSize, err = DecodeSymbolArrayField(remainingBytes)
		if err != nil {
			log.Println("Open Performative: Error occured decoding offered-capabilities Field")
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 8 {
		// desired-capabiliites
		op.DesiredCapabilities, fieldSize, err = DecodeSymbolArrayField(remainingBytes)
		if err != nil {
			log.Println("Open Performative: Error occured decoding desired-capabilities Field")
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 9 {
		// properties
		op.Properties, fieldSize, err = DecodeFieldsField(remainingBytes)
		if err != nil {
			log.Println("Open Performative: Error occured decoding properties Field")
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if len(remainingBytes) > 0 {
		err = errors.New("Disposition Performative: There should not be any bytes left")
	}
	return
}
