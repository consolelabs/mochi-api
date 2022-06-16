package guild_scheduled_event

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	UpsertOne(config *model.GuildScheduledEvent) error
	ListUncompleteByGuildID(guildID string) ([]model.GuildScheduledEvent, error)
}
