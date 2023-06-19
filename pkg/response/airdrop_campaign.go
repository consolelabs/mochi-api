package response

import (
	"github.com/defipod/mochi/pkg/model"
)

type AirdropCampaignResponse struct {
	Data *model.AirdropCampaign `json:"data"`
}

type AirdropCampaignsResponse struct {
	Data []model.AirdropCampaign `json:"data"`
}

type ProfileAirdropCampaignResponse struct {
	Data *model.ProfileAirdropCampaign `json:"data"`
}

type ProfileAirdropCampaignsResponse struct {
	Data []model.ProfileAirdropCampaign `json:"data"`
}

type AirdropCampaignStatResponse struct {
	Data []model.AirdropStatusCount `json:"data"`
}
