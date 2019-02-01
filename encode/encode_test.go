package encode

import (
	"encoding/json"
	"testing"
)

func TestByteToStruct(t *testing.T) {

	type jt struct {
		AA int
		B  int
		C  string
	}
	jts := jt{}

	//const oriStr = `{"aa": 1, "b": 2,  "c": "心心心情好"}`

	var jsonBlob = []byte("{\"aa\": 1, \"b\": 2,  \"c\": \"心心心情好\"}")

	err := json.Unmarshal(jsonBlob, &jts)
	if err != nil {
		println(err)
	}

	expectedA := 1
	expectedB := 2
	expectedC := "心心心情好"
	actualA := jts.AA
	actualB := jts.B
	actualC := jts.C

	if expectedA != actualA {
		//t.Errorf("ByteToMap error: expect-> %d, actual->%d", expectedA, actualA)
	}

	if expectedB != actualB {
		t.Errorf("ByteToMap error: expect->%d, actual->%d", expectedB, actualB)
	}
	if expectedC != actualC {
		t.Errorf("ByteToMap error: expect->%s, actual->%s", expectedC, actualC)
	}

}

func TestByteToMap(t *testing.T) {
	strByte := []byte("{\"account\":\"test\",\"password\":\"testingtesting\"}")

	m := map[string]interface{}{}
	err := json.Unmarshal(strByte, m)
	if err != nil {

	}

	if m["account"] != "test" {
		t.Errorf("TestByteToMap error resutl %s", m["account"])
	}
	if m["password"] != "testingtesting" {
		t.Errorf("TestByteToMap error resutl %s", m["password"])
	}
}
