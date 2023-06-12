package profileairdropcampaign

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	UpsertOne(*model.ProfileAirdropCampaign) (*model.ProfileAirdropCampaign, error)
	List(ListQuery) ([]model.ProfileAirdropCampaign, int64, error)
	Delete(*model.ProfileAirdropCampaign) (*model.ProfileAirdropCampaign, error)
}
