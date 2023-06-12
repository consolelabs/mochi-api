package request

import (
	"errors"
	"strings"
	"time"
)

const (
	StatusNew     = "new"
	StatusSkipped = "skipped"
	StatusDone    = "done"
	StatusSuccess = "success"
	StatusFailure = "failure"
)

var validStatuses = map[string]bool{
	StatusNew:     true,
	StatusSkipped: true,
	StatusDone:    true,
	StatusSuccess: true,
	StatusFailure: true,
}

type CreateAirdropCampaignRequest struct {
	Title                 string     `json:"title,omitempty"`
	Detail                string     `json:"detail,omitempty"`
	PrevAirdropCampaignId *int       `json:"prev_airdrop_campaign_id,omitempty"`
	DeadlineAt            *time.Time `json:"deadline_at,omitempty"`
}

type CreateProfileAirdropCampaignRequest struct {
	ProfileId         string `json:"-"`
	AirdropCampaignId int    `json:"airdrop_campaign_id"`
	Status            string `json:"status"`
	IsFavorite        bool   `json:"is_favorite"`
}

func (r *CreateProfileAirdropCampaignRequest) Validate() error {
	if r.ProfileId == "" {
		return errors.New("invalid profile_id")
	}
	if _, ok := validStatuses[strings.ToLower(r.Status)]; !ok {
		r.Status = StatusNew
	}

	return nil
}

type RemoveProfileAirdropCampaignRequest struct {
	ProfileId         string `form:"profile_id" json:"profile_id"`
	AirdropCampaignId string `form:"airdrop_campaign_id" json:"airdrop_campaign_id"`
}

type GetProfileAirdropCampaignsRequest struct {
	ProfileId  string `form:"-" json:"-"`
	Status     string `form:"status" json:"status"`
	IsFavorite *bool  `form:"is_favorite" json:"is_favorite"`
	PaginationRequest
}

func (r *GetProfileAirdropCampaignsRequest) Validate() error {
	if r.ProfileId == "" {
		return errors.New("invalid user_id")
	}

	if r.Status != "" {
		if _, ok := validStatuses[strings.ToLower(r.Status)]; !ok {
			return errors.New("invalid status")
		}
	}

	return nil
}

type DeleteProfileAirdropCampaignRequest struct {
	ProfileId         string `uri:"id" binding:"required"`
	AirdropCampaignId int    `uri:"airdrop_campaign_id" binding:"required"`
}

func (r *DeleteProfileAirdropCampaignRequest) Validate() error {
	if r.ProfileId == "" {
		return errors.New("invalid profile id")
	}
	if r.AirdropCampaignId <= 0 {
		return errors.New("invalid airdrop campaign id")
	}
	return nil
}
