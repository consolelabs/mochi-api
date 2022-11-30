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
	mock_chain "github.com/defipod/mochi/pkg/repo/chain/mocks"
	mock_guild_config_sales_tracker "github.com/defipod/mochi/pkg/repo/guild_config_sales_tracker/mocks"
	mock_nft_collection "github.com/defipod/mochi/pkg/repo/nft_collection/mocks"
	mock_nft_sales_tracker "github.com/defipod/mochi/pkg/repo/nft_sales_tracker/mocks"
	"github.com/defipod/mochi/pkg/repo/pg"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/service"
	"github.com/defipod/mochi/pkg/service/abi"
	mock_abi "github.com/defipod/mochi/pkg/service/abi/mocks"
	mock_discord "github.com/defipod/mochi/pkg/service/discord/mocks"
	"github.com/defipod/mochi/pkg/service/indexer"
	mock_indexer "github.com/defipod/mochi/pkg/service/indexer/mocks"
	"github.com/defipod/mochi/pkg/service/marketplace"
	mock_marketplace "github.com/defipod/mochi/pkg/service/marketplace/mocks"
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
	log := logger.NewLogrusLogger()
	salesTracker := mock_nft_sales_tracker.NewMockStore(ctrl)
	configSalesTracker := mock_guild_config_sales_tracker.NewMockStore(ctrl)
	r.NFTSalesTracker = salesTracker
	r.GuildConfigSalesTracker = configSalesTracker

	address, _ := util.ConvertToChecksumAddr("0x7aCeE5D0acC520222222")
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
				addr:     "0x7aCeE5D0acC520222222",
				platform: "ethereum",
				guildID:  "863278424433229854",
			},
			wantErr: false,
		},
		{
			name: "invalid guild id",
			fields: fields{
				repo: r,
				log:  log,
			},
			args: args{
				addr:     "0x7aCeE5D0acC520222222",
				platform: "ethereum",
				guildID:  "123",
			},
			wantErr: true,
		},
	}
	correctSalesTracker := model.InsertNFTSalesTracker{
		ContractAddress: address,
		Platform:        "ethereum",
		SalesConfigID:   "abab53eb-c13f-4f6a-b617-27d38a08e519",
	}
	invalidSalesTracker := model.InsertNFTSalesTracker{
		ContractAddress: address,
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

// func TestEntity_GetNewListedNFTCollection(t *testing.T) {
// 	type fields struct {
// 		repo        *repo.Repo
// 		store       repo.Store
// 		log         logger.Logger
// 		dcwallet    discordwallet.IDiscordWallet
// 		discord     *discordgo.Session
// 		cache       cache.Cache
// 		svc         *service.Service
// 		cfg         config.Config
// 		indexer     indexer.Service
// 		abi         abi.Service
// 		marketplace marketplace.Service
// 	}
// 	type args struct {
// 		interval string
// 		page     string
// 		size     string
// 	}
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	cfg := config.Config{
// 		DBUser: "postgres",
// 		DBPass: "postgres",
// 		DBHost: "localhost",
// 		DBPort: "5434",
// 		DBName: "mochi_local",

// 		InDiscordWalletMnemonic: "holiday frequent toy bachelor auto use style result recycle crumble glue blouse",
// 		FantomRPC:               "sample",
// 		FantomScan:              "sample",
// 		FantomScanAPIKey:        "sample",

// 		EthereumRPC:        "sample",
// 		EthereumScan:       "sample",
// 		EthereumScanAPIKey: "sample",

// 		BscRPC:        "sample",
// 		BscScan:       "sample",
// 		BscScanAPIKey: "sample",

// 		DiscordToken: "sample",

// 		RedisURL: "redis://localhost:6379/0",
// 	}

// 	s := pg.NewPostgresStore(&cfg)
// 	r := pg.NewRepo(s.DB())

// 	nftCollection := mock_nft_collection.NewMockStore(ctrl)
// 	r.NFTCollection = nftCollection
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    *response.NFTNewListedResponse
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name: "get new listed NFT successfully",
// 			fields: fields{
// 				repo: r,
// 			},
// 			args: args{
// 				interval: "7",
// 				page:     "0",
// 				size:     "2",
// 			},
// 			want: &response.NFTNewListedResponse{
// 				Data: &response.NFTNewListed{
// 					Metadata: util.Pagination{
// 						Page:  int64(0),
// 						Size:  int64(2),
// 						Total: int64(2),
// 					},
// 					Data: []model.NewListedNFTCollection{
// 						{
// 							ID:         util.GetNullUUID("05b1a563-1499-437f-b1e8-da4e630ab3ad"),
// 							Address:    "0x7aCeE5D0acC520faB33b3Ea25D4FEEF1FfebDE79",
// 							Name:       "neko",
// 							Symbol:     "neko",
// 							ChainID:    "250",
// 							Chain:      "Fantom Opera",
// 							ERCFormat:  "721",
// 							IsVerified: true,
// 							CreatedAt:  time.Date(2022, 6, 24, 1, 2, 3, 4, time.UTC),
// 						},
// 						{
// 							ID:         util.GetNullUUID("42970b6d-e141-4162-8529-7f961baf9fce"),
// 							Address:    "0x7aCeE5D0acC520faB33b3Ea25D4FEEF1FfebDE78",
// 							Name:       "neko",
// 							Symbol:     "neko",
// 							ChainID:    "250",
// 							Chain:      "Fantom Opera",
// 							ERCFormat:  "721",
// 							IsVerified: true,
// 							CreatedAt:  time.Date(2022, 6, 22, 1, 2, 3, 4, time.UTC),
// 						},
// 					}},
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "invalid API params",
// 			fields: fields{
// 				repo: r,
// 			},
// 			args: args{
// 				interval: "a",
// 				page:     "b",
// 				size:     "c",
// 			},
// 			want: &response.NFTNewListedResponse{
// 				Data: &response.NFTNewListed{
// 					Metadata: util.Pagination{
// 						Page:  int64(0),
// 						Size:  int64(0),
// 						Total: int64(0),
// 					},
// 					Data: []model.NewListedNFTCollection{},
// 				},
// 			},
// 			wantErr: false,
// 		},
// 	}
// 	repoReturn := []model.NewListedNFTCollection{
// 		{
// 			ID:         util.GetNullUUID("05b1a563-1499-437f-b1e8-da4e630ab3ad"),
// 			Address:    "0x7aCeE5D0acC520faB33b3Ea25D4FEEF1FfebDE79",
// 			Name:       "neko",
// 			Symbol:     "neko",
// 			ChainID:    "250",
// 			Chain:      "Fantom Opera",
// 			ERCFormat:  "721",
// 			IsVerified: true,
// 			CreatedAt:  time.Date(2022, 6, 24, 1, 2, 3, 4, time.UTC),
// 		},
// 		{
// 			ID:         util.GetNullUUID("42970b6d-e141-4162-8529-7f961baf9fce"),
// 			Address:    "0x7aCeE5D0acC520faB33b3Ea25D4FEEF1FfebDE78",
// 			Name:       "neko",
// 			Symbol:     "neko",
// 			ChainID:    "250",
// 			Chain:      "Fantom Opera",
// 			ERCFormat:  "721",
// 			IsVerified: true,
// 			CreatedAt:  time.Date(2022, 6, 22, 1, 2, 3, 4, time.UTC),
// 		},
// 	}
// 	emptyReturn := []model.NewListedNFTCollection{}
// 	nftCollection.EXPECT().GetNewListed(7, 0, 2).Return(repoReturn, int64(2), nil).AnyTimes()
// 	nftCollection.EXPECT().GetNewListed(0, 0, 0).Return(emptyReturn, int64(0), nil).AnyTimes()

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			e := &Entity{
// 				repo:        tt.fields.repo,
// 				store:       tt.fields.store,
// 				log:         tt.fields.log,
// 				dcwallet:    tt.fields.dcwallet,
// 				discord:     tt.fields.discord,
// 				cache:       tt.fields.cache,
// 				svc:         tt.fields.svc,
// 				cfg:         tt.fields.cfg,
// 				indexer:     tt.fields.indexer,
// 				abi:         tt.fields.abi,
// 				marketplace: tt.fields.marketplace,
// 			}
// 			got, err := e.GetNewListedNFTCollection(tt.args.interval, tt.args.page, tt.args.size)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Entity.GetNewListedNFTCollection() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Entity.GetNewListedNFTCollection() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestEntity_GetNFTDetail(t *testing.T) {
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
		symbol  string
		tokenID string
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
	log := logger.NewLogrusLogger()

	nftCollection := mock_nft_collection.NewMockStore(ctrl)
	mockIndexer := mock_indexer.NewMockService(ctrl)

	r.NFTCollection = nftCollection
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.IndexerGetNFTTokenDetailResponseWithSuggestions
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "query nft successfully",
			fields: fields{
				repo:    r,
				indexer: mockIndexer,
			},
			args: args{
				symbol:  "rabby",
				tokenID: "1",
				guildID: "12312123",
			},
			want: &response.IndexerGetNFTTokenDetailResponseWithSuggestions{
				Data: &response.IndexerNFTTokenDetailData{
					TokenID:           "1",
					CollectionAddress: "0x7D1070fdbF0eF8752a9627a79b00221b53F231fA",
					Name:              "Cyber Rabby #1",
					Description:       "The true pioneer Omnichain NFT to be minted on Ethereum and transferred across chains",
					Amount:            "1",
					Image:             "pic.png",
					ImageCDN:          "pic.png",
					ThumbnailCDN:      "thumb.png",
					ImageContentType:  "",
					RarityRank:        0,
					RarityScore:       "",
					RarityTier:        "",
					Attributes:        []response.IndexerNFTTokenAttribute{},
					Rarity:            &response.IndexerNFTTokenRarity{},
					MetadataID:        "",
				},
				Suggestions: []response.CollectionSuggestions{},
			},
			wantErr: false,
		},
		{
			name: "nft not in sync - indexer return isSync=false",
			fields: fields{
				repo:    r,
				indexer: mockIndexer,
				log:     log,
			},
			args: args{
				symbol:  "neko",
				tokenID: "1",
				guildID: "12312123",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "nft not found in db",
			fields: fields{
				repo:    r,
				indexer: mockIndexer,
				log:     log,
			},
			args: args{
				symbol:  "doggo",
				tokenID: "1",
				guildID: "12312123",
			},
			want:    nil,
			wantErr: true,
		},
		// TODO(trkhoi): re-do whole test

		// {
		// 	name: "token not found - symbol in db and indexer but ID not found",
		// 	fields: fields{
		// 		repo:    r,
		// 		indexer: mockIndexer,
		// 		log:     log,
		// 	},
		// 	args: args{
		// 		symbol:  "rabby",
		// 		tokenID: "99999999",
		// 		guildID: "12312123",
		// 	},
		// 	want:    nil,
		// 	wantErr: true,
		// },
	}
	validNFTCollectionRabby := []model.NFTCollection{
		{
			ID:         util.GetNullUUID("0905f61e-aaf5-4e82-82ef-4c5b929915ed"),
			Address:    "0x7D1070fdbF0eF8752a9627a79b00221b53F231fA",
			Name:       "rabby",
			Symbol:     "rabby",
			ChainID:    "250",
			ERCFormat:  "721",
			IsVerified: true,
			CreatedAt:  time.Date(2022, 6, 20, 1, 2, 3, 4, time.UTC),
		},
	}
	validNFTCollectionNeko := []model.NFTCollection{
		{
			ID:         util.GetNullUUID("05b1a563-1499-437f-b1e8-da4e630ab3ad"),
			Address:    "0x7aCeE5D0acC520faB33b3Ea25D4FEEF1FfebDE79",
			Name:       "neko",
			Symbol:     "neko",
			ChainID:    "250",
			ERCFormat:  "721",
			IsVerified: true,
			CreatedAt:  time.Date(2022, 6, 20, 1, 2, 3, 4, time.UTC),
		},
	}
	validIndexerResponse := &response.IndexerGetNFTTokenDetailResponse{
		Data: response.IndexerNFTTokenDetailData{
			TokenID:           "1",
			CollectionAddress: "0x7D1070fdbF0eF8752a9627a79b00221b53F231fA",
			Name:              "Cyber Rabby #1",
			Description:       "The true pioneer Omnichain NFT to be minted on Ethereum and transferred across chains",
			Amount:            "1",
			Image:             "pic.png",
			ImageCDN:          "pic.png",
			ThumbnailCDN:      "thumb.png",
			ImageContentType:  "",
			RarityRank:        0,
			RarityScore:       "",
			RarityTier:        "",
			Attributes:        []response.IndexerNFTTokenAttribute{},
			Rarity:            &response.IndexerNFTTokenRarity{},
			MetadataID:        "",
		},
	}
	// success
	nftCollection.EXPECT().GetBySymbolorName("rabby").Return(validNFTCollectionRabby, nil).AnyTimes()
	mockIndexer.EXPECT().GetNFTDetail("0x7D1070fdbF0eF8752a9627a79b00221b53F231fA", "1").Return(validIndexerResponse, nil).AnyTimes()

	// fail - sync data in progress
	nftCollection.EXPECT().GetBySymbolorName("neko").Return(validNFTCollectionNeko, nil).AnyTimes()
	mockIndexer.EXPECT().GetNFTDetail("0x7aCeE5D0acC520faB33b3Ea25D4FEEF1FfebDE79", "1").Return(nil, errors.New("data not in sync")).AnyTimes()

	// fail - collection has not been added
	nftCollection.EXPECT().GetBySymbolorName("doggo").Return(nil, errors.New("record not found")).AnyTimes()
	nftCollection.EXPECT().GetSuggestionsBySymbolorName("doggo", 2).Return(nil, errors.New("record not found")).AnyTimes()
	// function does not call indexer if record not found in database

	// fail - token not found
	nftCollection.EXPECT().GetBySymbolorName("rabby").Return(validNFTCollectionRabby, nil).AnyTimes()
	mockIndexer.EXPECT().GetNFTDetail("0x7D1070fdbF0eF8752a9627a79b00221b53F231fA", "99999999").Return(nil, errors.New("token not found")).AnyTimes()

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
			got, err := e.GetNFTDetail(tt.args.symbol, tt.args.tokenID, tt.args.guildID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.GetNFTDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entity.GetNFTDetail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntity_CreateNFTCollection(t *testing.T) {
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
		req request.CreateNFTCollectionRequest
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
	log := logger.NewLogrusLogger()

	nftCollection := mock_nft_collection.NewMockStore(ctrl)
	mockIndexer := mock_indexer.NewMockService(ctrl)
	mockAbi := mock_abi.NewMockService(ctrl)
	mockMarketplace := mock_marketplace.NewMockService(ctrl)
	mockDiscord := mock_discord.NewMockService(ctrl)

	svc, _ := service.NewService(cfg, log)
	svc.Discord = mockDiscord
	r.NFTCollection = nftCollection

	tests := []struct {
		name              string
		fields            fields
		args              args
		wantNftCollection *model.NFTCollection
		wantErr           bool
	}{
		// TODO: Add test cases.
		{
			name: "add collection successfully",
			fields: fields{
				repo:        r,
				abi:         mockAbi,
				indexer:     mockIndexer,
				marketplace: mockMarketplace,
				log:         log,
				svc:         svc,
			},
			args: args{
				request.CreateNFTCollectionRequest{
					Address: "0x7D1070fdbF0eF8752a9627a79b00221b53F231fA",
					Chain:   "Fantom",
					ChainID: "ftm",
					Author:  "catngh",
					GuildID: "863278424433229854",
				},
			},
			wantNftCollection: &model.NFTCollection{
				ID:         util.GetNullUUID("0905f61e-aaf5-4e82-82ef-4c5b929915ed"),
				Address:    "0x7D1070fdbF0eF8752a9627a79b00221b53F231fA",
				Symbol:     "rabby",
				Name:       "Cyber Rabby",
				ChainID:    "250",
				ERCFormat:  "ERC721",
				IsVerified: true,
				CreatedAt:  time.Date(2022, 7, 1, 1, 2, 3, 4, time.UTC),
				Image:      "",
				Author:     "catngh",
			},
			wantErr: false,
		},
		{
			name: "duplicated entry",
			fields: fields{
				repo:        r,
				abi:         mockAbi,
				indexer:     mockIndexer,
				marketplace: mockMarketplace,
				log:         log,
				svc:         svc,
			},
			args: args{
				request.CreateNFTCollectionRequest{
					Address: "0x7ACeE5d0ACC520Fab33b3ea25d4fEEf1FfEBdE79",
					Chain:   "Fantom",
					ChainID: "ftm",
					Author:  "catngh",
					GuildID: "863278424433229854",
				},
			},
			wantNftCollection: nil,
			wantErr:           true,
		},
		{
			name: "invalid chain id",
			fields: fields{
				repo:        r,
				abi:         mockAbi,
				indexer:     mockIndexer,
				marketplace: mockMarketplace,
				log:         log,
				svc:         svc,
			},
			args: args{
				request.CreateNFTCollectionRequest{
					Address: "0x23581767a106ae21c074b2276D25e5C3e136a68b",
					Chain:   "Etheabc",
					ChainID: "abc",
					Author:  "catngh",
					GuildID: "863278424433229854",
				},
			},
			wantNftCollection: nil,
			wantErr:           true,
		},
		{
			name: "abi contract not found",
			fields: fields{
				repo:        r,
				abi:         mockAbi,
				indexer:     mockIndexer,
				marketplace: mockMarketplace,
				log:         log,
				svc:         svc,
			},
			args: args{
				request.CreateNFTCollectionRequest{
					Address: "0x23581767a106ae21c074b2276D25e5C3e136a68c",
					Chain:   "Ethereum",
					ChainID: "11111",
					Author:  "catngh",
					GuildID: "863278424433229854",
				},
			},
			wantNftCollection: nil,
			wantErr:           true,
		},
		{
			name: "invalid address",
			fields: fields{
				repo:        r,
				abi:         mockAbi,
				indexer:     mockIndexer,
				marketplace: mockMarketplace,
				log:         log,
				svc:         svc,
			},
			args: args{
				request.CreateNFTCollectionRequest{
					Address: "0xabc",
					Chain:   "Ethereum",
					ChainID: "eth",
					Author:  "catngh",
					GuildID: "863278424433229854",
				},
			},
			wantNftCollection: nil,
			wantErr:           true,
		},
	}

	validCollection := model.NFTCollection{
		Address:    "0x7D1070fdbF0eF8752a9627a79b00221b53F231fA",
		Symbol:     "rabby",
		Name:       "Cyber Rabby",
		ChainID:    "250",
		ERCFormat:  "ERC721",
		IsVerified: true,
		Image:      "",
		Author:     "catngh",
	}
	returnedValidCollection := model.NFTCollection{
		ID:         util.GetNullUUID("0905f61e-aaf5-4e82-82ef-4c5b929915ed"),
		Address:    "0x7D1070fdbF0eF8752a9627a79b00221b53F231fA",
		Symbol:     "rabby",
		Name:       "Cyber Rabby",
		ChainID:    "250",
		ERCFormat:  "ERC721",
		IsVerified: true,
		CreatedAt:  time.Date(2022, 7, 1, 1, 2, 3, 4, time.UTC),
		Image:      "",
		Author:     "catngh",
	}
	nftReturnedByCheckExist := &model.NFTCollection{
		ID:         util.GetNullUUID("05b1a563-1499-437f-b1e8-da4e630ab3ad"),
		Address:    "0x7ACeE5d0ACC520Fab33b3ea25d4fEEf1FfEBdE79",
		Name:       "Cyber Neko",
		Symbol:     "NEKO",
		ChainID:    "250",
		ERCFormat:  "ERC721",
		IsVerified: true,
		CreatedAt:  time.Date(2022, 7, 1, 1, 2, 3, 4, time.UTC),
		Image:      "",
		Author:     "catngh",
	}
	syncedContract := &response.IndexerContract{
		ID:              3,
		LastUpdateTime:  time.Date(2022, 7, 1, 1, 2, 3, 4, time.UTC),
		LastUpdateBlock: 1000,
		CreationBlock:   1,
		CreatedTime:     time.Date(2022, 7, 1, 1, 2, 3, 4, time.UTC),
		Address:         "0x7ACeE5d0ACC520Fab33b3ea25d4fEEf1FfEBdE79",
		ChainID:         250,
		Type:            "ERC721",
		IsProxy:         false,
		LogicAddress:    "",
		Protocol:        "",
		GRPCAddress:     "indexer-grpc:80",
		IsSynced:        true,
	}
	paintswapCollection := &response.PaintswapCollectionResponse{
		Collection: response.PaintswapCollection{
			Image: "",
		},
	}
	// ########## Case 1: SUCCESSFUL
	mockMarketplace.EXPECT().HandleMarketplaceLink("0x7D1070fdbF0eF8752a9627a79b00221b53F231fA", "ftm").Return("0x7D1070fdbF0eF8752a9627a79b00221b53F231fA").AnyTimes()
	//---convert to checksum - tested
	nftCollection.EXPECT().GetByAddress("0x7D1070fdbF0eF8752a9627a79b00221b53F231fA").Return(nil, errors.New("record not found")).AnyTimes() //repo call for checkexist
	//---collection not existed so skip sync check
	//---convert chain to chain id
	mockAbi.EXPECT().GetNameAndSymbol("0x7D1070fdbF0eF8752a9627a79b00221b53F231fA", int64(250)).Return("Cyber Rabby", "rabby", nil).AnyTimes()
	mockIndexer.EXPECT().CreateERC721Contract(indexer.CreateERC721ContractRequest{Address: "0x7D1070fdbF0eF8752a9627a79b00221b53F231fA", ChainID: 250, Name: "Cyber Rabby", Symbol: "rabby"}).Return(nil).AnyTimes()
	mockMarketplace.EXPECT().GetCollectionFromPaintswap("0x7D1070fdbF0eF8752a9627a79b00221b53F231fA").Return(paintswapCollection, nil) //marketplace call for get image
	nftCollection.EXPECT().Create(validCollection).Return(&returnedValidCollection, nil).AnyTimes()
	mockDiscord.EXPECT().NotifyAddNewCollection("863278424433229854", "Cyber Rabby", "rabby", "ftm", "").Return(nil).AnyTimes()
	//####################

	// ########## Case 2: FAIL - duplicated entry
	mockMarketplace.EXPECT().HandleMarketplaceLink("0x7ACeE5d0ACC520Fab33b3ea25d4fEEf1FfEBdE79", "ftm").Return("0x7ACeE5d0ACC520Fab33b3ea25d4fEEf1FfEBdE79").AnyTimes()
	//---convert to checksum - tested
	nftCollection.EXPECT().GetByAddress("0x7ACeE5d0ACC520Fab33b3ea25d4fEEf1FfEBdE79").Return(nftReturnedByCheckExist, nil).AnyTimes() //repo call for checkexist
	mockIndexer.EXPECT().GetNFTContract("0x7ACeE5d0ACC520Fab33b3ea25d4fEEf1FfEBdE79").Return(syncedContract, nil).AnyTimes()          //repo call for check is sync
	// function stops here and return error
	// ####################

	// ########## Case 3: FAIL - invalid chain id
	mockMarketplace.EXPECT().HandleMarketplaceLink("0x23581767a106ae21c074b2276D25e5C3e136a68b", "abc").Return("0x23581767a106ae21c074b2276D25e5C3e136a68b").AnyTimes()
	//---convert to checksum - tested
	nftCollection.EXPECT().GetByAddress("0x23581767a106ae21c074b2276D25e5C3e136a68b").Return(nil, errors.New("record not found")).AnyTimes() //repo call for checkexist
	//---collection not existed so skip sync check
	// failed to convert chain id 'abc' and return error
	// ####################

	// ########## Case 4: FAIL - abi contract not found
	mockMarketplace.EXPECT().HandleMarketplaceLink("0x23581767a106ae21c074b2276D25e5C3e136a68c", "11111").Return("0x23581767a106ae21c074b2276D25e5C3e136a68c").AnyTimes()
	//---convert to checksum - tested
	nftCollection.EXPECT().GetByAddress("0x23581767A106Ae21C074B2276d25e5c3e136a68c").Return(nil, errors.New("record not found")).AnyTimes() //repo call for checkexist
	//---collection not existed so skip sync check
	mockAbi.EXPECT().GetNameAndSymbol("0x23581767A106Ae21C074B2276d25e5c3e136a68c", int64(11111)).Return("", "", errors.New("contract not found")).AnyTimes()
	// failed to find contract and return error
	// ####################

	// ########## Case 5: FAIL - invalid address
	mockMarketplace.EXPECT().HandleMarketplaceLink("0xabc", "eth").Return("0xabc").AnyTimes()
	// failed convert to checksum and return error
	// ####################

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
			gotNftCollection, err := e.CreateEVMNFTCollection(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.CreateNFTCollection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotNftCollection, tt.wantNftCollection) {
				t.Errorf("Entity.CreateNFTCollection() = %v, want %v", gotNftCollection, tt.wantNftCollection)
			}
		})
	}
}
func TestEntity_CheckIsSync(t *testing.T) {
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
		address string
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	log := logger.NewLogrusLogger()
	mockIndexer := mock_indexer.NewMockService(ctrl)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "check is sync successfully",
			fields: fields{
				indexer: mockIndexer,
			},
			args: args{
				address: "0x7D1070fdbF0eF8752a9627a79b00221b53F231fA",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "address invalid",
			fields: fields{
				indexer: mockIndexer,
				log:     log,
			},
			args: args{
				address: "0x7D1070fdbF0eF8752a9627a79b0022abc",
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "valid address, is not synced",
			fields: fields{
				indexer: mockIndexer,
				log:     log,
			},
			args: args{
				address: "0x7ACeE5d0ACC520Fab33b3ea25d4fEEf1FfEBdE79",
			},
			want:    false,
			wantErr: false,
		},
	}
	syncedContract := &response.IndexerContract{
		ID:              3,
		LastUpdateTime:  time.Date(2022, 7, 1, 1, 2, 3, 4, time.UTC),
		LastUpdateBlock: 1000,
		CreationBlock:   1,
		CreatedTime:     time.Date(2022, 7, 1, 1, 2, 3, 4, time.UTC),
		Address:         "0x7D1070fdbF0eF8752a9627a79b00221b53F231fA",
		ChainID:         250,
		Type:            "ERC721",
		IsProxy:         false,
		LogicAddress:    "",
		Protocol:        "",
		GRPCAddress:     "indexer-grpc:80",
		IsSynced:        true,
	}
	notSyncedContract := &response.IndexerContract{
		ID:              3,
		LastUpdateTime:  time.Date(2022, 7, 1, 1, 2, 3, 4, time.UTC),
		LastUpdateBlock: 1000,
		CreationBlock:   1,
		CreatedTime:     time.Date(2022, 7, 1, 1, 2, 3, 4, time.UTC),
		Address:         "0x7ACeE5d0ACC520Fab33b3ea25d4fEEf1FfEBdE79",
		ChainID:         250,
		Type:            "ERC721",
		IsProxy:         false,
		LogicAddress:    "",
		Protocol:        "",
		GRPCAddress:     "indexer-grpc:80",
		IsSynced:        false,
	}
	mockIndexer.EXPECT().GetNFTContract("0x7D1070fdbF0eF8752a9627a79b00221b53F231fA").Return(syncedContract, nil).AnyTimes()
	mockIndexer.EXPECT().GetNFTContract("0x7D1070fdbF0eF8752a9627a79b0022abc").Return(nil, errors.New("invalid address")).AnyTimes()
	mockIndexer.EXPECT().GetNFTContract("0x7ACeE5d0ACC520Fab33b3ea25d4fEEf1FfEBdE79").Return(notSyncedContract, nil).AnyTimes()

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
			got, err := e.CheckIsSync(tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.CheckIsSync() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Entity.CheckIsSync() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestEntity_GetNFTCollectionTickers(t *testing.T) {
// 	type fields struct {
// 		repo        *repo.Repo
// 		store       repo.Store
// 		log         logger.Logger
// 		dcwallet    discordwallet.IDiscordWallet
// 		discord     *discordgo.Session
// 		cache       cache.Cache
// 		svc         *service.Service
// 		cfg         config.Config
// 		indexer     indexer.Service
// 		abi         abi.Service
// 		marketplace marketplace.Service
// 	}
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	cfg := config.Config{
// 		DBUser: "postgres",
// 		DBPass: "postgres",
// 		DBHost: "localhost",
// 		DBPort: "5434",
// 		DBName: "mochi_local",

// 		InDiscordWalletMnemonic: "holiday frequent toy bachelor auto use style result recycle crumble glue blouse",
// 		FantomRPC:               "sample",
// 		FantomScan:              "sample",
// 		FantomScanAPIKey:        "sample",

// 		EthereumRPC:        "sample",
// 		EthereumScan:       "sample",
// 		EthereumScanAPIKey: "sample",

// 		BscRPC:        "sample",
// 		BscScan:       "sample",
// 		BscScanAPIKey: "sample",

// 		DiscordToken: "sample",

// 		RedisURL: "redis://localhost:6379/0",
// 	}

// 	s := pg.NewPostgresStore(&cfg)
// 	r := pg.NewRepo(s.DB())
// 	log := logger.NewLogrusLogger()

// 	nftCollection := mock_nft_collection.NewMockStore(ctrl)
// 	mockIndexer := mock_indexer.NewMockService(ctrl)

// 	r.NFTCollection = nftCollection

// 	type args struct {
// 		symbol   string
// 		rawQuery string
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    *response.IndexerNFTCollectionTickersResponse
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name: "get tickers successfully",
// 			fields: fields{
// 				repo:    r,
// 				indexer: mockIndexer,
// 				log:     log,
// 			},
// 			args: args{
// 				symbol:   "neko",
// 				rawQuery: "from=1658206545000&to=1658292945000",
// 			},
// 			want: &response.IndexerNFTCollectionTickersResponse{
// 				Data: &response.IndexerNFTCollectionTickersData{
// 					Tickers: &response.IndexerTickers{},
// 					FloorPrice: &response.IndexerPrice{
// 						Amount: "10",
// 					},
// 					Name:         "Neko",
// 					Address:      "0x23581767a106ae21c074b2276D25e5C3e136a68h",
// 					Chain:        &response.IndexerChain{Name: "eth"},
// 					Marketplaces: []string{"abc"},
// 					TotalVolume: &response.IndexerPrice{
// 						Amount: "10",
// 					},
// 					Items:           23,
// 					Owners:          100,
// 					CollectionImage: "image.png",
// 				},
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "failed to query repo - invalid symbol",
// 			fields: fields{
// 				repo:    r,
// 				indexer: mockIndexer,
// 				log:     log,
// 			},
// 			args: args{
// 				symbol:   "abc",
// 				rawQuery: "from=1658206545000&to=1658292945000",
// 			},
// 			want:    nil,
// 			wantErr: true,
// 		},
// 		{
// 			name: "failed to query indexer - invalid raw query",
// 			fields: fields{
// 				repo:    r,
// 				indexer: mockIndexer,
// 				log:     log,
// 			},
// 			args: args{
// 				symbol:   "neko",
// 				rawQuery: "from=1658206545&to=1658292945",
// 			},
// 			want:    nil,
// 			wantErr: true,
// 		},
// 	}
// 	repoCollection := &model.NFTCollection{
// 		Address: "0x23581767a106ae21c074b2276D25e5C3e136a68h",
// 	}
// 	indexerTicker := &response.IndexerNFTCollectionTickersResponse{
// 		Data: &response.IndexerNFTCollectionTickersData{
// 			Tickers: &response.IndexerTickers{},
// 			FloorPrice: &response.IndexerPrice{
// 				Amount: "10",
// 			},
// 			Name:         "Neko",
// 			Address:      "0x23581767a106ae21c074b2276D25e5C3e136a68h",
// 			Chain:        &response.IndexerChain{Name: "eth"},
// 			Marketplaces: []string{"abc"},
// 			TotalVolume: &response.IndexerPrice{
// 				Amount: "10",
// 			},
// 			Items:           23,
// 			Owners:          100,
// 			CollectionImage: "image.png",
// 		},
// 	}

// 	// case success
// 	nftCollection.EXPECT().GetBySymbol("neko").Return(repoCollection, nil).AnyTimes()
// 	mockIndexer.EXPECT().GetNFTCollectionTickers("0x23581767a106ae21c074b2276D25e5C3e136a68h", "from=1658206545000&to=1658292945000").Return(indexerTicker, nil).AnyTimes()

// 	// case fail repo
// 	nftCollection.EXPECT().GetBySymbol("abc").Return(nil, errors.New("invalid symbol")).AnyTimes()

// 	// case fail indexer
// 	mockIndexer.EXPECT().GetNFTCollectionTickers("0x23581767a106ae21c074b2276D25e5C3e136a68h", "from=1658206545&to=1658292945").Return(nil, errors.New("invalid query")).AnyTimes()

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			e := &Entity{
// 				repo:        tt.fields.repo,
// 				store:       tt.fields.store,
// 				log:         tt.fields.log,
// 				dcwallet:    tt.fields.dcwallet,
// 				discord:     tt.fields.discord,
// 				cache:       tt.fields.cache,
// 				svc:         tt.fields.svc,
// 				cfg:         tt.fields.cfg,
// 				indexer:     tt.fields.indexer,
// 				abi:         tt.fields.abi,
// 				marketplace: tt.fields.marketplace,
// 			}
// 			got, err := e.GetNFTCollectionTickers(tt.args.symbol, tt.args.rawQuery)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Entity.GetNFTCollectionTickers() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Entity.GetNFTCollectionTickers() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestEntity_GetNFTCollections(t *testing.T) {
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
		p string
		s string
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
	log := logger.NewLogrusLogger()

	nftCollection := mock_nft_collection.NewMockStore(ctrl)
	r.NFTCollection = nftCollection

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.NFTCollectionsResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "successful",
			fields: fields{
				repo: r,
				log:  log,
			},
			args: args{
				p: "0",
				s: "2",
			},
			want: &response.NFTCollectionsResponse{
				Data: response.NFTCollectionsData{
					Metadata: util.Pagination{
						Page:  int64(0),
						Size:  int64(2),
						Total: int64(2),
					},
					Data: []model.NFTCollection{
						{
							Address:   "0x7aCeE5D0acC520faB33b3Ea25D4FEEF1FfebDE73",
							Name:      "Cyber Neko",
							Symbol:    "NEKO",
							ChainID:   "250",
							ERCFormat: "ERC721",
						},
						{
							Address:   "0x7aCeE5D0acC520faB33b3Ea25D4FEEF1FfebDE73",
							Name:      "Cyber Neko",
							Symbol:    "NEKO",
							ChainID:   "250",
							ERCFormat: "ERC721",
						},
					},
				},
			},
			wantErr: false,
		},
	}
	repoNFTList := []model.NFTCollection{
		{
			Address:   "0x7aCeE5D0acC520faB33b3Ea25D4FEEF1FfebDE73",
			Name:      "Cyber Neko",
			Symbol:    "NEKO",
			ChainID:   "250",
			ERCFormat: "ERC721",
		},
		{
			Address:   "0x7aCeE5D0acC520faB33b3Ea25D4FEEF1FfebDE73",
			Name:      "Cyber Neko",
			Symbol:    "NEKO",
			ChainID:   "250",
			ERCFormat: "ERC721",
		},
	}
	nftCollection.EXPECT().ListAllWithPaging(0, 2).Return(repoNFTList, int64(2), nil).AnyTimes()
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
			got, err := e.GetNFTCollections(tt.args.p, tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.GetNFTCollections() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entity.GetNFTCollections() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntity_GetCollectionCount(t *testing.T) {
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
	log := logger.NewLogrusLogger()

	nftCollection := mock_nft_collection.NewMockStore(ctrl)
	chainMock := mock_chain.NewMockStore(ctrl)
	r.NFTCollection = nftCollection
	r.Chain = chainMock
	tests := []struct {
		name    string
		fields  fields
		want    *response.NFTCollectionCount
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "successful",
			fields: fields{
				repo: r,
				log:  log,
			},
			want: &response.NFTCollectionCount{
				Total: 10,
				Data: []response.NFTChainCollectionCount{
					{
						Chain: model.Chain{
							ID: 1,
						},
						Count: 4,
					},
					{
						Chain: model.Chain{
							ID: 250,
						},
						Count: 3,
					},
					{
						Chain: model.Chain{
							ID: 10,
						},
						Count: 3,
					},
				},
			},
			wantErr: false,
		},
	}
	chains := []model.Chain{
		{
			ID: 1,
		},
		{
			ID: 250,
		},
		{
			ID: 10,
		},
	}
	chainMock.EXPECT().GetAll().Return(chains, nil).AnyTimes()
	nftCollection.EXPECT().GetByChain(1).Return(nil, 4, nil).AnyTimes()
	nftCollection.EXPECT().GetByChain(250).Return(nil, 3, nil).AnyTimes()
	nftCollection.EXPECT().GetByChain(10).Return(nil, 3, nil).AnyTimes()

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
			got, err := e.GetCollectionCount()
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.GetCollectionCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entity.GetCollectionCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
