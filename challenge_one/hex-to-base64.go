package challenge_one

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func HexToBase64(hex string) (string, error) {
	b, err := decodeHex(hex)
	if err != nil {
		fmt.Printf("failed to decode hex: %s", err)
		return "", err
	}
	s := encodeBase64(b)
	fmt.Printf("successfully encoded base64: %s", s)
	return s, nil
}

func decodeHex(input string) ([]byte, error) {
	b, err := hex.DecodeString(input)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func encodeBase64(input []byte) string {
	s := base64.RawStdEncoding.EncodeToString(input)
	return s
}
