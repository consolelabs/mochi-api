package model

import (
	"time"

	"github.com/consolelabs/mochi-typeset/api/typeset"
)

type ProductBotCommand struct {
	Id              int64                          `json:"id"`
	Code            string                         `json:"code"`
	DiscordCommand  string                         `json:"discord_command"`
	TelegramCommand string                         `json:"telegram_command"`
	Scope           typeset.ProductBotCommandScope `json:"scope"`
	Description     string                         `json:"description"`
	DiscordAlias    string                         `json:"discord_alias"`
	TelegramAlias   string                         `json:"telegram_alias"`
	CreatedAt       time.Time                      `json:"created_at"`
	UpdatedAt       time.Time                      `json:"updated_at"`
}

type ProductChangelogs struct {
	Product      string    `json:"product"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	GithubUrl    string    `json:"github_url"`
	ThumbnailUrl string    `json:"thumbnail_url"`
	FileName     string    `json:"file_name"`
	IsExpired    bool      `json:"is_expired"`
	Version      string    `json:"version"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type ProductChangelogView struct {
	Key           string    `json:"key"`
	ChangelogName string    `json:"changelog_name"`
	Metadata      *[]byte   `json:"metadata" gorm:"metadata type:jsonb;default:'[]';not null"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type ProductChangelogSnapshot struct {
	Filename  string    `json:"filename"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
