package tip

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
	"github.com/defipod/mochi/pkg/service"
	mock_coingecko "github.com/defipod/mochi/pkg/service/coingecko/mocks"
	"github.com/defipod/mochi/pkg/util/testhelper"
)

func TestHandler_GetUserBalances(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db := testhelper.LoadTestDB("../../../migrations/test_seed")
	repo := pg.NewRepo(db)
	cfg := config.LoadTestConfig()
	log := logger.NewLogrusLogger()
	svc, _ := service.NewService(cfg, log)
	coingeckoMock := mock_coingecko.NewMockService(ctrl)
	svc.CoinGecko = coingeckoMock

	entity := entities.New(cfg, log, repo, nil, nil, nil, nil, svc, nil, nil, nil)
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
			name:             "success get bals",
			query:            "user_id=463379262620041226",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/offchain_tip_bot/200-balances.json",
		},
		{
			name:             "success get bals zero",
			query:            "user_id=11111",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/offchain_tip_bot/200-zero-balance.json",
		},
		{
			name:             "missing userID",
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/400-missing-userID.json",
		},
	}
	coinPrice := map[string]float64{
		"":                  0.01,
		"pancakeswap-token": 3,
	}
	coingeckoMock.EXPECT().GetCoinPrice(gomock.Any(), "usd").Return(coinPrice, nil).AnyTimes()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/offchain-tip-bot/balances?%s", tt.query), nil)

			h.GetUserBalances(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.GetUserBalances] response mismatched")
		})
	}
}

func TestHandler_TransferToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db := testhelper.LoadTestDB("../../../migrations/test_seed")
	repo := pg.NewRepo(db)
	cfg := config.LoadTestConfig()
	log := logger.NewLogrusLogger()
	svc, _ := service.NewService(cfg, log)
	coingeckoMock := mock_coingecko.NewMockService(ctrl)
	svc.CoinGecko = coingeckoMock

	entity := entities.New(cfg, log, repo, nil, nil, nil, nil, svc, nil, nil, nil)
	h := &Handler{
		entities: entity,
		log:      log,
	}

	tests := []struct {
		name             string
		req              request.OffchainTransferRequest
		wantCode         int
		wantResponsePath string
	}{
		{
			name: "tip 1 user success",
			req: request.OffchainTransferRequest{
				Sender:       "463379262620041226",
				Recipients:   []string{"760874365037314100"},
				GuildID:      "462663954813157376",
				ChannelID:    "1003381172178530494",
				Amount:       1.5,
				Token:        "CAKE",
				FullCommand:  "tip <@760874365037314100> 1.5 cake",
				TransferType: "tip",
			},
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/offchain_tip_bot/200-tip-one-user.json",
		},
		{
			name: "airdrop success",
			req: request.OffchainTransferRequest{
				Sender:       "463379262620041226",
				Recipients:   []string{"760874365037314100", "580788681967665173"},
				GuildID:      "462663954813157376",
				ChannelID:    "1003381172178530494",
				Amount:       2,
				Token:        "CAKE",
				FullCommand:  "drop 2 cake",
				TransferType: "airdrop",
			},
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/offchain_tip_bot/200-airdrop.json",
		},
		{
			name: "tip users success",
			req: request.OffchainTransferRequest{
				Sender:       "463379262620041226",
				Recipients:   []string{"760874365037314100", "580788681967665173"},
				GuildID:      "462663954813157376",
				ChannelID:    "1003381172178530494",
				Amount:       2,
				Token:        "CAKE",
				FullCommand:  "tip <@760874365037314100> <@580788681967665173> 2 cake",
				TransferType: "tip",
			},
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/offchain_tip_bot/200-tip-multiple-users.json",
		},
		{
			name: "tip each",
			req: request.OffchainTransferRequest{
				Sender:       "463379262620041226",
				Recipients:   []string{"760874365037314100", "580788681967665173"},
				GuildID:      "462663954813157376",
				ChannelID:    "1003381172178530494",
				Amount:       4,
				Each:         true,
				Token:        "CAKE",
				FullCommand:  "tip <@760874365037314100> <@580788681967665173> 2 cake each",
				TransferType: "tip",
			},
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/offchain_tip_bot/200-tip-each.json",
		},
		{
			name: "tip all",
			req: request.OffchainTransferRequest{
				Sender:       "463379262620041226",
				Recipients:   []string{"760874365037314100"},
				GuildID:      "462663954813157376",
				ChannelID:    "1003381172178530494",
				Amount:       0,
				All:          true,
				Token:        "CAKE",
				FullCommand:  "tip <@760874365037314100> all cake",
				TransferType: "tip",
			},
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/offchain_tip_bot/200-tip-all.json",
		},
		{
			name: "insufficient bal",
			req: request.OffchainTransferRequest{
				Sender:       "463379262620041226",
				Recipients:   []string{"760874365037314100"},
				GuildID:      "462663954813157376",
				ChannelID:    "1003381172178530494",
				Amount:       50,
				Token:        "CAKE",
				FullCommand:  "tip <@760874365037314100> 3 cake",
				TransferType: "tip",
			},
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/offchain_tip_bot/400-insufficient-bal.json",
		},
		{
			name: "token not support",
			req: request.OffchainTransferRequest{
				Sender:       "463379262620041226",
				Recipients:   []string{"760874365037314100"},
				GuildID:      "462663954813157376",
				ChannelID:    "1003381172178530494",
				Amount:       3,
				Token:        "ALT",
				FullCommand:  "tip <@760874365037314100> 3 ALT",
				TransferType: "tip",
			},
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/offchain_tip_bot/400-token-not-support.json",
		},
		{
			name: "user not have balance",
			req: request.OffchainTransferRequest{
				Sender:       "929673422198411304",
				Recipients:   []string{"381469582176813056"},
				GuildID:      "462663954813157376",
				ChannelID:    "1003381172178530494",
				Amount:       3,
				Token:        "CAKE",
				FullCommand:  "tip <@381469582176813056> 3 cake",
				TransferType: "tip",
			},
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/offchain_tip_bot/400-insufficient-bal.json",
		},
	}

	coinPrice := map[string]float64{
		"pancakeswap-token": 3,
	}
	coingeckoMock.EXPECT().GetCoinPrice(gomock.Any(), "usd").Return(coinPrice, nil).AnyTimes()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, err := json.Marshal(tt.req)
			if err != nil {
				t.Error(err)
				return
			}
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodPost, "/api/v1/offchain-tip-bot/transfer", bytes.NewBuffer(body))

			h.TransferToken(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.TransferToken] response mismatched")
		})
	}
}
