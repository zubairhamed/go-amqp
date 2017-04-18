package performatives

import (
	"github.com/zubairhamed/go-amqp/types"
)

type Performative interface {
	types.AMQPType
}

type BasePerformative struct {
	types.BaseAMQPType
}

