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
