package uuid

import (
	"testing"
	"fmt"
)

func TestGetV4(t *testing.T) {

	u,err:=GetV4()

	if err != nil {
		t.Error(err)

	} else {
		fmt.Printf("%s", u)
	}
}


func TestGetV1(t *testing.T) {

	u,err:=GetV1()

	if err != nil {
		t.Error(err)

	} else {
		fmt.Printf("%s", u)
	}
}