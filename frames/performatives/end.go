package performatives

import (
	"errors"
	. "github.com/zubairhamed/go-amqp/types"
)

func NewEndPerformative() *PerformativeEnd {
	return &PerformativeEnd{}
}

/*
<type name="end" class="composite" source="list" provides="frame">
    <descriptor name="amqp:end:list" code="0x00000000:0x00000017"/>
    <field name="error" type="error"/>
</type>
*/

type PerformativeEnd struct {
	BasePerformative
	Error *Error
}

func (b *PerformativeEnd) GetType() Type {
	return TYPE_PERFORMATIVE_END
}

func (p *PerformativeEnd) Encode() (enc []byte, l uint, err error) {
	return
}

func (b *PerformativeEnd) Stringify() string {
	return "Stringify: Performative End"
}

func DecodeEndPerformative(b []byte) (op *PerformativeEnd, err error) {
	op = NewEndPerformative()
	frameData, listCount, err := HandleBasePerformative(b, TYPE_PERFORMATIVE_END)
	if err != nil {
		return
	}

	remainingBytes := frameData
	var fieldSize uint

	if listCount > 0 {
		// error
		op.Error, fieldSize, err = DecodeErrorField(remainingBytes)
		if err != nil {
			return
		}
		remainingBytes = remainingBytes[fieldSize:]
	}

	if len(remainingBytes) > 0 {
		err = errors.New("End Performative: There should not be any bytes left")
	}

	return
}
