package entities

import (
	"time"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	ac "github.com/defipod/mochi/pkg/repo/airdrop_campaign"
	pac "github.com/defipod/mochi/pkg/repo/profile_airdrop_campaign"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
)

func (e *Entity) CreateAirdropCampaign(req *request.CreateAirdropCampaignRequest) (*model.AirdropCampaign, error) {
	ac := model.AirdropCampaign{
		Title:             req.Title,
		Detail:            req.Detail,
		RewardAmount:      req.RewardAmount,
		RewardTokenSymbol: req.RewardTokenSymbol,
		Status:            req.Status,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
		DeadlineAt:        req.DeadlineAt,
	}

	if req.PrevAirdropCampaignId != nil {
		ac.PrevAirdropCampaignId = req.PrevAirdropCampaignId
	}

	if req.DeadlineAt != nil {
		ac.DeadlineAt = req.DeadlineAt
	}

	if req.Id != nil {
		ac.Id = req.Id
	}

	earn, err := e.repo.AirdropCampaign.Upsert(&ac)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateAirdropCampaign] - e.repo.AirdropCampaign.Create failed")
		return nil, err
	}

	return earn, nil
}

func (e *Entity) GetAirdropCampaigns(req request.GetAirdropCampaignsRequest) (*response.AirdropCampaignsResponse, error) {
	acs, total, err := e.repo.AirdropCampaign.List(ac.ListQuery{
		Offset: int(req.Page * req.Size),
		Limit:  int(req.Size),
		Status: req.Status,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.GetAirdropCampaigns] - e.repo.AirdropCampaign.List failed")
		return nil, err
	}

	return &response.AirdropCampaignsResponse{
		Data:  acs,
		Page:  int(req.Page),
		Size:  int(req.Size),
		Total: total,
	}, nil
}

func (e *Entity) GetAirdropCampaignStats(req request.GetAirdropCampaignStatus) (*response.AirdropCampaignStatResponse, error) {
	resp := &response.AirdropCampaignStatResponse{}
	stats, err := e.repo.AirdropCampaign.CountStat()
	if err != nil {
		return nil, err
	}

	resp.Data = stats

	if req.ProfileId != "" {
		profileStats, err := e.repo.ProfileAirdropCampaign.CountStat(pac.StatQuery{
			ProfileId: req.ProfileId,
			Status:    req.Status,
		})
		if err != nil {
			return nil, err
		}
		resp.Data = append(resp.Data, profileStats...)
	}
	return resp, nil
}

func (e *Entity) CreateProfileAirdropCampaign(req *request.CreateProfileAirdropCampaignRequest) (*model.ProfileAirdropCampaign, error) {
	profileAirdropCampaign, err := e.repo.ProfileAirdropCampaign.UpsertOne(&model.ProfileAirdropCampaign{
		ProfileId:         req.ProfileId,
		AirdropCampaignId: req.AirdropCampaignId,
		Status:            req.Status,
		IsFavorite:        req.IsFavorite,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateProfileAirdropCampaign] - e.repo.ProfileAirdropCampaign.UpsertOne failed")
		return nil, err
	}

	return profileAirdropCampaign, nil
}

func (e *Entity) GetProfileAirdropCampaigns(req request.GetProfileAirdropCampaignsRequest) (*response.ProfileAirdropCampaignsResponse, error) {
	q := pac.ListQuery{
		ProfileId: req.ProfileId,
		Status:    req.Status,
		Limit:     int(req.Size),
		Offset:    int(req.Size * req.Page),
	}

	if req.IsFavorite != nil {
		q.IsFavorite = req.IsFavorite
	}

	acs, total, err := e.repo.ProfileAirdropCampaign.List(q)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.GetProfileAirdropCampaigns] - e.repo.ProfileAirdropCampaign.List failed")
		return nil, err
	}

	return &response.ProfileAirdropCampaignsResponse{
		Data:  acs,
		Page:  int(req.Page),
		Size:  int(req.Size),
		Total: total,
	}, nil

}

func (e *Entity) RemoveProfileAirdropCampaign(req request.DeleteProfileAirdropCampaignRequest) error {
	_, err := e.repo.ProfileAirdropCampaign.Delete(&model.ProfileAirdropCampaign{
		ProfileId:         req.ProfileId,
		AirdropCampaignId: req.AirdropCampaignId,
	})

	return err
}
