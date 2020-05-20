package maper


import (
	"fmt"
	"testing"
)

func TestGetMapIntKeys(t *testing.T) {

	fmt.Println("test 1 int map keys ")
	v := map[int]string{1: "a", 2: "b"}

	keys, err := GetMapIntKeys(v)
	if err != nil {
		t.Errorf("GetMapKeys error : %v \n", err)
	} else {
		fmt.Println(keys)
	}

	fmt.Println("test 2 string map keys ")
	m := map[string]string{"1": "a", "2": "b"}
	keys, err = GetMapIntKeys(m)
	if err != nil {
		t.Errorf("GetMapKeys error : %v \n", err)
	} else {
		fmt.Println(keys)
	}

	fmt.Println("test 3 empty int  map keys ")
	m2 := map[int]int{}
	keys, err = GetMapIntKeys(m2)
	if err != nil {
		t.Errorf("GetMapKeys error : %v \n", err)
	} else {
		fmt.Printf("len of keys:%d,keys:%v \n", len(keys), keys)
	}
}

func TestGetMapStringKeys(t *testing.T) {

	fmt.Println("test 2 string map keys ")
	v := map[string]string{"a": "a", "b": "b"}

	keys, err := GetMapStringKeys(v)
	if err != nil {
		t.Errorf("GetMapKeys error : %v \n", err)
	} else {
		fmt.Println("keys:", keys)
	}

	fmt.Println("test 2 int map keys ")
	m := map[int]string{0: "a", 1: "b"}

	keys, err = GetMapStringKeys(m)
	if err != nil {
		t.Errorf("GetMapKeys error : %v \n", err)
	} else {
		fmt.Println("keys:", keys)
	}
}
