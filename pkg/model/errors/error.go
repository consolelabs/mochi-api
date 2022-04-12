package errors

import "strconv"

// Error in server
type Error struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func (e Error) Error() string {
	return e.Message + " " + strconv.Itoa(e.Code)
}

// NewStringError new a error with message
func NewStringError(msg string, code int) error {
	return Error{
		Code:    code,
		Message: msg,
	}
}
