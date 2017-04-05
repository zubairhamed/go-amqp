package performatives

import (
	"errors"
	"log"
	. "github.com/zubairhamed/go-amqp/types"
	. "github.com/zubairhamed/go-amqp/frames"
)


func NewOpenPerformative(containerId, hostname string) *PerformativeOpen {
	return &PerformativeOpen{
		ContainerId: NewString(containerId),
		Hostname:    NewString(hostname),
	}
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
	ContainerId         *String
	Hostname            *String
	MaxFrameSize        *UInt
	ChannelMax          *UShort
	IdleTimeout         *UInt
	OutgoingLocales     []*Symbol
	IncomingLocales     []*Symbol
	OfferedCapabilities []*Symbol
	DesiredCapabilities []*Symbol
	Properties          *Map
}

func (p *PerformativeOpen) Encode() (enc []byte, err error) {
	var bodyFieldBytes []byte = []byte{}
	var bodyFieldLength uint = 0
	var encField []byte
	var fieldLen uint

	encField, fieldLen, err = EncodeField(p.ContainerId)
	if err != nil {
		return
	}
	bodyFieldLength += fieldLen
	bodyFieldBytes = append(bodyFieldBytes, encField... )

	encField, fieldLen, err = EncodeField(p.Hostname)
	if err != nil {
		return
	}
	bodyFieldLength += fieldLen
	bodyFieldBytes = append(bodyFieldBytes, encField... )

	encField, fieldLen, err = EncodeField(p.MaxFrameSize)
	if err != nil {
		return
	}
	bodyFieldLength += fieldLen
	bodyFieldBytes = append(bodyFieldBytes, encField... )

	encField, fieldLen, err = EncodeField(p.ChannelMax)
	if err != nil {
		return
	}
	bodyFieldLength += fieldLen
	bodyFieldBytes = append(bodyFieldBytes, encField... )

	encField, fieldLen, err = EncodeField(p.IdleTimeout)
	if err != nil {
		return
	}
	bodyFieldLength += fieldLen
	bodyFieldBytes = append(bodyFieldBytes, encField... )

	encField, fieldLen, err = EncodeSymbolArrayField(p.OutgoingLocales)
	if err != nil {
		return
	}
	bodyFieldLength += fieldLen
	bodyFieldBytes = append(bodyFieldBytes, encField... )

	encField, fieldLen, err = EncodeSymbolArrayField(p.IncomingLocales)
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
		0x10,
		0xC0,
		byte(bodyFieldLength),
	}

	performativeBytes = append(performativeBytes, bodyFieldBytes...)

	log.Println("Encoded Open Performative", performativeBytes)

	return performativeBytes, nil

	//fieldContainerBytes := append([]byte{0xA1, byte(len(p.ContainerId.Value()))}, []byte(p.ContainerId.Value())...)
	//fieldHostnameBytes := append([]byte{0xA1, byte(len(p.Hostname.Value()))}, []byte(p.Hostname.Value())...)
	//fieldOtherFieldsBytes := []byte{0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40}
	//
	//performativeFieldSize := 1 + len(fieldContainerBytes) + len(fieldHostnameBytes) + len(fieldOtherFieldsBytes)
	//
	//performativeBytes := []byte{
	//	0x00,
	//	0x53,
	//	0x10,
	//	0xC0,
	//	byte(performativeFieldSize),
	//	0x0A,
	//}
	//
	//performativeBytes = append(performativeBytes, fieldContainerBytes...)
	//performativeBytes = append(performativeBytes, fieldHostnameBytes...)
	//performativeBytes = append(performativeBytes, fieldOtherFieldsBytes...)
	//
	//frameSize := uint32(8 + len(performativeBytes))
	//frameSizeBytes := make([]byte, 4)
	//binary.BigEndian.PutUint32(frameSizeBytes, frameSize)
	//
	//out := []byte{}
	//
	//// Header
	//out = append(out, frameSizeBytes...)
	//out = append(out, byte(0x02), byte(0x00), byte(0x00), byte(0x00))
	//out = append(out, performativeBytes...)
	//
	//return out, nil
}

func (p *PerformativeOpen) Decode(b []byte) (err error) {
	log.Println("Decode Performative Open", b)
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

	if Type(frameBytes[2]) != TYPE_PERFORMATIVE_OPEN {
		err = errors.New("Malformed or unexpected frame. Expecting Open Performative.")
		return
	}

	if Type(frameBytes[3]) != TYPE_LIST_8 {
		err = errors.New("Malformed or unexpected frame. Expecting list 8")
		return
	}

	listBytes := int(frameBytes[4])
	// listCount := frameBytes[5]
	frameData := frameBytes[6:]

	if len(frameData)+1 != listBytes {
		err = errors.New("Malformed or unexpected frame. list size not equal or expected")
		return
	}
	remainingBytes := frameData

	var fieldSize uint

	// container-id
	p.ContainerId, fieldSize, err = DecodeStringField(remainingBytes)
	if err != nil {
		return
	}
	remainingBytes = remainingBytes[fieldSize:]

	// hostname
	p.Hostname, fieldSize, err = DecodeStringField(remainingBytes)
	if err != nil {
		return
	}
	remainingBytes = remainingBytes[fieldSize:]

	// max-frame-size
	p.MaxFrameSize, fieldSize, err = DecodeUIntField(remainingBytes)
	if err != nil {
		return
	}
	remainingBytes = remainingBytes[fieldSize:]

	// channel-max
	p.ChannelMax, fieldSize, err = DecodeUShortField(remainingBytes)
	if err != nil {
		return
	}
	remainingBytes = remainingBytes[fieldSize:]

	// idle-time-out
	p.IdleTimeout, fieldSize, err = DecodeUIntField(remainingBytes)
	if err != nil {
		return
	}
	remainingBytes = remainingBytes[fieldSize:]

	// outgoing-locales
	p.OutgoingLocales, fieldSize, err = DecodeSymbolArrayField(remainingBytes)
	if err != nil {
		return
	}
	remainingBytes = remainingBytes[fieldSize:]

	// incoming-locales
	p.IncomingLocales, fieldSize, err = DecodeSymbolArrayField(remainingBytes)
	if err != nil {
		return
	}
	remainingBytes = remainingBytes[fieldSize:]

	// offered-capabilities
	p.OfferedCapabilities, fieldSize, err = DecodeSymbolArrayField(remainingBytes)
	if err != nil {
		return
	}
	remainingBytes = remainingBytes[fieldSize:]

	// desired-capabiliites
	p.DesiredCapabilities, fieldSize, err = DecodeSymbolArrayField(remainingBytes)
	if err != nil {
		return
	}
	remainingBytes = remainingBytes[fieldSize:]

	// properties
	p.Properties, fieldSize, err = DecodeMapField(remainingBytes)
	if err != nil {
		return
	}
	remainingBytes = remainingBytes[fieldSize:]

	if len(remainingBytes) > 0 {
		log.Fatal("There should not be any bytes left")
	}

	log.Println("Completed parsing frame")

	return
}
