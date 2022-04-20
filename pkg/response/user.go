package response

type GetUserResponse struct {
	ID                     string                  `json:"id"`
	Username               string                  `json:"username"`
	InDiscordWalletAddress *string                 `json:"in_discord_wallet_address"`
	InDiscordWalletNumber  *int64                  `json:"in_discord_wallet_number"`
	GuildUsers             []*GetGuildUserResponse `json:"guild_users"`
}

type GetGuildUserResponse struct {
	GuildID   string  `json:"guild_id"`
	UserID    string  `json:"user_id"`
	Nickname  *string `json:"nickname"`
	InvitedBy *string `json:"invited_by"`
}
