package entities

import (
	"reflect"
	"testing"

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
	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
)

// func TestEntity_TokenCompare(t *testing.T) {
// 	type fields struct {
// 		repo     *repo.Repo
// 		store    repo.Store
// 		log      logger.Logger
// 		dcwallet discordwallet.IDiscordWallet
// 		discord  *discordgo.Session
// 		cache    cache.Cache
// 		svc      *service.Service
// 		cfg      config.Config
// 	}

// 	type args struct {
// 		sourceSymbolInfo [][]float32
// 		targetSymbolInfo [][]float32
// 	}

// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	cfg := config.Config{
// 		DBUser: "postgres",
// 		DBPass: "postgres",
// 		DBHost: "localhost",
// 		DBPort: "5434",
// 		DBName: "mochi_local",

// 		InDiscordWalletMnemonic: "holiday frequent toy bachelor auto use style result recycle crumble glue blouse",
// 		FantomRPC:               "https://rpc.ftm.tools",
// 		FantomScan:              "https://api.ftmscan.com/api?",
// 		FantomScanAPIKey:        "XEKSVDF5VWQDY5VY6ZNT6AK9QPQRH483EF",

// 		EthereumRPC:        "https://mainnet.infura.io/v3/5b389eb75c514cf6b1711d70084b0114",
// 		EthereumScan:       "https://api.etherscan.io/api?",
// 		EthereumScanAPIKey: "SM5BHYSNIRZ1HEWJ1JPHVTMJS95HRA6DQF",

// 		BscRPC:        "https://bsc-dataseed.binance.org",
// 		BscScan:       "https://api.bscscan.com/api?",
// 		BscScanAPIKey: "VTKF4RG4HP6WXQ5QTAJ8MHDDIUFYD6VZHC",

// 		DiscordToken: "OTcxNjMyNDMzMjk0MzQ4Mjg5.G5BEgF.rv-16ZuTzzqOv2W76OljymFxxnNpjVjCnOkn98",

// 		RedisURL: "redis://localhost:6379/0",
// 	}

// 	log := logger.NewLogrusLogger()
// 	svc, _ := service.NewService(cfg, log)

// 	uCompare := mock_coingecko.NewMockService(ctrl)
// 	svc.CoinGecko = uCompare
// 	tests := []struct {
// 		name                string
// 		fields              fields
// 		args                args
// 		wantTokenCompareRes *response.TokenCompareReponse
// 		wantErr             bool
// 	}{
// 		{
// 			name: "test return successfully",
// 			fields: fields{
// 				svc: svc,
// 			},
// 			args: args{
// 				sourceSymbolInfo: [][]float32{
// 					{
// 						1652803200000,
// 						2096.77,
// 						2110.55,
// 						2062.53,
// 						2062.53,
// 					},
// 					{
// 						1652817600000,
// 						2063.61,
// 						2063.88,
// 						2029.1,
// 						2029.1,
// 					},
// 				},
// 				targetSymbolInfo: [][]float32{
// 					{
// 						1652803200000,
// 						30555.49,
// 						30651.91,
// 						30117.68,
// 						30117.68,
// 					},
// 					{
// 						1652817600000,
// 						30159.76,
// 						30203.91,
// 						29835.27,
// 						29835.27,
// 					},
// 				},
// 			},
// 			wantTokenCompareRes: &response.TokenCompareReponse{
// 				PriceCompare: []float32{
// 					0.06862171,
// 					0.06842263,
// 				},
// 				Times: []string{
// 					"2022-05-17T16:00:39Z",
// 					"2022-05-17T20:00:57Z",
// 				},
// 			},
// 			wantErr: false,
// 		},
// 	}

// 	sourceSymbolInfoParam := [][]float32{
// 		{
// 			1652803200000,
// 			2096.77,
// 			2110.55,
// 			2062.53,
// 			2062.53,
// 		},
// 		{
// 			1652817600000,
// 			2063.61,
// 			2063.88,
// 			2029.1,
// 			2029.1,
// 		},
// 	}

// 	targetSymbolInfoParam := [][]float32{
// 		{
// 			1652803200000,
// 			30555.49,
// 			30651.91,
// 			30117.68,
// 			30117.68,
// 		},
// 		{
// 			1652817600000,
// 			30159.76,
// 			30203.91,
// 			29835.27,
// 			29835.27,
// 		},
// 	}

// 	resList := &response.TokenCompareReponse{
// 		PriceCompare: []float32{
// 			0.06862171,
// 			0.06842263,
// 		},
// 		Times: []string{
// 			"2022-05-17T16:00:39Z",
// 			"2022-05-17T20:00:57Z",
// 		},
// 	}

// 	uCompare.EXPECT().TokenCompare(sourceSymbolInfoParam, targetSymbolInfoParam).Return(resList, nil).AnyTimes()
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			e := &Entity{
// 				repo:     tt.fields.repo,
// 				store:    tt.fields.store,
// 				log:      tt.fields.log,
// 				dcwallet: tt.fields.dcwallet,
// 				discord:  tt.fields.discord,
// 				cache:    tt.fields.cache,
// 				svc:      tt.fields.svc,
// 				cfg:      tt.fields.cfg,
// 			}
// 			gotTokenCompareRes, err := e.CompareToken(tt.args.sourceSymbolInfo, tt.args.targetSymbolInfo)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Entity.TokenCompare() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(gotTokenCompareRes, tt.wantTokenCompareRes) {
// 				t.Errorf("Entity.TokenCompare() = %v, want %v", gotTokenCompareRes, tt.wantTokenCompareRes)
// 			}
// 		})
// 	}
// }

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
		name      string
		req       request.GetUserWatchlistRequest
		userItems []model.UserWatchlistItem
		itemsStr  []string
		want      []response.CoinMarketItemData
		wantErr   bool
	}{
		{
			name: "success - default tokens",
			req: request.GetUserWatchlistRequest{
				UserID: "319132138849173123",
				Size:   7,
			},
			itemsStr: []string{"bitcoin", "ethereum", "binancecoin", "fantom", "internet-computer", "solana", "avalanche-2", "matic-network"},
			want:     coingeckoDefaultTickers,
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
			want: []response.CoinMarketItemData{
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
			want: []response.CoinMarketItemData{
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserWatchlistItems.EXPECT().List(userwatchlistitem.UserWatchlistQuery{
				UserID: tt.req.UserID,
				Offset: tt.req.Page * tt.req.Size,
				Limit:  tt.req.Size,
			}).Return(tt.userItems, int64(len(tt.userItems)), nil).AnyTimes()
			mockServiceCoingecko.EXPECT().GetCoinsMarketData(tt.itemsStr).Return(tt.want, nil, 200).AnyTimes()

			got, err := e.GetUserWatchlist(tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.GetUserWatchlist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, &tt.want) {
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
			want: response.AddToWatchlistResponse{
				Data: &response.AddToWatchlistResponseData{
					BaseSuggestions: []model.CoingeckoSupportedTokens{
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
			coingeckoSupportedTokenError: nil,
			coingeckoSupportedTokenFound: model.CoingeckoSupportedTokens{
				ID:     "pancakeswap-token",
				Symbol: "cake",
				Name:   "PancakeSwap",
			},
			want: response.AddToWatchlistResponse{
				Data: nil,
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
