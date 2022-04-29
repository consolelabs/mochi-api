package guild_config_default_roles

import (
	"github.com/defipod/mochi/pkg/model"
)

type Store interface {
	GetAllByGuildID(guildID string) (model.GuildConfigDefaultRole, error)
	CreateDefaultRoleIfNotExist(config model.GuildConfigDefaultRole) error
	DeleteByGuildID(guildID string) error
}
