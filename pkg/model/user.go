package model

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User struct {
	ID                     int64          `json:"id" gorm:"primary_key"`
	Username               string         `json:"username"`
	Nickname               JSONNullString `json:"nickname"`
	JoinDate               *time.Time     `json:"join_date"`
	InDiscordWalletAddress JSONNullString `json:"in_discord_wallet_address"`
	InDiscordWalletNumber  JSONNullInt64  `json:"in_discord_wallet_number"`

	GuildUsers []*GuildUser `json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	cols := []clause.Column{}
	colsNames := []string{}
	for _, field := range tx.Statement.Schema.PrimaryFields {
		cols = append(cols, clause.Column{Name: field.DBName})
		colsNames = append(colsNames, field.DBName)
	}
	tx.Statement.AddClause(clause.OnConflict{
		Columns:   cols,
		DoNothing: true,
	})
	return nil
}
