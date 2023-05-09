package entities

import (
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/response"
)

func (e *Entity) GetApiKeyByDiscordId(discordId string) (*response.ProfileApiKeyResponse, error) {
	profile, err := e.svc.MochiProfile.GetByDiscordID(discordId)
	if err != nil {
		e.log.Fields(logger.Fields{"discordId": discordId}).Error(err, "[mochi-profile.GetByDiscordID] - cannot get profile")
		return nil, err
	}

	result, err := e.svc.MochiProfile.GetApiKeyByProfileID(profile.ID)
	if err != nil {
		e.log.Fields(logger.Fields{"discordId": discordId}).Error(err, "[mochi-profile.GetByDiscordID] - cannot get profile")
		return nil, err
	}

	return &response.ProfileApiKeyResponse{
		ProfileId: result.ProfileId,
		ApiKey:    result.ApiKey,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}, nil
}

func (e *Entity) CreateApiKeyByDiscordId(discordId string) (*response.ProfileApiKeyResponse, error) {
	profile, err := e.svc.MochiProfile.GetByDiscordID(discordId)
	if err != nil {
		e.log.Fields(logger.Fields{"discordId": discordId}).Error(err, "[mochi-profile.GetByDiscordID] - cannot get profile")
		return nil, err
	}

	result, err := e.svc.MochiProfile.CreateProfileApiKey(profile.ID)
	if err != nil {
		e.log.Fields(logger.Fields{"discordId": discordId}).Error(err, "[mochi-profile.CreateProfileApiKey] - cannot create profile apiKey")
		return nil, err
	}

	return &response.ProfileApiKeyResponse{
		ProfileId: result.ProfileId,
		ApiKey:    result.ApiKey,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}, nil
}
