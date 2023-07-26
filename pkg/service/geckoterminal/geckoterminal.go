package geckoterminal

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/response"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/stealth"
)

const (
	searchApi         = "https://app.geckoterminal.com/api/p1/search?query=%s"
	getPoolApi        = "https://api.geckoterminal.com/api/v2/networks/%s/pools/%s?include=base_token,quote_token,dex"
	getPoolApiP1      = "https://app.geckoterminal.com/api/p1/%s/pools/%s?base_token=0&include=pairs&fields[pool]=pairs&fields[pair]=base_price_in_usd,base_price_in_quote,quote_price_in_usd,quote_price_in_base"
	getCandlestickApi = "https://app.geckoterminal.com/api/p1/candlesticks/%s/%s?resolution=15&from_timestamp=%d&to_timestamp=%d"
)

type GeckoTerminal struct {
	chromeHost string
}

func NewService(cfg *config.Config) Service {
	return &GeckoTerminal{
		chromeHost: cfg.ChromeHost,
	}
}

func (g *GeckoTerminal) Search(query string) (*Search, error) {
	browser := rod.New().ControlURL(launcher.MustResolveURL(g.chromeHost)).MustConnect()
	defer browser.MustClose()
	page := stealth.MustPage(browser).MustNavigate(fmt.Sprintf(searchApi, query))

	data := page.MustElement("body").MustText()

	var search *Search
	if err := json.Unmarshal([]byte(data), &search); err != nil {
		return nil, err
	}

	return search, nil
}

func (g *GeckoTerminal) GetPool(network, poolAddr string) (*response.GetCoinResponse, error) {
	var pool *Pool
	browser := rod.New().ControlURL(launcher.MustResolveURL(g.chromeHost)).MustConnect()
	defer browser.MustClose()

	page := stealth.MustPage(browser).MustNavigate(fmt.Sprintf(getPoolApi, network, poolAddr))
	data := page.MustElement("body").MustText()

	if err := json.Unmarshal([]byte(data), &pool); err != nil {
		return nil, err
	}

	search, err := g.Search(poolAddr)
	if err != nil {
		return nil, err
	}

	imageUrl := ""
	searchPool := SearchPoolElement{}
	if len(search.Data.Attributes.Pools) > 0 {
		searchPool = search.Data.Attributes.Pools[0]

		for _, token := range searchPool.Tokens {
			if token.IsBaseToken {
				if _, err := url.Parse(token.ImageURL); err != nil {
					imageUrl = token.ImageURL
				}
				break
			}
		}
	}

	baseTokenPriceUsd, err := strconv.ParseFloat(pool.Data.Attributes.BaseTokenPriceUsd, 64)
	if err != nil {
		baseTokenPriceUsd = 0
	}

	fdvUsd, err := strconv.ParseFloat(pool.Data.Attributes.FdvUsd, 64)
	if err != nil {
		fdvUsd = 0
	}

	marketCapUsd, err := strconv.ParseFloat(pool.Data.Attributes.ReserveInUsd, 64)
	if err != nil {
		marketCapUsd = 0
	}

	totalVolumeUsd, err := strconv.ParseFloat(pool.Data.Attributes.VolumeUsd.H24, 64)
	if err != nil {
		totalVolumeUsd = 0
	}

	priceChangePercentage1h, err := strconv.ParseFloat(pool.Data.Attributes.PriceChangePercentage.H1, 64)
	if err != nil {
		priceChangePercentage1h = 0
	}

	priceChangePercentage24h, err := strconv.ParseFloat(pool.Data.Attributes.PriceChangePercentage.H24, 64)
	if err != nil {
		priceChangePercentage24h = 0
	}

	coingeckoId := ""
	if len(pool.Included) > 0 && pool.Included[0].Attributes.CoingeckoCoinID != nil {
		coingeckoId = *pool.Included[0].Attributes.CoingeckoCoinID
	}

	baseToken := SearchToken{}
	if len(searchPool.Tokens) > 0 {
		for _, token := range searchPool.Tokens {
			if token.IsBaseToken {
				baseToken = token
				break
			}
		}
	}

	coinResp := &response.GetCoinResponse{
		ID:              fmt.Sprintf("geckoterminal_%s_%s", network, poolAddr),
		CoingeckoId:     coingeckoId,
		Name:            baseToken.Name,
		Symbol:          baseToken.Symbol,
		AssetPlatformID: searchPool.Dex.Identifier,
		AssetPlatform: &response.AssetPlatformResponseData{
			ID:   searchPool.Dex.Identifier,
			Name: fmt.Sprintf("%s (DEX)", searchPool.Dex.Name),
		},
		Image: response.CoinImage{
			Thumb: imageUrl,
			Small: imageUrl,
			Large: imageUrl,
		},
		MarketData: response.MarketData{
			CurrentPrice: map[string]float64{
				"usd": baseTokenPriceUsd,
			},
			MarketCap: map[string]float64{
				"usd": marketCapUsd,
			},
			MarketCapChange24hInCurrency: map[string]float64{
				"usd": 0,
			},
			MarketCapChangePercentage24hInCurrency: map[string]float64{
				"usd": 0,
			},
			TotalMarketCap: map[string]float64{
				"usd": marketCapUsd,
			},
			TotalVolume: map[string]float64{
				"usd": totalVolumeUsd,
			},
			FullyDilutedValuation: map[string]float64{
				"usd": fdvUsd,
			},
			PriceChange24hInCurrency: map[string]float64{
				"usd": 0,
			},
			PriceChangePercentage1h:  priceChangePercentage1h,
			PriceChangePercentage24h: priceChangePercentage24h,
			PriceChangePercentage1hInCurrency: map[string]float64{
				"usd": priceChangePercentage1h,
			},
			PriceChangePercentage24hInCurrency: map[string]float64{
				"usd": priceChangePercentage24h,
			},
			PriceChangePercentage7dInCurrency: map[string]float64{
				"usd": 0,
			},
			PriceChangePercentage14dInCurrency: map[string]float64{
				"usd": 0,
			},
			PriceChangePercentage30dInCurrency: map[string]float64{
				"usd": 0,
			},
			PriceChangePercentage60dInCurrency: map[string]float64{
				"usd": 0,
			},
			PriceChangePercentage200dInCurrency: map[string]float64{
				"usd": 0,
			},
			PriceChangePercentage1yInCurrency: map[string]float64{
				"usd": 0,
			},
		},
		Tickers:         []response.TickerData{},
		ContractAddress: pool.Data.Attributes.Address,
	}

	return coinResp, nil
}

func (g *GeckoTerminal) GetHistoricalMarketData(network, poolAddr string, from, to int64) (*response.HistoricalMarketChartResponse, error) {
	browser := rod.New().ControlURL(launcher.MustResolveURL(g.chromeHost)).MustConnect()
	defer browser.MustClose()

	page := stealth.MustPage(browser).MustNavigate(fmt.Sprintf(getPoolApiP1, network, poolAddr))
	data := page.MustElement("body").MustText()

	pool := &PoolP1{}
	if err := json.Unmarshal([]byte(data), &pool); err != nil {
		return nil, err
	}

	if len(pool.Data.Relationships.Pairs.Data) == 0 {
		return nil, fmt.Errorf("no pair found")
	}

	baseToken := pool.Data.ID
	quoteToken := pool.Data.Relationships.Pairs.Data[0].ID

	page = stealth.MustPage(browser).MustNavigate(fmt.Sprintf(getCandlestickApi, baseToken, quoteToken, from, to))
	data = page.MustElement("body").MustText()

	candlesticks := &Candlesticks{}

	if err := json.Unmarshal([]byte(data), &candlesticks); err != nil {
		return nil, err
	}

	prices := [][]float64{}

	for _, candlestick := range candlesticks.Data {
		price := candlestick.C
		ts, err := time.Parse("2006-01-02T15:04:05.000Z", candlestick.Dt)
		if err != nil {
			return nil, err
		}

		prices = append(prices, []float64{float64(ts.UnixMilli()), price})
	}

	resp := &response.HistoricalMarketChartResponse{
		Prices: prices,
	}

	return resp, nil
}
