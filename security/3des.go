package security

import (
	"crypto/des"
)

// TripleDesEcbEncrypt returns the encrypt string of the plantext by ECB 3DES.
// The input key must bigger or equal to 8 bytes and must be a multiple of 8.
// According by https://en.wikipedia.org/wiki/Triple_DES to implement.
func TripleDesEcbEncrypt(src, key []byte) ([]byte, error) {
	tripleDESKey := make([]byte, 0, 24)

	if len(key) == 8 {
		tripleDESKey = append(tripleDESKey, key[:8]...)
		tripleDESKey = append(tripleDESKey, key[:8]...)
		tripleDESKey = append(tripleDESKey, key[:8]...)
	} else if len(key) == 16 {
		tripleDESKey = append(tripleDESKey, key[:16]...)
		tripleDESKey = append(tripleDESKey, key[:8]...)
	} else {
		tripleDESKey = append(tripleDESKey, key[:]...)
	}

	td, err := des.NewTripleDESCipher(tripleDESKey)
	if err != nil {
		panic(err)
	}

	mod := len(src) % td.BlockSize()
	v := td.BlockSize() - mod

	data := []byte(src)
	for i := 0; i < v; i++ {
		data = append(data, byte(v))
	}

	n := len(data) / td.BlockSize()
	var rb []byte
	for i := 0; i < n; i++ {
		dst := make([]byte, td.BlockSize())
		td.Encrypt(dst, data[i*8:(i+1)*8])
		rb = append(rb, dst[:]...)
	}
	return rb, nil
}

// TripleDesEcbDecrypt returns the decrypt string of the crypted string by ECB 3DES.
// The input key must bigger or equal to 8 bytes.
// According by https://en.wikipedia.org/wiki/Triple_DES to implement.
func TripleDesEcbDecrypt(src, key []byte) ([]byte, error) {
	tripleDESKey := make([]byte, 0, 24)

	if len(key) == 8 {
		tripleDESKey = append(tripleDESKey, key[:8]...)
		tripleDESKey = append(tripleDESKey, key[:8]...)
		tripleDESKey = append(tripleDESKey, key[:8]...)
	} else if len(key) == 16 {
		tripleDESKey = append(tripleDESKey, key[:16]...)
		tripleDESKey = append(tripleDESKey, key[:8]...)
	} else {
		tripleDESKey = append(tripleDESKey, key[:]...)
	}

	td, err := des.NewTripleDESCipher(tripleDESKey)
	if err != nil {
		panic(err)
	}

	n := len(src) / td.BlockSize()
	var rb []byte
	for i := 0; i < n; i++ {
		dst := make([]byte, td.BlockSize())
		td.Decrypt(dst, src[i*8:(i+1)*8])
		rb = append(rb, dst[:]...)
	}

	lastValue := int(rb[len(rb)-1])
	return rb[0 : len(rb)-lastValue], nil
}
