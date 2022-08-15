package guild_config_level_role

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetHighest(guildID string, level int) (*model.GuildConfigLevelRole, error)
	GetByGuildID(guildID string) ([]model.GuildConfigLevelRole, error)
	GetByRoleID(guildID, roleID string) (*model.GuildConfigLevelRole, error)
	UpsertOne(config model.GuildConfigLevelRole) error
	DeleteOne(guildID string, level int) error
}
