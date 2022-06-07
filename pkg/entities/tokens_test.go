package entities

import (
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
					Id:                  10,
					Address:             "",
					Symbol:              "12",
					ChainID:             1234,
					Decimals:            10,
					DiscordBotSupported: true,
					CoinGeckoID:         "1234",
					Name:                "nminhdai",
					GuildDefault:        false,
					GuildID:             "980100825579917343",
				},
			},
			wantErr: false,
		},
	}

	tokenParam := model.Token{
		ID:                  10,
		Address:             "",
		Symbol:              "12",
		ChainID:             1234,
		Decimals:            10,
		DiscordBotSupported: true,
		CoinGeckoID:         "1234",
		Name:                "NMINHDAI",
		GuildDefault:        false,
	}

	guildConfigTokenParam := model.GuildConfigToken{
		GuildID: "980100825579917343",
		TokenID: 10,
		Active:  true,
	}

	uToken.EXPECT().UpsertOne(tokenParam).Return(nil).AnyTimes()
	uGuildConfigToken.EXPECT().UpsertMany(guildConfigTokenParam).Return(nil).AnyTimes()

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
			if err := e.UpsertGuildCustomTokenConfig(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("Entity.UpsertGuildCustomTokenConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
