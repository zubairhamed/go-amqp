package amqp

import (
	"github.com/stretchr/testify/assert"
	"github.com/zubairhamed/go-amqp/types"
	"testing"
)

// array
// binary
// boolean
// byte
// char
// decimal
// double
// float
// int
// list
// long
// map
// null
// short

func TestTypeString(t *testing.T) {
	b, l, e := types.EncodeStringField(nil)
	assert.Nil(t, e)
	assert.Equal(t, uint(1), l)
	assert.Equal(t, []byte{byte(types.TYPE_NULL)}, b)

	var s *types.String

	s = types.NewString("test")
	b, l, e = types.EncodeStringField(s)
	assert.Nil(t, e)
	assert.Equal(t, uint(6), l)

	s = types.NewString("Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.")
	b, l, e = types.EncodeStringField(s)
	assert.Nil(t, e)
	assert.Equal(t, uint(579), l)
}

func TestTypeymbol(t *testing.T) {
	b, l, e := types.EncodeSymbolField(nil)
	assert.Nil(t, e)
	assert.Equal(t, uint(1), l)
	assert.Equal(t, []byte{byte(types.TYPE_NULL)}, b)

	var s *types.Symbol

	s = types.NewSymbol("test")
	b, l, e = types.EncodeSymbolField(s)
	assert.Nil(t, e)
	assert.Equal(t, uint(6), l)

	s = types.NewSymbol("Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.")
	b, l, e = types.EncodeSymbolField(s)
	assert.Nil(t, e)
	assert.Equal(t, uint(579), l)
}

// timestamp
// ubyte
func TestTypeUInt(t *testing.T) {
	b, l, e := types.EncodeUIntField(nil)
	assert.Nil(t, e)
	assert.Equal(t, uint(1), l)
	assert.Equal(t, []byte{byte(types.TYPE_UINT_0)}, b)

	var s *types.UInt

	s = types.NewUInt(254)
	b, l, e = types.EncodeUIntField(s)
	assert.Nil(t, e)
	assert.Equal(t, uint(2), l)

	s = types.NewUInt(65535)
	b, l, e = types.EncodeUIntField(s)
	assert.Nil(t, e)
	assert.Equal(t, uint(5), l)
}

func TestEncodeULong(t *testing.T) {
	b, l, e := types.EncodeULongField(nil)
	assert.Nil(t, e)
	assert.Equal(t, uint(1), l)
	assert.Equal(t, []byte{byte(types.TYPE_ULONG_0)}, b)

	var s *types.ULong

	s = types.NewULong(254)
	b, l, e = types.EncodeULongField(s)
	assert.Nil(t, e)
	assert.Equal(t, uint(2), l)

	s = types.NewULong(65535)
	b, l, e = types.EncodeULongField(s)
	assert.Nil(t, e)
	assert.Equal(t, uint(9), l)
}

func TestEncodeUShort(t *testing.T) {
	b, l, e := types.EncodeUShortField(nil)
	assert.Nil(t, e)
	assert.Equal(t, uint(3), l)
	assert.Equal(t, []byte{byte(types.TYPE_USHORT), 0, 0}, b)

	var s *types.UShort

	s = types.NewUShort(254)
	b, l, e = types.EncodeUShortField(s)
	assert.Nil(t, e)
	assert.Equal(t, uint(3), l)

	s = types.NewUShort(65535)
	b, l, e = types.EncodeUShortField(s)
	assert.Nil(t, e)
	assert.Equal(t, uint(3), l)
}

// uuid
