package performatives

import (
	"errors"
	"log"
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

func (p *PerformativeBegin) Decode(b []byte) (err error) {
	log.Println("PerformativeBegin bytes", b)

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

	log.Println("Type Performative",  Type(frameBytes[2]))
	if Type(frameBytes[2]) != TYPE_PERFORMATIVE_BEGIN {
		err = errors.New("Malformed or unexpected frame. Expecting Begin Performative.")
		log.Println(frameBytes[2])
		return
	}

	return
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
		0x11,
		0xC0,
		byte(bodyFieldLength),
	}

	performativeBytes = append(performativeBytes, bodyFieldBytes...)

	return performativeBytes, nil
}

/*
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
 */


