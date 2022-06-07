package request

type UpsertCustomTokenConfigRequest struct {
	Id       int    `json:"id"`
	Address  string `json:"address"`
	Symbol   string `json:"symbol"`
	ChainID  int    `json:"chain_id"`
	Decimals int    `json:"decimals"`

	DiscordBotSupported bool   `json:"discord_bot_supported"`
	CoinGeckoID         string `json:"coin_gecko_id"`
	GuildID             string `json:"guild_id"`

	Name         string `json:"name"`
	GuildDefault bool   `json:"guild_default"`
	Active       bool   `json:"active"`
}
