package model

import "time"

type OnchainTipBotTransaction struct {
	ID                 int       `json:"id"`
	SenderDiscordID    string    `json:"sender"`
	RecipientDiscordID string    `json:"recipients"`
	RecipientAddress   *string   `json:"recipient_address"`
	GuildID            string    `json:"guild_id"`
	ChannelID          string    `json:"channel_id"`
	Amount             float64   `json:"amount"`
	TokenSymbol        string    `json:"token_symbol"`
	Each               bool      `json:"each"`
	All                bool      `json:"all"`
	TransferType       string    `json:"transfer_type"`
	FullCommand        string    `json:"full_command"`
	Duration           int       `json:"duration"`
	Message            string    `json:"message"`
	Image              string    `json:"image"`
	TxHash             string    `json:"tx_hash"`
	Status             string    `json:"string"` // (pending, claimed)
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	ClaimedAt          time.Time `json:"claimed_at"`
}

func (OnchainTipBotTransaction) TableName() string {
	return "onchain_tip_bot_transactions"
}
