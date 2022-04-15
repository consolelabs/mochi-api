package guild_users

type Store interface {
	Update(guildId, userId string, field string, value interface{}) error
	CountByGuildUser(guildId, userId string) (int64, error)
}
