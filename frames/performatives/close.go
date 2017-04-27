package performatives

import (
	"errors"
	. "github.com/zubairhamed/go-amqp/types"
)

func NewClosePerformative() *PerformativeClose {
	return &PerformativeClose{}
}

/*
<type name="close" class="composite" source="list" provides="frame">
    <descriptor name="amqp:close:list" code="0x00000000:0x00000018"/>
    <field name="error" type="error"/>
</type>
*/
type PerformativeClose struct {
	BasePerformative
	Error *Error
}

func (b *PerformativeClose) GetType() Type {
	return TYPE_PERFORMATIVE_CLOSE
}

func (p *PerformativeClose) Encode() (enc []byte, l uint, err error) {
	return
}

func (b *PerformativeClose) Stringify() string {
	return "Stringify: Performative Close"
}

func DecodeClosePerformative(b []byte) (op *PerformativeClose, err error) {
	op = NewClosePerformative()
	frameData, listCount, err := HandleBasePerformative(b, TYPE_PERFORMATIVE_CLOSE)
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
		err = errors.New("Close Performative: There should not be any bytes left")
	}

	return
}
