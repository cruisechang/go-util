package security

import (
	"bytes"
	"testing"
)

func TestBlowFish(t *testing.T) {
	plaintext := []byte("this is the plaintext string")
	secretkey := []byte("1234567890abcdefghijklmnopqrstuvwxyz")
	encryptedtext := BlowfishEncrypt(plaintext, secretkey)
	decryptedtext := BlowfishDecrypt(encryptedtext, secretkey)

	if !bytes.Equal(plaintext, decryptedtext) {
		t.Errorf("TestBlowFish fail oriStr:%s decStr:%s", decryptedtext, decryptedtext)
	}
}
