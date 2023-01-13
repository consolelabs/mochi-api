package model

type UserTokenBalance struct {
	UserAddress string         `json:"user_address"`
	ChainType   JSONNullString `json:"chain_type"`
	TokenID     int            `json:"token_id"`
	Balance     float64        `json:"balance"`
}

type UserTokenBalancesByGuild struct {
	UserDiscordId string `json:"user_discord_id"`
	TotalBalance  int64  `json:"total_balance"`
}
