package amqp

import (
	"bufio"
	"github.com/zubairhamed/go-amqp/frames/performatives"
	"github.com/zubairhamed/go-amqp/types"
	"log"
	"net"
	"reflect"
	"time"
)

func DescribeField(t types.AMQPType) string {
	val := reflect.ValueOf(t)
	if val.IsNil() {
		return "nil"
	}
	return t.Stringify()
}

func DescribeType(t types.AMQPType) {
	switch t.GetType() {
	case types.TYPE_PERFORMATIVE_OPEN:
		p := t.(*performatives.PerformativeOpen)
		log.Println()
		log.Println("-- OPEN PERFORMATIVE --")
		log.Println("container-id:", DescribeField(p.ContainerId))
		log.Println("hostname:", DescribeField(p.Hostname))
		log.Println("max-frame-size:", DescribeField(p.MaxFrameSize))
		log.Println("channel-max:", DescribeField(p.ChannelMax))
		log.Println("idle-time-out:", DescribeField(p.IdleTimeout))
		log.Println("outgoing-locales:", DescribeField(p.OutgoingLocales))
		log.Println("incoming-locales:", DescribeField(p.IncomingLocales))
		log.Println("offered-capabilities:", DescribeField(p.OfferedCapabilities))
		log.Println("desired-capabilities:", DescribeField(p.DesiredCapabilities))
		log.Println("properties:", DescribeField(p.Properties))
		log.Println()

	case types.TYPE_PERFORMATIVE_BEGIN:
		p := t.(*performatives.PerformativeBegin)
		log.Println()
		log.Println("-- BEGIN PERFORMATIVE --")
		log.Println("remote-channel:", DescribeField(p.RemoteChannel))
		log.Println("next-outgoing-id:", DescribeField(p.NextOutgoingId))
		log.Println("incoming-window:", DescribeField(p.IncomingWindow))
		log.Println("outgoing-window:", DescribeField(p.OutgoingWindow))
		log.Println("handle-max:", DescribeField(p.HandleMax))
		log.Println("offered-capabilities:", DescribeField(p.OfferedCapabilities))
		log.Println("desired-capabilities:", DescribeField(p.DesiredCapabilities))
		log.Println("properties:", DescribeField(p.Properties))
		log.Println()

	case types.TYPE_PERFORMATIVE_ATTACH:
		p := t.(*performatives.PerformativeAttach)
		log.Println()
		log.Println("-- ATTACH PERFORMATIVE --")
		log.Println("name:", DescribeField(p.Name))
		log.Println("handle:", DescribeField(p.Handle))
		log.Println("role:", DescribeField(p.Role))
		log.Println("snd-settle-mode:", DescribeField(p.SenderSettleMode))
		log.Println("rcv-settle-mode:", DescribeField(p.ReceiverSettleMode))
		log.Println("source:", DescribeField(p.Source))
		log.Println("target:", DescribeField(p.Target))
		log.Println("unsettled:", DescribeField(p.Unsettled))
		log.Println("incomplete-unsettled:", DescribeField(p.IncompleteUnsettled))
		log.Println("initial-delivery-count:", DescribeField(p.InitialDeliveryCount))
		log.Println("max-message-size:", DescribeField(p.MaxMessageSize))
		log.Println("offered-capabilities:", DescribeField(p.OfferedCapabilities))
		log.Println("desired-capabilities:", DescribeField(p.DesiredCapabilities))
		log.Println("properties:", DescribeField(p.Properties))
		log.Println()

	case types.TYPE_PERFORMATIVE_CLOSE:
		// p := t.(*performatives.PerformativeClose)
		log.Println()
		log.Println("-- CLOSE PERFORMATIVE --")
		/*
			error
		*/
		log.Println()

	case types.TYPE_PERFORMATIVE_DETACH:
		// p := t.(*performatives.PerformativeDetach)
		log.Println()
		log.Println("-- DETATCH PERFORMATIVE --")
		/*
			handle
			closed
			error
		*/
		log.Println()

	case types.TYPE_PERFORMATIVE_DISPOSITION:
		// p := t.(*performatives.PerformativeDisposition)
		log.Println()
		log.Println("-- DISPOSITION PERFORMATIVE --")
		/*
			role
			first
			last
			settled
			state
			batchable
		*/
		log.Println()

	case types.TYPE_PERFORMATIVE_END:
		// p := t.(*performatives.PerformativeEnd)
		log.Println()
		log.Println("-- END PERFORMATIVE --")
		/*
			name
		*/
		log.Println()

	case types.TYPE_PERFORMATIVE_FLOW:
		// p := t.(*performatives.PerformativeFlow)
		log.Println()
		log.Println("-- FLOW PERFORMATIVE --")
		/*
			next-incoming-id
			incoming-window
			next-outgoing-id
			outgoing-window
			handle
			delivery-count
			link-credit
			available
			drain
			echo
			properties
		*/
		log.Println()

	case types.TYPE_PERFORMATIVE_TRANSFER:
		// p := t.(*performatives.PerformativeTransfer)
		log.Println()
		log.Println("-- TRANSFER PERFORMATIVE --")
		/*
			handle
			delivery-id
			delivery-tag
			message-format
			settled
			more
			rcv-settle-mode
			state
			resume
			aborted
			batchable
		*/
		log.Println()
	}
}

func DescribeTypeValue(t types.AMQPType) {

}

func ReadFromConnection(c net.Conn) ([]byte, error) {
	readBuf := make([]byte, 1500)

	_, err := bufio.NewReader(c).Read(readBuf)

	return readBuf, err
}

func LogIn(perf, name string) {

	log.Println(time.Now(), " - [", name, "] <<", perf)
}

func LogOut(perf, name string) {
	log.Println(time.Now(), " - [", name, "] >>", perf)
}
