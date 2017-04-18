package types

func NewTransferNumber(v uint32) *TransferNumber {
	return &TransferNumber{
		SequenceNumber: NewSequenceNumber(v),
	}
}

type TransferNumber struct {
	*SequenceNumber
}

func DecodeTransferNumberField(v []byte) (val *TransferNumber, fieldLength uint, err error) {
	iVal, fieldLength, err := DecodeSequenceNumber(v)

	val = &TransferNumber{
		SequenceNumber: iVal,
	}
	return
}

