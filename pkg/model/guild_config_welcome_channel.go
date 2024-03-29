package model

import (
	"github.com/google/uuid"
)

type GuildConfigWelcomeChannel struct {
	ID             uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	GuildID        string        `json:"guild_id"`
	ChannelID      string        `json:"channel_id"`
	WelcomeMessage string        `json:"welcome_message"`
}
