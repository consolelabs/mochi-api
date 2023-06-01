package entities

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/golang/mock/gomock"
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/cache"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/discordwallet"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/repo"
	mock_chain "github.com/defipod/mochi/pkg/repo/chain/mocks"
	coingeckosupportedtokens "github.com/defipod/mochi/pkg/repo/coingecko_supported_tokens"
	mock_coingeckosupportedtokens "github.com/defipod/mochi/pkg/repo/coingecko_supported_tokens/mocks"
	mock_guild_config_token "github.com/defipod/mochi/pkg/repo/guild_config_token/mocks"
	"github.com/defipod/mochi/pkg/repo/pg"
	mock_token "github.com/defipod/mochi/pkg/repo/token/mocks"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/service"
	mock_coingecko "github.com/defipod/mochi/pkg/service/coingecko/mocks"
	mock_covalent "github.com/defipod/mochi/pkg/service/covalent/mocks"
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

	log := logger.NewLogrusLogger()
	svc, _ := service.NewService(cfg, log)

	uToken := mock_token.NewMockStore(ctrl)
	chainMock := mock_chain.NewMockStore(ctrl)
	cgTokensMock := mock_coingeckosupportedtokens.NewMockStore(ctrl)
	gctMock := mock_guild_config_token.NewMockStore(ctrl)
	cgMock := mock_coingecko.NewMockService(ctrl)
	covalentMock := mock_covalent.NewMockService(ctrl)

	r.Token = uToken
	r.GuildConfigToken = gctMock
	r.Chain = chainMock
	r.CoingeckoSupportedTokens = cgTokensMock
	svc.CoinGecko = cgMock
	svc.Covalent = covalentMock

	tests := []struct {
		name       string
		fields     fields
		args       args
		coinIds    []string
		coinPrices map[string]float64
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			name: "Create invite successfully",
			fields: fields{
				repo: r,
				svc:  svc,
			},
			args: args{
				request.UpsertCustomTokenConfigRequest{
					Address: "0x0e09fabb73bd3ade0a17ecc321fd13a19e81ce82",
					Symbol:  "cake",
					Chain:   "bsc",
					Name:    "PancakeSwap",
					GuildID: "980100825579917343",
					Active:  false,
				},
			},
			coinIds:    []string{"pancakeswap-token"},
			coinPrices: map[string]float64{"pancakeswap-token": 1.7},
			wantErr:    false,
		},
	}

	token := model.Token{
		ID:                  1,
		Address:             "0x0e09fabb73bd3ade0a17ecc321fd13a19e81ce82",
		Symbol:              "cake",
		ChainID:             56,
		Decimals:            18,
		DiscordBotSupported: true,
		CoinGeckoID:         "pancakeswap-token",
		Name:                "PancakeSwap",
		Chain: &model.Chain{
			ID:        56,
			Currency:  "bnb",
			ShortName: "bsc",
		},
	}
	chain := model.Chain{
		ID:          1,
		ShortName:   "bsc",
		Currency:    "bnb",
		CoinGeckoID: "binance-smart-chain",
	}

	gct := model.GuildConfigToken{
		GuildID: "980100825579917343",
		TokenID: token.ID,
		Active:  true,
	}
	coin := response.GetCoinResponse{
		ID:              token.CoinGeckoID,
		Name:            token.Name,
		Symbol:          token.Symbol,
		AssetPlatformID: chain.CoinGeckoID,
	}
	htp := response.HistoricalTokenPricesResponse{
		Data: []response.HistoricalTokenPrice{
			{
				Name:     token.Name,
				Decimals: 18,
				Symbol:   token.Symbol,
				Address:  token.Address,
			},
		},
	}
	tokens := []model.CoingeckoSupportedTokens{{ID: "pancakeswap-token", Symbol: "cake", Name: "PancakeSwap"}}

	chainMock.EXPECT().GetByShortName(chain.ShortName).Return(&chain, nil).AnyTimes()
	cgTokensMock.EXPECT().GetOne(token.Symbol).Return(nil, gorm.ErrRecordNotFound).AnyTimes()
	cgTokensMock.EXPECT().List(coingeckosupportedtokens.ListQuery{Symbol: token.Symbol}).Return(tokens, nil).AnyTimes()
	cgMock.EXPECT().GetCoin(tokens[0].ID).Return(&coin, nil, http.StatusOK).AnyTimes()
	covalentMock.EXPECT().GetHistoricalTokenPrices(chain.ID, chain.Currency, token.Address).Return(&htp, nil, http.StatusOK).AnyTimes()
	uToken.EXPECT().GetOneBySymbol(token.Symbol).Return(&token, nil).AnyTimes()
	uToken.EXPECT().CreateOne(token).Return(nil).AnyTimes()
	gctMock.EXPECT().GetByGuildIDAndTokenID(gct.GuildID, gct.TokenID).Return(nil, nil).AnyTimes()
	gctMock.EXPECT().CreateOne(gct).Return(nil).AnyTimes()

	for _, tt := range tests {
		if tt.coinIds != nil && len(tt.coinIds) != 0 {
			for _, coinId := range tt.coinIds {
				cgMock.EXPECT().GetCoinPrice([]string{coinId}, "usd").Return(map[string]float64{coinId: tt.coinPrices[coinId]}, nil).AnyTimes()
			}
		}

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
		guildId string
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
				guildId: "1234",
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
	}

	uToken.EXPECT().GetSupportedTokenByGuildId("1234").Return(resList, nil).AnyTimes()
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
			got, err := e.GetAllSupportedToken(tt.args.guildId)
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
