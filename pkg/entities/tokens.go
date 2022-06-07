package entities

import (
	"strings"

	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
)

func (e *Entity) UpsertCustomToken(req request.UpsertCustomTokenConfigRequest) error {
	err := e.repo.Token.UpsertOne(model.Token{
		Address:             req.Address,
		Symbol:              req.Symbol,
		ChainID:             req.ChainID,
		Decimals:            req.Decimals,
		DiscordBotSupported: req.DiscordBotSupported,
		CoinGeckoID:         req.CoinGeckoID,
		Name:                strings.ToUpper(req.Name),
		GuildDefault:        req.GuildDefault,
	})
	if err != nil {
		return err
	}

	return nil
}

func (e *Entity) GetTokenBySymbol(symbol string, flag bool) (int, error) {
	token, err := e.repo.Token.GetBySymbol(symbol, flag)
	if err != nil {
		return 0, err
	}
	return token.ID, nil
}

func (e *Entity) UpsertGuildCustomTokenConfig(req request.UpsertCustomTokenConfigRequest) error {
	err := e.repo.GuildConfigToken.UpsertOne(model.GuildConfigToken{
		GuildID: req.GuildID,
		TokenID: req.Id,
		Active:  req.Active,
	})
	if err != nil {
		return err
	}

	return nil
}
