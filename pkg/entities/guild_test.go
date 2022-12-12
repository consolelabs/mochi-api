package entities

import (
	"errors"
	"reflect"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/golang/mock/gomock"

	"github.com/defipod/mochi/pkg/cache"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/discordwallet"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/repo"
	mock_discord_guilds "github.com/defipod/mochi/pkg/repo/discord_guilds/mocks"
	"github.com/defipod/mochi/pkg/repo/pg"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/service"
	"github.com/defipod/mochi/pkg/service/abi"
	"github.com/defipod/mochi/pkg/service/indexer"
	"github.com/defipod/mochi/pkg/service/marketplace"
)

func TestEntity_GetGuild(t *testing.T) {
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
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cfg := config.LoadTestConfig()
	s := pg.NewPostgresStore(&cfg)
	r := pg.NewRepo(s.DB())
	dGuilds := mock_discord_guilds.NewMockStore(ctrl)
	r.DiscordGuilds = dGuilds

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.GetGuildResponse
		wantErr bool
	}{
		{
			name: "test get guils succesfully",
			fields: fields{
				repo: r,
			},
			args: args{
				guildID: "981128899280908299",
			},
			want: &response.GetGuildResponse{
				ID:        "981128899280908299",
				Name:      "a",
				BotScopes: model.JSONArrayString{"*"},
				GlobalXP:  false,
				Active:    true,
			},
			wantErr: false,
		},
		{
			name: "case guildId not exists, cannot find",
			fields: fields{
				repo: r,
			},
			args: args{
				guildID: "abc",
			},
			want:    nil,
			wantErr: true,
		},
	}

	discordGuild := model.DiscordGuild{
		ID:        "981128899280908299",
		Name:      "a",
		BotScopes: model.JSONArrayString{"*"},
		GlobalXP:  false,
	}

	dGuilds.EXPECT().GetByID("981128899280908299").Return(&discordGuild, nil).AnyTimes()
	dGuilds.EXPECT().GetByID("abc").Return(nil, errors.New("cannot find guild id"))

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
			got, err := e.GetGuild(tt.args.guildID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.GetGuild() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entity.GetGuild() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestEntity_GetGuilds(t *testing.T) {
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
	dGuilds := mock_discord_guilds.NewMockStore(ctrl)
	r.DiscordGuilds = dGuilds

	tests := []struct {
		name    string
		fields  fields
		want    *response.GetGuildsResponse
		wantErr bool
	}{
		{
			name: "test get guils succesfully",
			fields: fields{
				repo: r,
			},
			want: &response.GetGuildsResponse{
				Data: []*response.GetGuildResponse{
					{
						ID:        "895659000996200508",
						Name:      "hnh",
						BotScopes: model.JSONArrayString{"*"},
						GlobalXP:  false,
					},
					{
						ID:        "981128899280908299",
						Name:      "a",
						BotScopes: model.JSONArrayString{"*"},
						GlobalXP:  false,
					},
				},
			},
			wantErr: false,
		},
	}

	discordGuildslst := []model.DiscordGuild{
		{
			ID:        "895659000996200508",
			Name:      "hnh",
			BotScopes: model.JSONArrayString{"*"},
			GlobalXP:  false,
		},
		{
			ID:        "981128899280908299",
			Name:      "a",
			BotScopes: model.JSONArrayString{"*"},
			GlobalXP:  false,
		},
	}

	dGuilds.EXPECT().Gets().Return(discordGuildslst, nil).AnyTimes()

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
			got, err := e.GetGuilds()
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.GetGuilds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entity.GetGuilds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntity_UpdateGuild(t *testing.T) {
	type fields struct {
		repo        *repo.Repo
		store       repo.Store
		log         logger.Logger
		dcwallet    discordwallet.IDiscordWallet
		discord     *discordgo.Session
		cache       cache.Cache
		svc         *service.Service
		cfg         config.Config
		indexer     indexer.Service
		abi         abi.Service
		marketplace marketplace.Service
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cfg := config.LoadTestConfig()
	s := pg.NewPostgresStore(&cfg)
	r := pg.NewRepo(s.DB())
	log := logger.NewLogrusLogger()
	dGuilds := mock_discord_guilds.NewMockStore(ctrl)
	r.DiscordGuilds = dGuilds

	type args struct {
		guildID string
		req     request.UpdateGuildRequest
	}

	guildIdExist := "552427722551459840"
	modelGuildExist := &model.DiscordGuild{
		ID:         guildIdExist,
		Name:       "bestKN",
		BotScopes:  make(model.JSONArrayString, 0),
		GlobalXP:   false,
		LogChannel: "895659000996200508",
		Active:     true,
	}

	guildIDFailed := "fail_guild_id"
	modelGuildFailed := &model.DiscordGuild{
		ID:         guildIDFailed,
		Name:       "fail_guild",
		BotScopes:  make(model.JSONArrayString, 0),
		GlobalXP:   false,
		LogChannel: "fail_log",
		Active:     true,
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test update guild succesfully",
			fields: fields{
				repo: r,
			},
			args: args{
				guildID: guildIdExist,
				req: request.UpdateGuildRequest{
					LogChannel: &modelGuildExist.LogChannel,
					GlobalXP:   &modelGuildExist.GlobalXP,
					Active:     &modelGuildExist.Active,
				},
			},
			wantErr: false,
		},
		{
			name: "test update guild failed",
			fields: fields{
				repo: r,
				log:  log,
			},
			args: args{
				guildID: guildIDFailed,
				req: request.UpdateGuildRequest{
					LogChannel: &modelGuildFailed.LogChannel,
					GlobalXP:   &modelGuildFailed.GlobalXP,
					Active:     &modelGuildFailed.Active,
				},
			},
			wantErr: true,
		},
		{
			name: "test update guild not found",
			fields: fields{
				repo: r,
			},
			args: args{
				guildID: "552427722551459841",
				req: request.UpdateGuildRequest{
					LogChannel: &modelGuildExist.LogChannel,
					GlobalXP:   &modelGuildExist.GlobalXP,
					Active:     &modelGuildExist.Active,
				},
			},
			wantErr: true,
		},
	}

	// update guild successfully
	dGuilds.EXPECT().GetByID(guildIdExist).Return(&model.DiscordGuild{
		ID:         guildIdExist,
		Name:       "bestKN",
		BotScopes:  make(model.JSONArrayString, 0),
		GlobalXP:   false,
		LogChannel: "895659000996200508",
		Active:     true,
	}, nil).AnyTimes()
	dGuilds.EXPECT().Update(modelGuildExist).Return(nil)

	// update guild failed
	dGuilds.EXPECT().GetByID(guildIDFailed).Return(&model.DiscordGuild{
		ID:         guildIDFailed,
		Name:       "fail_guild",
		BotScopes:  make(model.JSONArrayString, 0),
		GlobalXP:   false,
		LogChannel: "fail_log",
		Active:     true,
	}, nil).AnyTimes()
	dGuilds.EXPECT().Update(modelGuildFailed).Return(errors.New("fail update guild")).AnyTimes()

	// get guild not exist
	dGuilds.EXPECT().GetByID("552427722551459841").Return(nil, errors.New("guild not exist")).AnyTimes()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Entity{
				repo:        tt.fields.repo,
				store:       tt.fields.store,
				log:         tt.fields.log,
				dcwallet:    tt.fields.dcwallet,
				discord:     tt.fields.discord,
				cache:       tt.fields.cache,
				svc:         tt.fields.svc,
				cfg:         tt.fields.cfg,
				indexer:     tt.fields.indexer,
				abi:         tt.fields.abi,
				marketplace: tt.fields.marketplace,
			}
			if err := e.UpdateGuild(tt.args.guildID, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("Entity.UpdateGuild() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
