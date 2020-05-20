package security

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"testing"
)

func TestAESCfb(t *testing.T) {
	expect := []byte("This is a TEST!!")
	key := []byte("a1c2e3g4i5k6m7o8q9s0u1w2qes234df")
	actual, _ := AESCfbEncrypt(expect, key)

	t.Log(len(actual))
	t.Log(base64.StdEncoding.EncodeToString(actual))
	t.Log(hex.EncodeToString(actual))

	actual, _ = AESCfbDecrypt(actual, key)
	if !bytes.Equal(actual, expect) {
		t.Errorf("TestAESCfb: expect->%q, actual->%q", expect, actual)
	}
}

func TestAESCBC(t *testing.T) {


	tt := struct {
		Account  string `json:"account"`
		Password string `json:"password"`
	}{
		"123", "123",}

	expect, _ := json.Marshal(tt)

	t.Log(string(expect))

	key := []byte("0123456789abcdef")
	actual, _ := AESCbcPkcs7PaddingEncrypt(expect, key)

	b:=base64.StdEncoding.EncodeToString(actual)
	t.Log(b)

	actual,_=AESCbcPkcs7PaddingDecrypt(actual, key)
	if !bytes.Equal(actual, expect) {
		t.Errorf("TestAECBC: expect->%q, actual->%q", expect, actual)
	}

}


func TestAESECB(t *testing.T) {


	tt := struct {
		Account  string `json:"account"`
		Password string `json:"password"`
	}{
		"123", "123",}

	expect, _ := json.Marshal(tt)

	t.Log(string(expect))

	key := []byte("0123456789abcdef") //16 bytes
	actual, _ := AESEcbPkcs7PaddingEncrypt(expect, key)

	//b:=base64.StdEncoding.EncodeToString(actual)
	//t.Log(b)

	actual,_=AESEcbPkcs7PaddingDecrypt(actual, key)
	if !bytes.Equal(actual, expect) {
		t.Errorf("TestAESECB: expect->%q, actual->%q", expect, actual)
	}

}


/*
func TestGenerateAES256(t *testing.T) {
	expect := []byte("This is a TEST!!")

	keyByte, _ := GenerateAES256RandomBites()
	actual, _ := AESCfbEncrypt(expect, keyByte)
	actual, _ = AESCfbDecrypt(actual, keyByte)
	if !bytes.Equal(actual, expect) {
		t.Errorf("TestGenerateAES256: expect->%q, actual->%q", expect, actual)
	}

	keyString, _ := GenerateAES256RandomBase64String()
	key, _ := base64.StdEncoding.DecodeString(keyString)
	actual, _ = AESCfbEncrypt(expect, key)
	actual, _ = AESCfbDecrypt(actual, key)
	if !bytes.Equal(actual, expect) {
		t.Errorf("TestGenerateAES256: expect->%q, actual->%q", expect, actual)
	}
}
*/

func TestAESCBCPKCS7Padding(t *testing.T) {
	oriStr := "TestAESCBCPKCS7Padding Test!"

	key := []byte("qwertyuiasdfghjk12345678!@#$%^&*")
	result, _ := AESCbcPkcs7PaddingEncrypt([]byte(oriStr), key)

	desData, _ := AESCbcPkcs7PaddingDecrypt(result, key)

	if oriStr != string(desData) {
		t.Errorf("TestAESCBCPKCS5Padding error: ori:%s new:%s", oriStr, string(desData))
	}
}
