package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

// GenerateAES256RandomBites returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateAES256RandomBytes() ([]byte, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GenerateAES256RandomBase64String returns a base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateAES256RandomBase64String() (string, error) {
	b, err := GenerateAES256RandomBytes()
	return base64.StdEncoding.EncodeToString(b), err
}

// AESCfbEncrypt returns the encrypt string of the crypted string by CFB AES.
// According by https://en.wikipedia.org/wiki/Advanced_Encryption_Standard
func AESCfbEncrypt(text, key []byte) (ciphertext []byte, err error) {
	var block cipher.Block
	if block, err = aes.NewCipher(key); err != nil {
		return
	}
	ciphertext = make([]byte, aes.BlockSize+len(text))
	iv := ciphertext[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], text)
	return
}

// AESCfbDecrypt returns the decrypt string of the crypted string by CFB AES.
// According by https://en.wikipedia.org/wiki/Advanced_Encryption_Standard
func AESCfbDecrypt(ciphertext, key []byte) (plaintext []byte, err error) {
	var block cipher.Block
	if block, err = aes.NewCipher(key); err != nil {
		return
	}
	if len(ciphertext) < aes.BlockSize {
		err = errors.New("ciphertext too short")
		return
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(ciphertext, ciphertext)
	plaintext = ciphertext
	return
}

//AESCbcPkcs7PaddingEncrypt AES encrypttion with CBC and PKCS7Padding
func AESCbcPkcs7PaddingEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	iv := []byte("0000000000000000")
	blockSize := block.BlockSize()
	origData, _ = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

//AESCbcPkcs7PaddingDecrypt AES decrypttion with CBC and PKCS7Padding
// Key length: 16, 24, 32 bytes => AES-128, AES-192, AES-256
func AESCbcPkcs7PaddingDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	iv := []byte("0000000000000000")
	//blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	return PKCS7UnPadding(origData)
}

