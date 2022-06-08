package request

type UpsertCustomTokenConfigRequest struct {
	Id       int    `json:"id"`
	Address  string `json:"address"`
	Symbol   string `json:"symbol"`
	ChainID  int    `json:"chain_id" gorm:"default:0"`
	Decimals int    `json:"decimals" gorm:"default:18"`

	DiscordBotSupported bool   `json:"discord_bot_supported" gorm:" default:true"`
	CoinGeckoID         string `json:"coin_gecko_id" gorm:"default:''"`
	GuildID             string `json:"guild_id"`

	Name         string `json:"name" gorm:"default:''"`
	GuildDefault bool   `json:"guild_default" gorm:"default:0"`
	Active       bool   `json:"active" gorm:"default:false"`
}
