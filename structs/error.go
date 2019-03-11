package structs

import (
	"fmt"
)

// Error represents an error response from the API.
type Error struct {
	ErrorMessage     string `json:"error,omitempty"`
	ErrorCode        int64  `json:"error_code,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}

func (e Error) String() (s string) {
	return e.format()
}

func (e Error) Error() (s string) {
	return e.format()
}

func (e Error) format() (s string) {
	s = fmt.Sprintf("error: %v, error_code: %v, error_decription: %v",
		e.ErrorMessage, e.ErrorCode, e.ErrorDescription)
	return s
}
