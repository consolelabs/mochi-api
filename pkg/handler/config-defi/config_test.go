package configdefi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

func TestHandler_ListAllCustomToken(t *testing.T) {
	db := testhelper.LoadTestDB("../../../migrations/test_seed")
	repo := pg.NewRepo(db)
	cfg := config.LoadTestConfig()
	log := logger.NewLogrusLogger()
	entity := entities.New(cfg, log, repo, nil, nil, nil, nil, nil, nil, nil, nil, nil)

	h := &Handler{
		entities: entity,
		log:      log,
	}

	tests := []struct {
		name             string
		param            string
		wantCode         int
		wantResponsePath string
	}{
		{
			name:             "200_empty_guild_id",
			param:            "testt",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/200_data_empty_slice.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/guilds/%s/custom-tokens", tt.param), nil)

			h.ListAllCustomToken(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.ListAllCustomToken] response mismatched")
		})
	}
}

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

	entityMock := entities.New(cfg, log, repo, s, nil, nil, nil, svc, nil, nil, nil, nil)
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
		coinIds               []string
		coinPrices            map[string]float64
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
			coinIds:          []string{"pancakeswap-token"},
			coinPrices:       map[string]float64{"pancakeswap-token": 1.7},
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
			coinIds:               []string{"ethereum"},
			coinPrices:            map[string]float64{"pancakeswap-token": 1700},
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

			if tt.coinIds != nil && len(tt.coinIds) != 0 {
				for _, coinId := range tt.coinIds {
					coingeckoMock.EXPECT().GetCoinPrice([]string{coinId}, "usd").Return(map[string]float64{coinId: tt.coinPrices[coinId]}, nil).AnyTimes()
				}
			}

			h.HandlerGuildCustomTokenConfig(ctx)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)
			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.HandlerGuildCustomTokenConfig] response mismatched")
		})
	}
}

func TestHandler_GetMonikerByGuildID(t *testing.T) {
	db := testhelper.LoadTestDB("../../../migrations/test_seed")
	repo := pg.NewRepo(db)
	cfg := config.LoadTestConfig()
	log := logger.NewLogrusLogger()
	entity := entities.New(cfg, log, repo, nil, nil, nil, nil, nil, nil, nil, nil, nil)

	h := &Handler{
		entities: entity,
		log:      log,
	}

	tests := []struct {
		name             string
		query            string
		wantCode         int
		wantResponsePath string
	}{
		{
			name:             "400_empty_guild_id",
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/400-missing-guildID.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/config-defi/monikers?%s", tt.query), nil)

			h.GetMonikerByGuildID(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.GetMonikerByGuildID] response mismatched")
		})
	}
}

func TestHandler_GetDefaultMoniker(t *testing.T) {
	db := testhelper.LoadTestDB("../../../migrations/test_seed")
	repo := pg.NewRepo(db)
	cfg := config.LoadTestConfig()
	log := logger.NewLogrusLogger()
	entity := entities.New(cfg, log, repo, nil, nil, nil, nil, nil, nil, nil, nil, nil)

	h := &Handler{
		entities: entity,
		log:      log,
	}

	tests := []struct {
		name             string
		query            string
		wantCode         int
		wantResponsePath string
	}{
		// {
		// 	name:             "400_empty_guild_id",
		// 	wantCode:         http.StatusBadRequest,
		// 	wantResponsePath: "testdata/400-missing-guildID.json",
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/config-defi/monikers?%s", tt.query), nil)

			h.GetDefaultMoniker(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.GetDefaultMoniker] response mismatched")
		})
	}
}

func TestHandler_UpsertMonikerConfig(t *testing.T) {
	db := testhelper.LoadTestDB("../../../migrations/test_seed")
	repo := pg.NewRepo(db)
	cfg := config.LoadTestConfig()
	log := logger.NewLogrusLogger()
	entity := entities.New(cfg, log, repo, nil, nil, nil, nil, nil, nil, nil, nil, nil)

	h := &Handler{
		entities: entity,
		log:      log,
	}

	tests := []struct {
		name             string
		args             interface{}
		wantCode         int
		wantResponsePath string
	}{
		{
			name: "200_ok",
			args: request.UpsertMonikerConfigRequest{
				GuildID: "895659000996200508",
				Moniker: "Phela",
				Plural:  "",
				Amount:  10,
				Token:   "icy",
			},
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/200-message-ok.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, err := json.Marshal(tt.args)
			if err != nil {
				t.Error(err)
				return
			}

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodPost, "/api/v1/config-defi/monikers", bytes.NewBuffer(body))

			h.UpsertMonikerConfig(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.UpsertMonikerConfig] response mismatched")
		})
	}
}

func TestHandler_DeleteMonikerConfig(t *testing.T) {
	db := testhelper.LoadTestDB("../../../migrations/test_seed")
	repo := pg.NewRepo(db)
	cfg := config.LoadTestConfig()
	log := logger.NewLogrusLogger()
	entity := entities.New(cfg, log, repo, nil, nil, nil, nil, nil, nil, nil, nil, nil)

	h := &Handler{
		entities: entity,
		log:      log,
	}

	tests := []struct {
		name             string
		args             interface{}
		wantCode         int
		wantResponsePath string
	}{
		{
			name: "200_ok",
			args: request.DeleteMonikerConfigRequest{
				GuildID: "895659000996200508",
				Moniker: "Phela",
			},
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/200-message-ok.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, err := json.Marshal(tt.args)
			if err != nil {
				t.Error(err)
				return
			}

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodDelete, "/api/v1/config-defi/monikers", bytes.NewBuffer(body))

			h.DeleteMonikerConfig(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.DeleteMonikerConfig] response mismatched")
		})
	}
}

func TestHandler_GetGuildDefaultCurrency(t *testing.T) {
	db := testhelper.LoadTestDB("../../../migrations/test_seed")
	repo := pg.NewRepo(db)
	cfg := config.LoadTestConfig()
	log := logger.NewLogrusLogger()
	entity := entities.New(cfg, log, repo, nil, nil, nil, nil, nil, nil, nil, nil, nil)

	h := &Handler{
		entities: entity,
		log:      log,
	}

	tests := []struct {
		name             string
		param            string
		wantCode         int
		wantResponsePath string
	}{
		{
			name:             "500_record_not_found",
			param:            "testt",
			wantCode:         http.StatusInternalServerError,
			wantResponsePath: "testdata/404_record_not_found.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/config-defi/default-currency?guild_id=%s", tt.param), nil)

			h.GetGuildDefaultCurrency(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.GetGuildDefaultCurrency] response mismatched")
		})
	}
}

func TestHandler_UpsertGuildDefaultCurrency(t *testing.T) {
	db := testhelper.LoadTestDB("../../../migrations/test_seed")
	repo := pg.NewRepo(db)
	cfg := config.LoadTestConfig()
	log := logger.NewLogrusLogger()
	entity := entities.New(cfg, log, repo, nil, nil, nil, nil, nil, nil, nil, nil, nil)

	h := &Handler{
		entities: entity,
		log:      log,
	}

	tests := []struct {
		name             string
		args             interface{}
		wantCode         int
		wantResponsePath string
	}{
		{
			name: "200_ok",
			args: request.UpsertGuildDefaultCurrencyRequest{
				GuildID: "895659000996200508",
				Symbol:  "cake",
			},
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/200-message-ok.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, err := json.Marshal(tt.args)
			if err != nil {
				t.Error(err)
				return
			}

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodPost, "/api/v1/config-defi/default-currency", bytes.NewBuffer(body))

			h.UpsertGuildDefaultCurrency(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.UpsertGuildDefaultCurrency] response mismatched")
		})
	}
}

func TestHandler_DeleteGuildDefaultCurrency(t *testing.T) {
	db := testhelper.LoadTestDB("../../../migrations/test_seed")
	repo := pg.NewRepo(db)
	cfg := config.LoadTestConfig()
	log := logger.NewLogrusLogger()
	entity := entities.New(cfg, log, repo, nil, nil, nil, nil, nil, nil, nil, nil, nil)

	h := &Handler{
		entities: entity,
		log:      log,
	}

	tests := []struct {
		name             string
		args             interface{}
		wantCode         int
		wantResponsePath string
	}{
		{
			name: "200_ok",
			args: request.GuildIDRequest{
				GuildID: "895659000996200508",
			},
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/200-message-ok.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, err := json.Marshal(tt.args)
			if err != nil {
				t.Error(err)
				return
			}

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodDelete, "/api/v1/config-defi/default-currency", bytes.NewBuffer(body))

			h.DeleteGuildDefaultCurrency(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.DeleteGuildDefaultCurrency] response mismatched")
		})
	}
}
