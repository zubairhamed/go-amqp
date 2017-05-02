package performatives

import (
	"errors"
	. "github.com/zubairhamed/go-amqp/types"
)

func NewTransferPerformative() *PerformativeTransfer {
	return &PerformativeTransfer{}
}

/*
<type name="transfer" class="composite" source="list" provides="frame">
    <descriptor name="amqp:transfer:list" code="0x00000000:0x00000014"/>
    <field name="handle" type="handle" mandatory="true"/>
    <field name="delivery-id" type="delivery-number"/>
    <field name="delivery-tag" type="delivery-tag"/>
    <field name="message-format" type="message-format"/>
    <field name="settled" type="boolean"/>
    <field name="more" type="boolean" default="false"/>
    <field name="rcv-settle-mode" type="receiver-settle-mode"/>
    <field name="state" type="*" requires="delivery-state"/>
    <field name="resume" type="boolean" default="false"/>
    <field name="aborted" type="boolean" default="false"/>
    <field name="batchable" type="boolean" default="false"/>
</type>
*/

type PerformativeTransfer struct {
	BasePerformative
	Handle             *Handle
	DeliveryId         *DeliveryNumber
	DeliveryTag        *DeliveryTag
	MessageFormat      *MessageFormat
	Settled            *Boolean
	More               *Boolean
	ReceiverSettleMode *ReceiverSettleMode
	State              AMQPType
	Resume             *Boolean
	Aborted            *Boolean
	Batchable          *Boolean
	Value              *AMQPValue
}

func (b *PerformativeTransfer) GetType() Type {
	return TYPE_PERFORMATIVE_TRANSFER
}

func (p *PerformativeTransfer) Encode() (enc []byte, l uint, err error) {
	bodyFieldBytes, bodyFieldLength, err := EncodeFields(
		p.Handle,
		p.DeliveryId,
		p.DeliveryTag,
		p.MessageFormat,
		p.Settled,
		p.More,
		p.ReceiverSettleMode,
		p.State,
		p.Resume,
		p.Aborted,
		p.Batchable,
		p.Value,
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

func DecodeTransferPerformative(b []byte) (op *PerformativeTransfer, err error) {
	op = NewTransferPerformative()
	frameData, listCount, err := HandleBasePerformative(b, TYPE_PERFORMATIVE_TRANSFER)
	if err != nil {
		return
	}

	remainingBytes := frameData
	var fieldSize uint

	if listCount > 0 {
		// handle
		op.Handle, fieldSize, err = DecodeHandleField(remainingBytes)

		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 1 {
		// delivery-id
		op.DeliveryId, fieldSize, err = DecodeDeliveryNumberField(remainingBytes)
		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 2 {
		// delivery-tag
		op.DeliveryTag, fieldSize, err = DecodeDeliveryTagField(remainingBytes)
		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 3 {
		// message-format
		op.MessageFormat, fieldSize, err = DecodeMessageFormatField(remainingBytes)
		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 4 {
		// settled
		op.Settled, fieldSize, err = DecodeBooleanField(remainingBytes)

		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 5 {
		// more
		op.More, fieldSize, err = DecodeBooleanField(remainingBytes)

		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 6 {
		// rcv-settled-mode
		op.ReceiverSettleMode, fieldSize, err = DecodeReceiverSettleModeField(remainingBytes)
		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 7 {
		// state
		op.State, fieldSize, err = DecodeField(remainingBytes)
		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 8 {
		// resume
		op.Resume, fieldSize, err = DecodeBooleanField(remainingBytes)

		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 9 {
		// aborted
		op.Aborted, fieldSize, err = DecodeBooleanField(remainingBytes)
		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if listCount > 10 {
		// batchable
		op.Batchable, fieldSize, err = DecodeBooleanField(remainingBytes)
		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if len(remainingBytes) > 0 {
		err = errors.New("Transfer Performative: There should not be any bytes left")
	}

	return
}

func (b *PerformativeTransfer) Stringify() string {
	return "Stringify: Performative Transfer"
}
