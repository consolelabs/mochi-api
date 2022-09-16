package entities

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
)

func (e *Entity) GetTelegramByDiscordID(telegramID string) (*response.GetLinkedTelegramResponse, error) {
	data, err := e.repo.UserTelegramDiscordAssociation.GetOneByTelegramID(telegramID)
	if err != nil {
		e.log.Error(err, "[entity.GetTelegramByDiscordID] repo.UserTelegramDiscordAssociation.GetOneByTelegramID() failed")
		return nil, err
	}
	return &response.GetLinkedTelegramResponse{
		Data: data,
	}, nil
}

func (e *Entity) LinkUserTelegramWithDiscord(req request.LinkUserTelegramWithDiscordRequest) (*response.LinkUserTelegramWithDiscordResponse, error) {
	dcUser, err := e.discord.User(req.DiscordID)
	if err != nil {
		e.log.Error(err, "[entity.LinkUserTelegramWithDiscord] discord.User() failed")
		return nil, err
	}

	model := model.UserTelegramDiscordAssociation{
		DiscordID:  req.DiscordID,
		TelegramID: req.TelegramID,
	}
	err = e.repo.UserTelegramDiscordAssociation.Upsert(&model)
	if err != nil {
		e.log.Error(err, "[entity.LinkUserTelegramWithDiscord] repo.UserTelegramDiscordAssociation.Upsert() failed")
		return nil, err
	}
	return &response.LinkUserTelegramWithDiscordResponse{
		Data: &response.LinkUserTelegramWithDiscordResponseData{
			DiscordUsername:                dcUser.Username,
			UserTelegramDiscordAssociation: model,
		},
	}, nil
}
