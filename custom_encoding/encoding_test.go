package custom_encoding

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestCustomEncoding(t *testing.T) {
  t.Run("Encode to Base33", func (t *testing.T) {
    charset := "123456789ABCDEFGHJKLMNPQRSTUVWXYZ"

    ce := (&CustomEncoding{}).SetCharset(charset)
    actual := ce.Encode(uint64(8))
    expected := "9"

    assert.Equal(t, expected, actual)
  })
} 