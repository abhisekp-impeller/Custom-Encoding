package custom_encoding

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCustomEncoding(t *testing.T) {
	t.Run("Encode to Base33", func(t *testing.T) {
		charset := "123456789ABCDEFGHJKLMNPQRSTUVWXYZ"

		ce := (&CustomEncoding{}).SetCharset(charset)
		actual := ce.Encode(uint64(8))
		expected := "9"

		assert.Equal(t, expected, actual)
	})

	t.Run("Decode from Base33", func(t *testing.T) {
		charset := "123456789ABCDEFGHJKLMNPQRSTUVWXYZ"

		ce := (&CustomEncoding{}).SetCharset(charset)
		actual := ce.Decode("9")
		expected := "9"

		assert.Equal(t, expected, actual)
	})
}
