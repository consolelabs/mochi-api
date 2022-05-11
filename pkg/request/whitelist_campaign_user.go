package request

type AddWhitelistCampaignUser struct {
	Address             string `json:"address"`
	DiscordID           string `json:"discord_id"`
	Notes               string `json:"notes"`
	WhitelistCampaignId string `json:"whitelist_campaign_id"`
}

type AddWhitelistCampaignUserRequest struct {
	Users []AddWhitelistCampaignUser `json:"users"`
}
