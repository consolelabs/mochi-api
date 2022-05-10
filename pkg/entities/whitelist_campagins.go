package entities

import (
	"errors"

	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"gorm.io/gorm"
	"time"
)

func (e *Entity) CreateWhitelistCampaign(req request.CreateWhitelistCampaignRequest) error {
	err := e.repo.WhitelistCampaigns.CreateIfNotExists(model.WhitelistCampaign{
		Name:      req.Name,
		GuildID:   req.GuildID,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return err
	}

	return nil
}

func (e *Entity) GetWhitelistCampaigns() ([]model.WhitelistCampaign, error) {
	data, err := e.repo.WhitelistCampaigns.Gets()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (e *Entity) GetWhitelistCampaign(campaignId string) (*model.WhitelistCampaign, error) {
	data, err := e.repo.WhitelistCampaigns.GetByID(campaignId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}

	return data, nil
}

func (e *Entity) AddWhitelistCampaignUsers(req request.AddWhitelistCampaignUserRequest) error {
	for _, r := range req.Users {
		err := e.repo.WhitelistCampaignUsers.UpsertOne(model.WhitelistCampaignUser{
			Address:             r.Address,
			DiscordID:           r.DiscordID,
			Notes:               r.Notes,
			WhitelistCampaignId: r.WhitelistCampaignId,
			CreatedAt:           time.Now(),
		})
		if err != nil {
			return err
		}

	}

	return nil
}

func (e *Entity) GetWhitelistCampaignUsers() ([]model.WhitelistCampaignUser, error) {
	data, err := e.repo.WhitelistCampaignUsers.Gets()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}

	return data, nil
}

func (e *Entity) GetWhitelistCampaignUser(campaignId, address string) (*model.WhitelistCampaignUser, error) {
	data, err := e.repo.WhitelistCampaignUsers.GetByCampaignIdAddress(campaignId, address)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}

	return data, nil
}
