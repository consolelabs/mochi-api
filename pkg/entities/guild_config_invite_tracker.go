package entities

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
)

func (e *Entity) CreateOrUpdateInviteTrackerLogChannel(req request.ConfigureInviteRequest) error {
	config := model.GuildConfigInviteTracker{
		GuildID:    req.GuildID,
		ChannelID:  req.LogChannel,
		WebhookURL: model.JSONNullString{},
	}

	return e.repo.GuildConfigInviteTracker.Upsert(&config)
}
