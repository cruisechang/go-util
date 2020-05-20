package security

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestDes(t *testing.T) {
	key := []byte("12345678")

	ori:=[]byte("ABCDEFG")

	//encrypt
	result, err := DESEncrypt(ori, key)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	fmt.Println(base64.StdEncoding.EncodeToString(result))

	//decrypt
	origData, err := DESDecrypt(result, key)
	if err != nil {
		panic(err)
	}

	if string(ori)!=string(origData){
		t.Errorf("des error got=%s,want=%s",string(ori),string(origData))
	}
	//fmt.Printf("origin=%s\n",string(origData))
}

func TestDes2(t *testing.T) {
	key := []byte("47bce5c7")

	//ori64:="rh9+rQ9irvM9lF8iNnXYIROrwETUX40rxs/09NfaPvRvTq9U3iYmnFA8/7xbGKJ6zJQeUHLkyjM="
	//oriStr:="rh9+rQ9irvMkmqivTbV+kbe0qpk7J0kE0I5S2Cgziuau1dgn176Ubu8+r2QPFIxti58gG2E8rk35M9h2l0m8Oq+X9YeSjQWPFZNXYQAuJJIv6H2vVxCosQ=="
	ori64:="rh9+rQ9irvM9lF8iNnXYIROrwETUX40rxs/09NfaPvSSGGvzMN4d4ZJWdkjhPLLrcL43oOzX8OPoNfE47gp3HEUIOGOAPX6xvzxiEYo/ZhLjEnG9azmtow=="
	ori,err:=base64.StdEncoding.DecodeString(ori64)

	if err!=nil{
		t.Errorf("err=%s",err.Error())
	}



	//decrypt
	origData, err := DESDecrypt(ori, key)
	if err != nil {
		panic(err)
	}

	t.Logf("oriData=%s",origData)

}


