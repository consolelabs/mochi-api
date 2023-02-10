package errors

var (
	//ErrInvokeValidatorEngineFailed ...
	ErrInvokeValidatorEngineFailed = NewStringError("Fail to invoke validator engine", 500)

	//ErrTypeAssertionFailed ...
	ErrTypeAssertionFailed = NewStringError("Fail to type assertion", 400)

	ErrInvalidDiscordChannelID = NewStringError("Invalid Discord channel ID", 400)
	ErrInvalidDiscordGuildID   = NewStringError("Invalid Discord guild ID", 400)
	ErrInvalidDiscordUserID    = NewStringError("Invalid Discord user ID", 400)
	ErrInvalidChain            = NewStringError("Invalid chain", 400)
	ErrInvalidTokenContract    = NewStringError("Invalid token contract", 400)
	ErrInvalidProposalType     = NewStringError("Invalid proposal type", 400)
	ErrInvalidProposalID       = NewStringError("Invalid proposal ID", 400)
	ErrInvalidVoteID           = NewStringError("Invalid vote ID", 400)
	ErrInvalidVoteChoice       = NewStringError("Invalid vote choice", 400)
	ErrInvalidAuthorityType    = NewStringError("Invalid authority type", 400)
	ErrTokenNotFound           = NewStringError("Token not found", 404)
	ErrXPRoleExisted           = NewStringError("XP role config already existed", 400)
	ErrMixRoleExisted          = NewStringError("Mix role config already existed", 400)
)
