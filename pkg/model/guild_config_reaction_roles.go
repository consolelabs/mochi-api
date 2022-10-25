package model

import (
	"github.com/google/uuid"
)

type GuildConfigReactionRole struct {
	ID            uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	MessageID     string        `json:"message_id"`
	GuildID       string        `json:"guild_id"`
	ChannelID     string        `json:"channel_id"`
	ReactionRoles string        `json:"reaction_roles"`
}
