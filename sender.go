package amqp

import (
	"github.com/zubairhamed/go-amqp/frames/performatives"
	"github.com/zubairhamed/go-amqp/types"
)

func NewSender(name string, ch chan *Event) *Sender {
	return &Sender{
		Client: Client{
			name: name,
			ch:   ch,
			role: ROLE_SENDER,
		},
	}
}

type Sender struct {
	Client
	//	session *Session
}

func (s *Sender) Send(msg *Message) {
	transfer := performatives.NewTransferPerformative()

	// 0
	transfer.Handle = types.NewHandle(0)

	// 0
	transfer.DeliveryId = types.NewDeliveryNumber(0)

	// 00000000
	transfer.DeliveryTag = types.NewDeliveryTag([]byte{0, 0, 0, 0, 0})

	// 0
	transfer.MessageFormat = types.NewMessageFormat(0)

	// false
	transfer.Settled = types.NewBoolean(false)

	// null
	transfer.More = nil

	// null
	transfer.ReceiverSettleMode = nil

	// ?
	transfer.State = nil

	// nil
	transfer.Resume = nil

	// nil
	transfer.Aborted = nil

	// true
	transfer.Batchable = types.NewBoolean(true)

	// 48 65 6c 6c 6f 20 57 6f 72 6c 64 21
	transfer.Value = types.NewAmqpValue(types.NewString("Hello World!"))

	LogOut("TRANSFER", "sender")

	s.conn.SendPerformative(transfer)
}
