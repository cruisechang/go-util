package security

import (
	"crypto/cipher"

	"golang.org/x/crypto/blowfish"
)

var blowfishIV = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

//BlowfishEncrypt blowfish encryption
func BlowfishEncrypt(dataToEncrypt, key []byte) []byte {
	// create the cipher
	ci, err := blowfish.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	s := cipher.NewCFBEncrypter(ci, blowfishIV)
	dst := make([]byte, len(dataToEncrypt))
	s.XORKeyStream(dst, dataToEncrypt)
	return dst
}

//BlowfishDecrypt blowfish decryption
func BlowfishDecrypt(dataToDecrypt, key []byte) []byte {
	// create the cipher
	ci, err := blowfish.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	s := cipher.NewCFBDecrypter(ci, blowfishIV)

	dst := make([]byte, len(dataToDecrypt))
	s.XORKeyStream(dst, dataToDecrypt)
	return dst
}
