package reaction_role_configs

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	Gets(guildId string) ([]model.ReactionRoleConfig, error)
}
