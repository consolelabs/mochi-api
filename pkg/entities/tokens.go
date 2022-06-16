package entities

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
)

type Response struct {
	Id     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

type TokenCompareReponse struct {
	PriceCompare []float32 `json:"price_compare"`
	Times        []string  `json:"times"`
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

func (e *Entity) CheckExistToken(symbol string) (bool, error) {
	listSymbol, err := e.repo.Token.GetAll()
	if err != nil {
		return false, err
	}

	for i, _ := range listSymbol {
		if strings.ToLower(symbol) == strings.ToLower(listSymbol[i].Symbol) {
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

func (e *Entity) GetAllSupportedToken(guildID string) (returnToken []model.Token, err error) {
	returnToken, err = e.repo.Token.GetAllSupportedToken(guildID)
	if err != nil {
		return returnToken, err
	}

	return returnToken, nil
}

func (e *Entity) GetHistoryInfo(sourceSymbol string, interval string) (res [][]float32, err error) {
	reqUrl := "https://api.coingecko.com/api/v3/coins/" + sourceSymbol + "/ohlc?days=" + interval + "&vs_currency=usd"
	println(reqUrl)
	resp, err := http.Get(reqUrl)
	// Get request
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}

	_ = json.Unmarshal([]byte(body), &res)

	return res, err
}

func (e *Entity) TokenCompare(sourceSymbolInfo [][]float32, targetSymbolInfo [][]float32) (tokenCompareRes TokenCompareReponse, err error) {
	for i := 0; i < len(sourceSymbolInfo); i++ {
		currentRatio := sourceSymbolInfo[i][1] / targetSymbolInfo[i][1]
		tokenCompareRes.PriceCompare = append(tokenCompareRes.PriceCompare, currentRatio)

		// get times and convert to string, ex: 2022-05-24T07:59:30Z
		getTime := int(sourceSymbolInfo[i][0])
		getStringTime := strconv.Itoa(getTime / 1000)
		convertTime, err := strconv.ParseInt(getStringTime, 10, 64)
		if err != nil {
			return tokenCompareRes, err
		}

		stringTime, err := time.Unix(convertTime, 0).UTC().MarshalText()
		if err != nil {
			return tokenCompareRes, err
		}

		tokenCompareRes.Times = append(tokenCompareRes.Times, string(stringTime))
	}

	return tokenCompareRes, err
}
