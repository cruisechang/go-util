package security

import (
	"bytes"
	"errors"
	"strconv"
)

const (
	// TripleDESKeyingOption1 means all three keys are independent
	TripleDESKeyingOption1 = 24
	// TripleDESKeyingOption2 means K1 and K2 are independent, and K3 = K1
	TripleDESKeyingOption2 = 16
	// TripleDESKeyingOption3 means All three keys are identical, i.e. K1 = K2 = K3
	TripleDESKeyingOption3 = 8
)

var (
	// ErrDivisionByZero returns because the input may cause division by zero
	ErrDivisionByZero = errors.New("division by zero")
	// ErrKeyLength returned because input key length is not qualified
	ErrKeyLength = errors.New("key length error, must be 16 or 24")
)

// KeyTo24Padding returns 24 bytes length of key.
func KeyTo24Padding(key []byte) ([]byte, error) {
	if len(key) == TripleDESKeyingOption1 {
		return []byte(key), nil
	}
	padKey := make([]byte, 0, TripleDESKeyingOption1)
	if len(key) == TripleDESKeyingOption2 {
		padKey = append(padKey, key[:TripleDESKeyingOption2]...)
		padKey = append(padKey, key[:TripleDESKeyingOption3]...)
	} else if len(key) == TripleDESKeyingOption3 {
		padKey = append(padKey, key[:]...)
		padKey = append(padKey, key[:]...)
		padKey = append(padKey, key[:]...)
	} else {
		return nil, ErrKeyLength
	}
	return padKey, nil
}

//填充主要有三种模式：
//ZeroPadding，数据长度不对齐时使用 0 填充，否则不填充。
//PKCS7Padding，假设数据长度需要填充 n(n>0) 个字节才对齐，那么填充n个字节，每个字节都是 n ;如果数据本身就已经对齐了，则填充一块长度为块大小的数据，每个字节都是块大小。
//PKCS5Padding，PKCS7Padding 的子集，block size大小固定为8字节。


func PKCS5Padding(data []byte, blockSize int) ([]byte, error) {
	if blockSize == 0 {
		return nil, ErrDivisionByZero
	}
	// if data is multiple to blockSize, means we do not need padding
	mod := len(data) % blockSize
	padLeng := blockSize - mod
	padCode := []byte(strconv.Itoa(padLeng))
	padText := bytes.Repeat(padCode, padLeng)
	return append(data, padText...), nil
}

// PKCS5UnPadding returns a list of bytes which has removed padding bytes in it.
func PKCS5UnPadding(data []byte) ([]byte, error) {
	padText := data[len(data)-1]
	unPadding, err := strconv.Atoi(string(padText))
	if err != nil {
		return nil, err
	}
	ret := data[:(len(data) - unPadding)]
	if string(ret) == "" {
		return data, nil
	}
	return ret, nil
}

//AESPKCS5Padding PKCS5padding for AES
func AESPKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//AESPKCS5UnPadding PKCS5UnPadding for AES
func AESPKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// remove the last byte
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// AESPKCS7Padding right-pads the given byte slice with 1 to n bytes, where
// n is the block size. The size of the result is x times n, where x
// is at least 1.
func PKCS7Padding(b []byte, blocksize int) ([]byte, error) {
	if blocksize <= 0 {
		return nil, errors.New("invalid blocksize")
	}
	if b == nil || len(b) == 0 {
		return nil, errors.New("invalid PKCS7 data (empty or not padded)")
	}
	n := blocksize - (len(b) % blocksize)
	pb := make([]byte, len(b)+n)
	copy(pb, b)
	copy(pb[len(b):], bytes.Repeat([]byte{byte(n)}, n))
	return pb, nil
}

// AESPKCS7UnPadding validates and unpads data from the given bytes slice.
// The returned value will be 1 to n bytes smaller depending on the
// amount of padding, where n is the block size.
func PKCS7UnPadding(b []byte) ([]byte, error) {
	//if blocksize <= 0 {
	//	return nil, errors.New("invalid blocksize")
	//}
	//if b == nil || len(b) == 0 {
	//	return nil, errors.New("invalid PKCS7 data (empty or not padded)")
	//}
	//if len(b)%blocksize != 0 {
	//	return nil, errors.New("invalid padding on input")
	//}
	c := b[len(b)-1]
	n := int(c)
	if n == 0 || n > len(b) {
		return nil, errors.New("invalid padding on input")
	}
	for i := 0; i < n; i++ {
		if b[len(b)-n+i] != c {
			return nil, errors.New("invalid padding on input")
		}
	}
	return b[:len(b)-n], nil
}


//zeroPadding
func ZerosPadding(src []byte, blockSize int) []byte {
	paddingCount := blockSize - len(src)%blockSize
	if paddingCount == 0 {
		return src
	} else {
		return append(src, bytes.Repeat([]byte{byte(0)}, paddingCount)...)
	}
}

func ZerosUnPadding(src []byte) []byte {
	for i := len(src) - 1; ; i-- {
		if src[i] != 0 {
			return src[:i+1]
		}
	}
	return nil
}