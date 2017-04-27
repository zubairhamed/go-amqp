package amqp

import "github.com/zubairhamed/go-amqp/frames/performatives"

func NewEvent(p performatives.Performative, t EventType) *Event {
	return &Event{
		Type:         t,
		Performative: p,
	}
}

type Event struct {
	Type         EventType
	Performative performatives.Performative
}
