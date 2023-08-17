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
	CreatedAt       time.Time                      `json:"created_at"`
	UpdatedAt       time.Time                      `json:"updated_at"`
}
