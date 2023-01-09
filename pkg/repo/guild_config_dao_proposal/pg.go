package guild_config_dao_proposal

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetByGuildIDAndGuideLineChannelID(guildId, guidelineChannelId string) (*model.GuildConfigDaoProposal, error)
	GetById(id int64) (*model.GuildConfigDaoProposal, error)
	GetByGuildId(guildId string) (*model.GuildConfigDaoProposal, error)
	DeleteById(id string) (*model.GuildConfigDaoProposal, error)
	Create(config model.GuildConfigDaoProposal) (*model.GuildConfigDaoProposal, error)
}
