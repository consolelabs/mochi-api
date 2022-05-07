package discord_guild_stats

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	UpsertOne(streak model.DiscordGuildStat) error
	GetByGuildID(guildID string) (*model.DiscordGuildStat, error)
}
