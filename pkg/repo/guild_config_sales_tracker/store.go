package guild_config_sales_tracker

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetByGuildID(guildID string) (*model.GuildConfigSalesTracker, error)
	UpsertOne(config *model.GuildConfigSalesTracker) error
}
