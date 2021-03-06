package amqp

import (
	. "github.com/zubairhamed/go-amqp/frames/performatives"
	. "github.com/zubairhamed/go-amqp/types"
	"log"
	"github.com/zubairhamed/go-amqp/util"
)

type Client struct {
	name           string
	ch             chan *Event
	conn           *Connection
	connectionInfo *ConnectInfo
	role           RoleType
}

func (c *Client) Dial(connInfo *ConnectInfo) (err error) {

	c.conn = NewConnection(connInfo)
	c.connectionInfo = connInfo
	conn := c.conn

	if !conn.connected {
		err = conn.doConnect(c.dispatchPerformative, c.name)
		if err != nil {
			return
		}
	}
	return nil
}

func (r *Client) Close() {
	log.Println("Client:Close")
}

func (c *Client) dispatchPerformative(b []byte) {
	perfByte := Type(b[2])

	switch perfByte {
	case TYPE_PERFORMATIVE_OPEN:
		LogIn("OPEN", c.name)
		perf, err := DecodeOpenPerformative(b)
		if err != nil {
			log.Println(err.Error())
			return
		}
		c.handlePerformativeOpen(perf)

	case TYPE_PERFORMATIVE_BEGIN:
		LogIn("BEGIN", c.name)

		perf, err := DecodeBeginPerformative(b)
		if err != nil {
			log.Println(err.Error())
			return
		}
		c.handleBeginPerformative(perf)

	case TYPE_PERFORMATIVE_ATTACH:
		LogIn("ATTACH", c.name)
		log.Println(string(perfByte))
		perf, err := DecodeAttachPerformative(b)
		if err != nil {
			log.Println(err.Error())
			return
		}
		c.handlePerformativeAttach(perf)

	case TYPE_PERFORMATIVE_DETACH:
		LogIn("DETACH", c.name)
		perf, err := DecodeDetachPerformative(b)
		if err != nil {
			log.Println(err.Error())
			return
		}
		c.handlePerformativeDetach(perf)

	case TYPE_PERFORMATIVE_CLOSE:
		perf, err := DecodeClosePerformative(b)
		if err != nil {
			log.Println(err.Error())
			return
		}
		c.handleClosePerformative(perf)

	case TYPE_PERFORMATIVE_FLOW:
		LogIn("FLOW", c.name)
		perf, err := DecodeFlowPerformative(b)
		if err != nil {
			log.Println(err.Error())
			return
		}
		c.handlePerformativeFlow(perf)

	case TYPE_PERFORMATIVE_TRANSFER:
		LogIn("TRANSFER", c.name)
		perf, err := DecodeTransferPerformative(b)
		if err != nil {
			log.Println(err.Error())
			return
		}
		c.handleTransferPerformative(perf)

	case TYPE_PERFORMATIVE_DISPOSITION:
		LogIn("DISPOSITION", c.name)
		perf, err := DecodeDispositionPerformative(b)
		if err != nil {
			log.Println(err.Error())
			return
		}
		c.handlePerformativeDisposition(perf)

	case TYPE_PERFORMATIVE_END:
		LogIn("END", c.name)
		perf, err := DecodeEndPerformative(b)
		if err != nil {
			log.Println(err.Error())
			return
		}
		c.handlePerformativeEnd(perf)
	}
}

func (c *Client) handlePerformativeOpen(p *PerformativeOpen) {
	beginPerformative := NewBeginPerformative()
	beginPerformative.NextOutgoingId = NewTransferNumber(4294967293)
	beginPerformative.IncomingWindow = NewUInt(2048)
	beginPerformative.OutgoingWindow = NewUInt(2048)
	beginPerformative.HandleMax = NewHandle(7)

	LogOut("BEGIN", c.name)
	c.conn.SendPerformative(beginPerformative)
}

func (c *Client) handleBeginPerformative(p *PerformativeBegin) {

	attach := NewAttachPerformative()

	attach.Name = NewString(c.name)
	attach.SenderSettleMode = nil
	attach.ReceiverSettleMode = nil

	if c.role == ROLE_RECEIVER {
		attach.Handle = NewHandle(1)
		attach.Role = NewRole(false)


		// TODO: Map qith Address
		attach.Source = NewFields(map[string]AMQPType{
			"Address": NewString("my_queue "),
		})
	} else if c.role == ROLE_SENDER {
		attach.Handle = NewHandle(0)
		attach.Role = NewRole(true)
		attach.Target = NewFields(map[string]AMQPType{
			// "Address": NewString(c.connectionInfo.nodeAddress),
			"Address": NewString("my_queue "),
		})
		attach.InitialDeliveryCount = NewSequenceNumber(0)

	} else {
		log.Println("ERROR: Unknown Role Type")
	}

	attach.InitialDeliveryCount = NewSequenceNumber(0)

	e, _, _ := attach.Encode()
	log.Println("Attach Sending ", util.ToHex(e))

	 LogOut("ATTACH", c.name)
	c.conn.SendPerformative(attach)
}

func (c *Client) handlePerformativeAttach(p *PerformativeAttach) {
	c.conn.beginWg.Done()

	log.Println("Describing Attach")
	DescribeType(p)
	c.ch <- NewEvent(p, EVENT_MSG_ATTACH)
}

func (c *Client) handlePerformativeDetach(p *PerformativeDetach) {
	c.ch <- NewEvent(p, EVENT_MSG_DETACH)
}

func (c *Client) handleClosePerformative(p *PerformativeClose) {
	LogIn("CLOSE", c.name)
	log.Println("Close", p.Error.Description.Stringify())
	c.ch <- NewEvent(p, EVENT_MSG_CLOSE)

}

func (c *Client) handlePerformativeFlow(p *PerformativeFlow) {
	c.ch <- NewEvent(p, EVENT_MSG_FLOW)
}

func (c *Client) handleTransferPerformative(p *PerformativeTransfer) {
	c.ch <- NewEvent(p, EVENT_MSG_TRANSFER)
}

func (c *Client) handlePerformativeDisposition(p *PerformativeDisposition) {
	c.ch <- NewEvent(p, EVENT_MSG_DISPOSITION)
}

func (c *Client) handlePerformativeEnd(p *PerformativeEnd) {
	c.ch <- NewEvent(p, EVENT_MSG_END)
}
