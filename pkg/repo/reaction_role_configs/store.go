package reaction_role_configs

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetByMessageID(guildId, mesageId string) (model.ReactionRoleConfig, error)
}
