package coingecko

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/defipod/mochi/pkg/cache"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
	"github.com/go-redis/redis/v8"
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

	brc20Cache     cache.Cache
	brc20KeyPrefix string
}

func NewService(cfg *config.Config) Service {
	apiKey := cfg.CoinGeckoAPIKey

	// brc20
	redisOpt, err := redis.ParseURL(cfg.RedisURL)
	if err != nil {
		log.Fatal(err, "failed to init redis")
	}

	cache, err := cache.NewRedisCache(redisOpt)
	if err != nil {
		log.Fatal(err, "failed to init redis cache")
	}

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

		brc20Cache:     cache,
		brc20KeyPrefix: "brc20Token:",
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

func (c *CoinGecko) GetTopLoserGainer(req request.TopGainerLoserRequest) (*response.GetTopGainerLoser, error) {
	var res response.GetTopGainerLoser
	url := fmt.Sprintf(c.getTopGainerLoser, req.Duration)
	status, err := util.FetchData(url, &res)
	if err != nil || status != http.StatusOK {
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

func (c *CoinGecko) GetCoinBRC20(coinId string) (*response.GetCoinResponse, error, int) {
	coinIdLower := strings.ToLower(coinId)
	coinName := strings.TrimPrefix(coinIdLower, "brc20")
	// get from cache
	coinData, err := c.brc20Cache.GetString(c.brc20KeyPrefix + strings.ToLower(coinName))
	if err != nil {
		return nil, err, 0
	}

	if coinData == "" {
		return nil, errors.New("coin not found"), 0
	}

	coinDataMap := make(map[string]interface{})
	if err := json.Unmarshal([]byte(coinData), &coinDataMap); err != nil {
		return nil, err, 0
	}

	// id := coinDataMap["id"].(int)
	name := coinDataMap["name"].(string)
	priceUsd := coinDataMap["priceUsd"].(float64)
	marketCapUsd := coinDataMap["marketCapUsd"].(float64)
	percent24h := coinDataMap["percent24h"].(float64)

	resp := &response.GetCoinResponse{
		ID:              name,
		Name:            name,
		Symbol:          name,
		MarketCapRank:   0,
		AssetPlatformID: "brc20",
		Image:           response.CoinImage{},
		MarketData: response.MarketData{
			CurrentPrice: map[string]float64{
				"usd": priceUsd,
			},
			MarketCap: map[string]float64{
				"usd": marketCapUsd,
			},
			PriceChangePercentage1hInCurrency: map[string]float64{},
			PriceChangePercentage24hInCurrency: map[string]float64{
				"usd": percent24h,
			},
			PriceChangePercentage7dInCurrency: map[string]float64{},
		},
		Tickers: []response.TickerData{},
		Description: response.CoinDescription{
			EngDescription: "BRC20 Token, data from brc-20.io",
		},
	}

	return resp, nil, 0
}
