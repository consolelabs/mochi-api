package model

import "time"

type VaultTransaction struct {
	Id          int64     `json:"id"`
	GuildId     string    `json:"guild_id"`
	VaultId     int64     `json:"vault_id"`
	VaultName   string    `json:"vault_name" gorm:"-"`
	Action      string    `json:"action"`
	FromAddress string    `json:"from_address"`
	ToAddress   string    `json:"to_address"`
	Target      string    `json:"target"`
	Amount      string    `json:"amount"`
	Token       string    `json:"token"`
	Threshold   string    `json:"threshold"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Sender      string    `json:"sender"`
}
