package performatives

import (
	"errors"
	"fmt"
	. "github.com/zubairhamed/go-amqp/types"
)

type Performative interface {
	AMQPType
}

type BasePerformative struct {
	BaseAMQPType
}

func HandleBasePerformative(b []byte, expectedPerformative Type) (frameData []byte, listCount int, err error) {
	frameBytes := b

	if Type(frameBytes[0]) != TYPE_CONSTRUCTOR {
		err = errors.New("Malformed or unexpected frame. Expecting constructor.")
		return
	}

	if Type(frameBytes[1]) != TYPE_ULONG_SMALL {
		err = errors.New("Malformed or unexpected frame. Expecting small ulong type")
		return
	}

	if Type(frameBytes[2]) != expectedPerformative {
		err = errors.New(fmt.Sprint("Malformed or unexpected frame. Expecting Performative: ", expectedPerformative))
		return
	}

	if Type(frameBytes[3]) != TYPE_LIST_8 {
		err = errors.New("Malformed or unexpected frame. Expecting list 8")
		return
	}

	listBytes := int(frameBytes[4])
	listCount = int(frameBytes[5])

	switch {
	case expectedPerformative == TYPE_PERFORMATIVE_OPEN && listCount > 10:
		err = errors.New("Open Performative should contain 10 or less fields.")
		return

	case expectedPerformative == TYPE_PERFORMATIVE_BEGIN && listCount > 8:
		err = errors.New("Begin Performative should contain 8 or less fields.")
		return

	case expectedPerformative == TYPE_PERFORMATIVE_ATTACH && listCount > 14:
		err = errors.New("Attach Performative should contain 14 or less fields.")
		return

	case expectedPerformative == TYPE_PERFORMATIVE_FLOW && listCount > 11:
		err = errors.New("Flow Performative should contain 11 or less fields.")
		return

	case expectedPerformative == TYPE_PERFORMATIVE_TRANSFER && listCount > 11:
		err = errors.New("Transfer Performative should contain 11 or less fields.")
		return

	case expectedPerformative == TYPE_PERFORMATIVE_DISPOSITION && listCount > 6:
		err = errors.New("Disposition Performative should contain 6 or less fields.")
		return

	case expectedPerformative == TYPE_PERFORMATIVE_DETACH && listCount > 3:
		err = errors.New("Detach Performative should contain 3 or less fields.")
		return

	case expectedPerformative == TYPE_PERFORMATIVE_END && listCount > 1:
		err = errors.New("Detach Performative should contain 1 or less fields.")
		return

	case expectedPerformative == TYPE_PERFORMATIVE_CLOSE && listCount > 1:
		err = errors.New("Invalid list count. Expecting 1 or less.")
		return
	}

	frameData = frameBytes[6:]

	if len(frameData)+1 != listBytes {
		err = errors.New("Malformed or unexpected frame. list size not equal or expected")
		return
	}

	return
}
