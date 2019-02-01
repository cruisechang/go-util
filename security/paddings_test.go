package security

import (
	"bytes"
	"testing"
)

func TestPKCS5Padding(t *testing.T) {
	actual, _ := PKCS5Padding([]byte{'1', '2', '3', '4', '5'}, 8)
	expect := []byte{'1', '2', '3', '4', '5', '3', '3', '3'}
	if !bytes.Equal(actual, expect) {
		t.Errorf("TestPKCS5Padding: expect->%q, actual->%q", expect, actual)
	}
	actual, _ = PKCS5Padding([]byte{'1', '2', '3', '4', '5', '6', '7', '8'}, 8)
	expect = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '8', '8', '8', '8', '8', '8', '8', '8'}
	if !bytes.Equal(actual, expect) {
		t.Errorf("TestPKCS5Padding: expect->%q, actual->%q", expect, actual)
	}
}

func TestPKCS5UnPadding(t *testing.T) {
	actual, _ := PKCS5UnPadding([]byte{'1', '2', '3', '4', '5', '3', '3', '3'})
	expect := []byte{'1', '2', '3', '4', '5'}
	if !bytes.Equal(actual, expect) {
		t.Errorf("TestPKCS5UnPadding: expect->%q, actual->%q", expect, actual)
	}
	actual, _ = PKCS5UnPadding([]byte{'1', '2', '3', '4', '5', '6', '7', '8', '8', '8', '8', '8', '8', '8', '8', '8'})
	expect = []byte{'1', '2', '3', '4', '5', '6', '7', '8'}
	if !bytes.Equal(actual, expect) {
		t.Errorf("TestPKCS5UnPadding: expect->%q, actual->%q", expect, actual)
	}
}
