package configdefi

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/repo/pg"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/service"
	mock_coingecko "github.com/defipod/mochi/pkg/service/coingecko/mocks"
	mock_covalent "github.com/defipod/mochi/pkg/service/covalent/mocks"
	"github.com/defipod/mochi/pkg/util"
	"github.com/defipod/mochi/pkg/util/testhelper"
)

func TestHandler_HandlerGuildCustomTokenConfig(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := config.LoadTestConfig()
	db := testhelper.LoadTestDB("../../../migrations/test_seed")
	repo := pg.NewRepo(db)
	log := logger.NewLogrusLogger()
	svc, _ := service.NewService(cfg, log)

	covalentMock := mock_covalent.NewMockService(ctrl)
	coingeckoMock := mock_coingecko.NewMockService(ctrl)
	svc.Covalent = covalentMock
	svc.CoinGecko = coingeckoMock
	s := pg.NewPostgresStore(&cfg)

	entityMock := entities.New(cfg, log, repo, s, nil, nil, nil, svc, nil, nil, nil)
	h := &Handler{
		entities: entityMock,
		log:      log,
	}
	tests := []struct {
		name                  string
		req                   request.UpsertCustomTokenConfigRequest
		coingeckoResponse     *response.GetCoinResponse
		coingeckoSupportToken *response.CoingeckoSupportedTokenResponse
		historicalTokenPrices *response.HistoricalTokenPricesResponse
		wantResponsePath      string
	}{
		{
			name: "success - add custom token",
			req: request.UpsertCustomTokenConfigRequest{
				Address: "0x3EE2200Efb3400fAbB9AacF31297cBdD1d435D47",
				Symbol:  "cake",
				Chain:   "ftm",
				GuildID: "863278424433229854",
			},
			coingeckoSupportToken: &response.CoingeckoSupportedTokenResponse{
				ID:     "pancakeswap-token",
				Symbol: "cake",
				Name:   "PancakeSwap",
			},
			coingeckoResponse: &response.GetCoinResponse{
				AssetPlatformID: "fantom",
			},
			historicalTokenPrices: &response.HistoricalTokenPricesResponse{
				Data: []response.HistoricalTokenPrice{
					{
						Decimals: 18,
					},
				},
			},
			wantResponsePath: "testdata/200-message-ok.json",
		},
		{
			name: "failed - chain not found",
			req: request.UpsertCustomTokenConfigRequest{
				Address: "0x3EE2200Efb3400fAbB9AacF31297cBdD1d435D47",
				Symbol:  "cake",
				Chain:   "ftmabc",
				GuildID: "863278424433229854",
			},
			coingeckoSupportToken: &response.CoingeckoSupportedTokenResponse{},
			coingeckoResponse:     &response.GetCoinResponse{},
			historicalTokenPrices: &response.HistoricalTokenPricesResponse{},
			wantResponsePath:      "testdata/custom-token/500-chain-not-found.json",
		},
		{
			name: "failed - coin not found",
			req: request.UpsertCustomTokenConfigRequest{
				Address: "0x3EE2200Efb3400fAbB9AacF31297cBdD1d435D47",
				Symbol:  "cakeee",
				Chain:   "ftm",
				GuildID: "863278424433229854",
			},
			coingeckoSupportToken: &response.CoingeckoSupportedTokenResponse{},
			coingeckoResponse:     &response.GetCoinResponse{},
			historicalTokenPrices: &response.HistoricalTokenPricesResponse{},
			wantResponsePath:      "testdata/custom-token/500-coin-not-found.json",
		},
		{
			name: "failed - coin not supported",
			req: request.UpsertCustomTokenConfigRequest{
				Address: "0x3EE2200Efb3400fAbB9AacF31297cBdD1d435D47",
				Symbol:  "eth",
				Chain:   "eth",
				GuildID: "863278424433229854",
			},
			coingeckoSupportToken: &response.CoingeckoSupportedTokenResponse{
				ID:     "ethereum",
				Symbol: "ETH",
				Name:   "Ethereum",
			},
			coingeckoResponse: &response.GetCoinResponse{
				AssetPlatformID: "fantom",
			},
			historicalTokenPrices: &response.HistoricalTokenPricesResponse{},
			wantResponsePath:      "testdata/custom-token/500-coin-not-supported.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest("POST", "/api/v1/configs/token", nil)
			util.SetRequestBody(ctx, tt.req)

			coingeckoMock.EXPECT().GetCoin(tt.coingeckoSupportToken.ID).Return(tt.coingeckoResponse, nil, 0).AnyTimes()
			covalentMock.EXPECT().GetHistoricalTokenPrices(250, "FTM", tt.req.Address).Return(tt.historicalTokenPrices, nil, 200).AnyTimes()

			h.HandlerGuildCustomTokenConfig(ctx)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)
			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.HandlerGuildCustomTokenConfig] response mismatched")
		})
	}
}
