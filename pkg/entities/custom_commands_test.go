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
	guildcustomcommand "github.com/defipod/mochi/pkg/repo/guild_custom_command"
	mock_guild_custom_command "github.com/defipod/mochi/pkg/repo/guild_custom_command/mocks"
	"github.com/defipod/mochi/pkg/repo/pg"
	"github.com/defipod/mochi/pkg/service"
	"github.com/golang/mock/gomock"
)

func TestEntity_ListCustomCommands(t *testing.T) {
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
		guildID string
		enabled *bool
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

	gcc := mock_guild_custom_command.NewMockStore(ctrl)
	r.GuildCustomCommand = gcc
	tTrue := true

	gccResp := []model.GuildCustomCommand{}

	gcc.EXPECT().GetAll(guildcustomcommand.GetAllQuery{GuildID: "981128899280908299", Enabled: &tTrue}).Return(gccResp, nil).AnyTimes()
	gcc.EXPECT().GetAll(guildcustomcommand.GetAllQuery{GuildID: "abc", Enabled: &tTrue}).Return(nil, errors.New("cannot get all custom commands")).AnyTimes()

	test := []struct {
		name    string
		fields  fields
		args    args
		want    []model.GuildCustomCommand
		wantErr bool
	}{
		{
			name: "test list custom command successfully",
			fields: fields{
				repo: r,
			},
			args: args{
				guildID: "981128899280908299",
				enabled: &tTrue,
			},
			want:    []model.GuildCustomCommand{},
			wantErr: false,
		},
		{
			name: "case user not exist, cannot get commands",
			fields: fields{
				repo: r,
			},
			args: args{
				guildID: "abc",
				enabled: &tTrue,
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			e := Entity{
				repo:     tt.fields.repo,
				store:    tt.fields.store,
				log:      tt.fields.log,
				dcwallet: tt.fields.dcwallet,
				discord:  tt.fields.discord,
				cache:    tt.fields.cache,
				svc:      tt.fields.svc,
				cfg:      tt.fields.cfg,
			}
			got, err := e.ListCustomCommands(tt.args.guildID, tt.args.enabled)
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
