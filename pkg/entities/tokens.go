package entities

import (
	"strings"

	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
)

func (e *Entity) UpsertGuildCustomTokenConfig(req request.UpsertCustomTokenConfigRequest) error {

	if err := e.repo.Token.UpsertOne(model.Token{
		ID:                  req.Id,
		Address:             req.Address,
		Symbol:              req.Symbol,
		ChainID:             req.ChainID,
		Decimals:            req.Decimals,
		DiscordBotSupported: req.DiscordBotSupported,
		CoinGeckoID:         req.CoinGeckoID,
		Name:                strings.ToUpper(req.Name),
		GuildDefault:        req.GuildDefault,
	}); err != nil {
		return err
	}

	if err := e.repo.GuildConfigToken.UpsertMany([]model.GuildConfigToken{{
		GuildID: req.GuildID,
		TokenID: req.Id,
		Active:  req.Active,
	}}); err != nil {
		return err
	}

	return nil
}
