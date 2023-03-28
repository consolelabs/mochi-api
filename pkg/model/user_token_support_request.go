package model

import "time"

type TokenSupportRequestStatus string

const (
	TokenSupportApproved TokenSupportRequestStatus = "approved"
	TokenSupportRejected TokenSupportRequestStatus = "rejected"
	TokenSupportPending  TokenSupportRequestStatus = "pending"
)

type UserTokenSupportRequest struct {
	ID            int                       `json:"id"`
	UserDiscordID string                    `json:"user_discord_id"`
	GuildID       string                    `json:"guild_id"`
	ChannelID     string                    `json:"channel_id"`
	MessageID     string                    `json:"message_id"`
	TokenAddress  string                    `json:"token_address"`
	TokenChainID  int                       `json:"token_chain_id"`
	Status        TokenSupportRequestStatus `json:"status"`
	UpdatedAt     time.Time                 `json:"updated_at"`
	CreatedAt     time.Time                 `json:"created_at"`
}
