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
	mock_guild_config_level_role "github.com/defipod/mochi/pkg/repo/guild_config_level_role/mocks"
	"github.com/defipod/mochi/pkg/repo/pg"
	"github.com/defipod/mochi/pkg/service"
	"github.com/defipod/mochi/pkg/service/abi"
	"github.com/defipod/mochi/pkg/service/indexer"
	"github.com/defipod/mochi/pkg/service/marketplace"
	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
)

func TestEntity_GetGuildLevelRoleConfigs(t *testing.T) {
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

	cfg := config.Config{
		DBUser: "postgres",
		DBPass: "postgres",
		DBHost: "localhost",
		DBPort: "5434",
		DBName: "mochi_local",

		InDiscordWalletMnemonic: "holiday frequent toy bachelor auto use style result recycle crumble glue blouse",
		FantomRPC:               "sample",
		FantomScan:              "sample",
		FantomScanAPIKey:        "sample",

		EthereumRPC:        "sample",
		EthereumScan:       "sample",
		EthereumScanAPIKey: "sample",

		BscRPC:        "sample",
		BscScan:       "sample",
		BscScanAPIKey: "sample",

		DiscordToken: "sample",

		RedisURL: "redis://localhost:6379/0",
	}

	s := pg.NewPostgresStore(&cfg)
	r := pg.NewRepo(s.DB())
	guildConfigLR := mock_guild_config_level_role.NewMockStore(ctrl)

	r.GuildConfigLevelRole = guildConfigLR

	type args struct {
		guildID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.GuildConfigLevelRole
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "successful",
			fields: fields{
				repo: r,
			},
			args: args{
				guildID: "863278424433229854",
			},
			want: []model.GuildConfigLevelRole{
				{
					GuildID: "863278424433229854",
					Level:   5,
					RoleID:  "mod",
					LevelConfig: &model.ConfigXpLevel{
						Level: 5,
						MinXP: 100,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "invalid guild id",
			fields: fields{
				repo: r,
			},
			args: args{
				guildID: "863278424433229abc",
			},
			want:    nil,
			wantErr: true,
		},
	}
	config := []model.GuildConfigLevelRole{
		{
			GuildID: "863278424433229854",
			Level:   5,
			RoleID:  "mod",
			LevelConfig: &model.ConfigXpLevel{
				Level: 5,
				MinXP: 100,
			},
		},
	}
	guildConfigLR.EXPECT().GetByGuildID("863278424433229854").Return(config, nil).AnyTimes()
	guildConfigLR.EXPECT().GetByGuildID("863278424433229abc").Return(nil, errors.New("invalid guild id")).AnyTimes()

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
			got, err := e.GetGuildLevelRoleConfigs(tt.args.guildID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.GetGuildLevelRoleConfigs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entity.GetGuildLevelRoleConfigs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntity_RemoveGuildLevelRoleConfig(t *testing.T) {
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

	cfg := config.Config{
		DBUser: "postgres",
		DBPass: "postgres",
		DBHost: "localhost",
		DBPort: "5434",
		DBName: "mochi_local",

		InDiscordWalletMnemonic: "holiday frequent toy bachelor auto use style result recycle crumble glue blouse",
		FantomRPC:               "sample",
		FantomScan:              "sample",
		FantomScanAPIKey:        "sample",

		EthereumRPC:        "sample",
		EthereumScan:       "sample",
		EthereumScanAPIKey: "sample",

		BscRPC:        "sample",
		BscScan:       "sample",
		BscScanAPIKey: "sample",

		DiscordToken: "sample",

		RedisURL: "redis://localhost:6379/0",
	}

	s := pg.NewPostgresStore(&cfg)
	r := pg.NewRepo(s.DB())
	guildConfigLR := mock_guild_config_level_role.NewMockStore(ctrl)

	r.GuildConfigLevelRole = guildConfigLR
	type args struct {
		guildID string
		level   int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "successful",
			fields: fields{
				repo: r,
			},
			args: args{
				guildID: "863278424433229854",
				level:   3,
			},
			wantErr: false,
		},
		{
			name: "invalid guild id",
			fields: fields{
				repo: r,
			},
			args: args{
				guildID: "863278424433229abc",
				level:   3,
			},
			wantErr: true,
		},
	}
	guildConfigLR.EXPECT().DeleteOne("863278424433229854", 3).Return(nil)
	guildConfigLR.EXPECT().DeleteOne("863278424433229abc", 3).Return(errors.New("record not found"))

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
			if err := e.RemoveGuildLevelRoleConfig(tt.args.guildID, tt.args.level); (err != nil) != tt.wantErr {
				t.Errorf("Entity.RemoveGuildLevelRoleConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEntity_GetUserRoleByLevel(t *testing.T) {
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

	cfg := config.Config{
		DBUser: "postgres",
		DBPass: "postgres",
		DBHost: "localhost",
		DBPort: "5434",
		DBName: "mochi_local",

		InDiscordWalletMnemonic: "holiday frequent toy bachelor auto use style result recycle crumble glue blouse",
		FantomRPC:               "sample",
		FantomScan:              "sample",
		FantomScanAPIKey:        "sample",

		EthereumRPC:        "sample",
		EthereumScan:       "sample",
		EthereumScanAPIKey: "sample",

		BscRPC:        "sample",
		BscScan:       "sample",
		BscScanAPIKey: "sample",

		DiscordToken: "sample",

		RedisURL: "redis://localhost:6379/0",
	}

	s := pg.NewPostgresStore(&cfg)
	r := pg.NewRepo(s.DB())
	guildConfigLR := mock_guild_config_level_role.NewMockStore(ctrl)

	r.GuildConfigLevelRole = guildConfigLR

	type args struct {
		guildID string
		level   int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "successful",
			fields: fields{
				repo: r,
			},
			args: args{
				guildID: "863278424433229854",
				level:   3,
			},
			want:    "abc",
			wantErr: false,
		},
		{
			name: "invalid guild id",
			fields: fields{
				repo: r,
			},
			args: args{
				guildID: "863278424433229abc",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "not found lr with highest level",
			fields: fields{
				repo: r,
			},
			args: args{
				guildID: "863278424433229854",
				level:   1,
			},
			want:    "",
			wantErr: true,
		},
	}
	config := model.GuildConfigLevelRole{
		GuildID: "863278424433229854",
		Level:   5,
		RoleID:  "abc",
		LevelConfig: &model.ConfigXpLevel{
			Level: 5,
			MinXP: 100,
		},
	}
	guildConfigLR.EXPECT().GetHighest("863278424433229854", 3).Return(&config, nil).AnyTimes()
	guildConfigLR.EXPECT().GetHighest("863278424433229abc", gomock.Any()).Return(nil, errors.New("invalid guild id")).AnyTimes()
	guildConfigLR.EXPECT().GetHighest(gomock.Any(), 1).Return(nil, gorm.ErrRecordNotFound)
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
			got, err := e.GetUserRoleByLevel(tt.args.guildID, tt.args.level)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.GetUserRoleByLevel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entity.GetUserRoleByLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
