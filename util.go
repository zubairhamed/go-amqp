package amqp

import (
	"encoding/hex"
	"log"
)

func PrintHex(b []byte) {
	str := hex.EncodeToString(b)
	idx := 0

	outStr := ""
	for idx < len(str) {
		outStr = outStr + str[idx:idx+2] + " "
		idx += 2
	}

	log.Println(outStr)
}
