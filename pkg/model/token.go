package model

type Token struct {
	ID                  int    `json:"id"`
	Address             string `json:"address"`
	Symbol              string `json:"symbol"`
	ChainID             int    `json:"chain_id"`
	Decimals            int    `json:"decimal"`
	DiscordBotSupported bool   `json:"discord_bot_supported"`
	CoinGeckoID         string `json:"coin_gecko_id"`
	Name                string `json:"name"`
	GuildDefault        bool   `json:"guild_default"`
	IsNative            bool   `json:"is_native"`
	Chain               *Chain `json:"chain" gorm:"foreignKey:ChainID"`
}
