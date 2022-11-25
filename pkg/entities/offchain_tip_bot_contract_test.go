package entities

import (
	"errors"
	"reflect"
	"testing"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	mock_offchain_tip_bot_user_balances "github.com/defipod/mochi/pkg/repo/offchain_tip_bot_user_balances/mocks"
	"github.com/defipod/mochi/pkg/repo/pg"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/service"
	mock_coingecko "github.com/defipod/mochi/pkg/service/coingecko/mocks"
	"github.com/golang/mock/gomock"
)

func TestEntity_GetUserBalances(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := config.LoadTestConfig()

	log := logger.NewLogrusLogger()
	svc, _ := service.NewService(cfg, log)
	mockServiceCoingecko := mock_coingecko.NewMockService(ctrl)
	mockUserBalances := mock_offchain_tip_bot_user_balances.NewMockStore(ctrl)

	s := pg.NewPostgresStore(&cfg)
	r := pg.NewRepo(s.DB())
	svc.CoinGecko = mockServiceCoingecko
	r.OffchainTipBotUserBalances = mockUserBalances
	e := &Entity{
		repo: r,
		log:  log,
		svc:  svc,
		cfg:  cfg,
	}

	type req struct {
		userID string
	}
	type balances struct {
		bals []model.OffchainTipBotUserBalance
		err  error
	}
	type coinGeckoResp struct {
		coinPrice map[string]float64
		err       error
	}

	tests := []struct {
		name      string
		req       req
		balances  balances
		coinPrice coinGeckoResp
		want      []response.GetUserBalances
		wantErr   bool
	}{
		{
			name: "get balances success",
			req: req{
				userID: "963641551416881183",
			},
			balances: balances{
				bals: []model.OffchainTipBotUserBalance{
					{
						Token:  &model.OffchainTipBotToken{TokenName: "Icy", TokenSymbol: "ICY", CoinGeckoID: ""},
						Amount: 10,
					},
					{
						Token:  &model.OffchainTipBotToken{TokenName: "Fantom", TokenSymbol: "FTM", CoinGeckoID: "fantom"},
						Amount: 11.456,
					},
				},
				err: nil,
			},
			coinPrice: coinGeckoResp{
				coinPrice: map[string]float64{
					"":       0.0001,
					"fantom": 1,
				},
				err: nil,
			},
			want: []response.GetUserBalances{
				{
					ID:            "",
					Name:          "Icy",
					Symbol:        "ICY",
					Balances:      10,
					RateInUSD:     0.0001,
					BalancesInUSD: 0.001,
				},
				{
					ID:            "fantom",
					Name:          "Fantom",
					Symbol:        "FTM",
					Balances:      11.456,
					RateInUSD:     1,
					BalancesInUSD: 11.456,
				},
			},
		},
		{
			name: "fail to get user balance",
			req: req{
				userID: "111111",
			},
			balances: balances{
				bals: nil,
				err:  errors.New("error bals"),
			},
			coinPrice: coinGeckoResp{
				coinPrice: map[string]float64{},
				err:       nil,
			},
			want:    []response.GetUserBalances{},
			wantErr: true,
		},
		{
			name: "fail to get coingecko",
			req: req{
				userID: "963641551416881183",
			},
			balances: balances{
				bals: []model.OffchainTipBotUserBalance{
					{
						Token:  &model.OffchainTipBotToken{TokenName: "Icy", TokenSymbol: "ICY", CoinGeckoID: ""},
						Amount: 10,
					},
					{
						Token:  &model.OffchainTipBotToken{TokenName: "Fantom", TokenSymbol: "FTM", CoinGeckoID: "fantom"},
						Amount: 11.456,
					},
				},
				err: nil,
			},
			coinPrice: coinGeckoResp{
				coinPrice: nil,
				err:       errors.New("error"),
			},
			want:    []response.GetUserBalances{},
			wantErr: true,
		},
	}
	mockServiceCoingecko.EXPECT().GetCoinPrice(gomock.Any(), gomock.Any()).Return(tests[0].coinPrice.coinPrice, tests[0].coinPrice.err).Times(1)
	mockServiceCoingecko.EXPECT().GetCoinPrice(gomock.Any(), gomock.Any()).Return(tests[2].coinPrice.coinPrice, tests[2].coinPrice.err).Times(1)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserBalances.EXPECT().GetUserBalances(tt.req.userID).Return(tt.balances.bals, tt.balances.err)

			got, err := e.GetUserBalances(tt.req.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.TransferToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entity.TransferToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
