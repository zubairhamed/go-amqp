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

	transfer.Handle = types.NewHandle(0)
	transfer.DeliveryId = types.NewDeliveryNumber(0)
	transfer.DeliveryTag = types.NewDeliveryTag([]byte{0, 0, 0, 0, 0})
	transfer.MessageFormat = types.NewMessageFormat(0)
	transfer.Settled = types.NewBoolean(false)
	transfer.More = nil
	transfer.ReceiverSettleMode = nil
	transfer.State = nil
	transfer.Resume = nil
	transfer.Aborted = nil
	transfer.Batchable = types.NewBoolean(true)
	transfer.Value = types.NewAmqpValue(types.NewString("HelloWorld"))


	LogOut("TRANSFER", "sender")
	s.conn.SendPerformative(transfer)
}
