package model

import (
	"github.com/google/uuid"
)

type ReactionRoleConfig struct {
	ID            uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()"`
	GuildID       string        `json:"guild_id"`
	ChannelID     string        `json:"channel_id"`
	Title         string        `json:"title"`
	TitleUrl      string        `json:"title_url"`
	ThumbnailUrl  string        `json:"thumbnail_url"`
	Description   string        `json:"description"`
	FooterImage   string        `json:"footer_image"`
	FooterMessage string        `json:"footer_message"`
	MessageID     string        `json:"message_id"`
	ReactionRoles string        `json:"reaction_roles"`
}
