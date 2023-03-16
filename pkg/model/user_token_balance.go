package model

type UserTokenBalance struct {
	UserDiscordID string `json:"user_discord_id"`
	TokenID       int    `json:"token_id"`
	Balance       string `json:"balance" gorm:"type:numeric"`
}

type UserTokenBalancesByGuild struct {
	UserDiscordId string `json:"user_discord_id"`
	TotalBalance  int64  `json:"total_balance"`
}
