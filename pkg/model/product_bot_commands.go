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
	Id           int64                     `json:"id"`
	Product      typeset.ProductChangeLogs `json:"product"`
	Title        string                    `json:"title"`
	Content      string                    `json:"content"`
	GithubUrl    string                    `json:"github_url"`
	ThumbnailUrl string                    `json:"thumbnail_url"`
	CreatedAt    time.Time                 `json:"created_at"`
	UpdatedAt    time.Time                 `json:"updated_at"`
}
