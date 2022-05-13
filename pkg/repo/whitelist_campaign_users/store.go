package whitelist_campaign_users

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetByCampaignId(campaignId string) ([]model.WhitelistCampaignUser, error)
	GetByDiscordIdCampaignId(discordId, campaignId string) (*model.WhitelistCampaignUser, error)
	UpsertOne(model.WhitelistCampaignUser) error
}
