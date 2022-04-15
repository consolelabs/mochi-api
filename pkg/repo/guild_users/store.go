package guild_users

type Store interface {
	Update(guildId, userId int64, field string, value interface{}) error
	CountByGuildUser(guildId, userId int64) (int64, error)
}
