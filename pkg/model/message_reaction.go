package model

import "github.com/google/uuid"

type MessageReaction struct {
	ID        uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	MessageID string        `json:"message_id"`
	GuildID   string        `json:"guild_id"`
	UserID    string        `json:"user_id"`
	Reaction  string        `json:"reaction"`
}
