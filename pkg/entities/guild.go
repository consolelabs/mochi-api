package entities

import (
	"errors"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	baseerrs "github.com/defipod/mochi/pkg/model/errors"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
)

func (e *Entity) CreateGuild(guild request.CreateGuildRequest) error {
	err := e.repo.DiscordGuilds.CreateOrReactivate(model.DiscordGuild{
		ID:   guild.ID,
		Name: guild.Name,
		BotScopes: model.JSONArrayString{
			"*",
		},
		Active: true,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"guildID": guild.ID}).Errorf(err, "[e.CreateGuild] repo.DiscordGuilds.CreateOrReactivate() failed")
		return err
	}

	guilds, err := e.repo.DiscordGuilds.Gets()
	if err != nil {
		e.log.Errorf(err, "[e.CreateGuild] failed to get all guilds")
		return err
	}

	// notifiy new guild to discord
	err = e.svc.Discord.NotifyNewGuild(guild.ID, len(guilds))
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
			LogChannelID: g.GuildConfigInviteTracker.ChannelID, // TODO: refactor (rename)
			LogChannel:   g.LogChannel,
			Active:       g.Active,
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
		LogChannel:   guild.LogChannel,
		LogChannelID: guild.GuildConfigInviteTracker.ChannelID, // TODO: refactor (rename)
		GlobalXP:     guild.GlobalXP,
		Active:       true,
	}, nil
}

func listDiscordGuilds(s *discordgo.Session) ([]*discordgo.UserGuild, error) {

	var (
		guilds  []*discordgo.UserGuild
		afterID string
	)

	for {
		tmp, err := s.UserGuilds(100, "", afterID)
		if err != nil {
			return nil, err
		}

		afterID = tmp[len(tmp)-1].ID
		guilds = append(guilds, tmp...)

		if len(tmp) < 100 {
			break
		}
	}

	return guilds, nil
}

func (e *Entity) ListMyDiscordGuilds(accessToken string) ([]response.DiscordGuildResponse, error) {
	s, err := discordgo.New("Bearer " + accessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to open discord session: %v", err.Error())
	}

	userGuilds, err := listDiscordGuilds(s)
	if err != nil {
		return nil, fmt.Errorf("failed to list user's discord guilds: %v", err.Error())
	}

	mochiGuilds, err := listDiscordGuilds(e.discord)
	if err != nil {
		return nil, fmt.Errorf("failed to list mochi's discord guilds: %v", err.Error())
	}

	mochiArrived := make(map[string]bool)

	for _, g := range mochiGuilds {
		mochiArrived[g.ID] = true
	}

	guilds := make([]response.DiscordGuildResponse, 0)
	for _, g := range userGuilds {
		// Check for guilds that user has ADMINISTRATOR or MANAGE_GUILD permission
		if (g.Permissions&0x8) == 0x8 || (g.Permissions&0x20) == 0x20 {
			guilds = append(guilds, response.DiscordGuildResponse{*g, true, mochiArrived[g.ID]})
		}
	}

	return guilds, nil
}

func (e *Entity) UpdateGuild(guildID string, req request.UpdateGuildRequest) error {
	guild, err := e.repo.DiscordGuilds.GetByID(guildID)
	if err == gorm.ErrRecordNotFound {
		return baseerrs.ErrRecordNotFound
	}
	if err != nil {
		return err
	}
	if req.GlobalXP != nil {
		guild.GlobalXP = *req.GlobalXP
	}
	if req.LogChannel != nil {
		guild.LogChannel = *req.LogChannel
	}
	if req.Active != nil {
		guild.Active = *req.Active
	}
	if err := e.repo.DiscordGuilds.Update(guild); err != nil {
		e.log.Errorf(err, "failed to update guild %s", guildID)
		return err
	}
	return nil
}

func (e *Entity) DeactivateGuild(req request.HandleGuildDeleteRequest) error {
	active := false
	err := e.UpdateGuild(req.GuildID, request.UpdateGuildRequest{Active: &active})
	e.sendGuildDeactivationLog(req.GuildID, req.GuildName, req.IconURL)
	return err
}

func (e *Entity) sendGuildDeactivationLog(guildID, guildName, iconURL string) {
	guilds, err := e.GetGuilds()
	var guildsLeft int
	if err != nil {
		e.log.Errorf(err, "e.GetGuilds() failed")
		guildsLeft = -1
	} else {
		guildsLeft = len(guilds.Data)
	}
	err = e.svc.Discord.NotifyGuildDelete(guildID, guildName, iconURL, guildsLeft)
	if err != nil {
		e.log.Fields(logger.Fields{"guildID": guildID, "guildName": guildName, "iconURL": iconURL}).Errorf(err, "svc.Discord.NotifyGuildDelete() failed")
	}
}
