package entities

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
)

func (e *Entity) GetInviteTrackerLogChannel(guildID string) (*model.GuildConfigInviteTracker, error) {
	config, err := e.repo.GuildConfigInviteTracker.GetOne(guildID)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (e *Entity) CreateOrUpdateInviteTrackerLogChannel(req request.ConfigureInviteRequest) error {
	config := model.GuildConfigInviteTracker{
		GuildID:    req.GuildID,
		ChannelID:  req.LogChannel,
		WebhookURL: model.JSONNullString{},
	}

	return e.repo.GuildConfigInviteTracker.Upsert(&config)
}
