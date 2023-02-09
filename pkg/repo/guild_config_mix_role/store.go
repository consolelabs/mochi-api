package guild_config_mix_role

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	Create(config *model.GuildConfigMixRole) error
	Get(id int) (model *model.GuildConfigMixRole, err error)
	ListByGuildID(guildID string) ([]model.GuildConfigMixRole, error)
	GetByRoleID(guildID, roleID string) (*model.GuildConfigMixRole, error)
	Delete(id int) error
	GetMemberCurrentRoles(guildID string) ([]model.MemberMixRole, error)
}
