package util

import "encoding/hex"

func ToHex(b []byte) string {
	str := hex.EncodeToString(b)
	idx := 0

	outStr := ""
	for idx < len(str) {
		outStr = outStr + str[idx:idx+2] + " "
		idx += 2
	}
	return outStr
}
