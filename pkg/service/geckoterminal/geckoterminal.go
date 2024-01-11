package geckoterminal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/stealth"

	"github.com/defipod/mochi/pkg/cache"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model/errors"
	errs "github.com/defipod/mochi/pkg/model/errors"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
)

const (
	searchApi    = "https://app.geckoterminal.com/api/p1/search?query=%s"
	getPoolApi   = "https://api.geckoterminal.com/api/v2/networks/%s/pools/%s?include=base_token,quote_token,dex"
	getPoolApiP1 = "https://app.geckoterminal.com/api/p1/%s/pools/%s?base_token=0&include=pairs&fields[pool]=pairs&fields[pair]=base_price_in_usd,base_price_in_quote,quote_price_in_usd,quote_price_in_base"
	// getCandlestickApi = "https://app.geckoterminal.com/api/p1/candlesticks/%s/%s?resolution=15&from_timestamp=%d&to_timestamp=%d"
	getCandlestickApi = "https://api.geckoterminal.com/api/v2/networks/%s/pools/%s/ohlcv/minute?aggregate=15&before_timestamp=%d&limit=1000&currency=usd&token=base"
	getTokenByAddress = "https://api.geckoterminal.com/api/v2/networks/%s/tokens/%s"

	getTokenByAddressCacheKey = "geckoterminal-get-token-by-address"
)

type GeckoTerminal struct {
	chromeHost string
	cache      cache.Cache
	logger     logger.Logger
}

func NewService(cfg *config.Config, l logger.Logger, cache cache.Cache) Service {
	return &GeckoTerminal{
		chromeHost: cfg.ChromeHost,
		cache:      cache,
		logger:     l,
	}
}

func (g *GeckoTerminal) Search(query string) (*Search, error) {
	browser := rod.New().ControlURL(launcher.MustResolveURL(g.chromeHost)).MustConnect()
	defer browser.MustClose()
	page := stealth.MustPage(browser).MustNavigate(fmt.Sprintf(searchApi, query))

	data := page.MustWaitStable().MustElement("body").MustText()

	var search *Search
	if err := json.Unmarshal([]byte(data), &search); err != nil {
		g.logger.Infof("invalid raw data from geckoterminal: %s", data)
		return nil, err
	}

	return search, nil
}

func (g *GeckoTerminal) GetPool(network, poolAddr string) (*response.GetCoinResponse, error) {
	var pool *Pool
	reqUrl := fmt.Sprintf(getPoolApi, network, poolAddr)

	browser := rod.New().ControlURL(launcher.MustResolveURL(g.chromeHost)).MustConnect()
	defer browser.MustClose()

	page := stealth.MustPage(browser).MustNavigate(reqUrl)
	data := page.MustElement("body").MustText()

	if err := json.Unmarshal([]byte(data), &pool); err != nil {
		return nil, err
	}

	// status, err := util.FetchData(reqUrl, pool)
	// if err != nil {
	// 	return nil, err
	// }

	// if status != 200 {
	// 	return nil, fmt.Errorf("status: %d", status)
	// }

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
		AssetPlatformID: searchPool.Network.Identifier,
		AssetPlatform: &response.AssetPlatformResponseData{
			ID:   searchPool.Network.Identifier,
			Name: searchPool.Network.Name,
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

func (g *GeckoTerminal) GetHistoricalMarketData(network, poolAddr string, before int64) (*response.HistoricalMarketChartResponse, error) {
	browser := rod.New().ControlURL(launcher.MustResolveURL(g.chromeHost)).MustConnect()
	defer browser.MustClose()

	// page := stealth.MustPage(browser).MustNavigate(fmt.Sprintf(getPoolApiP1, network, poolAddr))
	// data := page.MustElement("body").MustText()

	// pool := &PoolP1{}
	// if err := json.Unmarshal([]byte(data), &pool); err != nil {
	// 	return nil, err
	// }

	// if len(pool.Data.Relationships.Pairs.Data) == 0 {
	// 	return nil, fmt.Errorf("no pair found")
	// }

	// baseToken := pool.Data.ID
	// quoteToken := pool.Data.Relationships.Pairs.Data[0].ID

	page := stealth.MustPage(browser).MustNavigate(fmt.Sprintf(getCandlestickApi, network, poolAddr, before))
	data := page.MustElement("body").MustText()

	candlesticks := &Candlesticks{}

	if err := json.Unmarshal([]byte(data), &candlesticks); err != nil {
		return nil, err
	}

	if len(candlesticks.Data.Attributes.OhlcvList) == 0 {
		return nil, errors.ErrTokenNotSupportedYet
	}

	prices := [][]float64{}
	marketCaps := [][]float64{}

	for i := len(candlesticks.Data.Attributes.OhlcvList) - 1; i >= 0; i-- {
		points := candlesticks.Data.Attributes.OhlcvList[i]

		if len(points) != 6 {
			continue
		}

		tsMillis := points[0] * 1000
		open := points[1]
		// high := points[2]
		// low := points[3]
		// close := points[4]
		volume := points[5]

		prices = append(prices, []float64{tsMillis, open})
		marketCaps = append(marketCaps, []float64{tsMillis, volume})
	}

	resp := &response.HistoricalMarketChartResponse{
		Prices:     prices,
		MarketCaps: marketCaps,
	}

	return resp, nil
}

func (g *GeckoTerminal) GetTokenByAddress(chain, address string) (*response.GeckoTerminalTokensResponse, error) {
	var data response.GeckoTerminalTokensResponse

	cached, err := g.doCacheGetTokenByAddress(chain, address)
	if err == nil && cached != "" {
		g.logger.Infof("hit cache data geckoterminal-service, address: %s", address)
		go g.doNetworkGetTokenByAddress(chain, address)
		return &data, json.Unmarshal([]byte(cached), &data)
	}

	return g.doNetworkGetTokenByAddress(chain, address)
}

func (g *GeckoTerminal) doCacheGetTokenByAddress(chain, address string) (string, error) {
	return g.cache.GetString(fmt.Sprintf("%s-%s-%s", getTokenByAddressCacheKey, chain, strings.ToLower(address)))
}

func (c *GeckoTerminal) doNetworkGetTokenByAddress(chain, address string) (*response.GeckoTerminalTokensResponse, error) {
	var res response.GeckoTerminalTokensResponse
	url := fmt.Sprintf(getTokenByAddress, chain, address)
	status, err := util.FetchData(url, &res)
	if err != nil || status != http.StatusOK {
		if status != http.StatusNotFound {
			return nil, fmt.Errorf("failed to fetch token info with status %d: %v", status, err)
		}
		return nil, errs.ErrRecordNotFound
	}

	bytes, _ := json.Marshal(&res)
	c.cache.Set(fmt.Sprintf("%s-%s-%s", getTokenByAddressCacheKey, chain, strings.ToLower(address)), string(bytes), 7*24*time.Hour)

	return &res, nil
}
