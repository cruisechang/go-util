package security

import (
	"fmt"
	"encoding/base64"
	"testing"
)

func TestDes(t *testing.T) {
	key := []byte("12345678")

	ori:=[]byte("ABCDEFG")

	//encrypt
	result, err := DesEncrypt(ori, key)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	fmt.Println(base64.StdEncoding.EncodeToString(result))

	//decrypt
	origData, err := DesDecrypt(result, key)
	if err != nil {
		panic(err)
	}

	if string(ori)!=string(origData){
		t.Errorf("des error got=%s,want=%s",string(ori),string(origData))
	}
	//fmt.Printf("origin=%s\n",string(origData))
}

