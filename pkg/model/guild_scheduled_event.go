package model

import "github.com/bwmarrin/discordgo"

type GuildScheduledEvent struct {
	GuildID string                              `json:"guild_id"`
	EventID string                              `json:"event_id"`
	Status  discordgo.GuildScheduledEventStatus `json:"status"`
}
