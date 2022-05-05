package main

import (
	"fmt"
	. "github.com/abhisekp-impeller/Custom-Encoding/custom_encoding"
)

func main() {
	charset := "123456789ABCDEFGHJKLMNPQRSTUVWXYZ"

	fmt.Printf("Base%d Encoding\n\n", len(charset))

	num := uint64(90663368909653968)
	num = uint64(33)

	ce := (&CustomEncoding{}).SetCharset(charset)
	encoded := ce.Encode(num)
	fmt.Printf("Encoded '%d' as '%s'\n", num, encoded)

	decoded := ce.Decode(encoded)
	fmt.Printf("Decoded '%s' as '%s'\n", encoded, decoded)
}
