package custom_encoding

import (
	"fmt"
	"math"
)

func Convert(charset string, num uint64) string {
  encoded := ""
numBits :=  math.Floor(math.Log2(float64(num))) + 1

  fmt.Printf("No.of Bits: %d\n", int(numBits))

  bin := make([]byte, 0, int(numBits) / 8 + 1)
  fmt.Printf("Number of bytes: %d\n", cap(bin))

  for ;num != 0;  {
    r := num % uint64(len(charset))
    fmt.Printf("r = %d\n", r)
    num-=r
  }
  
  return encoded
}