package coingecko

import (
	"fmt"
	"net/http"
	"strconv"
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
	getHistoryInfo    string
}

func NewService() Service {
	return &CoinGecko{
		getMarketChartURL: "https://api.coingecko.com/api/v3/coins/%s/market_chart?vs_currency=%s&days=%d",
		searchCoinURL:     "https://api.coingecko.com/api/v3/search?query=%s",
		getCoinURL:        "https://api.coingecko.com/api/v3/coins/%s",
		getPriceURL:       "https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=%s",
		getHistoryInfo:    "https://api.coingecko.com/api/v3/coins/%s/ohlc?days=%s&vs_currency=usd",
	}
}

func (c *CoinGecko) SearchCoins(query string) ([]response.SearchedCoin, error, int) {
	resp := &response.SearchedCoinsListResponse{}
	statusCode, err := util.FetchData(fmt.Sprintf(c.searchCoinURL, query), resp)
	if err != nil || resp == nil || len(resp.Coins) == 0 {
		return nil, fmt.Errorf("failed to search for coins by query %s: %v", query, err), statusCode
	}

	var matches []response.SearchedCoin
	for _, coin := range resp.Coins {
		if strings.EqualFold(query, coin.Symbol) || strings.EqualFold(query, coin.Name) {
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
	statusCode, err = util.FetchData(fmt.Sprintf(c.getHistoryInfo, sourceSymbol, interval), &resp)
	if err != nil || statusCode != http.StatusOK {
		return nil, err, statusCode
	}

	return resp, nil, http.StatusOK
}

func (c *CoinGecko) TokenCompare(sourceSymbolInfo [][]float32, targetSymbolInfo [][]float32) (*response.TokenCompareReponse, error) {
	fmt.Println("aaaaaaaaaaaaaaaa")
	tokenCompareRes := &response.TokenCompareReponse{}
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

	return tokenCompareRes, nil
}
