package cryptopals

import (
	"encoding/base64"
	"encoding/hex"
)

func HexToBase64(hexStr string) (string, error) {
	var value, err = hex.DecodeString(hexStr)
	if err != nil {
		return "", err
	}
	var base64Str = base64.StdEncoding.EncodeToString(value)
	return base64Str, nil
}
