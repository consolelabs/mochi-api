package request

type CreateVaultRequest struct {
	GuildId   string `json:"guild_id"`
	Name      string `json:"name"`
	Threshold string `json:"threshold"`
}
