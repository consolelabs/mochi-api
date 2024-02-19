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
	case
		ErrRecordNotFound,
		ErrTokenNotFound:
		code = http.StatusNotFound
	case
		ErrConflict,
		ErrTokenRequestExisted,
		ErrAliasAlreadyExisted:
		code = http.StatusConflict
	case
		ErrInvalidChain,
		ErrInvalidDiscordChannelID,
		ErrInvalidDiscordGuildID, ErrInvalidTokenContract,
		ErrInvalidProposalType,
		ErrInvalidDiscordUserID,
		ErrInvalidProposalID,
		ErrInvalidVoteID,
		ErrInvalidVoteChoice,
		ErrInvalidAuthorityType,
		ErrXPRoleExisted,
		ErrMixRoleExisted,
		ErrApiKeyBinancePermissionReadingDisabled,
		ErrProfileNotLinkBinance,
		ErrProfile,
		ErrCoingeckoNotSupported,
		ErrInvalidCoingeckoSvcParam,
		ErrKyberRouteNotFound:
		code = http.StatusBadRequest
	case
		ErrConflict,
		ErrChainTypeConflict:
		code = http.StatusConflict
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
