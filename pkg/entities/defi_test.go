package entities

import (
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	coingeckosupportedtokens "github.com/defipod/mochi/pkg/repo/coingecko_supported_tokens"
	mock_coingeckosupportedtokens "github.com/defipod/mochi/pkg/repo/coingecko_supported_tokens/mocks"
	"github.com/defipod/mochi/pkg/repo/pg"
	userwatchlistitem "github.com/defipod/mochi/pkg/repo/user_watchlist_item"
	mock_userwatchlistitem "github.com/defipod/mochi/pkg/repo/user_watchlist_item/mocks"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/service"
	mock_coingecko "github.com/defipod/mochi/pkg/service/coingecko/mocks"
	"github.com/defipod/mochi/pkg/util/testhelper"
)

func TestEntity_GetUserWatchlist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := config.LoadTestConfig()

	log := logger.NewLogrusLogger()
	svc, _ := service.NewService(cfg, log)
	mockUserWatchlistItems := mock_userwatchlistitem.NewMockStore(ctrl)
	mockServiceCoingecko := mock_coingecko.NewMockService(ctrl)
	s := pg.NewPostgresStore(&cfg)
	r := pg.NewRepo(s.DB())
	r.UserWatchlistItem = mockUserWatchlistItems
	svc.CoinGecko = mockServiceCoingecko
	e := &Entity{
		repo: r,
		log:  log,
		svc:  svc,
		cfg:  cfg,
	}

	coingeckoDefaultTickers := testhelper.WatchlistCoingeckoDefaultTickers()

	tests := []struct {
		name         string
		req          request.GetUserWatchlistRequest
		userItems    []model.UserWatchlistItem
		itemsStr     []string
		coingeckoRes []response.CoinMarketItemData
		want         *response.GetWatchlistResponse
		wantErr      bool
	}{
		{
			name: "success - default tokens",
			req: request.GetUserWatchlistRequest{
				UserID: "319132138849173123",
				Size:   7,
			},
			itemsStr: []string{"bitcoin", "ethereum", "binancecoin", "fantom", "internet-computer", "solana", "avalanche-2", "matic-network"},
			want: &response.GetWatchlistResponse{
				Pagination: &response.PaginationResponse{
					Total: 0,
					Pagination: model.Pagination{
						Page: 0,
						Size: 7,
					},
				},
				Data: coingeckoDefaultTickers,
			},
			coingeckoRes: coingeckoDefaultTickers,
		},
		{
			name: "success - custom tokens",
			req: request.GetUserWatchlistRequest{
				UserID: "319132138849173555",
				Size:   7,
			},
			userItems: []model.UserWatchlistItem{
				{
					UserID:      "319132138849173555",
					Symbol:      "btc",
					CoinGeckoID: "bitcoin",
				},
			},
			itemsStr: []string{"bitcoin"},
			coingeckoRes: []response.CoinMarketItemData{
				{
					ID:           "bitcoin",
					Name:         "Bitcoin",
					Symbol:       "btc",
					CurrentPrice: 19238.62,
					Image:        "https://assets.coingecko.com/coins/images/1/large/bitcoin.png?1547033579",
					SparkLineIn7d: struct {
						Price []float64 "json:\"price\""
					}{
						[]float64{
							18861.78098812819, 19087.709424228065, 19165.140362664395,
						},
					},
					PriceChangePercentage24h:          0.33071,
					PriceChangePercentage7dInCurrency: 2.281861441149524,
					IsPair:                            false,
				},
			},
			want: &response.GetWatchlistResponse{
				Pagination: &response.PaginationResponse{
					Total: 1,
					Pagination: model.Pagination{
						Page: 0,
						Size: 7,
					},
				},
				Data: []response.CoinMarketItemData{
					{
						ID:           "bitcoin",
						Name:         "Bitcoin",
						Symbol:       "btc",
						CurrentPrice: 19238.62,
						Image:        "https://assets.coingecko.com/coins/images/1/large/bitcoin.png?1547033579",
						SparkLineIn7d: struct {
							Price []float64 "json:\"price\""
						}{
							[]float64{
								18861.78098812819, 19087.709424228065, 19165.140362664395,
							},
						},
						PriceChangePercentage24h:          0.33071,
						PriceChangePercentage7dInCurrency: 2.281861441149524,
						IsPair:                            false,
					},
				},
			},
		},
		{
			name: "success - custom tokens with pairs",
			req: request.GetUserWatchlistRequest{
				UserID: "319132138849173555",
				Size:   7,
			},
			userItems: []model.UserWatchlistItem{
				{
					UserID:      "319132138849173555",
					Symbol:      "btc/doge",
					CoinGeckoID: "bitcoin/dogecoin",
				},
			},
			itemsStr: []string{"bitcoin/doge"},
			coingeckoRes: []response.CoinMarketItemData{
				{
					ID:           "bitcoin",
					Name:         "Bitcoin",
					Symbol:       "btc",
					CurrentPrice: 19238.62,
					Image:        "https://assets.coingecko.com/coins/images/1/large/bitcoin.png?1547033579",
					SparkLineIn7d: struct {
						Price []float64 "json:\"price\""
					}{
						[]float64{
							18861.78098812819, 19087.709424228065, 19165.140362664395,
						},
					},
					PriceChangePercentage24h:          0.33071,
					PriceChangePercentage7dInCurrency: 2.281861441149524,
					IsPair:                            false,
				},
			},
			want: &response.GetWatchlistResponse{
				Pagination: &response.PaginationResponse{
					Total: 1,
					Pagination: model.Pagination{
						Page: 0,
						Size: 7,
					},
				},
				Data: []response.CoinMarketItemData{
					{
						ID:           "bitcoin",
						Name:         "Bitcoin",
						Symbol:       "btc",
						CurrentPrice: 19238.62,
						Image:        "https://assets.coingecko.com/coins/images/1/large/bitcoin.png?1547033579",
						SparkLineIn7d: struct {
							Price []float64 "json:\"price\""
						}{
							[]float64{
								18861.78098812819, 19087.709424228065, 19165.140362664395,
							},
						},
						PriceChangePercentage24h:          0.33071,
						PriceChangePercentage7dInCurrency: 2.281861441149524,
						IsPair:                            false,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserWatchlistItems.EXPECT().List(userwatchlistitem.UserWatchlistQuery{
				UserID: tt.req.UserID,
				Offset: tt.req.Page * tt.req.Size,
				Limit:  tt.req.Size,
			}).Return(tt.userItems, int64(len(tt.userItems)), nil).AnyTimes()
			mockServiceCoingecko.EXPECT().GetCoinsMarketData(tt.itemsStr, true, "1", "100").Return(tt.coingeckoRes, nil, 200).AnyTimes()

			got, err := e.GetUserWatchlist(tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.GetUserWatchlist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entity.GetUserWatchlist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntity_AddToWatchlist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := config.LoadTestConfig()

	log := logger.NewLogrusLogger()
	svc, _ := service.NewService(cfg, log)
	mockUserWatchlistItems := mock_userwatchlistitem.NewMockStore(ctrl)
	mockServiceCoingecko := mock_coingecko.NewMockService(ctrl)
	mockCoingeckoSupportToken := mock_coingeckosupportedtokens.NewMockStore(ctrl)
	s := pg.NewPostgresStore(&cfg)
	r := pg.NewRepo(s.DB())
	r.UserWatchlistItem = mockUserWatchlistItems
	r.CoingeckoSupportedTokens = mockCoingeckoSupportToken
	svc.CoinGecko = mockServiceCoingecko

	e := &Entity{
		repo: r,
		log:  log,
		svc:  svc,
		cfg:  cfg,
	}
	tests := []struct {
		name                         string
		req                          request.AddToWatchlistRequest
		want                         response.AddToWatchlistResponse
		coingeckoSupportedTokenFound model.CoingeckoSupportedTokens
		coingeckoSupportedTokenError error
		coingeckoSupportedTokens     []model.CoingeckoSupportedTokens
		coinIds                      []string
		coinPrices                   map[string]float64
		wantErr                      bool
	}{
		// TODO: Add test cases.
		{
			name: "success - found suggestions",
			req: request.AddToWatchlistRequest{
				UserID: "319132138849173555",
				Symbol: "doge",
			},
			coingeckoSupportedTokenError: gorm.ErrRecordNotFound,
			coingeckoSupportedTokens: []model.CoingeckoSupportedTokens{
				{
					ID:     "binance-peg-dogecoin",
					Symbol: "doge",
					Name:   "Binance-Peg Dogecoin",
				},
				{
					ID:     "dogecoin",
					Symbol: "doge",
					Name:   "Dogecoin",
				},
			},
			coinIds:    []string{"binance-peg-dogecoin", "dogecoin"},
			coinPrices: map[string]float64{"binance-peg-dogecoin": 0.1, "dogecoin": 0.2},
			want: response.AddToWatchlistResponse{
				Data: &response.AddToWatchlistResponseData{
					BaseSuggestions: []model.CoingeckoSupportedTokens{
						{
							ID:           "binance-peg-dogecoin",
							Symbol:       "doge",
							Name:         "Binance-Peg Dogecoin",
							CurrentPrice: 0.1,
						},
						{
							ID:           "dogecoin",
							Symbol:       "doge",
							Name:         "Dogecoin",
							CurrentPrice: 0.2,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "success - found one",
			req: request.AddToWatchlistRequest{
				UserID: "319132138849173555",
				Symbol: "cake",
			},
			coinIds:                      []string{"pancakeswap-token"},
			coinPrices:                   map[string]float64{"pancakeswap-token": 0.1},
			coingeckoSupportedTokenError: nil,
			coingeckoSupportedTokenFound: model.CoingeckoSupportedTokens{
				ID:     "pancakeswap-token",
				Symbol: "cake",
				Name:   "PancakeSwap",
			},
			want: response.AddToWatchlistResponse{
				Data: &response.AddToWatchlistResponseData{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCoingeckoSupportToken.EXPECT().GetOne(tt.req.Symbol).Return(&tt.coingeckoSupportedTokenFound, tt.coingeckoSupportedTokenError).AnyTimes()
			mockCoingeckoSupportToken.EXPECT().List(coingeckosupportedtokens.ListQuery{Symbol: tt.req.Symbol}).Return(tt.coingeckoSupportedTokens, nil).AnyTimes()
			mockUserWatchlistItems.EXPECT().List(userwatchlistitem.UserWatchlistQuery{CoinGeckoID: tt.coingeckoSupportedTokenFound.ID, UserID: tt.req.UserID}).Return(nil, int64(0), nil).AnyTimes()
			mockUserWatchlistItems.EXPECT().Create(&model.UserWatchlistItem{
				UserID:      tt.req.UserID,
				Symbol:      tt.req.Symbol,
				CoinGeckoID: tt.coingeckoSupportedTokenFound.ID,
			}).Return(nil).AnyTimes()

			if tt.coinIds != nil && len(tt.coinIds) != 0 {
				for _, coinId := range tt.coinIds {
					mockServiceCoingecko.EXPECT().GetCoinPrice([]string{coinId}, "usd").Return(map[string]float64{coinId: tt.coinPrices[coinId]}, nil).AnyTimes()
				}
			}

			// if found one only -> get coin price
			if tt.coingeckoSupportedTokenFound.ID != "" && !tt.wantErr {
				mockServiceCoingecko.EXPECT().GetCoin(tt.coingeckoSupportedTokenFound.ID).Return(nil, nil, 0).AnyTimes()
			}

			got, err := e.AddToWatchlist(tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.AddToWatchlist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, &tt.want) {
				t.Errorf("Entity.AddToWatchlist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntity_SearchCoins(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := config.LoadTestConfig()

	log := logger.NewLogrusLogger()
	svc, _ := service.NewService(cfg, log)
	mockUserWatchlistItems := mock_userwatchlistitem.NewMockStore(ctrl)
	mockServiceCoingecko := mock_coingecko.NewMockService(ctrl)
	mockCoingeckoSupportToken := mock_coingeckosupportedtokens.NewMockStore(ctrl)
	s := pg.NewPostgresStore(&cfg)
	r := pg.NewRepo(s.DB())
	r.UserWatchlistItem = mockUserWatchlistItems
	r.CoingeckoSupportedTokens = mockCoingeckoSupportToken
	svc.CoinGecko = mockServiceCoingecko

	e := &Entity{
		repo: r,
		log:  log,
		svc:  svc,
		cfg:  cfg,
	}

	tests := []struct {
		name                   string
		query                  string
		coinGeckoTokenFound    *model.CoingeckoSupportedTokens
		coinGeckoTokenError    error
		coinGeckoSuggestTokens []model.CoingeckoSupportedTokens
		want                   []model.CoingeckoSupportedTokens
		coinIds                []string
		coinPrices             map[string]float64
		wantErr                bool
	}{
		{
			name:  "success - get one token",
			query: "cake",
			coinGeckoTokenFound: &model.CoingeckoSupportedTokens{
				ID:     "pancakeswap-token",
				Symbol: "cake",
				Name:   "PancakeSwap",
			},
			coinGeckoTokenError: nil,
			want: []model.CoingeckoSupportedTokens{{
				ID:     "pancakeswap-token",
				Symbol: "cake",
				Name:   "PancakeSwap",
			}},
			coinIds:    []string{"pancakeswap-token"},
			coinPrices: map[string]float64{"pancakeswap-token": 0},
			wantErr:    false,
		},
		{
			name:                "success - get multiple tokens",
			query:               "eth",
			coinGeckoTokenError: gorm.ErrRecordNotFound,
			coinGeckoSuggestTokens: []model.CoingeckoSupportedTokens{
				{
					ID:     "ethereum",
					Symbol: "eth",
					Name:   "Ethereum",
				},
				{
					ID:     "ethereum-wormhole",
					Symbol: "eth",
					Name:   "Ethereum (Wormhole)",
				},
			},
			coinIds:    []string{"ethereum", "ethereum-wormhole"},
			coinPrices: map[string]float64{"ethereum": 1700, "ethereum-wormhole": 120},
			want: []model.CoingeckoSupportedTokens{
				{
					ID:           "ethereum",
					Symbol:       "eth",
					Name:         "Ethereum",
					CurrentPrice: 1700,
				},
				{
					ID:           "ethereum-wormhole",
					Symbol:       "eth",
					Name:         "Ethereum (Wormhole)",
					CurrentPrice: 120,
				}},
			wantErr: false,
		},
		{
			name:                   "failed - coin not found",
			query:                  "cakeabc",
			coinGeckoTokenError:    gorm.ErrRecordNotFound,
			coinGeckoSuggestTokens: []model.CoingeckoSupportedTokens{},
			want:                   []model.CoingeckoSupportedTokens{},
			wantErr:                false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCoingeckoSupportToken.EXPECT().GetOne(tt.query).Return(tt.coinGeckoTokenFound, tt.coinGeckoTokenError).AnyTimes()
			mockCoingeckoSupportToken.EXPECT().List(coingeckosupportedtokens.ListQuery{Symbol: tt.query}).Return(tt.coinGeckoSuggestTokens, nil).AnyTimes()

			if tt.coinIds != nil && len(tt.coinIds) != 0 {
				for _, coinId := range tt.coinIds {
					mockServiceCoingecko.EXPECT().GetCoinPrice([]string{coinId}, "usd").Return(map[string]float64{coinId: tt.coinPrices[coinId]}, nil).AnyTimes()
				}
			}

			got, err := e.SearchCoins(tt.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.SearchCoins() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entity.SearchCoins() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntity_GetCoinData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := config.LoadTestConfig()

	log := logger.NewLogrusLogger()
	svc, _ := service.NewService(cfg, log)
	mockServiceCoingecko := mock_coingecko.NewMockService(ctrl)
	s := pg.NewPostgresStore(&cfg)
	r := pg.NewRepo(s.DB())
	svc.CoinGecko = mockServiceCoingecko

	e := &Entity{
		repo: r,
		log:  log,
		svc:  svc,
		cfg:  cfg,
	}
	tests := []struct {
		name         string
		coinId       string
		want         *response.GetCoinResponse
		code         int
		coingeckoErr error
		wantErr      bool
	}{
		{
			name:    "success - get coin",
			coinId:  "ethereum",
			code:    200,
			wantErr: false,
			want: &response.GetCoinResponse{
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
						"aed": 4977.27,
						"ars": 201382,
						"aud": 2085.41,
					},
					MarketCap: map[string]float64{
						"aed": 601328808723,
						"ars": 24330025249694,
						"aud": 251948328091,
					},
					PriceChangePercentage1hInCurrency: map[string]float64{
						"aed": -0.04997,
						"ars": -0.02977,
						"aud": -0.00954,
					},
					PriceChangePercentage24hInCurrency: map[string]float64{
						"aed": 0.46996,
						"ars": 0.73306,
						"aud": 1.19342,
					},
					PriceChangePercentage7dInCurrency: map[string]float64{
						"aed": 1.79833,
						"ars": 3.21793,
						"aud": 0.74728,
					},
				},
				Tickers: []response.TickerData{
					{
						Base:         "ETH",
						Target:       "USDT",
						Last:         1356.13,
						CoinID:       "ethereum",
						TargetCoinID: "tether",
					},
				},
				Description: map[string]string{},
			},
		},
		{
			name:         "failed - no data",
			coinId:       "etheabc",
			code:         400,
			coingeckoErr: errors.New("record not found"),
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockServiceCoingecko.EXPECT().GetCoin(tt.coinId).Return(tt.want, tt.coingeckoErr, tt.code)

			got, err, code := e.GetCoinData(tt.coinId, false)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.GetCoinData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entity.GetCoinData() got = %v, want %v", got, tt.want)
			}
			if code != tt.code {
				t.Errorf("Entity.GetCoinData() got1 = %v, want %v", code, tt.code)
			}
		})
	}
}
