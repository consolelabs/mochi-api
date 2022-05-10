package guild_user_activity_xp

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	CreateOne(record model.GuildUserActivityXP) error
}
