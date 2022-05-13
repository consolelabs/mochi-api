package model

type GuildUserXP struct {
	GuildID   string `json:"guild_id"`
	UserID    string `json:"user_id"`
	TotalXP   int    `json:"total_xp"`
	Level     int    `json:"level"`
	User      *User  `json:"user" gorm:"foreignKey:UserID"`
	GuildRank int    `json:"guild_rank"`
}
