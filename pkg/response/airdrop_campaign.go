package response

import (
	"github.com/defipod/mochi/pkg/model"
)

type AirdropCampaignResponse struct {
	Data *model.AirdropCampaign `json:"data"`
}

type AirdropCampaignsResponse struct {
	Data  []model.AirdropCampaign `json:"data"`
	Page  int                     `json:"page"`
	Size  int                     `json:"size"`
	Total int64                   `json:"total"`
}

type ProfileAirdropCampaignResponse struct {
	Data *model.ProfileAirdropCampaign `json:"data"`
}

type ProfileAirdropCampaignsResponse struct {
	Data  []model.ProfileAirdropCampaign `json:"data"`
	Page  int                            `json:"page"`
	Size  int                            `json:"size"`
	Total int64                          `json:"total"`
}
