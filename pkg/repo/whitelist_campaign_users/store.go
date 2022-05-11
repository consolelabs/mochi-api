package whitelist_campaign_users

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	Gets() ([]model.WhitelistCampaignUser, error)
	GetByCampaignIdAddress(campaignId, address string) (*model.WhitelistCampaignUser, error)
	UpsertOne(model.WhitelistCampaignUser) error
}
