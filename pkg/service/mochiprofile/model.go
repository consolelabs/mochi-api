package mochiprofile

import "time"

type Platform string

const (
	PlatformDiscord Platform = "discord"
	PlatformSol     Platform = "solana-chain"
	PlatformEVM     Platform = "evm-chain"
	PlatformTON     Platform = "ton"
)

type GetProfileResponse struct {
	ID                 string              `json:"id"`
	AssociatedAccounts []AssociatedAccount `json:"associated_accounts"`
	CreatedAt          time.Time           `json:"created_at"`
	UpdatedAt          time.Time           `json:"updated_at"`
}

type AssociatedAccount struct {
	ID                 string           `json:"id"`
	ProfileID          string           `json:"profile_id"`
	Platform           Platform         `json:"platform"`
	PlatformIdentifier string           `json:"platform_identifier"`
	PlatformMetadata   PlatformMetadata `json:"platform_metadata"`
	CreatedAt          time.Time        `json:"created_at"`
	UpdatedAt          time.Time        `json:"updated_at"`
}

type PlatformMetadata struct {
	Username  string `json:"username"`
	ApiSecret string `json:"api_secret"`
}
type ErrorResponse struct {
	Msg        string `json:"msg"`
	StatusCode int    `json:"status_code"`
	ErrorCode  string `json:"error_code"`
}

type ProfileApiKeyResponseData struct {
	Data ProfileApiKeyResponse `json:"data"`
}

type ProfileApiKeyResponse struct {
	ProfileId string    `json:"profile_id"`
	ApiKey    string    `json:"api_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type EvmAssociatedAccount struct {
	ID                 string           `json:"id"`
	ProfileID          string           `json:"profile_id"`
	Platform           Platform         `json:"platform"`
	PlatformIdentifier string           `json:"platform_identifier"`
	PlatformMetadata   PlatformMetadata `json:"platform_metadata"`
	DiscordId          string           `json:"discord_id"`
	CreatedAt          time.Time        `json:"created_at"`
	UpdatedAt          time.Time        `json:"updated_at"`
}

type OnboardingStatusResponse struct {
	DidOnboardingTelegram bool `json:"did_onboarding_telegram"`
}
