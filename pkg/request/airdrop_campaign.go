package request

import (
	"errors"
	"strings"
	"time"
)

const (
	StatusLive      = "live"
	StatusEnded     = "ended"
	StatusClaimable = "claimable"
	StatusCancelled = "cancelled"

	StatusIgnored     = "ignored"
	StatusJoined      = "joined"
	StatusClaimed     = "claimed"
	StatusNotEligible = "not_eligible"
)

var validAirdropCampaignStatuses = map[string]bool{
	StatusLive:      true,
	StatusEnded:     true,
	StatusClaimable: true,
	StatusCancelled: true,
}

var validProfileAirdropCampaignStatuses = map[string]bool{
	StatusIgnored:     true,
	StatusJoined:      true,
	StatusClaimed:     true,
	StatusNotEligible: true,
}

type GetAirdropCampaignRequest struct {
	Id        int64  `uri:"id" binding:"required"`
	ProfileId string `form:"profile_id"`
}

func (r *GetAirdropCampaignRequest) Validate() error {
	if r.Id <= 0 {
		return errors.New("invalid id")
	}
	return nil
}

type GetAirdropCampaignsRequest struct {
	ProfileId string `form:"profile_id" json:"profile_id"`
	Status    string `form:"status" json:"status"`
	PaginationRequest
}

type CreateAirdropCampaignRequest struct {
	Id                    *int64     `json:"id,omitempty"`
	Title                 string     `json:"title,omitempty"`
	Detail                string     `json:"detail,omitempty"`
	PrevAirdropCampaignId *int       `json:"prev_airdrop_campaign_id,omitempty"`
	DeadlineAt            *time.Time `json:"deadline_at,omitempty"`
	Status                string     `json:"status,omitempty"`
	RewardAmount          int        `json:"reward_amount,omitempty"`
	RewardTokenSymbol     string     `json:"reward_token_symbol,omitempty"`
}

func (r *CreateAirdropCampaignRequest) Validate() error {
	if r.Id != nil && *r.Id <= 0 {
		return errors.New("invalid id")
	}

	if _, ok := validAirdropCampaignStatuses[strings.ToLower(r.Status)]; !ok {
		return errors.New("invalid status")
	}

	return nil
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

	if _, ok := validProfileAirdropCampaignStatuses[strings.ToLower(r.Status)]; !ok {
		return errors.New("invalid status")
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
		if _, ok := validProfileAirdropCampaignStatuses[strings.ToLower(r.Status)]; !ok {
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

type GetAirdropCampaignStatus struct {
	ProfileId string `form:"profile_id"`
	Status    string `form:"status"`
}

func (r *GetAirdropCampaignStatus) Validate() error {
	if r.ProfileId == "" {
		return errors.New("invalid user_id")
	}

	if r.Status == "" {
		r.Status = StatusIgnored
	}

	return nil
}
