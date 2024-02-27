package entities

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	baseerrs "github.com/defipod/mochi/pkg/model/errors"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
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

	// guilds, err := e.repo.DiscordGuilds.Gets()
	// if err != nil {
	// 	e.log.Errorf(err, "[e.CreateGuild] failed to get all guilds")
	// 	return err
	// }

	// notifiy new guild to discord
	// err = e.svc.Discord.NotifyNewGuild(guild.ID, len(guilds))
	// if err != nil {
	// 	e.log.Errorf(err, "failed to notify new guild %s to discord", guild.ID)
	// }

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
			ID:         g.ID,
			Name:       g.Name,
			BotScopes:  g.BotScopes,
			Alias:      g.Alias,
			LogChannel: g.LogChannel,
			Active:     g.Active,
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

	discordGuildInfo, err := e.svc.Discord.GetGuild(guildID)
	if err != nil {
		return nil, err
	}

	return &response.GetGuildResponse{
		ID:         guild.ID,
		Name:       guild.Name,
		BotScopes:  guild.BotScopes,
		Alias:      guild.Alias,
		LogChannel: guild.LogChannel,
		GlobalXP:   guild.GlobalXP,
		Active:     true,
		Icon:       discordGuildInfo.Icon,
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
	if req.LeftAt != nil {
		guild.LeftAt = req.LeftAt
	}
	if req.AvailableCMDs != nil {
		jsonData, err := json.Marshal(req.AvailableCMDs)
		if err != nil {
			e.log.Errorf(err, "failed to marshal available cmds")
			return err
		}
		guild.AvailableCMDs.NullString = sql.NullString{String: string(jsonData), Valid: true}
	}
	if err := e.repo.DiscordGuilds.Update(guild); err != nil {
		e.log.Errorf(err, "failed to update guild %s", guildID)
		return err
	}
	return nil
}

func (e *Entity) DeactivateGuild(req request.HandleGuildDeleteRequest) error {
	active := false
	t := time.Now()
	err := e.UpdateGuild(req.GuildID, request.UpdateGuildRequest{Active: &active, LeftAt: &t})
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

func (e *Entity) TotalServers() (*response.Metric, error) {
	discordGuilds, err := e.repo.DiscordGuilds.GetNonLeftGuilds()
	if err != nil {
		e.log.Error(err, "[entities.TotalServers] - cannot get total servers")
		return nil, err
	}

	return &response.Metric{TotalServers: int64(len(discordGuilds))}, nil
}

func (e *Entity) GetGuildRoles(guildID string) (*response.DiscordGuildRoles, error) {
	guildRoles, err := e.svc.Discord.GetGuildRoles(guildID)
	if err != nil {
		return nil, err
	}

	resp := response.DiscordGuildRoles{}

	for _, role := range guildRoles {
		icon := ""
		if role.Icon != "" {
			icon = fmt.Sprintf("https://cdn.discordapp.com/role-icons/%s/%s.png", role.ID, role.Icon)
		}

		resp = append(resp, &response.DiscordGuildRole{
			ID:           role.ID,
			Name:         role.Name,
			Color:        role.Color,
			Hoist:        role.Hoist,
			Icon:         icon,
			UnicodeEmoji: role.UnicodeEmoji,
			Position:     role.Position,
			Permissions:  role.Permissions,
			Managed:      role.Managed,
			Mentionable:  role.Mentionable,
		})
	}

	return &resp, nil
}

func (e *Entity) FetchAndSyncGuilds() ([]*response.GetGuildResponse, error) {
	res, err := e.GetGuilds()
	if err != nil {
		e.log.Error(err, "[entity.FetchAndSyncGuilds] GetGuilds() failed")
		return nil, err
	}

	dbGuilds := make(map[string]*response.GetGuildResponse)
	for _, g := range res.Data {
		dbGuilds[g.ID] = g
	}

	fetchedGuilds := make(map[string]*discordgo.Guild)
	e.discord.StateEnabled = true
	for _, g := range e.discord.State.Guilds {
		fetchedGuilds[g.ID] = g
	}

	result := make([]*response.GetGuildResponse, 0)

	for id, g := range dbGuilds {
		_, ok := fetchedGuilds[id]
		if !ok {
			err = e.DeactivateGuild(request.HandleGuildDeleteRequest{
				GuildID:   id,
				GuildName: g.Name,
				IconURL:   g.Icon,
			})
			if err != nil {
				e.log.Fields(logger.Fields{"guildID": id}).Error(err, "[entity.FetchAndSyncGuilds] DeactivateGuild() failed")
			}
			continue
		}

		result = append(result, g)
	}

	for id, g := range fetchedGuilds {
		dbGuild, ok := dbGuilds[id]
		if !ok || !dbGuild.Active {
			err = e.CreateGuild(request.CreateGuildRequest{
				ID:       id,
				Name:     g.Name,
				JoinedAt: g.JoinedAt,
			})
			if err != nil {
				e.log.Fields(logger.Fields{"guildID": id}).Error(err, "[entity.FetchAndSyncGuilds] DeactivateGuild() failed")
				continue
			}
		}
		result = append(result, dbGuild)
	}

	return result, nil
}

func (e *Entity) CreateGuildIfNotExists(guildID string) error {
	l := e.log.Fields(logger.Fields{"guildID": guildID})

	_, err := e.repo.DiscordGuilds.GetByID(guildID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		l.Error(err, "[entity.CreateGuildIfNotExists] repo.DiscordGuilds.GetByID() failed")
		return err
	}

	if err == nil {
		return nil
	}

	g, err := e.svc.Discord.GetGuild(guildID)
	if err != nil {
		l.Error(err, "[entity.CreateGuildIfNotExists] svc.Discord.GetGuild() failed")
		return err
	}

	err = e.CreateGuild(request.CreateGuildRequest{
		ID:       guildID,
		Name:     g.Name,
		JoinedAt: g.JoinedAt,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"guildID": guildID}).Error(err, "[entity.CreateGuildIfNotExists] DeactivateGuild() failed")
		return err
	}

	return nil
}

func (e *Entity) ValidateUser(ids []string, guildId string) ([]string, error) {
	members, err := e.repo.GuildUsers.GetUsersOfGuild(ids, guildId)
	if err != nil {
		e.log.Errorf(err, "[entity.ValidateUser] failed to get users of guild")
		return nil, nil
	}

	res := make([]string, 0)
	for _, member := range members {
		res = append(res, member.UserID)
	}

	return res, nil
}

func (e *Entity) GuildReportRoles(guildId string) (*response.GuildReportRoles, error) {
	guildInfo, err := e.svc.Discord.GuildWithCounts(guildId)
	if err != nil {
		e.log.Errorf(err, "[entity.Statistic] cannot get guild info from Discord")
		return nil, err
	}

	// Discord API not count number of members in each role
	// - need to get all guild member to see what roles they have and count
	// - only allow 1000 members per request, so need to loop until all members are counted
	after := ""
	limit := 1000
	countRole := make(map[string]int64, 0)

	for {
		guildMembers, err := e.discord.GuildMembers(guildId, after, limit)
		if err != nil {
			return nil, err
		}
		for _, member := range guildMembers {
			for _, role := range member.Roles {
				_, ok := countRole[role]
				if !ok {
					countRole[role] = 1
				} else {
					countRole[role] = countRole[role] + 1
				}

			}
		}

		if len(guildMembers) < limit {
			break
		}
		after = guildMembers[len(guildMembers)-1].User.ID
	}

	// mapping to response
	guildReportRoles := make([]response.GuildReportRoleDetail, 0)
	for _, role := range guildInfo.Roles {
		if role.ID != guildId {
			// change percentage: temp random value until implement database logic
			rand.Seed(time.Now().UnixNano())
			changePercentage := util.RandFloats(-100.0, 100.0)
			guildReportRoles = append(guildReportRoles, response.GuildReportRoleDetail{
				Id:               role.ID,
				Name:             role.Name,
				NrOfMember:       countRole[role.ID],
				ChangePercentage: changePercentage,
			})
		}
	}

	return &response.GuildReportRoles{
		Id:          guildInfo.ID,
		Name:        guildInfo.Name,
		LastUpdated: time.Now(),
		Roles:       guildReportRoles,
	}, nil
}

func (e *Entity) GuildReportMembers(guildId string) (*response.GuildReportMembers, error) {
	guildInfo, err := e.svc.Discord.GuildWithCounts(guildId)
	if err != nil {
		e.log.Errorf(err, "[entity.Statistic] cannot get guild info from Discord")
		return nil, err
	}

	rand.Seed(time.Now().UnixNano())
	changePercentage := util.RandFloats(-100.0, 100.0)

	return &response.GuildReportMembers{
		Id:               guildInfo.ID,
		Name:             guildInfo.Name,
		NrOfMember:       int64(guildInfo.ApproximateMemberCount),
		ChangePercentage: changePercentage,
		LastUpdated:      time.Now(),
	}, nil

}
