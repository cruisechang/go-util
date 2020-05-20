package goutil

import (
	"bytes"
	"encoding/json"
	"errors"
	"testing"
)

func TestReturnSuccess(t *testing.T) {
	actual, err := NewReturn(
		NewErr(ExternalOK, nil),
		"success!",
	)
	if err != nil {
		t.Fatal(err)
	}
	expect, err := json.Marshal(ReturnForm{
		Code: ReturnSuccess,
		Msg:  "",
		Data: "success!",
	})
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(expect, actual) {
		t.Error("expect and actual are not the same")
	}
}

func TestKumayCode(t *testing.T) {
	actual, err := NewReturn(
		NewErr(ReturnServerError, errors.New("testing code 9")),
		"",
	)
	if err != nil {
		t.Fatal(err)
	}
	expect, err := json.Marshal(ReturnForm{
		Code: ReturnServerError,
		Msg:  "testing code 9",
	})
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(expect, actual) {
		t.Error("expect and actual are not the same")
	}
}

func TestCodeSubOne(t *testing.T) {
	actual, err := NewReturn(
		NewErr(-1, errors.New("testing code -1")),
		"",
	)
	if err != nil {
		t.Fatal(err)
	}
	expect, err := json.Marshal(ReturnForm{
		Code: ReturnExternalServerError,
		Msg:  "testing code -1",
	})
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(expect, actual) {
		t.Error("expect and actual are not the same")
	}
}

func TestCodeSmallThanZero(t *testing.T) {
	actual, err := NewReturn(
		NewErr(-2, errors.New("testing code -2")),
		"",
	)
	if err != nil {
		t.Fatal(err)
	}
	expect, err := json.Marshal(ReturnForm{
		Code: ReturnParamsError,
		Msg:  "testing code -2",
	})
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(expect, actual) {
		t.Error("expect and actual are not the same")
	}
}
