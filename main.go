package main

import (
	"fmt"
	. "main/custom_encoding"
)

func main() {
	charset := "123456789ABCDEFGHJKLMNPQRSTUVWXYZ"

	fmt.Printf("Base%d Encoding\n\n", len(charset))

	num := uint64(83)

	ce := (&CustomEncoding{}).SetCharset(charset)
	encoded := ce.Encode(num)
	fmt.Printf("Encoded '%d' as '%s'\n", num, encoded)

	decoded := ce.Decode(encoded)
	fmt.Printf("Encoded '%s' as '%s'\n", encoded, decoded)
}
