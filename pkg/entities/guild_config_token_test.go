package entities

import (
	"reflect"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/defipod/mochi/pkg/cache"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/discordwallet"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/repo"
	mock_guild_config_token "github.com/defipod/mochi/pkg/repo/guild_config_token/mocks"
	"github.com/defipod/mochi/pkg/repo/pg"
	mock_token "github.com/defipod/mochi/pkg/repo/token/mocks"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/service"
	"github.com/golang/mock/gomock"
)

func TestEntity_CheckExistTokenConfig(t *testing.T) {
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
		tokenId int
		guildID string
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
	uGuildConfigToken := mock_guild_config_token.NewMockStore(ctrl)
	r.GuildConfigToken = uGuildConfigToken
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "test return successfully",
			fields: fields{
				repo: r,
			},
			args: args{
				tokenId: 10,
				guildID: "980100825579917343",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "test return not found",
			fields: fields{
				repo: r,
			},
			args: args{
				tokenId: 12,
				guildID: "980100825579917343",
			},
			want:    false,
			wantErr: false,
		},
	}

	resList := []model.GuildConfigToken{
		{
			TokenID: 10,
			GuildID: "980100825579917343",
		},
		{
			TokenID: 13,
			GuildID: "980100825579917343",
		},
	}
	uGuildConfigToken.EXPECT().GetAll().Return(resList, nil).AnyTimes()
	uGuildConfigToken.EXPECT().GetAll().Return(resList, nil).AnyTimes()

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

			got, err := e.CheckExistTokenConfig(tt.args.tokenId, tt.args.guildID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.CheckExistTokenConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Entity.CheckExistTokenConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntity_UpsertGuildCustomTokenConfig(t *testing.T) {
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
		req request.UpsertCustomTokenConfigRequest
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

	uToken := mock_token.NewMockStore(ctrl)
	uGuildConfigToken := mock_guild_config_token.NewMockStore(ctrl)

	r.Token = uToken
	r.GuildConfigToken = uGuildConfigToken

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
				request.UpsertCustomTokenConfigRequest{
					Id: 0,

					Address:             "0x6c021Ae822BEa943b2E66552bDe1D2696a53fbB6",
					Symbol:              "ftm",
					Chain:               "ftm",
					ChainID:             0,
					Decimals:            0,
					DiscordBotSupported: true,
					CoinGeckoID:         "",

					GuildID:      "980100825579917343",
					Name:         "",
					GuildDefault: true,
					Active:       false,
				},
			},
			wantErr: false,
		},
	}

	tokenParam := model.Token{
		Address:             "0x6c021Ae822BEa943b2E66552bDe1D2696a53fbB6",
		Symbol:              "ftm",
		ChainID:             250,
		Decimals:            18,
		DiscordBotSupported: true,
		CoinGeckoID:         "fantom",
		Name:                "Fantom",
		GuildDefault:        false,
	}

	uToken.EXPECT().GetBySymbol("ftm", true).Return(tokenParam, nil).AnyTimes()
	guildConfigTokenParam := model.GuildConfigToken{
		GuildID: "980100825579917343",
		TokenID: tokenParam.ID,
		Active:  false,
	}

	uGuildConfigToken.EXPECT().CreateOne(guildConfigTokenParam).Return(nil).AnyTimes()

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
			if err := e.CreateGuildCustomTokenConfig(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("Entity.CreateGuildCustomTokenConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEntity_ListAllTokenID(t *testing.T) {
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

	uGuildConfigToken := mock_guild_config_token.NewMockStore(ctrl)
	r.GuildConfigToken = uGuildConfigToken
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []int
		wantErr bool
	}{
		{
			name: "test return successfully",
			fields: fields{
				repo: r,
			},
			args: args{
				guildID: "980100825579917343",
			},

			want: []int{
				1, 2,
			},
			wantErr: false,
		},
	}

	resList := []model.GuildConfigToken{
		{
			GuildID: "980100825579917343",
			TokenID: 1,
		},
		{
			GuildID: "980100825579917343",
			TokenID: 2,
		},
		{
			GuildID: "980100825579917322",
			TokenID: 1,
		},
	}

	uGuildConfigToken.EXPECT().GetAll().Return(resList, nil).AnyTimes()

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
			got, err := e.ListAllTokenID(tt.args.guildID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.ListAllTokenID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entity.ListAllTokenID() = %v, want %v", got, tt.want)
			}
		})
	}
}
