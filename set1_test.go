package cryptopals

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHexToBase64(t *testing.T) {
	var hexStr = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	var base64Str, err = HexToBase64(hexStr)
	assert.Nil(t, err, "Error Decoding Hex/Encoding Base64")
	assert.Equal(t, base64Str, "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t", "Incorrect Base64 Output")

	base64Str, err = HexToBase64("")
	assert.Nil(t, err, "Error Decoding Hex/Encoding Base64 Empty String")
	assert.Equal(t, base64Str, "", "Fails with empty string")

	base64Str, err = HexToBase64("BAD_INPUT")
	assert.NotNil(t, err, "Incorrect error handling")
}

func TestXorArrays(t *testing.T) {
	var empty []byte
	var zero = []byte{0}
	var one = []byte{1}
	var result []byte
	var err error

	result, err = XorArrays(one, one)
	assert.Nil(t, err)
	assert.Equal(t, result, zero, "{1} XOR {1} != {0}")

	var x, _ = hex.DecodeString("1c0111001f010100061a024b53535009181c")
	var y, _ = hex.DecodeString("686974207468652062756c6c277320657965")
	result, err = XorArrays(x, y)
	assert.Nil(t, err)
	assert.Equal(t, hex.EncodeToString(result), "746865206b696420646f6e277420706c6179")

	result, err = XorArrays(empty, empty)
	assert.Nil(t, err)
	assert.Equal(t, result, empty, "Fails with empty array")

	result, err = XorArrays(empty, zero)
	assert.NotNil(t, err, "Failed to reject different sized arrays")
}
