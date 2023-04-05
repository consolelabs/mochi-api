package mochiprofile

import "time"

type Platform string

const (
	PlatformDiscord Platform = "discord"
	PlatformSol     Platform = "solana-chain"
	PlatformEVM     Platform = "evm-chain"
)

type GetProfileByDiscordResponse struct {
	ID                 string              `json:"id"`
	AssociatedAccounts []AssociatedAccount `json:"associated_accounts"`
	CreatedAt          time.Time           `json:"created_at"`
	UpdatedAt          time.Time           `json:"updated_at"`
}

type AssociatedAccount struct {
	ID                 string    `json:"id"`
	ProfileID          string    `json:"profile_id"`
	Platform           Platform  `json:"platform"`
	PlatformIdentifier string    `json:"platform_identifier"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type ErrorResponse struct {
	Msg        string `json:"msg"`
	StatusCode int    `json:"status_code"`
	ErrorCode  string `json:"error_code"`
}
