package model

import (
	"github.com/google/uuid"
)

type GuildConfigInviteTracker struct {
	ID         uuid.NullUUID  `json:"id" gorm:"default:uuid_generate_v4()"`
	GuildID    string          `json:"guild_id"`
	ChannelID  string          `json:"user_id"`
	WebhookURL JSONNullString `json:"webhook_url"`
}
