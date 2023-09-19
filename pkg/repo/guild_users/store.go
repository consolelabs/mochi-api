package guild_users

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetGuildUsers(guildID string) ([]model.GuildUser, error)
	Update(guildId, userId string, field string, value interface{}) error
	CountByGuildUser(guildId, userId string) (int64, error)
	FirstOrCreate(guildUser *model.GuildUser) error
	Create(guildUser *model.GuildUser) error
	UpsertMany(guildUsers []model.GuildUser) error
	GetUsersOfGuild(ids []string, guildId string) (res []model.GuildUser, err error)
}
