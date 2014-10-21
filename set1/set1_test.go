package cryptopals

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHexToBase64(t *testing.T) {
	var hexStr = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	var base64Str, err = HexToBase64(hexStr)
	assert.Nil(t, err, "Error Decoding Hex/Encoding Base64")
	assert.Equal(t, base64Str, "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t", "Incorrect Base64 Output")
	base64Str, err = HexToBase64("BAD_INPUT")
	assert.NotNil(t, err, "Incorrect error handling")
}
