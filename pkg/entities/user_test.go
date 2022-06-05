package entities

import (
	"errors"
	"reflect"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/defipod/mochi/pkg/cache"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/discordwallet"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/repo"
	mock_guild_user_activity_log "github.com/defipod/mochi/pkg/repo/guild_user_activity_log/mocks"
	mock_guild_user_xp "github.com/defipod/mochi/pkg/repo/guild_user_xp/mocks"
	"github.com/defipod/mochi/pkg/repo/pg"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/service"
	"github.com/golang/mock/gomock"
)

func TestEntity_SendGiftXp(t *testing.T) {
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
		guildID      string
		userID       string
		earnedXp     int
		activityName string
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
	uXp := mock_guild_user_xp.NewMockStore(ctrl)
	uLog := mock_guild_user_activity_log.NewMockStore(ctrl)

	r.GuildUserXP = uXp
	r.GuildUserActivityLog = uLog
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.HandleUserActivityResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test gift xp successfully",
			fields: fields{
				repo: r,
			},
			args: args{
				guildID: "552427722551459840",
				userID: "973069332034752522",
				earnedXp: 10,
				activityName: "gifted",
			},
			want: &response.HandleUserActivityResponse{
				GuildID: "552427722551459840",
				UserID: "973069332034752522",
				Action: "gifted",
				CurrentXP: 20,
				CurrentLevel: 0,
				LevelUp: false,
			},
			wantErr: false,
		},
		{
			name: "case user not exist in server, cannot gift xp to this user",
			fields: fields{
				repo: r,
			},
			args: args{
				guildID: "abc",
				userID: "abc",
				earnedXp: 10,
				activityName: "gifted",
			},
			want: nil,
			wantErr: true,
		},
	}
	guildUserXpParam := model.GuildUserXP{
		GuildID: "552427722551459840",
		UserID: "973069332034752522",
		TotalXP: 20,
		Level: 0,
	}
	guildUserActivityLog := model.GuildUserActivityLog{
		GuildID: "552427722551459840",
		UserID: "973069332034752522",
		EarnedXP: 10,
		ActivityName: "gifted",
	}

	// case sucess send gift xp
	uXp.EXPECT().GetOne("552427722551459840", "973069332034752522").Return(&guildUserXpParam, nil).AnyTimes()
	uLog.EXPECT().CreateOne(guildUserActivityLog).Return(nil).AnyTimes()
	// case cannot find user, user not exist in server discord
	uXp.EXPECT().GetOne("abc", "abc").Return(nil, errors.New("Error cannot find user in server")).AnyTimes()

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
			got, err := e.SendGiftXp(tt.args.guildID, tt.args.userID, tt.args.earnedXp, tt.args.activityName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.SendGiftXp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entity.SendGiftXp() = %v, want %v", got, tt.want)
			}
		})
	}
}
