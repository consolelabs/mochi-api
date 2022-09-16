package entities

import (
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/model"
	baseerrs "github.com/defipod/mochi/pkg/model/errors"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
)

func (e *Entity) GetByTelegramUsername(TelegramUsername string) (*response.GetLinkedTelegramResponse, error) {
	data, err := e.repo.UserTelegramDiscordAssociation.GetOneByTelegramUsername(TelegramUsername)
	if err != nil && err != gorm.ErrRecordNotFound {
		e.log.Error(err, "[entity.GetByTelegramUsername] repo.UserTelegramDiscordAssociation.GetOneByTelegramUsername() failed")
		return nil, err
	}
	if err != nil {
		e.log.Error(err, "[entity.GetByTelegramUsername] no data found")
		return nil, baseerrs.ErrRecordNotFound
	}
	return &response.GetLinkedTelegramResponse{
		Data: data,
	}, nil
}

func (e *Entity) LinkUserTelegramWithDiscord(req request.LinkUserTelegramWithDiscordRequest) (*response.LinkUserTelegramWithDiscordResponse, error) {
	model := model.UserTelegramDiscordAssociation{
		DiscordID:        req.DiscordID,
		TelegramUsername: req.TelegramUsername,
	}
	err := e.repo.UserTelegramDiscordAssociation.Upsert(&model)
	if err != nil {
		e.log.Error(err, "[entity.LinkUserTelegramWithDiscord] repo.UserTelegramDiscordAssociation.Upsert() failed")
		return nil, err
	}
	return &response.LinkUserTelegramWithDiscordResponse{
		Data: nil,
	}, nil
}
