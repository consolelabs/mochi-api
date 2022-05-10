package request

type CreateWhitelistCampaignRequest struct {
	Name    string `json:"name"`
	GuildID string `json:"guild_id"`
}
