package entities

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
)

func (e *Entity) GetGmConfig(guildID string) (*model.GuildConfigGmGn, error) {
	config, err := e.repo.GuildConfigGmGn.GetByGuildID(guildID)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (e *Entity) UpsertGmConfig(req request.UpsertGmConfigRequest) error {
	if err := e.repo.GuildConfigGmGn.UpsertOne(&model.GuildConfigGmGn{
		GuildID:   req.GuildID,
		ChannelID: req.ChannelID,
	}); err != nil {
		return err
	}

	return nil
}

func (e *Entity) GetGuildTokens(guildID string) ([]model.GuildConfigToken, error) {
	guildTokens, err := e.repo.GuildConfigToken.GetByGuildID(guildID)
	if err != nil {
		return nil, err
	}

	return guildTokens, nil
}

func (e *Entity) UpsertGuildTokenConfig(req request.UpsertGuildTokenConfigRequest) error {
	token, err := e.repo.Token.GetBySymbol(req.Symbol, true)
	if err != nil {
		return err
	}

	if err := e.repo.GuildConfigToken.UpsertMany([]model.GuildConfigToken{{
		GuildID: req.GuildID,
		TokenID: token.ID,
		Active:  req.Active,
	}}); err != nil {
		return err
	}

	return nil
}
