package binance

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/defipod/mochi/pkg/response"
	bapdater "github.com/defipod/mochi/pkg/service/binance/adapter"
	"github.com/defipod/mochi/pkg/util"
)

type Binance struct {
	getExchangeInfoURL string
	getSymbolKlinesURL string
	getAvgPriceURL     string
	getTickerPriceURL  string
}

func NewService() Service {
	return &Binance{
		getExchangeInfoURL: "https://api.binance.com/api/v3/exchangeInfo",
		getSymbolKlinesURL: "https://api.binance.com/api/v3/uiKlines?symbol=%s&interval=1h&limit=168", // 168h = 7d
		getAvgPriceURL:     "https://api.binance.com/api/v3/avgPrice",
		getTickerPriceURL:  "https://api.binance.com/api/v3/ticker/price",
	}
}

func (b *Binance) GetExchangeInfo(symbol string) (*response.GetExchangeInfoResponse, error, int) {
	symbol = strings.Replace(symbol, "/", "", 1)
	res := &response.GetExchangeInfoResponse{}
	url := b.getExchangeInfoURL
	if symbol != "" {
		url = fmt.Sprintf("%s?symbol=%s", url, symbol)
	}
	statusCode, err := util.FetchData(url, res)
	if err != nil || statusCode != http.StatusOK {
		return nil, fmt.Errorf("binance.GetExchangeInfo() failed: %v", err), statusCode
	}
	return res, nil, http.StatusOK
}

func (b *Binance) GetTickerPrice(symbol string) (*response.GetTickerPriceResponse, error, int) {
	symbol = strings.Replace(symbol, "/", "", 1)
	res := &response.GetTickerPriceResponse{}
	url := b.getTickerPriceURL
	if symbol != "" {
		url = fmt.Sprintf("%s?symbol=%s", url, symbol)
	}
	statusCode, err := util.FetchData(url, res)
	if err != nil || statusCode != http.StatusOK {
		return nil, fmt.Errorf("binance.GetTickerPrice() failed: %v", err), statusCode
	}
	return res, nil, http.StatusOK
}

func (b *Binance) GetAvgPriceBySymbol(symbol string) (*response.GetAvgPriceBySymbolResponse, error, int) {
	symbol = strings.Replace(symbol, "/", "", 1)
	res := &response.GetAvgPriceBySymbolResponse{}
	url := b.getAvgPriceURL
	if symbol != "" {
		url = fmt.Sprintf("%s?symbol=%s", url, symbol)
	}
	statusCode, err := util.FetchData(url, res)
	if err != nil || statusCode != http.StatusOK {
		return nil, fmt.Errorf("binance.GetAvgPriceBySymbol() failed: %v", err), statusCode
	}
	return res, nil, http.StatusOK
}

func (b *Binance) GetKlinesBySymbol(symbol string) ([]response.GetKlinesDataResponse, error, int) {
	symbol = strings.Replace(symbol, "/", "", 1)
	data := make([][]interface{}, 0)
	statusCode, err := util.FetchData(fmt.Sprintf(b.getSymbolKlinesURL, symbol), &data)
	if err != nil || statusCode != http.StatusOK {
		return nil, fmt.Errorf("binance.GetKlinesBySymbol() failed: %v", err), statusCode
	}
	res := make([]response.GetKlinesDataResponse, 0, len(data))
	for _, item := range data {
		res = append(res, response.GetKlinesDataResponse{
			OPrice: item[1].(string),
			HPrice: item[2].(string),
			LPrice: item[3].(string),
			CPrice: item[4].(string),
		})
	}
	return res, nil, http.StatusOK
}

func (b *Binance) GetApiKeyPermission(apiKey, apiSecret string) (*response.BinanceApiKeyPermissionResponse, error) {
	permission, err := bapdater.GetApiKeyPermission(apiKey, apiSecret)
	if err != nil {
		return nil, err
	}

	return permission, nil
}

func (b *Binance) GetUserAsset(apiKey, apiSecret string) ([]response.BinanceUserAssetResponse, error) {
	asset, err := bapdater.GetUserAsset(apiKey, apiSecret)
	if err != nil {
		return nil, err
	}

	return asset, nil
}

func (b *Binance) GetFundingAsset(apiKey, apiSecret string) ([]response.BinanceUserAssetResponse, error) {
	asset, err := bapdater.GetFundingAsset(apiKey, apiSecret)
	if err != nil {
		return nil, err
	}

	return asset, nil
}
