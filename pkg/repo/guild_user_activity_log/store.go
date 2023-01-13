package guild_user_activity_log

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	CreateOne(record model.GuildUserActivityLog) error
	CreateOneNoGuild(record model.GuildUserActivityLog) error
	CreateBatch(records []model.GuildUserActivityLog) error
}
