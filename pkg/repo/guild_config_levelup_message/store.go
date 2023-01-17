package guild_config_levelup_message

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetByGuildId(guildId string) (*model.GuildConfigLevelupMessage, error)
	DeleteByGuildId(guildId string) error
	UpsertOne(config model.GuildConfigLevelupMessage) (*model.GuildConfigLevelupMessage, error)
}
