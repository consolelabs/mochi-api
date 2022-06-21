package entities

import (
	"errors"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/defipod/mochi/pkg/cache"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/discordwallet"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/repo"
	mock_guild_config_sales_tracker "github.com/defipod/mochi/pkg/repo/guild_config_sales_tracker/mocks"
	mock_nft_sales_tracker "github.com/defipod/mochi/pkg/repo/nft_sales_tracker/mocks"
	"github.com/defipod/mochi/pkg/repo/pg"
	"github.com/defipod/mochi/pkg/service"
	"github.com/defipod/mochi/pkg/service/indexer"
	"github.com/defipod/mochi/pkg/util"
	"github.com/golang/mock/gomock"
)

func TestEntity_CreateNFTSalesTracker(t *testing.T) {
	type fields struct {
		repo     *repo.Repo
		store    repo.Store
		log      logger.Logger
		dcwallet discordwallet.IDiscordWallet
		discord  *discordgo.Session
		cache    cache.Cache
		svc      *service.Service
		cfg      config.Config
		indexer  indexer.Service
	}
	type args struct {
		addr     string
		platform string
		guildID  string
	}

	ctrl:=gomock.NewController(t)
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
	salesTracker:= mock_nft_sales_tracker.NewMockStore(ctrl)
	configSalesTracker := mock_guild_config_sales_tracker.NewMockStore(ctrl)
	r.NFTSalesTracker = salesTracker
	r.GuildConfigSalesTracker = configSalesTracker

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "add sales tracker successfully",
			fields: fields{
				repo: r,
			},
			args: args{
				addr: "0xakjsdbajksbfqwueh182937h8123hbj1h2b3",
				platform: "ethereum",
				guildID: "863278424433229854",
			},
			wantErr: false,
		},
		{
			name: "invalid guild id",
			fields: fields{
				repo: r,
			},
			args: args{
				addr: "0xakjsdbajksbfqwueh182937h8123hbj1h2b3",
				platform: "ethereum",
				guildID: "123",
			},
			wantErr: true,
		},
	}
	correctSalesTracker:=model.InsertNFTSalesTracker{
		ContractAddress: "0xakjsdbajksbfqwueh182937h8123hbj1h2b3",
		Platform: "ethereum",
		SalesConfigID: "abab53eb-c13f-4f6a-b617-27d38a08e519",
	}
	invalidSalesTracker:=model.InsertNFTSalesTracker{
		ContractAddress: "0xakjsdbajksbfqwueh182937h8123hbj1h2b3",
		Platform: "ethereum",
		SalesConfigID: "abc",
	}
	config := model.GuildConfigSalesTracker{
		ID: util.GetNullUUID("abab53eb-c13f-4f6a-b617-27d38a08e519"),
		GuildID: "863278424433229854",
		ChannelID: "863278424433229854",
	}

	configSalesTracker.EXPECT().GetByGuildID("863278424433229854").Return(&config,nil).AnyTimes()
	configSalesTracker.EXPECT().GetByGuildID("123").Return(nil,errors.New("error")).AnyTimes()
	salesTracker.EXPECT().FirstOrCreate(&correctSalesTracker).Return(nil).AnyTimes()
	salesTracker.EXPECT().FirstOrCreate(&invalidSalesTracker).Return(errors.New("config id invalid")).AnyTimes()

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
				indexer:  tt.fields.indexer,
			}
			if err := e.CreateNFTSalesTracker(tt.args.addr, tt.args.platform, tt.args.guildID); (err != nil) != tt.wantErr {
				t.Errorf("Entity.CreateNFTSalesTracker() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}