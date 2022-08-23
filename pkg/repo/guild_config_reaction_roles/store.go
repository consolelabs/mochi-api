package guild_config_reaction_roles

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
)

type Store interface {
	ListAllByGuildID(guildID string) ([]model.GuildConfigReactionRole, error)
	GetByMessageID(guildID, messageID string) (model.GuildConfigReactionRole, error)
	GetByRoleID(guildID, roleID string) (*model.GuildConfigReactionRole, error)
	UpdateRoleConfig(req request.RoleReactionUpdateRequest, updateJson string) error
	CreateRoleConfig(req request.RoleReactionUpdateRequest, updateJson string) error
	ClearMessageConfig(guildID, messageID string) error
}
