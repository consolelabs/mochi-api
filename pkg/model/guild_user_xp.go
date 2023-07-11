package model

type GuildUserXP struct {
	GuildID          string                `json:"guild_id"`
	UserID           string                `json:"user_id"`
	ProfileID        string                `json:"profile_id"`
	TotalXP          int                   `json:"total_xp"`
	Level            int                   `json:"level"`
	User             *User                 `json:"user" gorm:"foreignKey:UserID"`
	Guild            *DiscordGuild         `json:"guild" gorm:"foreignKey:GuildID"`
	GuildRank        int                   `json:"guild_rank"`
	NrOfActions      int                   `json:"nr_of_actions"`
	Progress         float64               `json:"progress" gorm:"-"`
	CurrentLevelRole *GuildConfigLevelRole `json:"current_level_role" gorm:"-"`
	NextLevelRole    *GuildConfigLevelRole `json:"next_level_role" gorm:"-"`
}
