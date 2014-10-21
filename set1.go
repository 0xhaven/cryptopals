package cryptopals

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
)

func HexToBase64(hexStr string) (string, error) {
	var value, err = hex.DecodeString(hexStr)
	if err != nil {
		return "", err
	}
	var base64Str = base64.StdEncoding.EncodeToString(value)
	return base64Str, nil
}

func XorArrays(a, b []byte) ([]byte, error) {
	var result []byte

	if len(a) != len(b) {
		return nil, errors.New("Mismatched array sizes")
	}
	if len(a) == 0 {
		var empty []byte
		return empty, nil
	}
	result = make([]byte, len(a), len(a))
	for i := 0; i < len(a); i++ {
		result[i] = a[i] ^ b[i]
	}
	return result, nil
}
