package coingecko

import (
	"fmt"
	"net/http"
	"strings"
	"time"

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
}

func NewService() Service {
	return &CoinGecko{
		getMarketChartURL:  "https://api.coingecko.com/api/v3/coins/%s/market_chart?vs_currency=%s&days=%d",
		searchCoinURL:      "https://api.coingecko.com/api/v3/search?query=%s",
		getCoinURL:         "https://api.coingecko.com/api/v3/coins/%s",
		getPriceURL:        "https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=%s",
		getCoinOhlc:        "https://api.coingecko.com/api/v3/coins/%s/ohlc?days=%s&vs_currency=usd",
		getCoinsMarketData: "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&ids=%s&order=market_cap_desc&per_page=100&page=1&sparkline=true&price_change_percentage=7d",
		getSupportedCoins:  "https://api.coingecko.com/api/v3/coins/list",
	}
}

func (c *CoinGecko) SearchCoins(query string) ([]response.SearchedCoin, error, int) {
	// if query is valid coin ID then use GetCoin()
	coin, err, code := c.GetCoin(query)
	if err == nil {
		return []response.SearchedCoin{
			{
				ID:            coin.ID,
				Name:          coin.Name,
				Symbol:        coin.Symbol,
				MarketCapRank: coin.MarketCapRank,
				Thumb:         coin.Image.Thumb,
			},
		}, nil, http.StatusOK
	} else if code == http.StatusTooManyRequests {
		return nil, fmt.Errorf("too many requests GetCoin(%s): %v", query, err), code
	}

	// if not valid coin ID then search coins by symbol
	res := &response.SearchedCoinsListResponse{}
	code, err = util.FetchData(fmt.Sprintf(c.searchCoinURL, query), res)
	if err != nil || res == nil || len(res.Coins) == 0 {
		return nil, fmt.Errorf("failed to search for coins by query %s: %v", query, err), code
	}

	var matches []response.SearchedCoin
	for _, coin := range res.Coins {
		if strings.EqualFold(query, coin.Symbol) {
			matches = append(matches, coin)
		}
	}

	return matches, nil, http.StatusOK
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
	coinIDsArg := strings.Join(coinIDs, ",")
	statusCode, err := util.FetchData(fmt.Sprintf(c.getPriceURL, coinIDsArg, currency), resp)
	if err != nil || statusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch price data of %s: %v", coinIDs, err)
	}

	prices := make(map[string]float64)
	for k, v := range *resp {
		prices[k] = v[currency]
	}

	return prices, nil
}

func (c *CoinGecko) GetHistoryCoinInfo(sourceSymbol string, interval string) (resp [][]float32, err error, statusCode int) {
	statusCode, err = util.FetchData(fmt.Sprintf(c.getCoinOhlc, sourceSymbol, interval), &resp)
	if err != nil || statusCode != http.StatusOK {
		return nil, err, statusCode
	}

	return resp, nil, http.StatusOK
}

func (c *CoinGecko) GetCoinsMarketData(ids []string) ([]response.CoinMarketItemData, error, int) {
	var res []response.CoinMarketItemData
	statusCode, err := util.FetchData(fmt.Sprintf(c.getCoinsMarketData, strings.Join(ids, ",")), &res)
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
