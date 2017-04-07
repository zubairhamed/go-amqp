package amqp

import (
	"testing"
	"github.com/zubairhamed/go-amqp/frames/performatives"
	"github.com/stretchr/testify/assert"
	"github.com/zubairhamed/go-amqp/frames"
)

func TestOpenPerformative(t *testing.T) {
	sp := performatives.NewOpenPerformative("SendContainer", "localhost")

	b, err := sp.Encode()
	assert.Nil(t, err)

	rp := performatives.NewOpenPerformative("RcvContainer", "localhost")

	ef := frames.EncodeFrame(b)

	err = rp.Decode(ef)
	assert.Nil(t, err)

	assert.Equal(t, "SendContainer", rp.ContainerId.Value())
}