package model

type GuildConfigToken struct {
	GuildID   string `json:"guild_id"`
	TokenID   int    `json:"token_id"`
	Active    bool   `json:"active"`
	IsDefault bool   `json:"id_default"`
	Token     *Token `json:"token" gorm:"foreignKey:TokenID"`
}
