package guild_config_reaction_roles

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
)

type Store interface {
	GetByMessageID(guildId, mesageId string) (model.GuildConfigReactionRole, error)
	UpdateRoleConfig(req request.RoleReactionUpdateRequest, updateJson string) error
	CreateRoleConfig(req request.RoleReactionUpdateRequest, updateJson string) error
}
