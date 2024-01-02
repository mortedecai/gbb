package gbb

import (
	"fmt"
	"encoding/json"
)

type Error struct {
	Method  string      `json:"method"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func CreateError(method string, msg string, details interface{}) *Error {
	return &Error{Method: method, Message: msg, Details: details}
}

func (err *Error) Error() string {
	return fmt.Sprintf("[%s] %s [Details: |%v|]", err.Method, err.Message, err.Details)
}

func (err *Error) JSON() string {
	data, _ := json.Marshal(err)
	return string(data)
}