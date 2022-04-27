package reaction_role_configs

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
)

type Store interface {
	GetByMessageID(guildId, mesageId string) (model.ReactionRoleConfig, error)
	UpdateRoleConfig(req request.RoleReactionUpdateRequest, updateJson string) error
	CreateRoleConfig(req request.RoleReactionUpdateRequest, updateJson string) error
}
