package entities

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
)

func (e *Entity) CheckExistTokenConfig(tokenId int, guildID string) (bool, error) {
	listConfigToken, err := e.repo.GuildConfigToken.GetAll()
	if err != nil {
		return false, err
	}

	for i := 0; i < len(listConfigToken); i++ {
		if tokenId == listConfigToken[i].TokenID && guildID == listConfigToken[i].GuildID {
			return true, nil
		}
	}

	return false, nil
}

func (e *Entity) CreateGuildCustomTokenConfig(req request.UpsertCustomTokenConfigRequest) error {
	err := e.repo.GuildConfigToken.CreateOne(model.GuildConfigToken{
		GuildID: req.GuildID,
		TokenID: req.Id,
		Active:  req.Active,
	})
	if err != nil {
		return err
	}

	return nil
}
