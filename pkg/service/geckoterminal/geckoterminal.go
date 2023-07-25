package geckoterminal

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/response"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/stealth"
)

type GeckoTerminal struct {
	chromeHost  string
	searchApi   string
	getPoolApi  string
	getPoolPage string
}

func NewService(cfg *config.Config) Service {
	return &GeckoTerminal{
		chromeHost:  cfg.ChromeHost,
		searchApi:   "https://app.geckoterminal.com/api/p1/search?query=%s",
		getPoolApi:  "https://api.geckoterminal.com/api/v2/networks/%s/pools/%s?include=base_token,quote_token,dex",
		getPoolPage: "https://www.geckoterminal.com/%s/pools/%s",
	}
}

func (g *GeckoTerminal) Search(query string) (*Search, error) {
	browser := rod.New().ControlURL(launcher.MustResolveURL(g.chromeHost)).MustConnect()
	defer browser.MustClose()
	page := stealth.MustPage(browser).MustNavigate(fmt.Sprintf(g.searchApi, query))

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

	page := stealth.MustPage(browser).MustNavigate(fmt.Sprintf(g.getPoolApi, network, poolAddr))
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

	marketCapUsd, err := strconv.ParseFloat(pool.Data.Attributes.MarketCapUsd, 64)
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
	if len(pool.Included) > 0 {
		coingeckoId = *pool.Included[0].Attributes.CoingeckoCoinID
	}

	coinResp := &response.GetCoinResponse{
		ID:              fmt.Sprintf("geckoterminal_%s", pool.Data.ID),
		CoingeckoId:     coingeckoId,
		Name:            pool.Data.Attributes.Name,
		Symbol:          pool.Data.Attributes.Name,
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

func (g *GeckoTerminal) GetPoolInfo(network, pool string) (*Pool, error) {
	var poolResp *Pool
	browser := rod.New().ControlURL(launcher.MustResolveURL(g.chromeHost)).MustConnect()
	defer browser.MustClose()

	page := stealth.MustPage(browser).MustNavigate(fmt.Sprintf(g.getPoolApi, network, pool))
	data := page.MustElement("body").MustText()

	if err := json.Unmarshal([]byte(data), &poolResp); err != nil {
		return nil, err
	}

	return poolResp, nil
}
