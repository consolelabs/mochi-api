package guild_config_dao_tracker

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/response"
)

type Store interface {
	GetAllByGuildID(guildId string) (*[]model.GuildConfigDaoTracker, error)
	GetAllBySpace(space string) ([]model.GuildConfigDaoTracker, error)
	GetAllBySpaceAndSource(space, source string) ([]model.GuildConfigDaoTracker, error)
	GetUsageStatsWithPaging(page int, size int) (models []response.DaoTrackerSpaceCountData, total int64, err error)
	DeleteByID(id string) error
	Upsert(model.GuildConfigDaoTracker) error
}
