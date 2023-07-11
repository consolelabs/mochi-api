package guild_user_xp

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetOne(GetOneQuery) (*model.GuildUserXP, error)
	GetByGuildID(guildID string) ([]model.GuildUserXP, error)
	GetTopUsers(guildID, query, sort string, limit, offset int) ([]model.GuildUserXP, error)
	GetTotalTopUsersCount(guildID, query string) (int64, error)
}
