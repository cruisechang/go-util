package util

import (
	"encoding/json"
	"fmt"
)

const (
	// ExternalOK means the code from external server status code.
	ExternalOK = 1

	// ReturnSuccess = 0
	ReturnSuccess = -1 + iota
	// ReturnExternalServerError map to external server error code 1
	ReturnExternalServerError
	// ReturnParamsError map to external server error code 2
	ReturnParamsError
	// ReturnJSONError map to external server error code 3
	ReturnJSONError
	// ReturnAlreadyLogin map to external server error code 4
	ReturnAlreadyLogin
	// ReturnAccountOrPasswordError map to external server error code 5
	ReturnAccountOrPasswordError
	// ReturnSessionError map to external server error code 6
	ReturnSessionError
	// ReturnPasswordNotMatched map to external server error code 7
	ReturnPasswordNotMatched
	// ReturnAccountNotExist map to external server error code 8
	ReturnAccountNotExist

	// ReturnServerError is 503 server internal error. code 9
	ReturnServerError
	// ReturnUnauthorized is 401 Unauthorized error. code 10
	ReturnUnauthorized
)

// Err is the inner error tranport format.
type Err struct {
	Code int
	Msg  error
}

func (e *Err) Error() string {
	return e.Msg.Error()
}

// NewErr returns a Err instance.
func NewErr(code int, msg error) *Err {
	return &Err{
		Code: code,
		Msg:  msg,
	}
}

// ReturnForm is the kumay all server output form.
type ReturnForm struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

// NewReturn returns the unified output in JSON format.
func NewReturn(kumayErr error, data string) ([]byte, error) {
	e, ok := kumayErr.(*Err)
	if !ok {
		return nil, fmt.Errorf("cannot cast into kumay Err type")
	}
	var out []byte
	var err error

	if e.Code == ExternalOK || e.Code == ReturnSuccess {
		out, err = json.Marshal(&ReturnForm{
			Code: ReturnSuccess,
			Data: data,
		})
	} else if e.Code < 0 {
		out, err = json.Marshal(&ReturnForm{
			Code: -e.Code,
			Msg:  e.Msg.Error(),
		})
	} else {
		out, err = json.Marshal(&ReturnForm{
			Code: e.Code,
			Msg:  e.Msg.Error(),
		})
	}
	if err != nil {
		return nil, err
	}
	return out, nil
}
