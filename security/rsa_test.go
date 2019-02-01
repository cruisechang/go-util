package security

import (
	"bytes"
	"testing"
)

const (
	privateKeyFilepath = "rsa_private.pem"
	publicKeyFilepath  = "public.pem"
)

func TestRSA(t *testing.T) {
	err := GenerateRSAKeyPairToFile(privateKeyFilepath, publicKeyFilepath, 2048)
	if err != nil {
		t.Fatalf("TestRSA: %v", err)
	}
	expect := []byte("hongkong3345678")
	rsa, err := NewRSA("sha256")
	if err != nil {
		t.Fatalf("TestRSA: %v", err)
	}
	pub, err := ReadRSAPublicKeyFromFile(publicKeyFilepath)
	if err != nil {
		t.Fatalf("TestRSA: %v", err)
	}
	encrypted, err := rsa.Encrypt(pub, expect)
	if err != nil {
		t.Fatalf("TestRSA: %v", err)
	}
	actual, err := rsa.DecryptByReadRSAPrivateKeyFromFile(privateKeyFilepath, encrypted)
	if err != nil {
		t.Fatalf("TestRSA: %v", err)
	}
	if !bytes.Equal(expect, actual) {
		t.Errorf("TestRSA: expect->%v, actual->%v", expect, actual)
	}
}
