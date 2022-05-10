package model

import "time"

type WhitelistCampaignUser struct {
	Address             string    `json:"address"`
	DiscordID           string    `json:"discord_id"`
	Notes               string    `json:"notes"`
	WhitelistCampaignId string    `json:"whitelist_campaign_id"`
	CreatedAt           time.Time `json:"created_at"`
}
