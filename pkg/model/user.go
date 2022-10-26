package model

type User struct {
	ID                     string         `json:"id" gorm:"primary_key"`
	Username               string         `json:"username"`
	InDiscordWalletAddress JSONNullString `json:"in_discord_wallet_address"`
	InDiscordWalletNumber  JSONNullInt64  `json:"in_discord_wallet_number"`
	NrOfJoin               int64          `json:"nr_of_join"`
	IsMigrateBal           bool           `json:"is_migrate_bal"`

	GuildUsers []*GuildUser `json:"guild_users"`
}
