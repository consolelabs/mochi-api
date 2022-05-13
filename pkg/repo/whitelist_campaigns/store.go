package whitelist_campaigns

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetByGuildId(guildId string) ([]model.WhitelistCampaign, error)
	GetByID(id string) (*model.WhitelistCampaign, error)
	CreateIfNotExists(guild model.WhitelistCampaign) error
}
