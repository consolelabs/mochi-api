package vault

type ListQuery struct {
	GuildID       string
	UserDiscordID string
	EvmWallet     string
	SolanaWallet  string
	Threshold     string
}
