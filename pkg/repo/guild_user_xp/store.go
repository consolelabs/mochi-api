package guild_user_xp

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetOne(guildID, userID string) (*model.GuildUserXP, error)
	GetTopUsers(guildID string, limit, offset int) ([]model.GuildUserXP, error)
}
