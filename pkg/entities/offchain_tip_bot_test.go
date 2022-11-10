package entities

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	mock_offchain_tip_activity_logs "github.com/defipod/mochi/pkg/repo/offchain_tip_activity_logs/mocks"
	mock_offchain_tip_bot_tokens "github.com/defipod/mochi/pkg/repo/offchain_tip_bot_tokens/mocks"
	mock_offchain_tip_transfer_histories "github.com/defipod/mochi/pkg/repo/offchain_tip_transfer_histories/mocks"
	mock_offchain_tip_bot_user_balances "github.com/defipod/mochi/pkg/repo/offchain_tip_user_balances/mocks"
	"github.com/defipod/mochi/pkg/repo/pg"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/service"
	mock_coingecko "github.com/defipod/mochi/pkg/service/coingecko/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func TestEntity_TransferToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := config.LoadTestConfig()

	log := logger.NewLogrusLogger()
	svc, _ := service.NewService(cfg, log)
	mockServiceCoingecko := mock_coingecko.NewMockService(ctrl)
	mockActivityLogs := mock_offchain_tip_activity_logs.NewMockStore(ctrl)
	mockTokens := mock_offchain_tip_bot_tokens.NewMockStore(ctrl)
	mockUserBalances := mock_offchain_tip_bot_user_balances.NewMockStore(ctrl)
	mockTransferHistories := mock_offchain_tip_transfer_histories.NewMockStore(ctrl)

	s := pg.NewPostgresStore(&cfg)
	r := pg.NewRepo(s.DB())
	svc.CoinGecko = mockServiceCoingecko
	r.OffchainTipBotActivityLogs = mockActivityLogs
	r.OffchainTipBotTokens = mockTokens
	r.OffchainTipBotUserBalances = mockUserBalances
	r.OffchainTipBotTransferHistories = mockTransferHistories
	e := &Entity{
		repo: r,
		log:  log,
		svc:  svc,
		cfg:  cfg,
	}

	type tokenResp struct {
		token *model.OffchainTipBotToken
		err   error
	}

	tests := []struct {
		name              string
		req               request.OffchainTransferRequest
		userBal           float64
		coinPrice         map[string]float64
		token             tokenResp
		transferHistories []model.OffchainTipBotTransferHistory
		want              []response.OffchainTipBotTransferToken
		wantErr           bool
	}{
		{
			name: "transfer token success",
			req: request.OffchainTransferRequest{
				Sender:      "463379262620041226",
				Recipients:  []string{"760874365037314100"},
				GuildID:     "462663954813157376",
				ChannelID:   "1003381172178530494",
				Amount:      1.5,
				Token:       "FTM",
				FullCommand: "tip <@760874365037314100> 1.5 ftm",
			},
			userBal: 10,
			coinPrice: map[string]float64{
				"fantom": 4,
			},
			token: tokenResp{
				token: &model.OffchainTipBotToken{ID: uuid.UUID{}, TokenID: "1", TokenName: "Fantom", TokenSymbol: "FTM", CoinGeckoID: "fantom"},
				err:   nil,
			},
			transferHistories: []model.OffchainTipBotTransferHistory{
				{
					SenderID:   "463379262620041226",
					ReceiverID: "760874365037314100",
				},
			},
			want: []response.OffchainTipBotTransferToken{
				{
					SenderID:    "463379262620041226",
					RecipientID: "760874365037314100",
					Symbol:      "FTM",
					Amount:      1.5,
					AmountInUSD: 6,
				},
			},
		},
		{
			name: "transfer token multiple users success",
			req: request.OffchainTransferRequest{
				Sender:      "463379262620041226",
				Recipients:  []string{"760874365037314100", "580788681967665173"},
				GuildID:     "462663954813157376",
				ChannelID:   "1003381172178530494",
				Amount:      10,
				Token:       "CAKE",
				FullCommand: "tip <@760874365037314100> <@580788681967665173> 10 cake",
			},
			userBal: 10,
			coinPrice: map[string]float64{
				"pancake-swap": 4.25,
			},
			token: tokenResp{
				token: &model.OffchainTipBotToken{ID: uuid.UUID{}, TokenID: "1", TokenName: "Pancake Swap", TokenSymbol: "CAKE", CoinGeckoID: "pancake-swap"},
				err:   nil,
			},
			transferHistories: []model.OffchainTipBotTransferHistory{
				{
					SenderID:   "463379262620041226",
					ReceiverID: "760874365037314100",
				},
				{
					SenderID:   "463379262620041226",
					ReceiverID: "580788681967665173",
				},
			},
			want: []response.OffchainTipBotTransferToken{
				{
					SenderID:    "463379262620041226",
					RecipientID: "760874365037314100",
					Symbol:      "CAKE",
					Amount:      5,
					AmountInUSD: 21.25,
				},
				{
					SenderID:    "463379262620041226",
					RecipientID: "580788681967665173",
					Symbol:      "CAKE",
					Amount:      5,
					AmountInUSD: 21.25,
				},
			},
		},
		{
			name: "tip 1 user all token",
			req: request.OffchainTransferRequest{
				Sender:      "463379262620041226",
				Recipients:  []string{"760874365037314100"},
				GuildID:     "462663954813157376",
				ChannelID:   "1003381172178530494",
				Amount:      0,
				All:         true,
				Token:       "CAKE",
				FullCommand: "tip <@760874365037314100> all cake",
			},
			userBal: 10,
			coinPrice: map[string]float64{
				"pancake-swap": 4.25,
			},
			token: tokenResp{
				token: &model.OffchainTipBotToken{ID: uuid.UUID{}, TokenID: "1", TokenName: "Pancake Swap", TokenSymbol: "CAKE", CoinGeckoID: "pancake-swap"},
				err:   nil,
			},
			transferHistories: []model.OffchainTipBotTransferHistory{
				{
					SenderID:   "463379262620041226",
					ReceiverID: "760874365037314100",
				},
			},
			want: []response.OffchainTipBotTransferToken{
				{
					SenderID:    "463379262620041226",
					RecipientID: "760874365037314100",
					Symbol:      "CAKE",
					Amount:      10,
					AmountInUSD: 42.5,
				},
			},
		},
		{
			name: "tip multiple users all token",
			req: request.OffchainTransferRequest{
				Sender:      "463379262620041226",
				Recipients:  []string{"760874365037314100", "580788681967665173"},
				GuildID:     "462663954813157376",
				ChannelID:   "1003381172178530494",
				Amount:      0,
				All:         true,
				Token:       "CAKE",
				FullCommand: "tip <@760874365037314100> <@580788681967665173> all cake",
			},
			userBal: 10,
			coinPrice: map[string]float64{
				"pancake-swap": 4.25,
			},
			token: tokenResp{
				token: &model.OffchainTipBotToken{ID: uuid.UUID{}, TokenID: "1", TokenName: "Pancake Swap", TokenSymbol: "CAKE", CoinGeckoID: "pancake-swap"},
				err:   nil,
			},
			transferHistories: []model.OffchainTipBotTransferHistory{
				{
					SenderID:   "463379262620041226",
					ReceiverID: "760874365037314100",
				},
				{
					SenderID:   "463379262620041226",
					ReceiverID: "580788681967665173",
				},
			},
			want: []response.OffchainTipBotTransferToken{
				{
					SenderID:    "463379262620041226",
					RecipientID: "760874365037314100",
					Symbol:      "CAKE",
					Amount:      5,
					AmountInUSD: 21.25,
				},
				{
					SenderID:    "463379262620041226",
					RecipientID: "580788681967665173",
					Symbol:      "CAKE",
					Amount:      5,
					AmountInUSD: 21.25,
				},
			},
		},
		{
			name: "tip each",
			req: request.OffchainTransferRequest{
				Sender:      "463379262620041226",
				Recipients:  []string{"760874365037314100", "580788681967665173"},
				GuildID:     "462663954813157376",
				ChannelID:   "1003381172178530494",
				Amount:      2,
				Token:       "CAKE",
				FullCommand: "tip <@760874365037314100> <@580788681967665173> 1 cake each",
			},
			userBal: 10,
			coinPrice: map[string]float64{
				"pancake-swap": 4.25,
			},
			token: tokenResp{
				token: &model.OffchainTipBotToken{ID: uuid.UUID{}, TokenID: "1", TokenName: "Pancake Swap", TokenSymbol: "CAKE", CoinGeckoID: "pancake-swap"},
				err:   nil,
			},
			transferHistories: []model.OffchainTipBotTransferHistory{
				{
					SenderID:   "463379262620041226",
					ReceiverID: "760874365037314100",
				},
				{
					SenderID:   "463379262620041226",
					ReceiverID: "580788681967665173",
				},
			},
			want: []response.OffchainTipBotTransferToken{
				{
					SenderID:    "463379262620041226",
					RecipientID: "760874365037314100",
					Symbol:      "CAKE",
					Amount:      1,
					AmountInUSD: 4.25,
				},
				{
					SenderID:    "463379262620041226",
					RecipientID: "580788681967665173",
					Symbol:      "CAKE",
					Amount:      1,
					AmountInUSD: 4.25,
				},
			},
		},
		{
			name: "token not supported",
			req: request.OffchainTransferRequest{
				Sender:      "463379262620041226",
				Recipients:  []string{"760874365037314100", "580788681967665173"},
				GuildID:     "462663954813157376",
				ChannelID:   "1003381172178530494",
				Amount:      2,
				All:         true,
				Token:       "ALT",
				FullCommand: "tip <@760874365037314100> <@580788681967665173> 1 alt each",
			},
			userBal:   0,
			coinPrice: nil,
			token: tokenResp{
				token: nil,
				err:   gorm.ErrRecordNotFound,
			},
			want:    []response.OffchainTipBotTransferToken{},
			wantErr: true,
		},
		{
			name: "insufficient balance",
			req: request.OffchainTransferRequest{
				Sender:      "463379262620041226",
				Recipients:  []string{"760874365037314100", "580788681967665173"},
				GuildID:     "462663954813157376",
				ChannelID:   "1003381172178530494",
				Amount:      4,
				Token:       "ALT",
				FullCommand: "tip <@760874365037314100> <@580788681967665173> 2 alt each",
			},
			userBal: 1,
			coinPrice: map[string]float64{
				"pancake-swap": 4.25,
			},
			token: tokenResp{
				token: &model.OffchainTipBotToken{ID: uuid.UUID{}, TokenID: "1", TokenName: "Pancake Swap", TokenSymbol: "CAKE", CoinGeckoID: "pancake-swap"},
				err:   nil,
			},
			want:    []response.OffchainTipBotTransferToken{},
			wantErr: true,
		},
	}
	mockActivityLogs.EXPECT().CreateActivityLog(gomock.Any()).Return(&model.OffchainTipBotActivityLog{ID: uuid.UUID{}}, nil).AnyTimes()
	mockUserBalances.EXPECT().CreateIfNotExists(gomock.Any()).Return(nil).AnyTimes()
	mockUserBalances.EXPECT().UpdateListUserBalances(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
	mockUserBalances.EXPECT().UpdateUserBalance(gomock.Any()).Return(nil).AnyTimes()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockTokens.EXPECT().GetBySymbol(tt.req.Token).Return(tt.token.token, tt.token.err)
			if tt.coinPrice != nil {
				mockServiceCoingecko.EXPECT().GetCoinPrice(gomock.Any(), gomock.Any()).Return(tt.coinPrice, nil)
				mockUserBalances.EXPECT().GetUserBalanceByTokenID(tt.req.Sender, gomock.Any()).Return(&model.OffchainTipBotUserBalance{Amount: tt.userBal}, nil)
			}
			if tt.transferHistories != nil {
				mockTransferHistories.EXPECT().CreateTransferHistories(gomock.Any()).Return(tt.transferHistories, nil)
			}

			got, err := e.TransferToken(tt.req)
			fmt.Println(err)
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
