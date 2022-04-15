package entities

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/defipod/mochi/pkg/request"
	"github.com/gin-gonic/gin"
)

const (
	getBeefyLPPriceURL    = `https://api.beefy.finance/lps`
	getBeefyTokenPriceURL = `https://api.beefy.finance/prices`
)

type HistoricalMarketChartResponse struct {
	Prices [][]float64 `json:"prices"`
}

type SearchCoinsResponse struct {
	Coins []CoinDataResponse `json:"coins"`
}

type CoinDataResponse struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Symbol        string `json:"symbol"`
	MarketCapRank int    `json:"market_cap_rank"`
	Thumb         string `json:"thumb"`
	Large         string `json:"large"`
}

type CoinPriceHistoryResponse struct {
	Timestamps []string  `json:"timestamps"`
	Prices     []float64 `json:"prices"`
	From       string    `json:"from"`
	To         string    `json:"to"`
}

func fetchHistoricalMarketData(req *request.GetMarketChartRequest, result interface{}) (int, error) {
	client := &http.Client{Timeout: time.Second * 60}
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/%s/market_chart?vs_currency=%s&days=%d", req.CoinID, req.Currency, req.Days)
	resp, err := client.Get(url)
	if err != nil {
		return resp.StatusCode, err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, err
	}

	return resp.StatusCode, json.Unmarshal(bytes, result)
}

func searchForCoins(query, result interface{}) (int, error) {
	client := &http.Client{Timeout: time.Second * 60}
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/search?query=%s", query)
	resp, err := client.Get(url)
	if err != nil {
		return resp.StatusCode, err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, err
	}

	return resp.StatusCode, json.Unmarshal(bytes, result)
}

func (e *Entity) GetHistoricalMarketChart(c *gin.Context) (*CoinPriceHistoryResponse, error, int) {
	req, err := request.ValidateRequest(c)
	if err != nil {
		return nil, err, http.StatusBadRequest
	}

	var resp *HistoricalMarketChartResponse
	statusCode, err := fetchHistoricalMarketData(req, &resp)
	if err != nil || statusCode != 200 {
		if statusCode != 404 {
			return nil, fmt.Errorf("failed to fetch historical market data - coin %s: %v", req.CoinID, err), statusCode
		}

		var searchResp *SearchCoinsResponse
		_, err := searchForCoins(req.CoinID, &searchResp)
		if err != nil || searchResp == nil || len(searchResp.Coins) == 0 {
			return nil, fmt.Errorf("failed to search for coins by query %s: %v", req.CoinID, err), http.StatusNotFound
		}
		req.CoinID = searchResp.Coins[0].ID
		statusCode, err := fetchHistoricalMarketData(req, &resp)
		if err != nil || statusCode != 200 {
			return nil, fmt.Errorf("failed to fetch historical market data 2 - coin %s: %v", req.CoinID, err), statusCode
		}
	}

	data := CoinPriceHistoryResponse{}
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
