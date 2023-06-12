package model

import "time"

type AirdropCampaign struct {
	Id                    int64      `json:"id"`
	Title                 string     `json:"title"`
	Detail                string     `json:"detail"`
	PrevAirdropCampaignId *int       `json:"prev_airdrop_campaign_id,omitempty"`
	CreatedAt             time.Time  `json:"created_at"`
	UpdatedAt             time.Time  `json:"updated_at"`
	DeadlineAt            *time.Time `json:"deadline_at,omitempty"`
}
