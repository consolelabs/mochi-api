package entities

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
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

func (e *Entity) GetChainIdBySymbol(symbol string) (int, error) {
	mapChainChainId := map[string]string{
		"eth":    "1",
		"heco":   "128",
		"bsc":    "56",
		"matic":  "137",
		"op":     "10",
		"btt":    "199",
		"okt":    "66",
		"movr":   "1285",
		"celo":   "42220",
		"metis":  "1088",
		"cro":    "25",
		"xdai":   "0x64",
		"boba":   "288",
		"ftm":    "250",
		"avax":   "0xa86a",
		"arb":    "42161",
		"aurora": "1313161554",
	}

	if c, exist := mapChainChainId[strings.ToLower(symbol)]; exist {
		chainID, err := strconv.Atoi(c)
		if err != nil {
			err = fmt.Errorf("chain is not supported/invalid")
			return 0, err
		}
		return chainID, nil
	}

	err := fmt.Errorf("chain is not supported/invalid")
	return 0, err
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
