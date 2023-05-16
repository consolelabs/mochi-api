package guild_config_admin_role

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	Create(config []model.GuildConfigAdminRole) error
	ListByGuildID(guildID string) ([]model.GuildConfigAdminRole, error)
	Delete(id int) error
}
