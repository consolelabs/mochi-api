package entities

import (
	"fmt"
	"net/http"

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
		return nil, http.StatusInternalServerError, fmt.Errorf("failed to get user's gm streak: %v", err)
	}

	if err == gorm.ErrRecordNotFound {
		return nil, http.StatusBadRequest, fmt.Errorf("user has no gm streak")
	}

	return streak, http.StatusOK, nil
}

func (e *Entity) HandleUserActivities(req *request.HandleUserActivityRequest) (*response.HandleUserActivityResponse, error) {
	userXP, err := e.repo.GuildUserXP.GetOne(req.GuildID, req.UserID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	gConfigActivity, err := e.repo.GuildConfigActivity.GetOneByActivityName(req.GuildID, req.Action)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
		if err := e.InitGuildDefaultActivityConfigs(req.GuildID); err != nil {
			return nil, err
		}
		return e.HandleUserActivities(req)
	}

	if err := e.repo.GuildUserActivityLog.CreateOne(model.GuildUserActivityLog{
		GuildID:      req.GuildID,
		UserID:       req.UserID,
		ActivityName: gConfigActivity.Activity.Name,
		EarnedXP:     gConfigActivity.Activity.XP,
		CreatedAt:    req.Timestamp,
	}); err != nil {
		return nil, err
	}

	latestUserXP, err := e.repo.GuildUserXP.GetOne(req.GuildID, req.UserID)
	if err != nil {
		return nil, err
	}

	return &response.HandleUserActivityResponse{
		GuildID:      req.GuildID,
		UserID:       req.UserID,
		Action:       gConfigActivity.Activity.Name,
		AddedXP:      gConfigActivity.Activity.XP,
		CurrentXP:    latestUserXP.TotalXP,
		CurrentLevel: latestUserXP.Level,
		Timestamp:    req.Timestamp,
		LevelUp:      latestUserXP.Level > userXP.Level,
	}, nil
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
	top, err := e.repo.GuildUserXP.GetTopUsers(guildID, limit, offset)
	if err != nil {
		return nil, err
	}

	author, err := e.repo.GuildUserXP.GetOne(guildID, userID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &response.GetTopUsersResponse{
		Author:      author,
		Leaderboard: top,
	}, nil
}
