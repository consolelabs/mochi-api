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
	getMarketChartURL string
	searchCoinURL     string
	getCoinURL        string
	getPriceURL       string
}

func NewService() Service {
	return &CoinGecko{
		getMarketChartURL: "https://api.coingecko.com/api/v3/coins/%s/market_chart?vs_currency=%s&days=%d",
		searchCoinURL:     "https://api.coingecko.com/api/v3/search?query=%s",
		getCoinURL:        "https://api.coingecko.com/api/v3/coins/%s",
		getPriceURL:       "https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=%s",
	}
}

func (c *CoinGecko) searchCoin(query string) (string, error, int) {
	resp := &response.SearchedCoinsListResponse{}
	statusCode, err := util.FetchData(fmt.Sprintf(c.searchCoinURL, query), resp)
	if err != nil || resp == nil || len(resp.Coins) == 0 {
		return "", fmt.Errorf("failed to search for coins by query %s: %v", query, err), statusCode
	}

	return resp.Coins[0].ID, nil, http.StatusOK
}

func (c *CoinGecko) GetHistoricalMarketData(req *request.GetMarketChartRequest) (*response.CoinPriceHistoryResponse, error, int) {
	resp := &response.HistoricalMarketChartResponse{}
	statusCode, err := util.FetchData(fmt.Sprintf(c.getMarketChartURL, req.CoinID, req.Currency, req.Days), resp)
	if err != nil || statusCode != http.StatusOK {
		if statusCode != http.StatusNotFound {
			return nil, fmt.Errorf("failed to fetch historical market data - coin %s: %v", req.CoinID, err), statusCode
		}

		// search coin by name, id, symbol, ...  if coinID is invalid
		req.CoinID, err, statusCode = c.searchCoin(req.CoinID)
		if err != nil || statusCode != http.StatusOK {
			return nil, err, statusCode
		}
		return c.GetHistoricalMarketData(req)
	}

	data := response.CoinPriceHistoryResponse{}
	for _, p := range resp.Prices {
		timestamp := time.UnixMilli(int64(p[0])).Format("01-02")
		data.Timestamps = append(data.Timestamps, timestamp)
		data.Prices = append(data.Prices, p[1])
	}

	from := time.UnixMilli(int64(resp.Prices[0][0])).Format("January 02, 2006")
	data.From = from
	to := time.UnixMilli(int64(resp.Prices[len(resp.Prices)-1][0])).Format("January 02, 2006")
	data.To = to

	return &data, nil, http.StatusOK
}

func (c *CoinGecko) GetCoin(coinID string) (*response.GetCoinResponse, error, int) {
	resp := &response.GetCoinResponse{}
	statusCode, err := util.FetchData(fmt.Sprintf(c.getCoinURL, coinID), resp)
	if err != nil || statusCode != http.StatusOK {
		if statusCode != http.StatusNotFound {
			return nil, fmt.Errorf("failed to fetch coin data of %s: %v", coinID, err), statusCode
		}

		// search coin by name, id, symbol, ...  if coinID is invalid
		coinID, err, statusCode := c.searchCoin(coinID)
		if err != nil || statusCode != http.StatusOK {
			return nil, fmt.Errorf("failed to search for coins by query %s: %v", coinID, err), statusCode
		}
		return c.GetCoin(coinID)
	}

	return resp, nil, http.StatusOK
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
