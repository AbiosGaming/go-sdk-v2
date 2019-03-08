package structs

import (
	"fmt"
)

// ErrorStruct represents an error response from the API.
type ErrorStruct struct {
	ErrorMessage     string `json:"error,omitempty"`
	ErrorCode        int64  `json:"error_code,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

func (e ErrorStruct) String() (s string) {
	return e.format()
}

func (e ErrorStruct) Error() (s string) {
	return e.format()
}

func (e ErrorStruct) format() (s string) {
	s = fmt.Sprintf("error: %v, error_code: %v, error_decription: %v",
		e.Error, e.ErrorCode, e.ErrorDescription)
	return s
}
