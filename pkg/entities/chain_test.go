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
	mock_chain "github.com/defipod/mochi/pkg/repo/chain/mocks"
	"github.com/defipod/mochi/pkg/repo/pg"
	"github.com/defipod/mochi/pkg/service"
	"github.com/golang/mock/gomock"
)

func TestEntity_GetChainIdBySymbol(t *testing.T) {
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
		symbol string
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
	uChain := mock_chain.NewMockStore(ctrl)
	r.Chain = uChain
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.Chain
		wantErr bool
	}{
		{
			name: "test return successfully",
			fields: fields{
				repo: r,
			},
			args: args{
				symbol: "eth",
			},
			want: []model.Chain{
				{
					ID:       1,
					Name:     "Ethereum Mainnet",
					Currency: "ETH",
				},
			},
			wantErr: false,
		},
		{
			name: "test return not found",
			fields: fields{
				repo: r,
			},
			args: args{
				symbol: "ethaaaaaaaaaa",
			},
			want: []model.Chain{
				{
					ID:       1,
					Name:     "Ethereum Mainnet",
					Currency: "ETH",
				},
				{
					ID:       56,
					Name:     "Binance Smart Chain Mainnet",
					Currency: "BNB",
				},
				{
					ID:       250,
					Name:     "Fantom Opera",
					Currency: "FTM",
				},
			},
			wantErr: false,
		},
	}

	resList := []model.Chain{
		{
			ID:       1,
			Name:     "Ethereum Mainnet",
			Currency: "ETH",
		},
		{
			ID:       56,
			Name:     "Binance Smart Chain Mainnet",
			Currency: "BNB",
		},
		{
			ID:       250,
			Name:     "Fantom Opera",
			Currency: "FTM",
		},
	}
	uChain.EXPECT().GetAll().Return(resList, nil).AnyTimes()
	uChain.EXPECT().GetAll().Return(resList, nil).AnyTimes()

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
			got, err := e.GetChainIdBySymbol(tt.args.symbol)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.GetChainIdBySymbol() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entity.GetChainIdBySymbol() = %v, want %v", got, tt.want)
			}
		})
	}
}
