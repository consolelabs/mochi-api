package coingecko

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/model"
	errs "github.com/defipod/mochi/pkg/model/errors"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
)

type CoinGecko struct {
	getMarketChartURL                 string
	searchCoinURL                     string
	getCoinURL                        string
	getPriceURL                       string
	getCoinOhlc                       string
	getCoinsMarketData                string
	getSupportedCoins                 string
	getAssetPlatforms                 string
	getCoinByContract                 string
	getTrendingSearch                 string
	getTopGainerLoser                 string
	getHistoricalGlobalMarketChartURL string
	geckoTerminalSearchURL            string
	getGlobalCryptoDataURL            string
}

func NewService(cfg *config.Config) Service {
	apiKey := cfg.CoinGeckoAPIKey

	return &CoinGecko{
		getMarketChartURL:                 "https://pro-api.coingecko.com/api/v3/coins/%s/market_chart?vs_currency=%s&days=%d&x_cg_pro_api_key=" + apiKey,
		searchCoinURL:                     "https://pro-api.coingecko.com/api/v3/search?query=%s&x_cg_pro_api_key=" + apiKey,
		getCoinURL:                        "https://pro-api.coingecko.com/api/v3/coins/%s?x_cg_pro_api_key=" + apiKey,
		getPriceURL:                       "https://pro-api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=%s&x_cg_pro_api_key=" + apiKey,
		getCoinOhlc:                       "https://pro-api.coingecko.com/api/v3/coins/%s/ohlc?days=%s&vs_currency=usd&x_cg_pro_api_key=" + apiKey,
		getCoinsMarketData:                "https://pro-api.coingecko.com/api/v3/coins/markets?vs_currency=usd&ids=%s&order=market_cap_desc&per_page=%s&page=%s&sparkline=%t&price_change_percentage=1h,24h,7d&x_cg_pro_api_key=" + apiKey,
		getSupportedCoins:                 "https://pro-api.coingecko.com/api/v3/coins/list?x_cg_pro_api_key=" + apiKey,
		getAssetPlatforms:                 "https://pro-api.coingecko.com/api/v3/asset_platforms?x_cg_pro_api_key=" + apiKey,
		getCoinByContract:                 "https://pro-api.coingecko.com/api/v3/coins/%s/contract/%s?x_cg_pro_api_key=" + apiKey,
		getTrendingSearch:                 "https://pro-api.coingecko.com/api/v3/search/trending?x_cg_pro_api_key=" + apiKey,
		getTopGainerLoser:                 "https://pro-api.coingecko.com/api/v3/coins/top_gainers_losers?vs_currency=usd&duration=%s&top_coins=300&x_cg_pro_api_key=" + apiKey,
		getHistoricalGlobalMarketChartURL: "https://pro-api.coingecko.com/api/v3/global/market_cap_chart?days=%d&x_cg_pro_api_key=" + apiKey,
		getGlobalCryptoDataURL:            "https://pro-api.coingecko.com/api/v3/global?x_cg_pro_api_key=" + apiKey,
		geckoTerminalSearchURL:            "https://app.geckoterminal.com/api/p1/search",
	}
}

func (c *CoinGecko) GetHistoricalMarketData(coinID, currency string, days int) (*response.HistoricalMarketChartResponse, error, int) {
	resp := &response.HistoricalMarketChartResponse{}
	statusCode, err := util.FetchData(fmt.Sprintf(c.getMarketChartURL, coinID, currency, days), resp)
	if err != nil || statusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch historical market data - coin %s: %v", coinID, err), statusCode
	}
	return resp, nil, http.StatusOK
}

func (c *CoinGecko) GetCoin(coinID string) (*response.GetCoinResponse, error, int) {
	data := &response.GetCoinResponse{}
	statusCode, err := util.FetchData(fmt.Sprintf(c.getCoinURL, coinID), data)
	if err != nil || statusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch coin data of %s: %v", coinID, err), statusCode
	}

	data.CoingeckoId = coinID

	return data, nil, http.StatusOK
}

func (c *CoinGecko) GetCoinPrice(coinIDs []string, currency string) (map[string]float64, error) {
	resp := &response.CoinPriceResponse{}
	coinIDsArg := ""
	icyIsExist := false
	for _, v := range coinIDs {
		if v != "icy" {
			coinIDsArg = coinIDsArg + "," + v
		} else {
			icyIsExist = true
		}
	}

	prices := make(map[string]float64)
	if coinIDsArg != "" {
		statusCode, err := util.FetchData(fmt.Sprintf(c.getPriceURL, coinIDsArg, currency), resp)
		if err != nil || statusCode != http.StatusOK {
			return prices, fmt.Errorf("failed to fetch price data of %s: %v", coinIDs, err)
		}
	}

	for k, v := range *resp {
		prices[k] = v[currency]
	}
	if icyIsExist && currency == "usd" {
		prices["icy"] = 1.5
	}

	return prices, nil
}

func (c *CoinGecko) GetHistoryCoinInfo(sourceSymbol string, days string) (resp [][]float64, err error, statusCode int) {
	statusCode, err = util.FetchData(fmt.Sprintf(c.getCoinOhlc, sourceSymbol, days), &resp)
	if err != nil || statusCode != http.StatusOK {
		if statusCode == http.StatusBadRequest {
			return nil, errs.ErrCoingeckoNotSupported, statusCode
		}
		if statusCode == http.StatusNotFound {
			return nil, errs.ErrRecordNotFound, statusCode
		}
		return nil, err, statusCode
	}

	return resp, nil, http.StatusOK
}

func (c *CoinGecko) GetCoinsMarketData(ids []string, sparkline bool, page, pageSize string) ([]response.CoinMarketItemData, error, int) {
	res := make([]response.CoinMarketItemData, 0)
	var resTmp []response.CoinMarketItemDataRes

	statusCode, err := util.FetchData(fmt.Sprintf(c.getCoinsMarketData, strings.Join(ids, ","), pageSize, page, sparkline), &resTmp)
	if err != nil {
		return nil, err, statusCode
	}

	for _, r := range resTmp {
		res = append(res, r.ToCoinMarketItemData())
	}

	return res, nil, http.StatusOK
}

func (c *CoinGecko) GetSupportedCoins() ([]response.CoingeckoSupportedTokenResponse, error, int) {
	data := make([]response.CoingeckoSupportedTokenResponse, 0)
	statusCode, err := util.FetchData(c.getSupportedCoins, &data)
	if err != nil || statusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch supported coins list: %v", err), statusCode
	}
	return data, nil, http.StatusOK
}

func (c *CoinGecko) GetAssetPlatforms() ([]*response.AssetPlatformResponseData, error) {
	var res []*response.AssetPlatformResponseData
	status, err := util.FetchData(c.getAssetPlatforms, &res)
	if err != nil || status != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch asset platforms with status %d: %v", status, err)
	}
	return res, nil
}

func (c *CoinGecko) GetCoinByContract(platformId, contractAddress string) (*response.GetCoinByContractResponseData, error) {
	var res response.GetCoinByContractResponseData
	url := fmt.Sprintf(c.getCoinByContract, platformId, contractAddress)
	status, err := util.FetchData(url, &res)
	if err != nil || status != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch asset platforms with status %d: %v", status, err)
	}
	return &res, nil
}

func (c *CoinGecko) GetTrendingSearch() (*response.GetTrendingSearch, error) {
	var res response.GetTrendingSearch
	status, err := util.FetchData(c.getTrendingSearch, &res)
	if err != nil || status != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch trending search with status %d: %v", status, err)
	}
	return &res, nil
}

func (c *CoinGecko) GetTopLoserGainer(req request.TopGainerLoserRequest) (*response.GetTopGainerLoser, error) {
	var res response.GetTopGainerLoser
	url := fmt.Sprintf(c.getTopGainerLoser, req.Duration)
	status, err := util.FetchData(url, &res)
	if err != nil || status != http.StatusOK {
		if status == http.StatusBadRequest {
			return nil, errs.ErrInvalidCoingeckoSvcParam
		}
		return nil, fmt.Errorf("failed to fetch trending search with status %d: %v", status, err)
	}
	return &res, nil
}

func (c *CoinGecko) GetHistoricalGlobalMarketChart(days int) (*response.GetHistoricalGlobalMarketResponse, error) {
	res := &response.GetHistoricalGlobalMarketResponse{}
	url := fmt.Sprintf(c.getHistoricalGlobalMarketChartURL, days)
	status, err := util.FetchData(url, &res)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch global market chart with status %d: %v", status, err)
	}
	return res, nil
}

func (c *CoinGecko) GetGlobalData() (*response.GetGlobalDataResponse, error) {
	res := &response.GetGlobalDataResponse{}
	url := c.getGlobalCryptoDataURL
	status, err := util.FetchData(url, &res)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch global market chart with status %d: %v", status, err)
	}
	return res, nil
}

func (c *CoinGecko) SearchCoin(query string) (*response.SearchCoinResponse, error, int) {
	resp := &CoinGeckoSearchResponse{}

	req, err := http.NewRequest(http.MethodGet, c.searchCoinURL, nil)
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}

	q := req.URL.Query()
	q.Add("query", query)
	req.URL.RawQuery = q.Encode()

	statusCode, err := util.FetchData(req.URL.String(), resp)
	if err != nil || statusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to search coin %s: %v", query, err), statusCode
	}

	coins := make([]model.CoingeckoSupportedTokens, 0)
	for _, coin := range resp.Coins {
		if coin.ID != nil && coin.Name != nil && coin.Symbol != nil {
			coins = append(coins, model.CoingeckoSupportedTokens{
				ID:     *coin.ID,
				Name:   *coin.Name,
				Symbol: *coin.Symbol,
			})
		}
	}

	return &response.SearchCoinResponse{
		Data: coins,
	}, nil, http.StatusOK
}
