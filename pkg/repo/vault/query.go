package vault

type ListQuery struct {
	GuildID       string
	UserProfileID string
	EvmWallet     string
	SolanaWallet  string
	Threshold     string
	VaultIDs      []string
	GetArchived   bool
}
