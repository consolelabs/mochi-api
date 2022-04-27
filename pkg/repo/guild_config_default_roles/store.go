package guild_config_default_roles

import (
	"github.com/defipod/mochi/pkg/model"
)

type Store interface {
	GetAll() ([]model.GuildConfigDefaultRole, error)
	CreateDefaultRoleIfNotExist(config model.GuildConfigDefaultRole) error
}
