package entities

import (
	"fmt"
	"math"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"gorm.io/gorm"
)

func (e *Entity) CreateUser(req request.CreateUserRequest) error {

	user := &model.User{
		ID:       req.ID,
		Username: req.Username,
		GuildUsers: []*model.GuildUser{
			{
				GuildID:   req.GuildID,
				UserID:    req.ID,
				Nickname:  req.Nickname,
				InvitedBy: req.InvitedBy,
			},
		},
	}

	if err := e.repo.Users.Create(user); err != nil {
		e.log.Errorf(err, "[e.repo.Users.Create] failed to create user")
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (e *Entity) CreateUserIfNotExists(id, username string) error {
	user := &model.User{
		ID:       id,
		Username: username,
	}

	if err := e.repo.Users.FirstOrCreate(user); err != nil {
		e.log.Errorf(err, "[e.repo.Users.FirstOrCreate] failed to create if not exists user")
		return fmt.Errorf("failed to create if not exists user: %w", err)
	}

	return nil
}

func (e *Entity) GetUser(discordID string) (*response.GetUserResponse, error) {
	user, err := e.repo.Users.GetOne(discordID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrRecordNotFound
		}
		e.log.Errorf(err, "[e.repo.Users.GetOne] failed to get user")
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	guildUsers := []*response.GetGuildUserResponse{}
	for _, guildUser := range user.GuildUsers {
		guildUsers = append(guildUsers, &response.GetGuildUserResponse{
			GuildID:   guildUser.GuildID,
			UserID:    guildUser.UserID,
			Nickname:  guildUser.Nickname,
			InvitedBy: guildUser.InvitedBy,
		})
	}

	if user.InDiscordWalletAddress.String == "" {
		if err = e.generateInDiscordWallet(user); err != nil {
			e.log.Errorf(err, "[e.generateInDiscordWallet] cannot generate in-discord wallet")
			err = fmt.Errorf("cannot generate in-discord wallet: %v", err)
			return nil, err
		}
	}

	res := &response.GetUserResponse{
		ID:                     user.ID,
		Username:               user.Username,
		InDiscordWalletAddress: &user.InDiscordWalletAddress.String,
		InDiscordWalletNumber:  &user.InDiscordWalletNumber.Int64,
		GuildUsers:             guildUsers,
	}
	return res, nil
}

func (e *Entity) GetUserCurrentGMStreak(discordID, guildID string) (*model.DiscordUserGMStreak, int, error) {
	streak, err := e.repo.DiscordUserGMStreak.GetByDiscordIDGuildID(discordID, guildID)
	if err != nil && err != gorm.ErrRecordNotFound {
		e.log.Errorf(err, "[e.repo.DiscordUserGMStreak.GetByDiscordIDGuildID] cannot get user's gm streak by discordID %s and guildID %s", discordID, guildID)
		return nil, http.StatusInternalServerError, fmt.Errorf("failed to get user's gm streak: %v", err)
	}

	if err == gorm.ErrRecordNotFound {
		e.log.Errorf(err, "[e.repo.DiscordUserGMStreak.GetByDiscordIDGuildID] user has no gm streak")
		return nil, http.StatusBadRequest, fmt.Errorf("user has no gm streak")
	}

	return streak, http.StatusOK, nil
}

func (e *Entity) HandleUserActivities(req *request.HandleUserActivityRequest) (*response.HandleUserActivityResponse, error) {
	userXP, err := e.repo.GuildUserXP.GetOne(req.GuildID, req.UserID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	earnedXP := int(req.CustomXP)
	if earnedXP == 0 {
		gActivityConfig, err := e.GetGuildActivityConfig(req.GuildID, req.Action)
		if err != nil {
			e.log.Errorf(err, "[e.GetGuildActivityConfig] cannot get guild config activity by guildID %s and action %s", req.GuildID, req.Action)
			return nil, fmt.Errorf("failed to get guild config activity: %v", err.Error())
		}
		earnedXP = gActivityConfig.Activity.XP
	}

	if err := e.repo.GuildUserActivityLog.CreateOne(model.GuildUserActivityLog{
		GuildID:      req.GuildID,
		UserID:       req.UserID,
		ActivityName: req.Action,
		EarnedXP:     earnedXP,
		CreatedAt:    req.Timestamp,
	}); err != nil {
		return nil, err
	}

	latestUserXP, err := e.repo.GuildUserXP.GetOne(req.GuildID, req.UserID)
	if err != nil {
		return nil, err
	}

	res := &response.HandleUserActivityResponse{
		GuildID:      req.GuildID,
		ChannelID:    req.ChannelID,
		UserID:       req.UserID,
		Action:       req.Action,
		AddedXP:      earnedXP,
		CurrentXP:    latestUserXP.TotalXP,
		CurrentLevel: latestUserXP.Level,
		Timestamp:    req.Timestamp,
		LevelUp:      latestUserXP.Level > userXP.Level,
	}

	role, err := e.GetRoleByGuildLevelConfig(req.GuildID, req.UserID)
	if err != nil {
		e.log.Errorf(err, "[HandleUserActivities] - SendLevelUpMessage failed - guild %s, user %s", req.GuildID, req.UserID)
	} else {
		e.svc.Discord.SendLevelUpMessage(userXP.Guild.LogChannel, role, res)
	}
	return res, nil
}

func (e *Entity) InitGuildDefaultActivityConfigs(guildID string) error {
	activities, err := e.repo.Activity.GetDefaultActivities()
	if err != nil {
		return err
	}

	var configs []model.GuildConfigActivity
	for _, a := range activities {
		configs = append(configs, model.GuildConfigActivity{
			GuildID:    guildID,
			ActivityID: a.ID,
			Active:     true,
		})
	}

	return e.repo.GuildConfigActivity.UpsertMany(configs)
}

func (e *Entity) GetTopUsers(guildID, userID string, limit, page int) (*response.GetTopUsersResponse, error) {
	offset := page * limit
	leaderboard, err := e.repo.GuildUserXP.GetTopUsers(guildID, limit, offset)
	if err != nil {
		return nil, err
	}

	for i := range leaderboard {
		item := &leaderboard[i]
		currentLevel, err := e.repo.ConfigXPLevel.GetNextLevel(item.TotalXP, false)
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, err
		}

		nextLevel, err := e.repo.ConfigXPLevel.GetNextLevel(item.TotalXP, true)
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, err
		}
		item.Progress = math.Min(float64(item.TotalXP-currentLevel.MinXP)/float64(nextLevel.MinXP-currentLevel.MinXP), 1)
		if nextLevel.Level == 0 {
			item.Progress = 1
		}
	}

	author, err := e.repo.GuildUserXP.GetOne(guildID, userID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &response.GetTopUsersResponse{
		Author:      author,
		Leaderboard: leaderboard,
	}, nil
}

func (e *Entity) GetGuildUserXPs(guildID string) ([]model.GuildUserXP, error) {
	return e.repo.GuildUserXP.GetByGuildID(guildID)
}

func (e *Entity) GetGuildMember(guildID, userID string) (*discordgo.Member, error) {
	member, err := e.discord.GuildMember(guildID, userID)
	if err != nil {
		return nil, err
	}
	return member, nil
}

func (e *Entity) ListGuildMembers(guildID string) ([]*discordgo.Member, error) {

	var afterID string

	res := make([]*discordgo.Member, 0)
	for {
		members, err := e.discord.GuildMembers(guildID, afterID, 100)
		if err != nil {
			return nil, err
		}

		res = append(res, members...)

		if len(members) < 100 {
			break
		}
		afterID = members[len(members)-1].User.ID
	}

	return res, nil
}

func (e *Entity) GetUserProfile(guildID, userID string) (*response.GetUserProfileResponse, error) {
	gUserXP, err := e.repo.GuildUserXP.GetOne(guildID, userID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	currentLevel, err := e.repo.ConfigXPLevel.GetNextLevel(gUserXP.TotalXP, false)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	nextLevel, err := e.repo.ConfigXPLevel.GetNextLevel(gUserXP.TotalXP, true)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	progress := math.Min(float64(gUserXP.TotalXP-currentLevel.MinXP)/float64(nextLevel.MinXP-currentLevel.MinXP), 1)
	if nextLevel.Level == 0 {
		progress = 1
	}

	if gUserXP.Guild == nil {
		if gUserXP.Guild, err = e.repo.DiscordGuilds.GetByID(guildID); err != nil {
			return nil, err
		}
	}

	userWallet, err := e.repo.UserWallet.GetOneByDiscordIDAndGuildID(userID, guildID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &response.GetUserProfileResponse{
		ID:           userID,
		CurrentLevel: currentLevel,
		NextLevel:    nextLevel,
		GuildXP:      gUserXP.TotalXP,
		NrOfActions:  gUserXP.NrOfActions,
		Progress:     progress,
		Guild:        gUserXP.Guild,
		GuildRank:    gUserXP.GuildRank,
		UserWallet:   userWallet,
	}, nil
}

func (e *Entity) SendGiftXp(req request.GiftXPRequest) (*response.HandleUserActivityResponse, error) {
	res, err := e.HandleUserActivities(&request.HandleUserActivityRequest{
		GuildID:   req.GuildID,
		ChannelID: req.ChannelID,
		UserID:    req.UserDiscordID,
		Action:    "gifted",
		CustomXP:  int64(req.XPAmount),
	})
	if err != nil {
		e.log.Errorf(err, "[SendGiftXp] - HandleUserActivities failed: %v %v %v", req.GuildID, req.UserDiscordID, req.XPAmount)
		return nil, err
	}

	return res, nil
}

func (e *Entity) ListAllWalletAddresses() ([]model.WalletAddress, error) {
	was, err := e.repo.UserWallet.ListWalletAddresses("evm")
	if err != nil {
		e.log.Errorf(err, "[e.repo.UserWallet.ListWalletAddresses] cannot get wallet addresses")
		return nil, fmt.Errorf("failed to get wallet addresses: %v", err.Error())
	}
	return was, nil
}

func (e *Entity) GetRoleByGuildLevelConfig(guildID, userID string) (string, error) {
	if e.discord == nil {
		return "", nil
	}
	configs, err := e.repo.GuildConfigLevelRole.GetByGuildID(guildID)
	if err != nil {
		return "", err
	}

	gMember, err := e.discord.GuildMember(guildID, userID)
	if err != nil {
		return "", err
	}
	if gMember.Roles == nil {
		e.log.Errorf(err, "Member %s of guild %s has no role", userID, guildID)
		return "", fmt.Errorf("Member %s of guild %s has no role", userID, guildID)
	}

	for _, cfg := range configs {
		for _, memRole := range gMember.Roles {
			if cfg.RoleID == memRole {
				return fmt.Sprintf("<@&%s>", cfg.RoleID), nil
			}
		}
	}

	return "", nil
}
