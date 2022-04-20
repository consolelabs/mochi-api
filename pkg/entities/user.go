package entities

import (
	"database/sql"
	"fmt"

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
				GuildID:  req.GuildID,
				UserID:   req.ID,
				Nickname: model.JSONNullString{NullString: sql.NullString{String: *req.Nickname, Valid: *req.Nickname != ""}},
			},
		},
	}

	if err := e.repo.Users.Create(user); err != nil {
		return fmt.Errorf("failed to create user: %w", err)
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
			Nickname:  &guildUser.Nickname.String,
			InvitedBy: &guildUser.InvitedBy.String,
		})
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
