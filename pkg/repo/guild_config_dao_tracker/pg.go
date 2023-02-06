package guild_config_dao_tracker

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetAllByGuildID(guildId string) (*[]model.GuildConfigDaoTracker, error)
	GetAllBySpace(space string) ([]model.GuildConfigDaoTracker, error)
	DeleteByID(id string) error
	Upsert(model.GuildConfigDaoTracker) error
}
