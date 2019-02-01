package security

import (
	"bytes"
	"testing"
)

func TestTripleDesEcb(t *testing.T) {
	expect := []byte("This is a TEST!!")

	key24 := []byte("a1c2e3g4i5k6m7o8q9s0u1w2")
	actual, _ := TripleDesEcbEncrypt(expect, key24)
	actual, _ = TripleDesEcbDecrypt(actual, key24)
	if !bytes.Equal(actual, expect) {
		t.Errorf("TestTripleDesEcb 24bits Key: expect->%q, actual->%q", expect, actual)
	}

	key16 := []byte("a1c2e3g4i5k6m7o8")
	actual, _ = TripleDesEcbEncrypt(expect, key16)
	actual, _ = TripleDesEcbDecrypt(actual, key16)
	if !bytes.Equal(actual, expect) {
		t.Errorf("TestTripleDesEcb 16bits Key: expect->%q, actual->%q", expect, actual)
	}

	key8 := []byte("a1c2e3g4")
	actual, _ = TripleDesEcbEncrypt(expect, key8)
	actual, _ = TripleDesEcbDecrypt(actual, key8)
	if !bytes.Equal(actual, expect) {
		t.Errorf("TestTripleDesEcb 8bits Key: expect->%q, actual->%q", expect, actual)
	}
}
