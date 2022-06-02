package model

import (
	"time"

	"github.com/google/uuid"
)

type DiscordGuildStat struct {
	ID      uuid.NullUUID `json:"id" gorm:"default:uuid_generate_v4()"`
	GuildID string        `json:"guild_id"`

	NrOfMembers int `json:"nr_of_members"`
	NrOfUsers   int `json:"nr_of_users"`
	NrOfBots    int `json:"nr_of_bots"`

	NrOfChannels             int `json:"nr_of_channels"`
	NrOfTextChannels         int `json:"nr_of_text_channels"`
	NrOfVoiceChannels        int `json:"nr_of_voice_channels"`
	NrOfStageChannels        int `json:"nr_of_stage_channels"`
	NrOfCategories           int `json:"nr_of_categories"`
	NrOfAnnouncementChannels int `json:"nr_of_announcement_channels"`

	NrOfEmojis         int `json:"nr_of_emojis"`
	NrOfStaticEmojis   int `json:"nr_of_static_emojis"`
	NrOfAnimatedEmojis int `json:"nr_of_animated_emojis"`

	NrOfStickers       int `json:"nr_of_stickers"`
	NrOfCustomStickers int `json:"nr_of_custom_stickers"`
	NrOfServerStickers int `json:"nr_of_server_stickers"`

	NrOfRoles int `json:"nr_of_roles"`

	CreatedAt time.Time `json:"created_at"`
}
