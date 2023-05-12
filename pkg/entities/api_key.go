package entities

import (
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/response"
)

func (e *Entity) CreateApiKey(profileAccessToken string) (*response.ProfileApiKeyResponse, error) {
	result, err := e.svc.MochiProfile.CreateProfileApiKey(profileAccessToken)
	if err != nil {
		e.log.Fields(logger.Fields{"profileAccessToken": profileAccessToken}).Error(err, "[mochi-profile.CreateProfileApiKey] - cannot create profile apiKey")
		return nil, err
	}

	return &response.ProfileApiKeyResponse{
		ProfileId: result.ProfileId,
		ApiKey:    result.ApiKey,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}, nil
}
