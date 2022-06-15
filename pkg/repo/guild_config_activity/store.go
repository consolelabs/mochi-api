package guild_config_activity

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetOneByActivityName(guildID, activityName string) (*model.GuildConfigActivity, error)
	UpsertMany(configs []model.GuildConfigActivity) error
	ForkDefaulActivityConfigs(guildID string) error
	ListByActivity(activity string) ([]model.GuildConfigActivity, error)
	UpsertToggleActive(config *model.GuildConfigActivity) error
}
