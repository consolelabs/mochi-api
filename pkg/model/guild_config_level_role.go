package model

type GuildConfigLevelRole struct {
	GuildID     string         `json:"guild_id"`
	Level       int            `json:"level"`
	RoleID      string         `json:"role_id"`
	LevelConfig *ConfigXpLevel `json:"level_config" gorm:"foreignKey:Level"`
}
