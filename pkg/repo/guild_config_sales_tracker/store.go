package guild_config_sales_tracker

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetByGuildID(guildID string) (config []model.GuildConfigSalesTracker, err error)
	Create(config *model.GuildConfigSalesTracker) error
	GetAllSalesTrackerConfig() ([]model.GuildConfigSalesTracker, error)
}
