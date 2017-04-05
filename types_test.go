package amqp

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/zubairhamed/go-amqp/types"
)

func TestEncodeString(t *testing.T) {
	b, l, e := types.EncodeStringField(nil)
	assert.Nil(t, e)
	assert.Equal(t, uint(1), l)
	assert.Equal(t, []byte { byte(types.TYPE_NULL)}, b)

	var s *types.String

	s = types.NewString("test")
	b, l, e = types.EncodeStringField(s)
	assert.Nil(t, e)
	assert.Equal(t, uint(6), l)

	s = types.NewString("Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.")
	b, l, e = types.EncodeStringField(s)
	assert.Nil(t, e)
	assert.Equal(t, uint(577), l)
}