package guild_config_dao_proposal

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetById(id int64) (*model.GuildConfigDaoProposal, error)
}
