package security

import (
	"errors"
)

// CaesarEncrypt but it can handles numbers and any ASCII codes.
// recording by: https://en.wikipedia.org/wiki/Caesar_cipher
func CaesarEncrypt(n int, data []byte) ([]byte, error) {
	if n < -48 || n > 5 {
		return nil, errors.New("out of ASCII range")
	}
	out := make([]byte, len(data))
	for i, v := range data {
		out[i] = byte(int(v) + n)
	}
	return out, nil
}

// CaesarDecrypt but it can handles numbers and any ASCII codes.
// recording by: https://en.wikipedia.org/wiki/Caesar_cipher
func CaesarDecrypt(n int, data []byte) ([]byte, error) {
	if n < -48 || n > 5 {
		return nil, errors.New("out of ASCII range")
	}
	out := make([]byte, len(data))
	for i, v := range data {
		out[i] = byte(int(v) - n)
	}
	return out, nil
}
