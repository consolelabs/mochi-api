package model

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User struct {
	ID                     string         `json:"id" gorm:"primary_key"`
	Username               string         `json:"username"`
	InDiscordWalletAddress JSONNullString `json:"in_discord_wallet_address"`
	InDiscordWalletNumber  JSONNullInt64  `json:"in_discord_wallet_number"`

	GuildUsers []*GuildUser `json:"guild_users"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	cols := []clause.Column{}
	for _, field := range tx.Statement.Schema.PrimaryFields {
		cols = append(cols, clause.Column{Name: field.DBName})
	}
	tx.Statement.AddClause(clause.OnConflict{
		Columns:   cols,
		DoUpdates: clause.AssignmentColumns([]string{"in_discord_wallet_number", "in_discord_wallet_address"}),
	})
	return nil
}
