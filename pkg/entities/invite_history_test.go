package entities

import (
	"errors"
	"reflect"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/golang/mock/gomock"
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/cache"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/discordwallet"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/repo"
	mock_invite_histories "github.com/defipod/mochi/pkg/repo/invite_histories/mocks"
	"github.com/defipod/mochi/pkg/repo/pg"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/service"
)

func TestEntity_CreateInviteHistory(t *testing.T) {
	type fields struct {
		repo     *repo.Repo
		store    repo.Store
		log      logger.Logger
		dcwallet discordwallet.IDiscordWallet
		discord  *discordgo.Session
		cache    cache.Cache
		svc      *service.Service
		cfg      config.Config
	}

	type args struct {
		req request.CreateInviteHistoryRequest
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := config.Config{
		DBUser: "postgres",
		DBPass: "postgres",
		DBHost: "localhost",
		DBPort: "5434",
		DBName: "mochi_local",

		InDiscordWalletMnemonic: "holiday frequent toy bachelor auto use style result recycle crumble glue blouse",
		FantomRPC:               "https://rpc.ftm.tools",
		FantomScan:              "https://api.ftmscan.com/api?",
		FantomScanAPIKey:        "XEKSVDF5VWQDY5VY6ZNT6AK9QPQRH483EF",

		EthereumRPC:        "https://mainnet.infura.io/v3/5b389eb75c514cf6b1711d70084b0114",
		EthereumScan:       "https://api.etherscan.io/api?",
		EthereumScanAPIKey: "SM5BHYSNIRZ1HEWJ1JPHVTMJS95HRA6DQF",

		BscRPC:        "https://bsc-dataseed.binance.org",
		BscScan:       "https://api.bscscan.com/api?",
		BscScanAPIKey: "VTKF4RG4HP6WXQ5QTAJ8MHDDIUFYD6VZHC",

		DiscordToken: "OTcxNjMyNDMzMjk0MzQ4Mjg5.G5BEgF.rv-16ZuTzzqOv2W76OljymFxxnNpjVjCnOkn98",

		RedisURL: "redis://localhost:6379/0",
	}

	s := pg.NewPostgresStore(&cfg)
	r := pg.NewRepo(s.DB())

	uInv := mock_invite_histories.NewMockStore(ctrl)

	r.InviteHistories = uInv
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Create invite successfully",
			fields: fields{
				repo: r,
			},
			args: args{
				request.CreateInviteHistoryRequest{
					GuildID: "980100825579917343",
					Inviter: "382540937517334538",
					Invitee: "959993808970461205",
					Type:    "normal",
				},
			},
			wantErr: false,
		},
		{
			name: "Case user inviter not exits in server",
			fields: fields{
				repo: r,
			},
			args: args{
				request.CreateInviteHistoryRequest{
					GuildID: "980100825579917343",
					Inviter: "abc",
					Invitee: "959993808970461205",
					Type:    "normal",
				},
			},
			wantErr: true,
		},
		{
			name: "Case user invitee not exits in server",
			fields: fields{
				repo: r,
			},
			args: args{
				request.CreateInviteHistoryRequest{
					GuildID: "980100825579917343",
					Inviter: "959993808970461205",
					Invitee: "abc",
					Type:    "normal",
				},
			},
			wantErr: true,
		},
		{
			name: "Case guildID not exits in server",
			fields: fields{
				repo: r,
			},
			args: args{
				request.CreateInviteHistoryRequest{
					GuildID: "abc",
					Inviter: "959993808970461205",
					Invitee: "382540937517334538",
					Type:    "normal",
				},
			},
			wantErr: true,
		},
	}

	inviteHistoriesParam := model.InviteHistory{
		GuildID:   "980100825579917343",
		UserID:    "959993808970461205",
		InvitedBy: "382540937517334538",
		Type:      "normal",
	}

	inviteHistoriesFail1Param := model.InviteHistory{
		GuildID:   "980100825579917343",
		UserID:    "959993808970461205",
		InvitedBy: "abc",
		Type:      "normal",
	}

	inviteHistoriesFail2Param := model.InviteHistory{
		GuildID:   "980100825579917343",
		UserID:    "abc",
		InvitedBy: "959993808970461205",
		Type:      "normal",
	}

	inviteHistoriesFail3Param := model.InviteHistory{
		GuildID:   "abc",
		UserID:    "382540937517334538",
		InvitedBy: "959993808970461205",
		Type:      "normal",
	}

	// case success
	uInv.EXPECT().Create(&inviteHistoriesParam).Return(nil).AnyTimes()

	// case cannot find user inviter
	uInv.EXPECT().Create(&inviteHistoriesFail1Param).Return(errors.New("Error cannot find user in server")).AnyTimes()

	// case cannot find user invitee
	uInv.EXPECT().Create(&inviteHistoriesFail2Param).Return(errors.New("Error cannot find user in server")).AnyTimes()

	// case cannot find discord
	uInv.EXPECT().Create(&inviteHistoriesFail3Param).Return(errors.New("Error cannot find discord")).AnyTimes()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Entity{
				repo:     tt.fields.repo,
				store:    tt.fields.store,
				log:      tt.fields.log,
				dcwallet: tt.fields.dcwallet,
				discord:  tt.fields.discord,
				cache:    tt.fields.cache,
				svc:      tt.fields.svc,
				cfg:      tt.fields.cfg,
			}
			if err := e.CreateInviteHistory(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("Entity.CreateInviteHistory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEntity_GetInvitesLeaderboard(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mock repo
	inviteHistoriesRepo := mock_invite_histories.NewMockStore(ctrl)
	repo := &repo.Repo{
		InviteHistories: inviteHistoriesRepo,
	}

	// create entity
	cfg := config.LoadTestConfig()
	log := logger.NewLogrusLogger()
	e := &Entity{
		cfg:  cfg,
		log:  log,
		repo: repo,
	}

	type res struct {
		data []response.UserInvitesAggregation
		err  error
	}
	type args struct {
		guildID string
	}
	tests := []struct {
		name    string
		args    args
		res     res
		want    []response.UserInvitesAggregation
		wantErr bool
	}{
		{
			name: "happy_case",
			args: args{
				guildID: "962589711841525780",
			},
			res: res{
				data: []response.UserInvitesAggregation{
					{
						InviterID: "962592086849376266",
						Regular:   2,
						Fake:      0,
						Left:      0,
					},
				},
				err: nil,
			},
			want: []response.UserInvitesAggregation{
				{
					InviterID: "962592086849376266",
					Regular:   2,
					Fake:      0,
					Left:      0,
				},
			},
			wantErr: false,
		},
		{
			name: "record_not_found",
			args: args{
				guildID: "962589711841525123",
			},
			res: res{
				data: nil,
				err:  gorm.ErrRecordNotFound,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inviteHistoriesRepo.EXPECT().GetInvitesLeaderboard(tt.args.guildID).Return(tt.res.data, tt.res.err).Times(1)

			got, err := e.GetInvitesLeaderboard(tt.args.guildID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.GetInvitesLeaderboard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && tt.want != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entity.GetInvitesLeaderboard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntity_GetUserInvitesAggregation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mock repo
	inviteHistoriesRepo := mock_invite_histories.NewMockStore(ctrl)
	repo := &repo.Repo{
		InviteHistories: inviteHistoriesRepo,
	}

	// create entity
	cfg := config.LoadTestConfig()
	log := logger.NewLogrusLogger()
	e := &Entity{
		cfg:  cfg,
		log:  log,
		repo: repo,
	}

	data := response.UserInvitesAggregation{
		InviterID: "962592086849376266",
		Regular:   2,
		Fake:      0,
		Left:      0,
	}
	type res struct {
		data *response.UserInvitesAggregation
		err  error
	}
	type args struct {
		guildID string
		userID  string
	}
	tests := []struct {
		name    string
		args    args
		res     res
		want    *response.UserInvitesAggregation
		wantErr bool
	}{
		{
			name: "happy_case",
			args: args{
				guildID: "962589711841525780",
				userID:  "962592086849376266",
			},
			res: res{
				data: &data,
				err:  nil,
			},
			want:    &data,
			wantErr: false,
		},
		{
			name: "record_not_found",
			args: args{
				guildID: "962589711841525123",
				userID:  "962592086849376123",
			},
			res: res{
				data: nil,
				err:  gorm.ErrRecordNotFound,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inviteHistoriesRepo.EXPECT().GetUserInvitesAggregation(tt.args.guildID, tt.args.userID).Return(tt.res.data, tt.res.err).Times(1)

			got, err := e.GetUserInvitesAggregation(tt.args.guildID, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.GetUserInvitesAggregation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entity.GetUserInvitesAggregation() = %v, want %v", got, tt.want)
			}
		})
	}
}
