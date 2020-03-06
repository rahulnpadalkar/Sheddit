package utils

import "encoding/binary"

func ConvertIntToByte(number int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(number))
	return b
}
