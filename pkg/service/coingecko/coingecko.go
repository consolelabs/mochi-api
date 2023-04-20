package coingecko

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
)

type CoinGecko struct {
	getMarketChartURL  string
	searchCoinURL      string
	getCoinURL         string
	getPriceURL        string
	getCoinOhlc        string
	getCoinsMarketData string
	getSupportedCoins  string
	getAssetPlatforms  string
	getCoinByContract  string
	getTrendingSearch  string
}

func NewService(cfg *config.Config) Service {
	apiKey := cfg.CoinGeckoAPIKey
	return &CoinGecko{
		getMarketChartURL:  "https://pro-api.coingecko.com/api/v3/coins/%s/market_chart?vs_currency=%s&days=%d&x_cg_pro_api_key=" + apiKey,
		searchCoinURL:      "https://pro-api.coingecko.com/api/v3/search?query=%s&x_cg_pro_api_key=" + apiKey,
		getCoinURL:         "https://pro-api.coingecko.com/api/v3/coins/%s?x_cg_pro_api_key=" + apiKey,
		getPriceURL:        "https://pro-api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=%s&x_cg_pro_api_key=" + apiKey,
		getCoinOhlc:        "https://pro-api.coingecko.com/api/v3/coins/%s/ohlc?days=%s&vs_currency=usd&x_cg_pro_api_key=" + apiKey,
		getCoinsMarketData: "https://pro-api.coingecko.com/api/v3/coins/markets?vs_currency=usd&ids=%s&order=market_cap_desc&per_page=%s&page=%s&sparkline=%t&price_change_percentage=1h,24h,7d&x_cg_pro_api_key=" + apiKey,
		getSupportedCoins:  "https://pro-api.coingecko.com/api/v3/coins/list?x_cg_pro_api_key=" + apiKey,
		getAssetPlatforms:  "https://pro-api.coingecko.com/api/v3/asset_platforms?x_cg_pro_api_key=" + apiKey,
		getCoinByContract:  "https://pro-api.coingecko.com/api/v3/coins/%s/contract/%s?x_cg_pro_api_key=" + apiKey,
		getTrendingSearch:  "https://pro-api.coingecko.com/api/v3/search/trending?x_cg_pro_api_key=" + apiKey,
	}
}

func (c *CoinGecko) GetHistoricalMarketData(req *request.GetMarketChartRequest) (*response.CoinPriceHistoryResponse, error, int) {
	resp := &response.HistoricalMarketChartResponse{}
	statusCode, err := util.FetchData(fmt.Sprintf(c.getMarketChartURL, req.CoinID, req.Currency, req.Days), resp)
	if err != nil || statusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch historical market data - coin %s: %v", req.CoinID, err), statusCode
	}

	data := &response.CoinPriceHistoryResponse{}
	for _, p := range resp.Prices {
		timestamp := time.UnixMilli(int64(p[0])).Format("01-02")
		data.Times = append(data.Times, timestamp)
		data.Prices = append(data.Prices, p[1])
	}
	from := time.UnixMilli(int64(resp.Prices[0][0])).Format("January 02, 2006")
	to := time.UnixMilli(int64(resp.Prices[len(resp.Prices)-1][0])).Format("January 02, 2006")
	data.From = from
	data.To = to

	return data, nil, http.StatusOK
}

func (c *CoinGecko) GetCoin(coinID string) (*response.GetCoinResponse, error, int) {
	data := &response.GetCoinResponse{}
	statusCode, err := util.FetchData(fmt.Sprintf(c.getCoinURL, coinID), data)
	if err != nil || statusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch coin data of %s: %v", coinID, err), statusCode
	}

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
	if coinIDsArg != "" {
		statusCode, err := util.FetchData(fmt.Sprintf(c.getPriceURL, coinIDsArg, currency), resp)
		if err != nil || statusCode != http.StatusOK {
			return nil, fmt.Errorf("failed to fetch price data of %s: %v", coinIDs, err)
		}
	}

	prices := make(map[string]float64)
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
		return nil, err, statusCode
	}

	return resp, nil, http.StatusOK
}

func (c *CoinGecko) GetCoinsMarketData(ids []string, sparkline bool, page, pageSize string) ([]response.CoinMarketItemData, error, int) {
	var res []response.CoinMarketItemData
	statusCode, err := util.FetchData(fmt.Sprintf(c.getCoinsMarketData, strings.Join(ids, ","), pageSize, page, sparkline), &res)
	if err != nil {
		return nil, err, statusCode
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

func (c *CoinGecko) GetAssetPlatform(chainId int) (*response.AssetPlatformResponseData, error) {
	var res []response.AssetPlatformResponseData
	status, err := util.FetchData(c.getAssetPlatforms, &res)
	if err != nil || status != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch asset platforms with status %d: %v", status, err)
	}
	for _, p := range res {
		if p.ChainIdentifier != nil && *p.ChainIdentifier == int64(chainId) {
			return &p, nil
		}
	}
	return nil, errors.New("asset platform not found")
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
