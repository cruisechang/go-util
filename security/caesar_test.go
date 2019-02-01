package security

import (
	"bytes"
	"testing"
)

func TestCaesar(t *testing.T) {
	expect := []byte("This is a TEST!!")
	actual, _ := CaesarEncrypt(5, expect)
	actual, _ = CaesarDecrypt(5, actual)
	if !bytes.Equal(actual, expect) {
		t.Errorf("TestCaesar: expect->%q, actual->%q", expect, actual)
	}
}
