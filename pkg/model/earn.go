package model

import "time"

type AirdropCampaign struct {
	Id                    *int64     `json:"id"`
	Title                 string     `json:"title"`
	Detail                string     `json:"detail"`
	PrevAirdropCampaignId *int       `json:"prev_airdrop_campaign_id,omitempty"`
	Status                string     `json:"status"`
	RewardAmount          int        `json:"reward_amount"`
	RewardTokenSymbol     string     `json:"reward_token_symbol"`
	CreatedAt             time.Time  `json:"created_at"`
	UpdatedAt             time.Time  `json:"updated_at"`
	DeadlineAt            *time.Time `json:"deadline_at,omitempty"`

	ProfileCampaignStatus string `json:"profile_campaign_status,omitempty" gorm:"-"'`
}

type AirdropStatusCount struct {
	Status string `json:"status"`
	Count  int64  `json:"count"`
}
