package guild_config_xp_role

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	Create(config *model.GuildConfigXPRole) error
	Get(id int) (model *model.GuildConfigXPRole, err error)
	ListByGuildID(guildID string) ([]model.GuildConfigXPRole, error)
	GetByRoleID(guildID, roleID string) (*model.GuildConfigXPRole, error)
	Update(config *model.GuildConfigXPRole) error
	Delete(id int) error
	GetMemberCurrentRoles(guildID string) ([]model.MemberXPRole, error)
}
