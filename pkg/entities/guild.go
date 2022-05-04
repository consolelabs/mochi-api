package entities

import (
	"errors"

	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"gorm.io/gorm"
)

func (e *Entity) CreateGuild(guild request.CreateGuildRequest) error {
	err := e.repo.DiscordGuilds.CreateIfNotExists(model.DiscordGuild{
		ID:   guild.ID,
		Name: guild.Name,
		BotScopes: model.JSONArrayString{
			"*",
		},
	})
	if err != nil {
		return err
	}

	// notifiy new guild to discord
	err = e.svc.Discord.NotifyNewGuild(guild.ID)
	if err != nil {
		e.log.Errorf(err, "failed to notify new guild %s to discord", guild.ID)
	}

	return nil
}

func (e *Entity) GetGuilds() (*response.GetGuildsResponse, error) {
	guilds, err := e.repo.DiscordGuilds.Gets()
	if err != nil {
		return nil, err
	}

	var res response.GetGuildsResponse
	res.Data = make([]*response.GetGuildResponse, 0)
	for _, g := range guilds {
		res.Data = append(res.Data, &response.GetGuildResponse{
			ID:           g.ID,
			Name:         g.Name,
			BotScopes:    g.BotScopes,
			Alias:        g.Alias,
			LogChannelID: g.GuildConfigInviteTracker.ChannelID,
		})
	}

	return &res, nil
}

func (e *Entity) GetGuild(guildID string) (*response.GetGuildResponse, error) {
	guild, err := e.repo.DiscordGuilds.GetByID(guildID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}

	return &response.GetGuildResponse{
		ID:           guild.ID,
		Name:         guild.Name,
		BotScopes:    guild.BotScopes,
		Alias:        guild.Alias,
		LogChannelID: guild.GuildConfigInviteTracker.ChannelID,
	}, nil
}
