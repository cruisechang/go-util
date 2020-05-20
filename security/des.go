package security

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
)

func DESPKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
func DESPKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//with zero IV
func DESEncrypt(src, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)

	if err != nil {
		return nil, err
	}

	bs := block.BlockSize()
	data := DESPKCS5Padding(src, bs)
	//blockMode := cipher.NewCBCEncrypter(block, []byte{0, 0, 0, 0, 0, 0, 0, 0})
	blockMode := cipher.NewCBCEncrypter(block, []byte("00000000"))
	cryted := make([]byte, len(data))
	blockMode.CryptBlocks(cryted, data)

	return cryted, nil
}

//with zero IV
func DESDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//blockMode := cipher.NewCBCDecrypter(block, []byte{0, 0, 0, 0, 0, 0, 0, 0})
	blockMode := cipher.NewCBCDecrypter(block, []byte("00000000"))
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = DESPKCS5UnPadding(origData)
	return origData, nil
}




