package entities

import (
	"errors"
	"fmt"
	"strings"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"gorm.io/gorm"
)

func (e *Entity) CreateCustomToken(req request.UpsertCustomTokenConfigRequest) error {
	chain, err := e.repo.Chain.GetByShortName(req.Chain)
	if err != nil {
		e.log.Error(err, "[Entity][CreateCustomToken] repo.Chain.GetByShortName failed")
		return fmt.Errorf("error getting chain: %v", err)
	}

	coins, err := e.SearchCoins(req.Symbol)
	if err != nil {
		e.log.Error(err, "[Entity][CreateCustomToken] svc.CoinGecko.SearchCoins failed")
		return fmt.Errorf("error seaching coin: %v", err)
	}
	if len(coins) == 0 {
		return fmt.Errorf("cannot find token by symbol")
	}

	coin, err, _ := e.svc.CoinGecko.GetCoin(coins[0].ID)
	if err != nil {
		e.log.Error(err, "[Entity][CreateCustomToken] svc.CoinGecko.GetCoin failed")
		return fmt.Errorf("error getting coin: %v", err)
	}
	if coin.AssetPlatformID != chain.CoinGeckoID {
		e.log.
			Fields(logger.Fields{"asset_platform_id": coin.AssetPlatformID, "coin_gecko_id": chain.CoinGeckoID}).
			Error(nil, "[Entity][CreateCustomToken] token is not supported on this chain")
		return fmt.Errorf("token is not supported on this chain")
	}

	historicalTokenPrices, err, statusCode := e.svc.Covalent.GetHistoricalTokenPrices(chain.ID, chain.Currency, req.Address)
	if err != nil {
		e.log.Fields(logger.Fields{"statusCode": statusCode}).Error(err, "[Entity][CreateCustomToken] svc.Covalent.GetHistoricalTokenPrices failed")
		return err
	}

	token, err := e.repo.Token.GetOneBySymbol(req.Symbol)
	if err != nil {
		e.log.Error(err, "[Entity][CreateCustomToken] repo.Token.GetOneBySymbol failed")
	}

	if token == nil {
		token = &model.Token{
			ChainID:             chain.ID,
			Address:             req.Address,
			CoinGeckoID:         coins[0].ID,
			Name:                coins[0].Name,
			Symbol:              strings.ToUpper(coins[0].Symbol),
			Decimals:            historicalTokenPrices.Data[0].Decimals,
			DiscordBotSupported: true,
			Chain:               chain,
		}
		if err = e.repo.Token.CreateOne(token); err != nil {
			e.log.Error(err, "[Entity][CreateCustomToken] repo.Token.CreateOne failed")
			return err
		}
	}

	if !token.DiscordBotSupported {
		return errors.New("token is not supported to add to your server")
	}

	gct, err := e.repo.GuildConfigToken.GetByGuildIDAndTokenID(req.GuildID, token.ID)
	if err != nil && err != gorm.ErrRecordNotFound {
		e.log.Error(err, "[Entity][CreateCustomToken] repo.GuildConfigToken.GetByGuildIDAndTokenID failed")
		return err
	}

	if gct == nil {
		if err := e.repo.GuildConfigToken.CreateOne(model.GuildConfigToken{
			GuildID: req.GuildID,
			TokenID: token.ID,
			Active:  true,
		}); err != nil {
			e.log.Error(err, "[Entity][CreateCustomToken] repo.GuildConfigToken.CreateOne failed")
			return err
		}
	}

	return nil
}

func (e *Entity) GetAllSupportedToken(guildID string) (returnToken []model.Token, err error) {
	returnToken, err = e.repo.Token.GetAllSupportedToken(guildID)
	if err != nil {
		return returnToken, err
	}

	return returnToken, nil
}

func (e *Entity) GetDefaultToken(guildID string) (*model.Token, error) {
	if _, err := e.repo.DiscordGuilds.GetByID(guildID); err != nil {
		e.log.Fields(logger.Fields{"guild_id": guildID}).Error(err, "[Entity][GetDefaultToken] repo.DiscordGuilds.GetByID failed")
		return nil, err
	}

	token, err := e.repo.Token.GetDefaultTokenByGuildID(guildID)
	if err != nil && err != gorm.ErrRecordNotFound {
		e.log.Fields(logger.Fields{"guild_id": guildID}).Error(err, "[Entity][GetDefaultToken] repo.Token.GetDefaultTokenByGuildID failed")
		return nil, err
	}

	return &token, nil
}

func (e *Entity) SetDefaultToken(req request.ConfigDefaultTokenRequest) error {
	_, err := e.repo.DiscordGuilds.GetByID(req.GuildID)
	if err != nil {
		e.log.Fields(logger.Fields{"guild_id": req.GuildID}).Error(err, "[Entity][SetDefaultToken] repo.DiscordGuilds.GetByID failed")
		return err
	}

	token, err := e.repo.Token.GetBySymbol(req.Symbol, true)
	if err != nil {
		e.log.Fields(logger.Fields{"symbol": req.Symbol}).Error(err, "[Entity][SetDefaultToken] repo.Token.GetBySymbol failed")
		return err
	}

	if err = e.repo.GuildConfigToken.UpsertOne(model.GuildConfigToken{
		GuildID:   req.GuildID,
		TokenID:   token.ID,
		Active:    true,
		IsDefault: true,
	}); err != nil {
		e.log.Fields(logger.Fields{"guild_id": req.GuildID, "token_id": token.ID}).Error(err, "[Entity][SetDefaultToken] repo.GuildConfigToken.UpsertOne failed")
		return err
	}

	if err := e.repo.GuildConfigToken.UnsetOldDefaultToken(req.GuildID, token.ID); err != nil {
		e.log.Fields(logger.Fields{"guild_id": req.GuildID, "token_id": token.ID}).Error(err, "[Entity][SetDefaultToken] repo.GuildConfigToken.SetDefaultToken failed")
		return err
	}

	return nil
}

func (e *Entity) RemoveDefaultToken(guildID string) error {
	if _, err := e.repo.DiscordGuilds.GetByID(guildID); err != nil {
		e.log.Fields(logger.Fields{"guild_id": guildID}).Error(err, "[Entity][RemoveDefaultToken] repo.DiscordGuilds.GetByID failed")
		return err
	}

	if err := e.repo.GuildConfigToken.RemoveDefaultToken(guildID); err != nil {
		e.log.Fields(logger.Fields{"guild_id": guildID}).Error(err, "[Entity][RemoveDefaultToken] repo.GuildConfigToken.RemoveDefaultToken failed")
		return err
	}

	return nil
}
