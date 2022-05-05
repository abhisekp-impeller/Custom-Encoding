package custom_encoding

import (
	"fmt"
)

func Encode(charset string, num uint64) string {
	base := uint64(len(charset))
	encoded := ""

	for num != 0 {
		r := num % base
		fmt.Printf("r = %d\n", r)
		encoded += string(charset[r])
		num /= base
	}

	return encoded
}

func Decode(charset, encoded string) string {
	base := uint64(len(charset))
	decoded := ""

	base = base

	return decoded
}
