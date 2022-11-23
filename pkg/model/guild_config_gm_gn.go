package model

import (
	"github.com/google/uuid"
)

type GuildConfigGmGn struct {
	ID        uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	GuildID   string        `json:"guild_id"`
	ChannelID string        `json:"channel_id"`
	Msg       string        `json:"msg"`
	Emoji     string        `json:"emoji"`
	Sticker   string        `json:"sticker"`
}
