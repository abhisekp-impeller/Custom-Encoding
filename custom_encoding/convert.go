package custom_encoding

import (
	// "fmt"
	"math"
	"strconv"
)

func Encode(charset string, num uint64) string {
	base := uint64(len(charset))
	encoded := ""

	for num != 0 {
		r := num % base
		// fmt.Printf("r = %d\n", r)
		encoded += string(charset[r])
		num /= base
	}

	return encoded
}

func Decode(charset string, encoded string) string {
	base := uint64(len(charset))
	decoded := ""

	charsetMap := map[rune]int{}
	for i, c := range charset {
		charsetMap[c] = i
	}

	num := uint64(0)
	for i, c := range encoded {
		num += uint64(charsetMap[c] * int(math.Pow(float64(base), float64(i))))
	}

	decoded = strconv.FormatUint(num, 10)

	return decoded
}
