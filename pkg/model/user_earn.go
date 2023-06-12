package model

type ProfileAirdropCampaign struct {
	Id                int64  `json:"id"`
	ProfileId         string `json:"profile_id"`
	AirdropCampaignId int    `json:"airdrop_campaign_id"`
	Status            string `json:"status"`
	IsFavorite        bool   `json:"is_favorite"`

	AirdropCampaign *AirdropCampaign `json:"airdrop_campaign,omitempty" gorm:"ForeignKey:AirdropCampaignId;references:Id"`
}
