package entities

import (
	"errors"

	"github.com/defipod/mochi/pkg/model"
	"gorm.io/gorm"
)

type GetGuildsResponse struct {
	Data []*GetGuildResponse `json:"data"`
}

type GetGuildResponse struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	BotScopes    []string `json:"bot_scopes"`
	Alias        string   `json:"alias"`
	LogChannelID string   `json:"log_channel_id"`
}

type CreateGuildRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (e *Entity) CreateGuild(guild CreateGuildRequest) error {
	return e.repo.DiscordGuilds.CreateIfNotExists(model.DiscordGuild{
		ID:   guild.ID,
		Name: guild.Name,
		BotScopes: model.JSONArrayString{
			"*",
		},
	})
}

func (e *Entity) GetGuilds() (*GetGuildsResponse, error) {
	guilds, err := e.repo.DiscordGuilds.Gets()
	if err != nil {
		return nil, err
	}

	var response GetGuildsResponse
	response.Data = make([]*GetGuildResponse, 0)
	for _, g := range guilds {
		response.Data = append(response.Data, &GetGuildResponse{
			ID:           g.ID,
			Name:         g.Name,
			BotScopes:    g.BotScopes,
			Alias:        g.Alias,
			LogChannelID: g.GuildConfigInviteTracker.ChannelID,
		})
	}

	return &response, nil
}

func (e *Entity) GetGuild(guildID string) (*GetGuildResponse, error) {
	guild, err := e.repo.DiscordGuilds.GetByID(guildID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}

	return &GetGuildResponse{
		ID:           guild.ID,
		Name:         guild.Name,
		BotScopes:    guild.BotScopes,
		Alias:        guild.Alias,
		LogChannelID: guild.GuildConfigInviteTracker.ChannelID,
	}, nil
}
