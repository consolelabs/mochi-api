package errors

import (
	"net/http"
	"strconv"
)

// Error in server
type Error struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func (e Error) Error() string {
	return e.Message + " " + strconv.Itoa(e.Code)
}

func GetStatusCode(err error) int {
	var code int
	switch err {
	case ErrRecordNotFound:
		code = http.StatusNotFound
	case ErrConflict:
		code = http.StatusConflict
	case
		ErrInvalidChain,
		ErrInvalidDiscordChannelID,
		ErrInvalidDiscordGuildID, ErrInvalidTokenContract,
		ErrInvalidProposalType,
		ErrInvalidDiscordUserID,
		ErrInvalidProposalID:
		code = http.StatusBadRequest
	default:
		code = http.StatusInternalServerError
	}
	return code
}

// NewStringError new a error with message
func NewStringError(msg string, code int) error {
	return Error{
		Code:    code,
		Message: msg,
	}
}
