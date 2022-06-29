package entities

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/defipod/mochi/pkg/cache"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/discordwallet"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/repo"
	mock_guild_config_sales_tracker "github.com/defipod/mochi/pkg/repo/guild_config_sales_tracker/mocks"
	mock_nft_collection "github.com/defipod/mochi/pkg/repo/nft_collection/mocks"
	mock_nft_sales_tracker "github.com/defipod/mochi/pkg/repo/nft_sales_tracker/mocks"
	"github.com/defipod/mochi/pkg/repo/pg"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/service"
	"github.com/defipod/mochi/pkg/service/abi"
	"github.com/defipod/mochi/pkg/service/indexer"
	"github.com/defipod/mochi/pkg/service/marketplace"
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
	salesTracker := mock_nft_sales_tracker.NewMockStore(ctrl)
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
				addr:     "0xakjsdbajksbfqwueh182937h8123hbj1h2b3",
				platform: "ethereum",
				guildID:  "863278424433229854",
			},
			wantErr: false,
		},
		{
			name: "invalid guild id",
			fields: fields{
				repo: r,
			},
			args: args{
				addr:     "0xakjsdbajksbfqwueh182937h8123hbj1h2b3",
				platform: "ethereum",
				guildID:  "123",
			},
			wantErr: true,
		},
	}
	correctSalesTracker := model.InsertNFTSalesTracker{
		ContractAddress: "0xakjsdbajksbfqwueh182937h8123hbj1h2b3",
		Platform:        "ethereum",
		SalesConfigID:   "abab53eb-c13f-4f6a-b617-27d38a08e519",
	}
	invalidSalesTracker := model.InsertNFTSalesTracker{
		ContractAddress: "0xakjsdbajksbfqwueh182937h8123hbj1h2b3",
		Platform:        "ethereum",
		SalesConfigID:   "abc",
	}
	config := model.GuildConfigSalesTracker{
		ID:        util.GetNullUUID("abab53eb-c13f-4f6a-b617-27d38a08e519"),
		GuildID:   "863278424433229854",
		ChannelID: "863278424433229854",
	}

	configSalesTracker.EXPECT().GetByGuildID("863278424433229854").Return(&config, nil).AnyTimes()
	configSalesTracker.EXPECT().GetByGuildID("123").Return(nil, errors.New("error")).AnyTimes()
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

func TestEntity_GetAllNFTSalesTracker(t *testing.T) {
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
	salesTracker := mock_nft_sales_tracker.NewMockStore(ctrl)
	r.NFTSalesTracker = salesTracker

	tests := []struct {
		name    string
		fields  fields
		want    []response.NFTSalesTrackerResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "get all sales tracker successfully",
			fields: fields{
				repo: r,
			},
			want: []response.NFTSalesTrackerResponse{
				{ContractAddress: "0x33910f98642914a3cb0db10f0c0b062aca2ef552", Platform: "ethereum", GuildID: "863278424433229854", ChannelID: "986854719999864863"},
				{ContractAddress: "0x33910f98642914a3cb0db10f0c0b062aca2ef552", Platform: "ethereum", GuildID: "863278424433229854", ChannelID: "986854719999864863"},
			},
			wantErr: false,
		},
	}

	dataFromRepo := []model.NFTSalesTracker{
		{
			ID:              util.GetNullUUID("0097a657-f823-419b-8392-1b1c3010b63a"),
			ContractAddress: "0x33910f98642914a3cb0db10f0c0b062aca2ef552",
			Platform:        "ethereum",
			SalesConfigID:   "c9554707-5372-4853-a243-af922ce0fddc",
			GuildConfigSalesTracker: model.GuildConfigSalesTracker{
				ID:        util.GetNullUUID("c9554707-5372-4853-a243-af922ce0fddc"),
				GuildID:   "863278424433229854",
				ChannelID: "986854719999864863",
			},
		},
		{
			ID:              util.GetNullUUID("0097a657-f823-419b-8392-1b1c3010b63a"),
			ContractAddress: "0x33910f98642914a3cb0db10f0c0b062aca2ef552",
			Platform:        "ethereum",
			SalesConfigID:   "c9554707-5372-4853-a243-af922ce0fddc",
			GuildConfigSalesTracker: model.GuildConfigSalesTracker{
				ID:        util.GetNullUUID("c9554707-5372-4853-a243-af922ce0fddc"),
				GuildID:   "863278424433229854",
				ChannelID: "986854719999864863",
			},
		},
	}
	salesTracker.EXPECT().GetAll().Return(dataFromRepo, nil).AnyTimes()
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
			got, err := e.GetAllNFTSalesTracker()
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.GetAllNFTSalesTracker() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entity.GetAllNFTSalesTracker() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntity_CheckExistNftCollection(t *testing.T) {
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
		address string
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

	nftCollection := mock_nft_collection.NewMockStore(ctrl)
	r.NFTCollection = nftCollection
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "check exist successfully",
			fields: fields{
				repo: r,
			},
			args: args{
				address: "0x7aCeE5D0acC520faB33b3Ea25D4FEEF1FfebDE73",
			},
			want:    true,
			wantErr: false,
		},
	}

	nftCollectionParam := model.NFTCollection{
		Address:   "0x7aCeE5D0acC520faB33b3Ea25D4FEEF1FfebDE73",
		Name:      "Cyber Neko",
		Symbol:    "NEKO",
		ChainID:   "250",
		ERCFormat: "ERC721",
	}

	nftCollection.EXPECT().GetByAddress("0x7aCeE5D0acC520faB33b3Ea25D4FEEF1FfebDE73").Return(&nftCollectionParam, nil).AnyTimes()

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
			got, err := e.CheckExistNftCollection(tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.CheckExistNftCollection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Entity.CheckExistNftCollection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntity_GetNewListedNFTCollection(t *testing.T) {
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
	type args struct {
		interval string
		page     string
		size     string
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

	nftCollection := mock_nft_collection.NewMockStore(ctrl)
	r.NFTCollection = nftCollection
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.NFTNewListedResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "get new listed NFT successfully",
			fields: fields{
				repo: r,
			},
			args: args{
				interval: "7",
				page:     "1",
				size:     "2",
			},
			want: &response.NFTNewListedResponse{
				Pagination: util.Pagination{
					Page:  int64(1),
					Size:  int64(2),
					Total: int64(2),
				},
				Data: []model.NFTCollection{
					{
						ID:         util.GetNullUUID("05b1a563-1499-437f-b1e8-da4e630ab3ad"),
						Address:    "0x7aCeE5D0acC520faB33b3Ea25D4FEEF1FfebDE79",
						Name:       "neko",
						Symbol:     "neko",
						ChainID:    "250",
						ERCFormat:  "721",
						IsVerified: true,
						CreatedAt:  time.Date(2022, 6, 24, 1, 2, 3, 4, time.UTC),
					},
					{
						ID:         util.GetNullUUID("42970b6d-e141-4162-8529-7f961baf9fce"),
						Address:    "0x7aCeE5D0acC520faB33b3Ea25D4FEEF1FfebDE78",
						Name:       "neko",
						Symbol:     "neko",
						ChainID:    "250",
						ERCFormat:  "721",
						IsVerified: true,
						CreatedAt:  time.Date(2022, 6, 22, 1, 2, 3, 4, time.UTC),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "invalid API params",
			fields: fields{
				repo: r,
			},
			args: args{
				interval: "a",
				page:     "b",
				size:     "c",
			},
			want: &response.NFTNewListedResponse{
				Pagination: util.Pagination{
					Page:  int64(0),
					Size:  int64(0),
					Total: int64(0),
				},
				Data: []model.NFTCollection{},
			},
			wantErr: false,
		},
	}
	repoReturn := []model.NFTCollection{
		{
			ID:         util.GetNullUUID("05b1a563-1499-437f-b1e8-da4e630ab3ad"),
			Address:    "0x7aCeE5D0acC520faB33b3Ea25D4FEEF1FfebDE79",
			Name:       "neko",
			Symbol:     "neko",
			ChainID:    "250",
			ERCFormat:  "721",
			IsVerified: true,
			CreatedAt:  time.Date(2022, 6, 24, 1, 2, 3, 4, time.UTC),
		},
		{
			ID:         util.GetNullUUID("42970b6d-e141-4162-8529-7f961baf9fce"),
			Address:    "0x7aCeE5D0acC520faB33b3Ea25D4FEEF1FfebDE78",
			Name:       "neko",
			Symbol:     "neko",
			ChainID:    "250",
			ERCFormat:  "721",
			IsVerified: true,
			CreatedAt:  time.Date(2022, 6, 22, 1, 2, 3, 4, time.UTC),
		},
	}
	emptyReturn := []model.NFTCollection{}
	nftCollection.EXPECT().GetNewListed(7, 1, 2).Return(repoReturn, int64(2), nil).AnyTimes()
	nftCollection.EXPECT().GetNewListed(0, 0, 0).Return(emptyReturn, int64(0), nil).AnyTimes()
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
			got, err := e.GetNewListedNFTCollection(tt.args.interval, tt.args.page, tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.GetNewListedNFTCollection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entity.GetNewListedNFTCollection() = %v, want %v", got, tt.want)
			}
		})
	}
}
