package model

type GuildConfigActivity struct {
	GuildID    string    `json:"guild_id"`
	ActivityID int       `json:"activity_id"`
	Active     bool      `json:"active"`
	Activity   *Activity `json:"activity" gorm:"foreignKey:ActivityID"`
}
