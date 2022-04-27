package guild_users

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	Update(guildId, userId string, field string, value interface{}) error
	CountByGuildUser(guildId, userId string) (int64, error)
	FirstOrCreate(guildUser *model.GuildUser) error
}
