package model

import "time"

type MigrateBalance struct {
	ID                int       `json:"-"`
	Symbol            string    `json:"symbol"`
	Username          string    `json:"username"`
	UserDiscordID     string    `json:"user_discord_id"`
	Txhash            string    `json:"txhash"`
	Txurl             string    `json:"txurl"`
	CreatedAt         time.Time `json:"created_at"`
	Transferredamount float64   `json:"transferredamount"`
}

func (MigrateBalance) TableName() string {
	return "migrate_balances"
}
