package types

func NewMessageFormat(v uint32) *MessageFormat {
	return &MessageFormat{
		UInt: NewUInt(v),
	}
}

type MessageFormat struct {
	*UInt
}
