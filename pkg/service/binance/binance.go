package binance

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/defipod/mochi/pkg/cache"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/response"
	badapter "github.com/defipod/mochi/pkg/service/binance/adapter"
	"github.com/defipod/mochi/pkg/service/sentrygo"
	"github.com/defipod/mochi/pkg/util"
)

type Binance struct {
	getExchangeInfoURL string
	getSymbolKlinesURL string
	getAvgPriceURL     string
	getTickerPriceURL  string
	config             *config.Config
	logger             logger.Logger
	cache              cache.Cache
	sentry             sentrygo.Service
}

func NewService(cfg *config.Config, l logger.Logger, cache cache.Cache, sentry sentrygo.Service) Service {
	return &Binance{
		getExchangeInfoURL: "https://api.binance.com/api/v3/exchangeInfo",
		getSymbolKlinesURL: "https://api.binance.com/api/v3/uiKlines?symbol=%s&interval=1h&limit=168", // 168h = 7d
		getAvgPriceURL:     "https://api.binance.com/api/v3/avgPrice",
		getTickerPriceURL:  "https://api.binance.com/api/v3/ticker/price",
		config:             cfg,
		logger:             l,
		cache:              cache,
		sentry:             sentry,
	}
}

var (
	sentryTags = map[string]string{
		"type": "system",
	}
)

func (b *Binance) GetExchangeInfo(symbol string) (*response.GetExchangeInfoResponse, error, int) {
	b.logger.Debug("start binance.GetExchangeInfo()")
	defer b.logger.Debug("end binance.GetExchangeInfo()")

	symbol = strings.Replace(symbol, "/", "", 1)
	res := &response.GetExchangeInfoResponse{}

	// check if data cached
	key := fmt.Sprintf("binance-exchange-info-symbol-%s", strings.ToLower(symbol))
	cached, err := b.cache.GetString(key)
	if err == nil && cached != "" {
		b.logger.Infof("hit cache data binance-service, key: %s", key)
		return res, json.Unmarshal([]byte(cached), res), http.StatusOK
	}

	url := b.getExchangeInfoURL
	if symbol != "" {
		url = fmt.Sprintf("%s?symbol=%s", url, symbol)
	}

	statusCode, err := util.FetchData(url, res)
	if err != nil || statusCode != http.StatusOK {
		b.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
			Message: fmt.Sprintf("[API mochi] - Binance - GetExchangeInfo failed %v", err),
			Tags:    sentryTags,
			Extra: map[string]interface{}{
				"symbols": symbol,
			},
		})
		return nil, fmt.Errorf("binance.GetExchangeInfo() failed: %v", err), statusCode
	}

	// cache binance-exchange-info-symbol
	// if error occurs -> ignore
	bytes, _ := json.Marshal(res)
	b.logger.Infof("cache data binance-service, key: %s", key)
	b.cache.Set(key, string(bytes), 30*time.Minute)

	return res, nil, http.StatusOK
}

func (b *Binance) GetTickerPrice(symbol string) (*response.GetTickerPriceResponse, error, int) {
	b.logger.Debug("start binance.GetTickerPrice()")
	defer b.logger.Debug("end binance.GetTickerPrice()")

	symbol = strings.Replace(symbol, "/", "", 1)
	res := &response.GetTickerPriceResponse{}

	// check if data cached
	key := fmt.Sprintf("binance-ticker-price-symbol-%s", strings.ToLower(symbol))
	cached, err := b.cache.GetString(key)
	if err == nil && cached != "" {
		b.logger.Infof("hit cache data binance-service, key: %s", key)
		return res, json.Unmarshal([]byte(cached), res), http.StatusOK
	}

	url := b.getTickerPriceURL
	if symbol != "" {
		url = fmt.Sprintf("%s?symbol=%s", url, symbol)
	}

	statusCode, err := util.FetchData(url, res)
	if err != nil || statusCode != http.StatusOK {
		b.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
			Message: fmt.Sprintf("[API mochi] - Binance - GetTickerPrice failed %v", err),
			Tags:    sentryTags,
			Extra: map[string]interface{}{
				"symbols": symbol,
			},
		})
		return nil, fmt.Errorf("binance.GetTickerPrice() failed: %v", err), statusCode
	}

	// cache binance-ticker-price-symbol
	// if error occurs -> ignore
	bytes, _ := json.Marshal(res)
	b.logger.Infof("cache data binance-service, key: %s", key)
	b.cache.Set(key, string(bytes), 30*time.Minute)

	return res, nil, http.StatusOK
}

func (b *Binance) GetAvgPriceBySymbol(symbol string) (*response.GetAvgPriceBySymbolResponse, error, int) {
	b.logger.Debug("start binance.GetAvgPriceBySymbol()")
	defer b.logger.Debug("end binance.GetAvgPriceBySymbol()")
	symbol = strings.Replace(symbol, "/", "", 1)
	res := &response.GetAvgPriceBySymbolResponse{}

	// check if data cached
	key := fmt.Sprintf("binance-avg-price-symbol-%s", strings.ToLower(symbol))
	cached, err := b.cache.GetString(key)
	if err == nil && cached != "" {
		b.logger.Infof("hit cache data binance-service, key: %s", key)
		return res, json.Unmarshal([]byte(cached), res), http.StatusOK
	}

	url := b.getAvgPriceURL
	if symbol != "" {
		url = fmt.Sprintf("%s?symbol=%s", url, symbol)
	}

	statusCode, err := util.FetchData(url, res)
	if err != nil || statusCode != http.StatusOK {
		b.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
			Message: fmt.Sprintf("[API mochi] - Binance - GetAvgPriceBySymbol failed %v", err),
			Tags:    sentryTags,
			Extra: map[string]interface{}{
				"symbols": symbol,
			},
		})
		return nil, fmt.Errorf("binance.GetAvgPriceBySymbol() failed: %v", err), statusCode
	}

	// cache binance-avg-price-symbol
	// if error occurs -> ignore
	bytes, _ := json.Marshal(res)
	b.logger.Infof("cache data binance-service, key: %s", key)
	b.cache.Set(key, string(bytes), 30*time.Minute)

	return res, nil, http.StatusOK
}

func (b *Binance) GetKlinesBySymbol(symbol string) ([]response.GetKlinesDataResponse, error, int) {
	b.logger.Debug("start binance.GetKlinesBySymbol()")
	defer b.logger.Debug("end binance.GetKlinesBySymbol()")

	symbol = strings.Replace(symbol, "/", "", 1)
	data := make([][]interface{}, 0)
	var res []response.GetKlinesDataResponse

	// check if data cached
	key := fmt.Sprintf("binance-klines-symbol-%s", strings.ToLower(symbol))
	cached, err := b.cache.GetString(key)
	if err == nil && cached != "" {
		b.logger.Infof("hit cache data binance-service, key: %s", key)
		return res, json.Unmarshal([]byte(cached), res), http.StatusOK
	}

	statusCode, err := util.FetchData(fmt.Sprintf(b.getSymbolKlinesURL, symbol), &data)
	if err != nil || statusCode != http.StatusOK {
		b.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
			Message: fmt.Sprintf("[API mochi] - Binance - GetKlinesBySymbol failed %v", err),
			Tags:    sentryTags,
			Extra: map[string]interface{}{
				"symbols": symbol,
			},
		})
		return nil, fmt.Errorf("binance.GetKlinesBySymbol() failed: %v", err), statusCode
	}

	for _, item := range data {
		res = append(res, response.GetKlinesDataResponse{
			OPrice: item[1].(string),
			HPrice: item[2].(string),
			LPrice: item[3].(string),
			CPrice: item[4].(string),
		})
	}

	// cache binance-klines-symbol
	// if error occurs -> ignore
	bytes, _ := json.Marshal(res)
	b.logger.Infof("cache data binance-service, key: %s", key)
	b.cache.Set(key, string(bytes), 30*time.Minute)

	return res, nil, http.StatusOK
}

func (b *Binance) GetApiKeyPermission(apiKey, apiSecret string) (*response.BinanceApiKeyPermissionResponse, error) {
	permission, err := badapter.GetApiKeyPermission(apiKey, apiSecret)
	if err != nil {
		return nil, err
	}

	return permission, nil
}

func (b *Binance) GetUserAsset(apiKey, apiSecret string) ([]response.BinanceUserAssetResponse, error) {
	b.logger.Debug("start binance.GetUserAsset()")
	defer b.logger.Debug("end binance.GetUserAsset()")

	var res []response.BinanceUserAssetResponse
	// check if data cached
	key := fmt.Sprintf("binance-user-asset-apikey-%s", strings.ToLower(apiKey))
	cached, err := b.cache.GetString(key)
	if err == nil && cached != "" {
		b.logger.Infof("hit cache data binance-service, key: %s", key)
		return res, json.Unmarshal([]byte(cached), &res)
	}

	res, err = badapter.GetUserAsset(apiKey, apiSecret)
	if err != nil {
		b.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
			Message: fmt.Sprintf("[API mochi] - Binance - GetUserAsset failed %v", err),
			Tags:    sentryTags,
			Extra: map[string]interface{}{
				"apiKey":    apiKey,
				"apiSecret": apiSecret,
			},
		})
		return nil, err
	}

	// cache binance-user-asset-apikey
	// if error occurs -> ignore
	bytes, _ := json.Marshal(res)
	b.logger.Infof("cache data binance-service, key: %s", key)
	b.cache.Set(key, string(bytes), 30*time.Minute)

	return res, nil
}

func (b *Binance) GetFundingAsset(apiKey, apiSecret string) ([]response.BinanceUserAssetResponse, error) {
	b.logger.Debug("start binance.GetFundingAsset()")
	defer b.logger.Debug("end binance.GetFundingAsset()")

	var res []response.BinanceUserAssetResponse

	// check if data cached
	key := fmt.Sprintf("binance-funding-asset-apikey-%s", strings.ToLower(apiKey))
	cached, err := b.cache.GetString(key)
	if err == nil && cached != "" {
		b.logger.Infof("hit cache data binance-service, key: %s", key)
		return res, json.Unmarshal([]byte(cached), &res)
	}

	res, err = badapter.GetFundingAsset(apiKey, apiSecret)
	if err != nil {
		b.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
			Message: fmt.Sprintf("[API mochi] - Binance - GetFundingAsset failed %v", err),
			Tags:    sentryTags,
			Extra: map[string]interface{}{
				"apiKey":    apiKey,
				"apiSecret": apiSecret,
			},
		})
		return nil, err
	}

	// cache binance-funding-asset-apikey
	// if error occurs -> ignore
	bytes, _ := json.Marshal(res)
	b.logger.Infof("cache data binance-service, key: %s", key)
	b.cache.Set(key, string(bytes), 30*time.Minute)

	return res, nil
}

func (b *Binance) GetStakingProductPosition(apiKey, apiSecret string) ([]response.BinanceStakingProductPosition, error) {
	b.logger.Debug("start binance.GetStakingProductPosition()")
	defer b.logger.Debug("end binance.GetStakingProductPosition()")

	var res []response.BinanceStakingProductPosition

	// check if data cached
	key := fmt.Sprintf("binance-staking-production-position-apikey-%s", strings.ToLower(apiKey))
	cached, err := b.cache.GetString(key)
	if err == nil && cached != "" {
		b.logger.Infof("hit cache data binance-service, key: %s", key)
		return res, json.Unmarshal([]byte(cached), &res)
	}

	res, err = badapter.GetStakingProductPosition(apiKey, apiSecret)
	if err != nil {
		b.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
			Message: fmt.Sprintf("[API mochi] - Binance - GetStakingProductPosition failed %v", err),
			Tags:    sentryTags,
			Extra: map[string]interface{}{
				"apiKey":    apiKey,
				"apiSecret": apiSecret,
			},
		})
		return nil, err
	}

	// cache binance-staking-production-position-apikey
	// if error occurs -> ignore
	bytes, _ := json.Marshal(res)
	b.logger.Infof("cache data binance-service, key: %s", key)
	b.cache.Set(key, string(bytes), 30*time.Minute)

	return res, nil
}

func (b *Binance) GetLendingAccount(apiKey, apiSecret string) (*response.BinanceLendingAccount, error) {
	b.logger.Debug("start binance.GetLendingAccount()")
	defer b.logger.Debug("end binance.GetLendingAccount()")

	res := &response.BinanceLendingAccount{}

	// check if data cached
	key := fmt.Sprintf("binance-lending-account-apikey-%s", strings.ToLower(apiKey))
	cached, err := b.cache.GetString(key)
	if err == nil && cached != "" {
		b.logger.Infof("hit cache data binance-service, key: %s", key)
		return res, json.Unmarshal([]byte(cached), &res)
	}

	res, err = badapter.GetLendingAccount(apiKey, apiSecret)
	if err != nil {
		b.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
			Message: fmt.Sprintf("[API mochi] - Binance - GetLendingAccount failed %v", err),
			Tags:    sentryTags,
			Extra: map[string]interface{}{
				"apiKey":    apiKey,
				"apiSecret": apiSecret,
			},
		})
		return nil, err
	}

	// cache binance-lending-account-apikey
	// if error occurs -> ignore
	bytes, _ := json.Marshal(res)
	b.logger.Infof("cache data binance-service, key: %s", key)
	b.cache.Set(key, string(bytes), 30*time.Minute)

	return res, nil
}

func (b *Binance) GetSimpleEarn(apiKey, apiSecret string) (*response.BinanceSimpleEarnAccount, error) {
	b.logger.Debug("start binance.GetSimpleEarn()")
	defer b.logger.Debug("end binance.GetSimpleEarn()")

	res := &response.BinanceSimpleEarnAccount{}

	// check if data cached
	key := fmt.Sprintf("binance-simple-earn-apikey-%s", strings.ToLower(apiKey))
	cached, err := b.cache.GetString(key)
	if err == nil && cached != "" {
		b.logger.Infof("hit cache data binance-service, key: %s", key)
		return res, json.Unmarshal([]byte(cached), &res)
	}

	res, err = badapter.GetSimpleEarnAccount(apiKey, apiSecret)
	if err != nil {
		b.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
			Message: fmt.Sprintf("[API mochi] - Binance - GetSimpleEarn failed %v", err),
			Tags:    sentryTags,
			Extra: map[string]interface{}{
				"apiKey":    apiKey,
				"apiSecret": apiSecret,
			},
		})
		return nil, err
	}

	// cache binance-simple-earn-apikey
	// if error occurs -> ignore
	bytes, _ := json.Marshal(res)
	b.logger.Infof("cache data binance-service, key: %s", key)
	b.cache.Set(key, string(bytes), 30*time.Minute)

	return res, nil
}

func (b *Binance) GetFutureAccountBalance(apiKey, apiSecret string) ([]response.BinanceFutureAccountBalance, error) {
	b.logger.Debug("start binance.GetFutureAccountBalance()")
	defer b.logger.Debug("end binance.GetFutureAccountBalance()")

	res := []response.BinanceFutureAccountBalance{}

	// check if data cached
	key := fmt.Sprintf("binance-future-account-balance-apikey-%s", strings.ToLower(apiKey))
	cached, err := b.cache.GetString(key)
	if err == nil && cached != "" {
		b.logger.Infof("hit cache data binance-service, key: %s", key)
		return res, json.Unmarshal([]byte(cached), &res)
	}

	res, err = badapter.GetFutureAccountBalance(apiKey, apiSecret)
	if err != nil {
		b.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
			Message: fmt.Sprintf("[API mochi] - Binance - GetFutureAccountBalance failed %v", err),
			Tags:    sentryTags,
			Extra: map[string]interface{}{
				"apiKey":    apiKey,
				"apiSecret": apiSecret,
			},
		})
		return nil, err
	}

	// cache binance-future-account-balance-apikey
	// if error occurs -> ignore
	bytes, _ := json.Marshal(res)
	b.logger.Infof("cache data binance-service, key: %s", key)
	b.cache.Set(key, string(bytes), 30*time.Minute)

	return res, nil
}

func (b *Binance) GetFutureAccount(apiKey, apiSecret string) (*response.BinanceFutureAccount, error) {
	b.logger.Debug("start binance.GetFutureAccount()")
	defer b.logger.Debug("end binance.GetFutureAccount()")

	res := &response.BinanceFutureAccount{}

	// check if data cached
	key := fmt.Sprintf("binance-future-account-apikey-%s", strings.ToLower(apiKey))
	cached, err := b.cache.GetString(key)
	if err == nil && cached != "" {
		b.logger.Infof("hit cache data binance-service, key: %s", key)
		return res, json.Unmarshal([]byte(cached), &res)
	}

	res, err = badapter.GetFutureAccount(apiKey, apiSecret)
	if err != nil {
		b.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
			Message: fmt.Sprintf("[API mochi] - Binance - GetFutureAccount failed %v", err),
			Tags:    sentryTags,
			Extra: map[string]interface{}{
				"apiKey":    apiKey,
				"apiSecret": apiSecret,
			},
		})
		return nil, err
	}

	// cache binance-future-account-balance-apikey
	// if error occurs -> ignore
	bytes, _ := json.Marshal(res)
	b.logger.Infof("cache data binance-service, key: %s", key)
	b.cache.Set(key, string(bytes), 30*time.Minute)

	return res, nil
}

func (b *Binance) GetFutureAccountInfo(apiKey, apiSecret string) ([]response.BinanceFuturePositionInfo, error) {
	b.logger.Debug("start binance.GetFutureAccountInfo()")
	defer b.logger.Debug("end binance.GetFutureAccountInfo()")

	res := []response.BinanceFuturePositionInfo{}

	// check if data cached
	key := fmt.Sprintf("binance-future-account-info-apikey-%s", strings.ToLower(apiKey))
	cached, err := b.cache.GetString(key)
	if err == nil && cached != "" {
		b.logger.Infof("hit cache data binance-service, key: %s", key)
		return res, json.Unmarshal([]byte(cached), &res)
	}

	res, err = badapter.GetFutureAccountInfo(apiKey, apiSecret)
	if err != nil {
		b.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
			Message: fmt.Sprintf("[API mochi] - Binance - GetFutureAccountInfo failed %v", err),
			Tags:    sentryTags,
			Extra: map[string]interface{}{
				"apiKey":    apiKey,
				"apiSecret": apiSecret,
			},
		})
		return nil, err
	}

	// cache binance-future-account-info-apikey
	// if error occurs -> ignore
	bytes, _ := json.Marshal(res)
	b.logger.Infof("cache data binance-service, key: %s", key)
	b.cache.Set(key, string(bytes), 30*time.Minute)

	return res, nil
}
