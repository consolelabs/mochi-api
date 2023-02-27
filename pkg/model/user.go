package model

type User struct {
	ID         string       `json:"id" gorm:"primary_key"`
	Username   string       `json:"username"`
	NrOfJoin   int64        `json:"nr_of_join"`
	GuildUsers []*GuildUser `json:"guild_users"`
}
