package handler

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/repo/pg"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/service"
	mock_coingecko "github.com/defipod/mochi/pkg/service/coingecko/mocks"
	"github.com/defipod/mochi/pkg/util/testhelper"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestHandler_GetCoin(t *testing.T) {
	cfg := config.LoadTestConfig()
	db := testhelper.LoadTestDB("../../migrations/test_seed")
	repo := pg.NewRepo(db)
	log := logger.NewLogrusLogger()
	s := pg.NewPostgresStore(&cfg)
	svc, _ := service.NewService(cfg, log)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	coingeckoMock := mock_coingecko.NewMockService(ctrl)
	svc.CoinGecko = coingeckoMock
	entityMock := entities.New(cfg, log, repo, s, nil, nil, nil, svc, nil, nil, nil)
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
					Thumb:  "https://assets.coingecko.com/coins/images/279/thumb/ethereum.png?1595348880",
					Small:  "https://assets.coingecko.com/coins/images/279/small/ethereum.png?1595348880",
					Larget: "https://assets.coingecko.com/coins/images/279/large/ethereum.png?1595348880",
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
				Description: response.CoinDescription{
					EngDescription: "eth",
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
