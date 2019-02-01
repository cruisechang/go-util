package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"testing"

	kumayutil "wangzugames.com/kumay/util"
)

func TestReturnSuccess(t *testing.T) {
	actual, err := kumayutil.NewReturn(
		kumayutil.NewErr(kumayutil.ExternalOK, nil),
		"success!",
	)
	if err != nil {
		t.Fatal(err)
	}
	expect, err := json.Marshal(kumayutil.ReturnForm{
		Code: kumayutil.ReturnSuccess,
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
	actual, err := kumayutil.NewReturn(
		kumayutil.NewErr(kumayutil.ReturnServerError, errors.New("testing code 9")),
		"",
	)
	if err != nil {
		t.Fatal(err)
	}
	expect, err := json.Marshal(kumayutil.ReturnForm{
		Code: kumayutil.ReturnServerError,
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
	actual, err := kumayutil.NewReturn(
		kumayutil.NewErr(-1, errors.New("testing code -1")),
		"",
	)
	if err != nil {
		t.Fatal(err)
	}
	expect, err := json.Marshal(kumayutil.ReturnForm{
		Code: kumayutil.ReturnExternalServerError,
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
	actual, err := kumayutil.NewReturn(
		kumayutil.NewErr(-2, errors.New("testing code -2")),
		"",
	)
	if err != nil {
		t.Fatal(err)
	}
	expect, err := json.Marshal(kumayutil.ReturnForm{
		Code: kumayutil.ReturnParamsError,
		Msg:  "testing code -2",
	})
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(expect, actual) {
		t.Error("expect and actual are not the same")
	}
}
