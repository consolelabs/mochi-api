package model

import (
	"github.com/google/uuid"
)

type GuildConfigRepostReaction struct {
	ID              uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()" swaggertype:"string"`
	GuildID         string        `json:"guild_id"`
	Emoji           string        `json:"emoji"`
	Quantity        int           `json:"quantity"`
	RepostChannelID string        `json:"repost_channel_id"`
	EmojiStart      string        `json:"emoji_start"`
	EmojiStop       string        `json:"emoji_stop"`
	ReactionType    string        `json:"reaction_type"`
}
