package security

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"hash"
	"io/ioutil"
)

var (
	// ErrHashTypeNotAllowed represents that input has type is not allowed
	ErrHashTypeNotAllowed = errors.New("Hash algorithm must be one of md5, sha1, sha256, sha512")
	// ErrNotPEMEncoded represents that key are not encoded by PEM
	ErrNotPEMEncoded = errors.New("Not PEM-encoded")
	// ErrUnkownKeyType represents that key file type are not RSA
	ErrUnkownKeyType = errors.New("Unknown key type")
	// ErrNotRSAPublicKey represents that key is not an RSA public key format
	ErrNotRSAPublicKey = errors.New("Not an RSA public format")
)

// RSA is the RSA hash algorithm instance.
type RSA struct {
	hashType   hash.Hash
	hashCrypto crypto.Hash
}

// NewRSA returns the RSA instance and the encrypt, decrypt type will depends on input.
func NewRSA(hashType string) (*RSA, error) {
	var h hash.Hash
	var ha crypto.Hash

	switch hashType {
	case "md5":
		h = md5.New()
		ha = crypto.MD5
	case "sha1":
		h = sha1.New()
		ha = crypto.SHA1
	case "sha256":
		h = sha256.New()
		ha = crypto.SHA256
	case "sha512":
		h = sha512.New()
		ha = crypto.SHA512
	default:
		return nil, ErrHashTypeNotAllowed
	}
	return &RSA{
		hashType:   h,
		hashCrypto: ha,
	}, nil
}

// Encrypt encrypts the plain text by RSA.
func (r *RSA) Encrypt(publicKey *rsa.PublicKey, plaintext []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, publicKey, plaintext)
	// return rsa.EncryptOAEP(r.hashType, rand.Reader, publicKey, plaintext, nil)
}

// DecryptByReadRSAPrivateKeyFromFile decodes the RSA encrypted text by reading the private key from file.
func (r *RSA) DecryptByReadRSAPrivateKeyFromFile(privateKeyFilepath string, ciphertext []byte) ([]byte, error) {
	prikey, err := ReadRSAPrivateKeyFromFile(privateKeyFilepath)
	if err != nil {
		return nil, err
	}
	return r.Decrypt(prikey, ciphertext)
}

// Decrypt decrypts the RSA encrypted text.
func (r *RSA) Decrypt(privateKey *rsa.PrivateKey, ciphertext []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	// return rsa.DecryptOAEP(r.hashType, rand.Reader, privateKey, ciphertext, nil)
}

// ReadRSAPublicKeyFromFile returns the RSA public key from file.
func ReadRSAPublicKeyFromFile(publicKeyFilepath string) (*rsa.PublicKey, error) {
	b, err := ioutil.ReadFile(publicKeyFilepath)
	if err != nil {
		return nil, err
	}
	return ReadRSAPublicKeyFromByte(b)
}

// ReadRSAPublicKeyFromByte returns the RSA public key from byte array.
func ReadRSAPublicKeyFromByte(publicKey []byte) (*rsa.PublicKey, error) {
	blk, _ := pem.Decode(publicKey)
	if blk == nil {
		return nil, ErrNotPEMEncoded
	}
	key, err := x509.ParsePKIXPublicKey(blk.Bytes)
	if err != nil {
		return nil, err
	}
	pub, ok := key.(*rsa.PublicKey)
	if !ok {
		return nil, ErrNotRSAPublicKey
	}
	return pub, nil
}

// ReadRSAPrivateKeyFromFile returns the RSA private key from file.
func ReadRSAPrivateKeyFromFile(privateKeyFilepath string) (*rsa.PrivateKey, error) {
	b, err := ioutil.ReadFile(privateKeyFilepath)
	if err != nil {
		return nil, err
	}
	return ReadRSAPrivateKeyFromByte(b)
}

// ReadRSAPrivateKeyFromByte returns the RSA private key from byte array.
func ReadRSAPrivateKeyFromByte(privateKey []byte) (*rsa.PrivateKey, error) {
	blk, _ := pem.Decode(privateKey)
	if blk == nil {
		return nil, ErrNotPEMEncoded
	}
	if got, want := blk.Type, "RSA PRIVATE KEY"; got != want {
		return nil, ErrUnkownKeyType
	}
	key, err := x509.ParsePKCS1PrivateKey(blk.Bytes)
	if err != nil {
		return nil, err
	}
	return key, nil
}

// GenerateRSAKeyPairToFile creates random RSA private and public key pairs then saves them in files.
func GenerateRSAKeyPairToFile(privateKeyFilepath, publicKeyFilepath string, keySize int) error {
	pri, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		return err
	}
	privateBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(pri),
	}
	err = ioutil.WriteFile(privateKeyFilepath, pem.EncodeToMemory(privateBlock), 0644)
	if err != nil {
		return err
	}

	pub := &pri.PublicKey
	pubDer, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		return err
	}
	publicBlock := &pem.Block{
		Type:    "PUBLIC KEY",
		Headers: nil,
		Bytes:   pubDer,
	}
	err = ioutil.WriteFile(publicKeyFilepath, pem.EncodeToMemory(publicBlock), 0644)
	if err != nil {
		return err
	}
	return nil
}
