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
	"github.com/defipod/mochi/pkg/repo/pg"
	mock_token "github.com/defipod/mochi/pkg/repo/token/mocks"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/service"
	"github.com/golang/mock/gomock"
)

func TestEntity_UpsertCustomToken(t *testing.T) {
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
					Id: 0,

					Address:             "0x6c021Ae822BEa943b2E66552bDe1D2696a53fbB6",
					Symbol:              "ftm",
					Chain:               "ftm",
					ChainID:             250,
					Decimals:            18,
					DiscordBotSupported: true,
					CoinGeckoID:         "fantom",
					Name:                "Fantom",

					GuildID: "980100825579917343",

					GuildDefault: true,
					Active:       false,
				},
			},
			wantErr: false,
		},
	}

	tokenParam := model.Token{
		Address: "0x6c021Ae822BEa943b2E66552bDe1D2696a53fbB6",
		Symbol:  "FTM",

		ChainID:             250,
		Decimals:            18,
		DiscordBotSupported: true,
		CoinGeckoID:         "fantom",
		Name:                "FANTOM",
		GuildDefault:        true,
	}

	uToken.EXPECT().CreateOne(tokenParam).Return(nil).AnyTimes()

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
			if err := e.CreateCustomToken(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("Entity.UpsertCustomToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEntity_CheckExistToken(t *testing.T) {
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
	uToken := mock_token.NewMockStore(ctrl)
	r.Token = uToken
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
				symbol: "bnc",
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
				symbol: "ethaaaaaaaaaa",
			},
			want:    false,
			wantErr: false,
		},
	}

	resList := []model.Token{
		{
			ID:     1,
			Symbol: "BNC",
		},
		{
			ID:     56,
			Symbol: "BNCA",
		},
		{
			ID:     250,
			Symbol: "BNCAB",
		},
	}
	uToken.EXPECT().GetAll().Return(resList, nil).AnyTimes()
	uToken.EXPECT().GetAll().Return(resList, nil).AnyTimes()

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
			got, err := e.CheckExistToken(tt.args.symbol)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.CheckExistToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Entity.CheckExistToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntity_GetTokenBySymbol(t *testing.T) {
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
		flag   bool
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

	r.Token = uToken
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "test return successfully",
			fields: fields{
				repo: r,
			},
			args: args{
				symbol: "BNC",
				flag:   true,
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "test return not found",
			fields: fields{
				repo: r,
			},
			args: args{
				symbol: "BNCC",
				flag:   true,
			},
			want:    0,
			wantErr: false,
		},
	}

	resList := model.Token{

		ID:                  2,
		Symbol:              "BNC",
		DiscordBotSupported: true,
	}

	resList1 := model.Token{}
	uToken.EXPECT().GetBySymbol("BNC", true).Return(resList, nil).AnyTimes()
	uToken.EXPECT().GetBySymbol("BNCC", true).Return(resList1, nil).AnyTimes()
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
			got, err := e.GetTokenBySymbol(tt.args.symbol, tt.args.flag)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.GetTokenBySymbol() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Entity.GetTokenBySymbol() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntity_ListAllCustomToken(t *testing.T) {
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
		listTokenId []int
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
	r.Token = uToken

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.Token
		wantErr bool
	}{
		{
			name: "test return successfully",
			fields: fields{
				repo: r,
			},
			args: args{
				listTokenId: []int{
					1, 56,
				},
			},

			want: []model.Token{
				{
					ID:     1,
					Symbol: "BNC",
				},
				{
					ID:     56,
					Symbol: "BNCA",
				},
			},
			wantErr: false,
		},
	}

	resList := []model.Token{
		{
			ID:     1,
			Symbol: "BNC",
		},
		{
			ID:     56,
			Symbol: "BNCA",
		},
		{
			ID:     250,
			Symbol: "BNCAB",
		},
	}

	uToken.EXPECT().GetAll().Return(resList, nil).AnyTimes()
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
			got, err := e.ListAllCustomToken(tt.args.listTokenId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.ListAllCustomToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entity.ListAllCustomToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
