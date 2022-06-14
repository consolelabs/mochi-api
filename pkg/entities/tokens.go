package entities

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
)

type Response struct {
	Id     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

func (e *Entity) GetIDAndName(symbol string) (string, string, error) {
	resp, err := http.Get("https://api.coingecko.com/api/v3/coins/list")
	// Get request
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}
	var res []Response
	err1 := json.Unmarshal([]byte(body), &res)
	if err1 != nil {
		return "", "", err
	}
	for i := 0; i < len(res); i++ {
		if res[i].Symbol == symbol {
			return res[i].Id, res[i].Name, nil
		}
	}

	return "", "", err
}

func (e *Entity) GetChainIdBySymbol(symbol string) ([]model.Chain, error) {
	listChain, err := e.repo.Chain.GetAll()
	var returnChain []model.Chain
	if err != nil {
		return listChain, err
	}

	for i := 0; i < len(listChain); i++ {
		if strings.ToUpper(symbol) == listChain[i].Currency {
			returnChain = append(returnChain, listChain[i])
			return returnChain, nil
		}
	}

	return listChain, nil
}

func (e *Entity) CheckExistToken(symbol string) (bool, error) {
	listSymbol, err := e.repo.Token.GetAll()
	if err != nil {
		return false, err
	}

	for i := 0; i < len(listSymbol); i++ {
		if strings.ToUpper(symbol) == listSymbol[i].Symbol {
			return true, nil
		}
	}

	return false, nil
}

func (e *Entity) CreateCustomToken(req request.UpsertCustomTokenConfigRequest) error {
	err := e.repo.Token.CreateOne(model.Token{
		Address:             req.Address,
		Symbol:              strings.ToUpper(req.Symbol),
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
