package airdropcampaign

import (
	"github.com/defipod/mochi/pkg/model"
)

type Store interface {
	Create(*model.AirdropCampaign) (*model.AirdropCampaign, error)
	GetById(int64) (*model.AirdropCampaign, error)
	List(ListQuery) ([]model.AirdropCampaign, int64, error)
}
