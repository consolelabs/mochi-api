package airdropcampaign

import (
	"github.com/defipod/mochi/pkg/model"
)

type Store interface {
	Upsert(*model.AirdropCampaign) (*model.AirdropCampaign, error)
	GetById(int64) (*model.AirdropCampaign, error)
	List(ListQuery) ([]model.AirdropCampaign, int64, error)
	CountStat() (stats []model.AirdropStatusCount, err error)
}
