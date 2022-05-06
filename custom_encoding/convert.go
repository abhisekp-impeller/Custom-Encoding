package custom_encoding

import (
	// "fmt"
	"math"
	"strconv"
	"strings"
)

func Encode(charset string, num uint64) string {
	base := uint64(len(charset))
	var encoded strings.Builder

  // origNum = num

  if num == 0 {
    encoded.WriteByte(charset[0])
  }

	for num != 0 {
		r := num % base
		// fmt.Printf("r = %d\n", r)
		encoded.WriteByte(charset[r])
		num /= base
	}

	return encoded.String()
}

func Decode(charset string, encoded string) string {
	base := uint64(len(charset))
	decoded := ""

	charsetMap := map[rune]int{}
	for i, c := range charset {
		charsetMap[c] = i
	}

	num := uint64(0)
	for _, c := range encoded {
    charIdx := charsetMap[c]
		num += uint64(charIdx * int(math.Pow(float64(base), float64(charIdx))))
	}

	decoded = strconv.FormatUint(num, 10)

	return decoded
}
