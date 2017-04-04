package amqp

import (
	"encoding/binary"
	"errors"
	"log"
)

type Performative interface {
	Encode() ([]byte, error)
	Decode([]byte) error
}

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

func (p *PerformativeOpen) Encode() ([]byte, error) {
	fieldContainerBytes := append([]byte{0xA1, byte(len(p.ContainerId.Value()))}, []byte(p.ContainerId.Value())...)
	fieldHostnameBytes := append([]byte{0xA1, byte(len(p.Hostname.Value()))}, []byte(p.Hostname.Value())...)
	fieldOtherFieldsBytes := []byte{0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40}

	performativeFieldSize := 1 + len(fieldContainerBytes) + len(fieldHostnameBytes) + len(fieldOtherFieldsBytes)

	performativeBytes := []byte{
		0x00,
		0x53,
		0x10,
		0xC0,
		byte(performativeFieldSize),
		0x0A,
	}

	performativeBytes = append(performativeBytes, fieldContainerBytes...)
	performativeBytes = append(performativeBytes, fieldHostnameBytes...)
	performativeBytes = append(performativeBytes, fieldOtherFieldsBytes...)

	frameSize := uint32(8 + len(performativeBytes))
	frameSizeBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(frameSizeBytes, frameSize)

	out := []byte{}

	// Header
	out = append(out, frameSizeBytes...)
	out = append(out, byte(0x02), byte(0x00), byte(0x00), byte(0x00))
	out = append(out, performativeBytes...)

	return out, nil
}

func (p *PerformativeOpen) Decode(b []byte) (err error) {
	f, err := UnmarshalFrameHeader(b)
	if err != nil {
		return
	}

	doff := f.dataOffset
	if uint32(len(b)) < f.size {
		err = errors.New("Malformed frame. Invalid size")
		return
	}

	frameBytes := b[doff*4 : f.size]

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
		log.Print(frameBytes[4])
		return
	}

	listBytes := int(frameBytes[4])
	listCount := frameBytes[5]
	log.Print("List Count", listCount)
	frameData := frameBytes[6:]

	if len(frameData)+1 != listBytes {
		err = errors.New("Malformed or unexpected frame. list size not equal or expected")
		return
	}
	remainingBytes := frameData

	var fieldSize uint

	// container-id
	p.ContainerId, fieldSize, err = GetStringField(remainingBytes)
	if err != nil {
		return
	}
	remainingBytes = remainingBytes[fieldSize:]

	// hostname
	p.Hostname, fieldSize, err = GetStringField(remainingBytes)
	if err != nil {
		return
	}
	remainingBytes = remainingBytes[fieldSize:]

	// max-frame-size
	p.MaxFrameSize, fieldSize, err = GetUIntField(remainingBytes)
	if err != nil {
		return
	}
	remainingBytes = remainingBytes[fieldSize:]

	// channel-max
	p.ChannelMax, fieldSize, err = GetUShortField(remainingBytes)
	if err != nil {
		return
	}
	remainingBytes = remainingBytes[fieldSize:]

	// idle-time-out
	p.IdleTimeout, fieldSize, err = GetUIntField(remainingBytes)
	if err != nil {
		return
	}
	remainingBytes = remainingBytes[fieldSize:]

	// outgoing-locales
	p.OutgoingLocales, fieldSize, err = GetSymbolArrayField(remainingBytes)
	if err != nil {
		return
	}
	remainingBytes = remainingBytes[fieldSize:]

	// incoming-locales
	p.IncomingLocales, fieldSize, err = GetSymbolArrayField(remainingBytes)
	if err != nil {
		return
	}
	remainingBytes = remainingBytes[fieldSize:]

	// offered-capabilities
	p.OfferedCapabilities, fieldSize, err = GetSymbolArrayField(remainingBytes)
	if err != nil {
		return
	}
	remainingBytes = remainingBytes[fieldSize:]

	// desired-capabiliites
	p.DesiredCapabilities, fieldSize, err = GetSymbolArrayField(remainingBytes)
	if err != nil {
		return
	}
	remainingBytes = remainingBytes[fieldSize:]

	// properties
	p.Properties, fieldSize, err = GetMapField(remainingBytes)
	if err != nil {
		return
	}
	remainingBytes = remainingBytes[fieldSize:]

	if len(remainingBytes) > 0 {
		log.Println("left", len(remainingBytes))
		log.Fatal("There should not be any bytes left")
	}

	log.Println("Completed parsing frame")

	return
}

func UnmarshalFrameHeader(b []byte) (f *FrameHeader, err error) {
	bLen := len(b)
	if bLen < MINIMUM_FRAME_SIZE {
		err = errors.New("Malformed frame. Invalid frame size")
		return
	}

	log.Println(b)

	sz := binary.BigEndian.Uint32(b[:4])

	doff := b[4]
	if doff < 2 {
		err = errors.New("Malformed frame. Data offset less than 2")
		return
	}

	var ft FrameType
	ftByte := b[5]
	if ftByte == 0 {
		ft = FrameTypeAmqp
	} else {
		ft = FrameTypeSasl
	}

	f = NewFrameHeader(doff, ft, 0, sz)

	return
}
