package guild_config_tip_range

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetByGuildID(guildID string) (*model.GuildConfigTipRange, error)
	UpsertOne(config *model.GuildConfigTipRange) (*model.GuildConfigTipRange, error)
	Remove(guildID string) error
}
