package performatives

import (
	"errors"
	. "github.com/zubairhamed/go-amqp/types"
	"log"
)

func NewAttachPerformative() *PerformativeAttach {
	return &PerformativeAttach{}
}

/*
<type name="attach" class="composite" source="list" provides="frame">
    <descriptor name="amqp:attach:list" code="0x00000000:0x00000012"/>
    <field name="name" type="string" mandatory="true"/>
    <field name="handle" type="handle" mandatory="true"/>
    <field name="role" type="role" mandatory="true"/>
    <field name="snd-settle-mode" type="sender-settle-mode" default="mixed"/>
    <field name="rcv-settle-mode" type="receiver-settle-mode" default="first"/>
    <field name="source" type="*" requires="source"/>
    <field name="target" type="*" requires="target"/>
    <field name="unsettled" type="map"/>
    <field name="incomplete-unsettled" type="boolean" default="false"/>
    <field name="initial-delivery-count" type="sequence-no"/>
    <field name="max-message-size" type="ulong"/>
    <field name="offered-capabilities" type="symbol" multiple="true"/>
    <field name="desired-capabilities" type="symbol" multiple="true"/>
    <field name="properties" type="fields"/>
</type>
*/
type PerformativeAttach struct {
	BasePerformative
	Name                 *String
	Handle               *Handle
	Role                 *Role
	SenderSettleMode     *SenderSettleMode
	ReceiverSettleMode   *ReceiverSettleMode
	Source               AMQPType
	Target               AMQPType
	Unsettled            *Map
	IncompleteUnsettled  *Boolean
	InitialDeliveryCount *SequenceNumber
	MaxMessageSize       *ULong
	OfferedCapabilities  *SymbolArray
	DesiredCapabilities  *SymbolArray
	Properties           *Fields
}

func (p *PerformativeAttach) Encode() (enc []byte, l uint, err error) {
	bodyFieldBytes, bodyFieldLength, err := EncodeFields(
		p.Name,
		p.Handle,
		p.Role,
		p.SenderSettleMode,
		p.ReceiverSettleMode,
		p.Source,
		p.Target,
		p.Unsettled,
		p.IncompleteUnsettled,
		p.InitialDeliveryCount,
		p.MaxMessageSize,
		p.OfferedCapabilities,
		p.DesiredCapabilities,
		p.Properties,
	)

	performativeBytes := []byte{
		0x00,
		0x53,
		byte(TYPE_PERFORMATIVE_ATTACH),
		0xC0,
		byte(bodyFieldLength),
		0x08,
	}

	performativeBytes = append(performativeBytes, bodyFieldBytes...)

	return performativeBytes, uint(len(performativeBytes)), nil
}

func (b *PerformativeAttach) GetType() Type {
	return TYPE_PERFORMATIVE_ATTACH
}

func (b *PerformativeAttach) Stringify() string {
	return "Stringify: Performative Attach"
}

func DecodeAttachPerformative(b []byte) (op *PerformativeAttach, err error) {
	op = NewAttachPerformative()

	frameData, listCount, err := HandleBasePerformative(b, TYPE_PERFORMATIVE_ATTACH)
	if err != nil {
		return
	}

	remainingBytes := frameData
	var fieldSize uint

	if listCount > 0 {
		//  name
		op.Name, fieldSize, err = DecodeStringField(remainingBytes)
		if err != nil {
			log.Println("Open Performative: Error occured decoding name Field")
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 1 {
		// handle
		op.Handle, fieldSize, err = DecodeHandleField(remainingBytes)
		if err != nil {
			log.Println("Open Performative: Error occured decoding handle Field")
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 2 {
		// role
		op.Role, fieldSize, err = DecodeRoleField(remainingBytes)
		if err != nil {
			log.Println("Open Performative: Error occured decoding role Field")
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 3 {
		// snd-settle-mode
		op.SenderSettleMode, fieldSize, err = DecodeSenderSettleModeField(remainingBytes)
		if err != nil {
			log.Println("Open Performative: Error occured decoding snd-settle-mode Field")
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 4 {
		// rcv-settle-mode
		op.ReceiverSettleMode, fieldSize, err = DecodeReceiverSettleModeField(remainingBytes)
		if err != nil {
			log.Println("Open Performative: Error occured decoding rcv-settle-mode Field")
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 5 {
		// source
		op.Source, fieldSize, err = DecodeStringField(remainingBytes)
		if err != nil {
			log.Println("Open Performative: Error occured decoding source Field")
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 6 {
		// target
		op.Target, fieldSize, err = DecodeStringField(remainingBytes)
		if err != nil {
			log.Println("Open Performative: Error occured decoding target Field")
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 7 {
		// unsettled
		op.Unsettled, fieldSize, err = DecodeMapField(remainingBytes)

		if err != nil {
			log.Println("Open Performative: Error occured decoding unsettled Field")
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 8 {
		// incomplete-unsettled
		op.IncompleteUnsettled, fieldSize, err = DecodeBooleanField(remainingBytes)
		if err != nil {
			log.Println("Open Performative: Error occured decoding incomplete-unsettled Field")
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 9 {
		// initial-delivery-count
		op.InitialDeliveryCount, fieldSize, err = DecodeSequenceNumber(remainingBytes)
		if err != nil {
			log.Println("Open Performative: Error occured decoding initial-delivery-count Field")
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 10 {
		// max-message-size
		op.MaxMessageSize, fieldSize, err = DecodeULongField(remainingBytes)
		if err != nil {
			log.Println("Open Performative: Error occured decoding max-message-sizae Field")
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 11 {
		// offered-capabilities
		op.OfferedCapabilities, fieldSize, err = DecodeSymbolArrayField(remainingBytes)
		if err != nil {
			log.Println("Open Performative: Error occured decoding offered-capabilities Field")
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 12 {
		// desired-capabilities
		op.DesiredCapabilities, fieldSize, err = DecodeSymbolArrayField(remainingBytes)
		if err != nil {
			log.Println("Open Performative: Error occured decoding desired-capabilities Field")
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 13 {
		// properties
		op.Properties, fieldSize, err = DecodeFieldsField(remainingBytes)
		if err != nil {
			log.Println("Open Performative: Error occured decoding properties field")
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if len(remainingBytes) > 0 {
		log.Println(remainingBytes)
		err = errors.New("Attach Performative: There should not be any bytes left")
	}
	return
}
