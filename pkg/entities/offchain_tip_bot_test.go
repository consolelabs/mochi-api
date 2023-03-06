package entities

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/consts"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	mock_offchain_tip_activity_logs "github.com/defipod/mochi/pkg/repo/offchain_tip_bot_activity_logs/mocks"
	mock_offchain_tip_bot_tokens "github.com/defipod/mochi/pkg/repo/offchain_tip_bot_tokens/mocks"
	mock_offchain_tip_transfer_histories "github.com/defipod/mochi/pkg/repo/offchain_tip_bot_transfer_histories/mocks"
	mock_offchain_tip_bot_user_balances "github.com/defipod/mochi/pkg/repo/offchain_tip_bot_user_balances/mocks"
	"github.com/defipod/mochi/pkg/repo/pg"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/service"
	mock_coingecko "github.com/defipod/mochi/pkg/service/coingecko/mocks"
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
	var tip = "tip"
	var duration = 0
	fullCmd1 := "tip <@760874365037314100> 1.5 ftm"
	fullCmd2 := "tip <@760874365037314100> <@580788681967665173> 10 cake"
	fullCmd3 := "tip <@760874365037314100> all cake"
	fullCmd4 := "tip <@760874365037314100> <@580788681967665173> all cake"
	fullCmd5 := "tip <@760874365037314100> <@580788681967665173> 1 cake each"
	fullCmd6 := "tip <@760874365037314100> <@580788681967665173> 1 alt each"
	fullCmd7 := "tip <@760874365037314100> <@580788681967665173> 2 cake each"
	senderID := "463379262620041226"
	guildID := "462663954813157376"
	channelID := "1003381172178530494"

	tests := []struct {
		name              string
		req               request.OffchainTransferRequest
		userBal           float64
		leftBal           float64
		coinPrice         map[string]float64
		token             tokenResp
		transferHistories []model.OffchainTipBotTransferHistory
		activitiesLogs    model.OffchainTipBotActivityLog
		want              []response.OffchainTipBotTransferToken
		wantErr           bool
	}{
		{
			name: "transfer token success",
			req: request.OffchainTransferRequest{
				Sender:       senderID,
				GuildID:      guildID,
				ChannelID:    channelID,
				TransferType: tip,
				Recipients:   []string{"760874365037314100"},
				Amount:       1.5,
				Token:        "FTM",
				FullCommand:  "tip <@760874365037314100> 1.5 ftm",
			},
			userBal: 10,
			leftBal: 8.5,
			coinPrice: map[string]float64{
				"fantom": 4,
			},
			token: tokenResp{
				token: &model.OffchainTipBotToken{ID: uuid.UUID{}, TokenID: "1", TokenName: "Fantom", TokenSymbol: "FTM", CoinGeckoID: "fantom"},
				err:   nil,
			},
			transferHistories: []model.OffchainTipBotTransferHistory{
				{
					SenderID:   &senderID,
					ReceiverID: "760874365037314100",
					GuildID:    guildID,
					LogID:      uuid.UUID{}.String(),
					Status:     consts.OffchainTipBotTrasferStatusSuccess,
					Amount:     1.5,
					Token:      "FTM",
					Action:     tip,
				},
			},
			activitiesLogs: model.OffchainTipBotActivityLog{
				UserID:          &senderID,
				GuildID:         &guildID,
				ChannelID:       &channelID,
				Action:          &tip,
				TokenID:         uuid.UUID{}.String(),
				Receiver:        []string{"760874365037314100"},
				NumberReceivers: 1,
				Duration:        &duration,
				Amount:          1.5,
				Status:          consts.OffchainTipBotTrasferStatusSuccess,
				FullCommand:     &fullCmd1,
			},
			want: []response.OffchainTipBotTransferToken{
				{
					SenderID:    senderID,
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
				Sender:       senderID,
				GuildID:      guildID,
				ChannelID:    channelID,
				TransferType: tip,
				Recipients:   []string{"760874365037314100", "580788681967665173"},
				Amount:       10,
				Token:        "CAKE",
				FullCommand:  "tip <@760874365037314100> <@580788681967665173> 10 cake",
			},
			userBal: 10,
			leftBal: 0,
			coinPrice: map[string]float64{
				"pancake-swap": 4.25,
			},
			token: tokenResp{
				token: &model.OffchainTipBotToken{ID: uuid.UUID{}, TokenID: "1", TokenName: "Pancake Swap", TokenSymbol: "CAKE", CoinGeckoID: "pancake-swap"},
				err:   nil,
			},
			transferHistories: []model.OffchainTipBotTransferHistory{
				{
					SenderID:   &senderID,
					ReceiverID: "760874365037314100",
					GuildID:    "462663954813157376",
					LogID:      uuid.UUID{}.String(),
					Status:     consts.OffchainTipBotTrasferStatusSuccess,
					Amount:     5,
					Token:      "CAKE",
					Action:     "tip",
				},
				{
					SenderID:   &senderID,
					ReceiverID: "580788681967665173",
					GuildID:    "462663954813157376",
					LogID:      uuid.UUID{}.String(),
					Status:     consts.OffchainTipBotTrasferStatusSuccess,
					Amount:     5,
					Token:      "CAKE",
					Action:     "tip",
				},
			},
			activitiesLogs: model.OffchainTipBotActivityLog{
				UserID:          &senderID,
				GuildID:         &guildID,
				ChannelID:       &channelID,
				Action:          &tip,
				TokenID:         uuid.UUID{}.String(),
				Receiver:        []string{"760874365037314100", "580788681967665173"},
				NumberReceivers: 2,
				Duration:        &duration,
				Amount:          5,
				Status:          consts.OffchainTipBotTrasferStatusSuccess,
				FullCommand:     &fullCmd2,
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
				Sender:       senderID,
				GuildID:      guildID,
				ChannelID:    channelID,
				TransferType: tip,
				Recipients:   []string{"760874365037314100"},
				Amount:       0,
				All:          true,
				Token:        "CAKE",
				FullCommand:  "tip <@760874365037314100> all cake",
			},
			userBal: 10,
			leftBal: 0,
			coinPrice: map[string]float64{
				"pancake-swap": 4.25,
			},
			token: tokenResp{
				token: &model.OffchainTipBotToken{ID: uuid.UUID{}, TokenID: "1", TokenName: "Pancake Swap", TokenSymbol: "CAKE", CoinGeckoID: "pancake-swap"},
				err:   nil,
			},
			transferHistories: []model.OffchainTipBotTransferHistory{
				{
					SenderID:   &senderID,
					ReceiverID: "760874365037314100",
					GuildID:    "462663954813157376",
					LogID:      uuid.UUID{}.String(),
					Status:     consts.OffchainTipBotTrasferStatusSuccess,
					Amount:     10,
					Token:      "CAKE",
					Action:     "tip",
				},
			},
			activitiesLogs: model.OffchainTipBotActivityLog{
				UserID:          &senderID,
				GuildID:         &guildID,
				ChannelID:       &channelID,
				Action:          &tip,
				TokenID:         uuid.UUID{}.String(),
				Receiver:        []string{"760874365037314100"},
				NumberReceivers: 1,
				Duration:        &duration,
				Amount:          10,
				Status:          consts.OffchainTipBotTrasferStatusSuccess,
				FullCommand:     &fullCmd3,
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
				Sender:       senderID,
				GuildID:      guildID,
				ChannelID:    channelID,
				TransferType: tip,
				Recipients:   []string{"760874365037314100", "580788681967665173"},
				Amount:       0,
				All:          true,
				Token:        "CAKE",
				FullCommand:  "tip <@760874365037314100> <@580788681967665173> all cake",
			},
			userBal: 10,
			leftBal: 0,
			coinPrice: map[string]float64{
				"pancake-swap": 4.25,
			},
			token: tokenResp{
				token: &model.OffchainTipBotToken{ID: uuid.UUID{}, TokenID: "1", TokenName: "Pancake Swap", TokenSymbol: "CAKE", CoinGeckoID: "pancake-swap"},
				err:   nil,
			},
			transferHistories: []model.OffchainTipBotTransferHistory{
				{
					SenderID:   &senderID,
					ReceiverID: "760874365037314100",
					GuildID:    "462663954813157376",
					LogID:      uuid.UUID{}.String(),
					Status:     consts.OffchainTipBotTrasferStatusSuccess,
					Amount:     5,
					Token:      "CAKE",
					Action:     "tip",
				},
				{
					SenderID:   &senderID,
					ReceiverID: "580788681967665173",
					GuildID:    "462663954813157376",
					LogID:      uuid.UUID{}.String(),
					Status:     consts.OffchainTipBotTrasferStatusSuccess,
					Amount:     5,
					Token:      "CAKE",
					Action:     "tip",
				},
			},
			activitiesLogs: model.OffchainTipBotActivityLog{
				UserID:          &senderID,
				GuildID:         &guildID,
				ChannelID:       &channelID,
				Action:          &tip,
				TokenID:         uuid.UUID{}.String(),
				Receiver:        []string{"760874365037314100", "580788681967665173"},
				NumberReceivers: 2,
				Duration:        &duration,
				Amount:          5,
				Status:          consts.OffchainTipBotTrasferStatusSuccess,
				FullCommand:     &fullCmd4,
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
				Sender:       senderID,
				GuildID:      guildID,
				ChannelID:    channelID,
				TransferType: tip,
				Recipients:   []string{"760874365037314100", "580788681967665173"},
				Amount:       2,
				Token:        "CAKE",
				FullCommand:  "tip <@760874365037314100> <@580788681967665173> 1 cake each",
			},
			userBal: 10,
			leftBal: 8,
			coinPrice: map[string]float64{
				"pancake-swap": 4.25,
			},
			token: tokenResp{
				token: &model.OffchainTipBotToken{ID: uuid.UUID{}, TokenID: "1", TokenName: "Pancake Swap", TokenSymbol: "CAKE", CoinGeckoID: "pancake-swap"},
				err:   nil,
			},
			transferHistories: []model.OffchainTipBotTransferHistory{
				{
					SenderID:   &senderID,
					ReceiverID: "760874365037314100",
					GuildID:    "462663954813157376",
					LogID:      uuid.UUID{}.String(),
					Status:     consts.OffchainTipBotTrasferStatusSuccess,
					Amount:     1,
					Token:      "CAKE",
					Action:     "tip",
				},
				{
					SenderID:   &senderID,
					ReceiverID: "580788681967665173",
					GuildID:    "462663954813157376",
					LogID:      uuid.UUID{}.String(),
					Status:     consts.OffchainTipBotTrasferStatusSuccess,
					Amount:     1,
					Token:      "CAKE",
					Action:     "tip",
				},
			},
			activitiesLogs: model.OffchainTipBotActivityLog{
				UserID:          &senderID,
				GuildID:         &guildID,
				ChannelID:       &channelID,
				Action:          &tip,
				TokenID:         uuid.UUID{}.String(),
				Receiver:        []string{"760874365037314100", "580788681967665173"},
				NumberReceivers: 2,
				Duration:        &duration,
				Amount:          1,
				Status:          consts.OffchainTipBotTrasferStatusSuccess,
				FullCommand:     &fullCmd5,
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
				Sender:       senderID,
				GuildID:      guildID,
				ChannelID:    channelID,
				TransferType: tip,
				Recipients:   []string{"760874365037314100", "580788681967665173"},
				Amount:       2,
				Token:        "ALT",
				FullCommand:  "tip <@760874365037314100> <@580788681967665173> 1 alt each",
			},
			userBal:   0,
			coinPrice: nil,
			token: tokenResp{
				token: nil,
				err:   gorm.ErrRecordNotFound,
			},
			activitiesLogs: model.OffchainTipBotActivityLog{
				UserID:          &senderID,
				GuildID:         &guildID,
				ChannelID:       &channelID,
				Action:          &tip,
				Receiver:        []string{"760874365037314100", "580788681967665173"},
				NumberReceivers: 2,
				Duration:        &duration,
				Amount:          1,
				Status:          consts.OffchainTipBotTrasferStatusFail,
				FailReason:      consts.OffchainTipBotFailReasonTokenNotSupported,
				FullCommand:     &fullCmd6,
			},
			want:    []response.OffchainTipBotTransferToken{},
			wantErr: true,
		},
		{
			name: "insufficient balance",
			req: request.OffchainTransferRequest{
				Sender:       senderID,
				GuildID:      guildID,
				ChannelID:    channelID,
				TransferType: tip,
				Recipients:   []string{"760874365037314100", "580788681967665173"},
				Amount:       4,
				Each:         true,
				Token:        "CAKE",
				FullCommand:  "tip <@760874365037314100> <@580788681967665173> 2 cake each",
			},
			userBal: 1,
			coinPrice: map[string]float64{
				"pancake-swap": 4.25,
			},
			token: tokenResp{
				token: &model.OffchainTipBotToken{ID: uuid.UUID{}, TokenID: "1", TokenName: "Pancake Swap", TokenSymbol: "CAKE", CoinGeckoID: "pancake-swap"},
				err:   nil,
			},
			activitiesLogs: model.OffchainTipBotActivityLog{
				UserID:          &senderID,
				GuildID:         &guildID,
				ChannelID:       &channelID,
				Action:          &tip,
				TokenID:         uuid.UUID{}.String(),
				Receiver:        []string{"760874365037314100", "580788681967665173"},
				NumberReceivers: 2,
				Duration:        &duration,
				Amount:          2,
				Status:          consts.OffchainTipBotTrasferStatusFail,
				FailReason:      consts.OffchainTipBotFailReasonNotEnoughBalance,
				FullCommand:     &fullCmd7,
			},
			want:    []response.OffchainTipBotTransferToken{},
			wantErr: true,
		},
	}

	for _, tt := range tests[:5] {
		t.Run(tt.name, func(t *testing.T) {
			mockTokens.EXPECT().GetBySymbol(tt.req.Token).Return(tt.token.token, tt.token.err)
			if tt.token.token != nil {
				mockUserBalances.EXPECT().GetUserBalanceByTokenID(tt.req.Sender, tt.token.token.ID).Return(&model.OffchainTipBotUserBalance{Amount: tt.userBal}, nil)
			}
			mockActivityLogs.EXPECT().CreateActivityLog(gomock.Any()).Return(nil)
			if !tt.wantErr {
				mockTransferHistories.EXPECT().CreateTransferHistories(gomock.Eq(tt.transferHistories)).Return(tt.transferHistories, nil)
				tokenID := uuid.UUID{}
				totalAmount := tt.activitiesLogs.Amount * float64(len(tt.req.Recipients))
				upsertBatch := []model.OffchainTipBotUserBalance{{UserID: tt.req.Sender, TokenID: tokenID, ChangedAmount: -totalAmount}}
				for _, r := range tt.req.Recipients {
					upsertBatch = append(upsertBatch, model.OffchainTipBotUserBalance{UserID: r, TokenID: tokenID, ChangedAmount: tt.activitiesLogs.Amount, Amount: tt.activitiesLogs.Amount})
				}
				mockUserBalances.EXPECT().UpsertBatch(upsertBatch).Return(nil)
				mockServiceCoingecko.EXPECT().GetCoinPrice(gomock.Any(), gomock.Any()).Return(tt.coinPrice, nil)
			}

			got, err := e.TransferToken(tt.req)
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
