package defi

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/stealth"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/repo/pg"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/service"
	mock_coingecko "github.com/defipod/mochi/pkg/service/coingecko/mocks"
	"github.com/defipod/mochi/pkg/util/testhelper"
)

func TestHandler_GetCoin(t *testing.T) {
	cfg := config.LoadTestConfig()
	db := testhelper.LoadTestDB("../../../migrations/test_seed")
	repo := pg.NewRepo(db)
	log := logger.NewLogrusLogger()
	s := pg.NewPostgresStore(&cfg)
	svc, _ := service.NewService(cfg, log)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	coingeckoMock := mock_coingecko.NewMockService(ctrl)
	svc.CoinGecko = coingeckoMock

	// rod browser
	browser := rod.New().Timeout(time.Minute).MustConnect()
	launcher.NewBrowser().MustGet()
	page := stealth.MustPage(browser)

	entityMock := entities.New(cfg, log, repo, s, nil, nil, nil, svc, nil, nil, nil, page)
	tests := []struct {
		name              string
		coinID            string
		wantCoingeckoResp *response.GetCoinResponse
		wantError         error
		wantCode          int
		wantResponsePath  string
	}{
		{
			name:   "get coin data successful",
			coinID: "ethereum",
			wantCoingeckoResp: &response.GetCoinResponse{
				ID:              "ethereum",
				Name:            "Ethereum",
				Symbol:          "eth",
				MarketCapRank:   2,
				AssetPlatformID: "",
				Image: response.CoinImage{
					Thumb: "https://assets.coingecko.com/coins/images/279/thumb/ethereum.png?1595348880",
					Small: "https://assets.coingecko.com/coins/images/279/small/ethereum.png?1595348880",
					Large: "https://assets.coingecko.com/coins/images/279/large/ethereum.png?1595348880",
				},
				MarketData: response.MarketData{
					CurrentPrice: map[string]float64{
						"usd": 1447.1,
						"vef": 144.9,
						"vnd": 33905901,
					},
					MarketCap: map[string]float64{
						"usd": 174483196724,
						"vef": 17471002488,
						"vnd": 4088172593861496,
					},
					PriceChangePercentage1hInCurrency: map[string]float64{
						"usd": -0.50332,
						"vef": -0.50332,
						"vnd": -0.50332,
					},
					PriceChangePercentage24hInCurrency: map[string]float64{
						"usd": -2.66891,
						"vef": -2.66891,
						"vnd": -2.62661,
					},
					PriceChangePercentage7dInCurrency: map[string]float64{
						"usd": -10.90085,
						"vef": -10.90085,
						"vnd": -10.8054,
					},
				},
				Tickers: []response.TickerData{
					{
						Base:         "ETH",
						Target:       "USDT",
						Last:         1444.99,
						CoinID:       "ethereum",
						TargetCoinID: "tether",
					},
					{
						Base:         "ETH",
						Target:       "USDT",
						Last:         1446.49,
						CoinID:       "ethereum",
						TargetCoinID: "tether",
					},
					{
						Base:         "ETH",
						Target:       "USDT",
						Last:         1444.8,
						CoinID:       "ethereum",
						TargetCoinID: "tether",
					},
				},
				Description: map[string]string{
					"en": "eth",
				},
			},
			wantError:        nil,
			wantCode:         200,
			wantResponsePath: "testdata/get_coin/200-success-eth.json",
		},
		{
			name:             "cannot find coin id",
			coinID:           "etherea",
			wantError:        errors.New("failed to fetch coin data of etherea: <nil>"),
			wantCode:         404,
			wantResponsePath: "testdata/get_coin/404-coin-not-found.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				entities: entityMock,
				log:      log,
			}
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest("GET", fmt.Sprintf("/api/v1/defi/coins/%s", tt.coinID), nil)
			ctx.Params = []gin.Param{
				{
					Key:   "id",
					Value: tt.coinID,
				},
			}
			coingeckoMock.EXPECT().GetCoin(tt.coinID).Return(tt.wantCoingeckoResp, tt.wantError, tt.wantCode)

			h.GetCoin(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)
			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.GetNFTDetail] response mismatched")
		})
	}
}

// func TestHandler_GetUserWatchlist(t *testing.T) {
// 	cfg := config.LoadTestConfig()
// 	db := testhelper.LoadTestDB("../../../migrations/test_seed")
// 	repo := pg.NewRepo(db)
// 	log := logger.NewLogrusLogger()
// 	s := pg.NewPostgresStore(&cfg)
// 	svc, _ := service.NewService(cfg, log)

// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	coingeckoMock := mock_coingecko.NewMockService(ctrl)
// 	svc.CoinGecko = coingeckoMock
// 	entityMock := entities.New(cfg, log, repo, s, nil, nil, nil, svc, nil, nil, nil, nil)

// 	h := &Handler{
// 		entities: entityMock,
// 		log:      log,
// 	}

// 	coingeckoDefaultTickers := testhelper.WatchlistCoingeckoDefaultTickers()
// 	tests := []struct {
// 		name                            string
// 		userId                          string
// 		defaultTickers                  []string
// 		coinGeckoDefaultTickersResponse []response.CoinMarketItemData
// 		wantError                       error
// 		wantCode                        int
// 		wantResponsePath                string
// 	}{
// 		{
// 			name:                            "success - default tickers only",
// 			userId:                          "319132138849173123",
// 			coinGeckoDefaultTickersResponse: coingeckoDefaultTickers,
// 			defaultTickers: []string{
// 				"bitcoin", "ethereum", "binancecoin", "fantom", "internet-computer", "solana", "avalanche-2", "matic-network",
// 			},
// 			wantCode:         200,
// 			wantResponsePath: "testdata/user_watchlist/get-200-ok-default.json",
// 		},
// 		{
// 			name:   "success - custom tickers",
// 			userId: "319132138849173505",
// 			coinGeckoDefaultTickersResponse: []response.CoinMarketItemData{
// 				{
// 					ID:           "dogecoin",
// 					Name:         "Dogecoin",
// 					Symbol:       "doge",
// 					CurrentPrice: 0.060343,
// 					Image:        "https://assets.coingecko.com/coins/images/5/large/dogecoin.png?1547792256",
// 					SparkLineIn7d: struct {
// 						Price []float64 "json:\"price\""
// 					}{
// 						[]float64{
// 							0.06083622964800758, 0.06103991568695497, 0.061043175854199744,
// 						},
// 					},
// 					PriceChangePercentage24h:          0.53837,
// 					PriceChangePercentage7dInCurrency: -1.6547840760734636,
// 					MarketCap:                         0,
// 					MarketCapRank:                     0,
// 					IsPair:                            false,
// 				},
// 			},
// 			defaultTickers: []string{
// 				"dogecoin",
// 			},
// 			wantCode:         200,
// 			wantResponsePath: "testdata/user_watchlist/get-200-ok-custom.json",
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			w := httptest.NewRecorder()
// 			ctx, _ := gin.CreateTestContext(w)
// 			ctx.Request = httptest.NewRequest("GET", fmt.Sprintf("/api/v1/defi/watchlist?user_id=%s&size=8", tt.userId), nil)
// 			coingeckoMock.EXPECT().GetCoinsMarketData(tt.defaultTickers, true, "1", "100").Return(tt.coinGeckoDefaultTickersResponse, nil, 200).AnyTimes()

// 			h.GetUserWatchlist(ctx)
// 			require.Equal(t, tt.wantCode, w.Code)
// 			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
// 			require.NoError(t, err)
// 			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.GetUserWatchlist] response mismatched")
// 		})
// 	}
// }

// func TestHandler_AddToWatchlist(t *testing.T) {
// 	cfg := config.LoadTestConfig()
// 	db := testhelper.LoadTestDB("../../../migrations/test_seed")
// 	repo := pg.NewRepo(db)
// 	log := logger.NewLogrusLogger()
// 	s := pg.NewPostgresStore(&cfg)
// 	svc, _ := service.NewService(cfg, log)

// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	coingeckoMock := mock_coingecko.NewMockService(ctrl)
// 	svc.CoinGecko = coingeckoMock
// 	entityMock := entities.New(cfg, log, repo, s, nil, nil, nil, svc, nil, nil, nil, nil)
// 	h := &Handler{
// 		entities: entityMock,
// 		log:      log,
// 	}
// 	tests := []struct {
// 		name                     string
// 		req                      request.AddToWatchlistRequest
// 		coingeckoSupportedTokens []model.CoingeckoSupportedTokens
// 		coinIds                  []string
// 		coinPrices               map[string]float64
// 		wantError                error
// 		wantCode                 int
// 		wantResponsePath         string
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name: "success - found suggestions",
// 			req: request.AddToWatchlistRequest{
// 				UserID: "319132138849173505",
// 				Symbol: "doge",
// 			},
// 			coingeckoSupportedTokens: []model.CoingeckoSupportedTokens{
// 				{
// 					ID:     "binance-peg-dogecoin",
// 					Symbol: "doge",
// 					Name:   "Binance-Peg Dogecoin",
// 				},
// 				{
// 					ID:     "dogecoin",
// 					Symbol: "doge",
// 					Name:   "Dogecoin",
// 				},
// 			},
// 			coinIds:          []string{"binance-peg-dogecoin", "dogecoin"},
// 			coinPrices:       map[string]float64{"binance-peg-dogecoin": 0.1, "dogecoin": 0.2},
// 			wantCode:         200,
// 			wantResponsePath: "testdata/user_watchlist/post-200-ok-suggest.json",
// 		},
// 		{
// 			name: "success - found single",
// 			req: request.AddToWatchlistRequest{
// 				UserID: "319132138849173505",
// 				Symbol: "cake",
// 			},
// 			coingeckoSupportedTokens: []model.CoingeckoSupportedTokens{
// 				{
// 					ID:     "pancakeswap-token",
// 					Symbol: "cake",
// 					Name:   "PancakeSwap",
// 				},
// 			},
// 			coinIds:          []string{"pancakeswap-token"},
// 			coinPrices:       map[string]float64{"pancakeswap-token": 1.7},
// 			wantCode:         200,
// 			wantResponsePath: "testdata/user_watchlist/200-add-single.json",
// 		},
// 		{
// 			name: "failed - token not found",
// 			req: request.AddToWatchlistRequest{
// 				UserID: "319132138849173505",
// 				Symbol: "xyzabc",
// 			},
// 			coingeckoSupportedTokens: []model.CoingeckoSupportedTokens{},
// 			wantCode:                 404,
// 			wantResponsePath:         "testdata/user_watchlist/404-not-found.json",
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			w := httptest.NewRecorder()
// 			ctx, _ := gin.CreateTestContext(w)
// 			ctx.Request = httptest.NewRequest("POST", "/api/v1/defi/watchlist", nil)
// 			util.SetRequestBody(ctx, tt.req)

// 			if tt.coinIds != nil && len(tt.coinIds) != 0 {
// 				for _, coinId := range tt.coinIds {
// 					coingeckoMock.EXPECT().GetCoinPrice([]string{coinId}, "usd").Return(map[string]float64{coinId: tt.coinPrices[coinId]}, nil).AnyTimes()
// 				}
// 			}

// 			if tt.wantCode == 200 && len(tt.coingeckoSupportedTokens) == 1 {
// 				coingeckoMock.EXPECT().GetCoin(tt.coingeckoSupportedTokens[0].ID).Return(nil, nil, 0).AnyTimes()
// 			}

// 			h.AddToWatchlist(ctx)
// 			require.Equal(t, tt.wantCode, w.Code)
// 			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
// 			require.NoError(t, err)
// 			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.AddToWatchlist] response mismatched")
// 		})
// 	}
// }

// func TestHandler_RemoveFromWatchlist(t *testing.T) {
// 	cfg := config.LoadTestConfig()
// 	db := testhelper.LoadTestDB("../../../migrations/test_seed")
// 	repo := pg.NewRepo(db)
// 	log := logger.NewLogrusLogger()
// 	s := pg.NewPostgresStore(&cfg)
// 	svc, _ := service.NewService(cfg, log)

// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	coingeckoMock := mock_coingecko.NewMockService(ctrl)
// 	svc.CoinGecko = coingeckoMock
// 	entityMock := entities.New(cfg, log, repo, s, nil, nil, nil, svc, nil, nil, nil, nil)
// 	h := &Handler{
// 		entities: entityMock,
// 		log:      log,
// 	}
// 	tests := []struct {
// 		name             string
// 		req              *request.RemoveFromWatchlistRequest
// 		wantError        error
// 		wantCode         int
// 		wantResponsePath string
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name: "success - found and delete",
// 			req: &request.RemoveFromWatchlistRequest{
// 				UserID: "319132138849173555",
// 				Symbol: "doge",
// 			},
// 			wantCode:         200,
// 			wantResponsePath: "testdata/200-data-null.json",
// 		},
// 		{
// 			name: "failed - not found",
// 			req: &request.RemoveFromWatchlistRequest{
// 				UserID: "319132138849173555",
// 				Symbol: "dogeee",
// 			},
// 			wantCode:         404,
// 			wantResponsePath: "testdata/user_watchlist/404-not-found.json",
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			w := httptest.NewRecorder()
// 			ctx, _ := gin.CreateTestContext(w)
// 			ctx.Request = httptest.NewRequest("DELETE", fmt.Sprintf("/api/v1/defi/watchlist?user_id=%s&symbol=%s", tt.req.UserID, tt.req.Symbol), nil)

// 			h.RemoveFromWatchlist(ctx)
// 			require.Equal(t, tt.wantCode, w.Code)
// 			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
// 			require.NoError(t, err)
// 			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.RemoveFromWatchlist] response mismatched")
// 		})
// 	}
// }

func TestHandler_SearchCoins(t *testing.T) {
	cfg := config.LoadTestConfig()
	db := testhelper.LoadTestDB("../../../migrations/test_seed")
	repo := pg.NewRepo(db)
	log := logger.NewLogrusLogger()
	s := pg.NewPostgresStore(&cfg)
	svc, _ := service.NewService(cfg, log)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	coingeckoMock := mock_coingecko.NewMockService(ctrl)
	svc.CoinGecko = coingeckoMock
	entityMock := entities.New(cfg, log, repo, s, nil, nil, nil, svc, nil, nil, nil, nil)
	h := &Handler{
		entities: entityMock,
		log:      log,
	}
	tests := []struct {
		name             string
		query            string
		wantCode         int
		wantResponsePath string
		coinIds          []string
		coinPrices       map[string]float64
	}{
		{
			name:             "success - get one coin",
			query:            "cake",
			wantCode:         200,
			wantResponsePath: "testdata/search_coin/200-ok-single.json",
			coinIds:          []string{"pancakeswap-token"},
			coinPrices:       map[string]float64{"pancakeswap-token": 1.7},
		},
		{
			name:             "success - get multiple coins",
			query:            "doge",
			wantCode:         200,
			wantResponsePath: "testdata/search_coin/200-ok-multiple.json",
			coinIds:          []string{"binance-peg-dogecoin", "dogecoin"},
			coinPrices:       map[string]float64{"binance-peg-dogecoin": 0.1, "dogecoin": 0.2},
		},
		{
			name:             "success - not  found",
			query:            "cakeabc",
			wantCode:         200,
			wantResponsePath: "testdata/search_coin/200-empty.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest("GET", fmt.Sprintf("/api/v1/defi/coins?query=%s", tt.query), nil)

			if tt.coinIds != nil && len(tt.coinIds) != 0 {
				for _, coinId := range tt.coinIds {
					coingeckoMock.EXPECT().GetCoinPrice([]string{coinId}, "usd").Return(map[string]float64{coinId: tt.coinPrices[coinId]}, nil).AnyTimes()
				}
			}

			h.SearchCoins(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)
			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.RemoveFromWatchlist] response mismatched")
		})
	}
}
