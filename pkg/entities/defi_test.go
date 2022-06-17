package entities

import (
	"reflect"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/defipod/mochi/pkg/cache"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/discordwallet"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/repo"

	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/service"
	mock_coingecko "github.com/defipod/mochi/pkg/service/coingecko/mocks"
	"github.com/golang/mock/gomock"
)

func TestEntity_TokenCompare(t *testing.T) {
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
		sourceSymbolInfo [][]float32
		targetSymbolInfo [][]float32
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

	svc, _ := service.NewService(cfg)

	uCompare := mock_coingecko.NewMockService(ctrl)
	svc.CoinGecko = uCompare
	tests := []struct {
		name                string
		fields              fields
		args                args
		wantTokenCompareRes *response.TokenCompareReponse
		wantErr             bool
	}{
		{
			name: "test return successfully",
			fields: fields{
				svc: svc,
			},
			args: args{
				sourceSymbolInfo: [][]float32{
					{
						1652803200000,
						2096.77,
						2110.55,
						2062.53,
						2062.53,
					},
					{
						1652817600000,
						2063.61,
						2063.88,
						2029.1,
						2029.1,
					},
				},
				targetSymbolInfo: [][]float32{
					{
						1652803200000,
						30555.49,
						30651.91,
						30117.68,
						30117.68,
					},
					{
						1652817600000,
						30159.76,
						30203.91,
						29835.27,
						29835.27,
					},
				},
			},
			wantTokenCompareRes: &response.TokenCompareReponse{
				PriceCompare: []float32{
					0.06862171,
					0.06842263,
				},
				Times: []string{
					"2022-05-17T16:00:39Z",
					"2022-05-17T20:00:57Z",
				},
			},
			wantErr: false,
		},
	}

	sourceSymbolInfoParam := [][]float32{
		{
			1652803200000,
			2096.77,
			2110.55,
			2062.53,
			2062.53,
		},
		{
			1652817600000,
			2063.61,
			2063.88,
			2029.1,
			2029.1,
		},
	}

	targetSymbolInfoParam := [][]float32{
		{
			1652803200000,
			30555.49,
			30651.91,
			30117.68,
			30117.68,
		},
		{
			1652817600000,
			30159.76,
			30203.91,
			29835.27,
			29835.27,
		},
	}

	resList := &response.TokenCompareReponse{
		PriceCompare: []float32{
			0.06862171,
			0.06842263,
		},
		Times: []string{
			"2022-05-17T16:00:39Z",
			"2022-05-17T20:00:57Z",
		},
	}

	uCompare.EXPECT().TokenCompare(sourceSymbolInfoParam, targetSymbolInfoParam).Return(resList, nil).AnyTimes()
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
			gotTokenCompareRes, err := e.TokenCompare(tt.args.sourceSymbolInfo, tt.args.targetSymbolInfo)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.TokenCompare() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTokenCompareRes, tt.wantTokenCompareRes) {
				t.Errorf("Entity.TokenCompare() = %v, want %v", gotTokenCompareRes, tt.wantTokenCompareRes)
			}
		})
	}
}
