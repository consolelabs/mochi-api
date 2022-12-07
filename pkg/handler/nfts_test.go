package handler

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	mock_nftcollection "github.com/defipod/mochi/pkg/repo/nft_collection/mocks"
	"github.com/defipod/mochi/pkg/repo/pg"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/service"
	mock_abi "github.com/defipod/mochi/pkg/service/abi/mocks"
	mock_discord "github.com/defipod/mochi/pkg/service/discord/mocks"
	"github.com/defipod/mochi/pkg/service/indexer"
	mock_indexer "github.com/defipod/mochi/pkg/service/indexer/mocks"
	mock_marketplace "github.com/defipod/mochi/pkg/service/marketplace/mocks"
	"github.com/defipod/mochi/pkg/util"
	"github.com/defipod/mochi/pkg/util/testhelper"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

// func TestHandler_GetNewListedNFTCollection(t *testing.T) {
// 	cfg := config.LoadTestConfig()
// 	db := testhelper.LoadTestDB("../../migrations/test_seed")
// 	repo := pg.NewRepo(db)
// 	log := logger.NewLogrusLogger()
// 	s := pg.NewPostgresStore(&cfg)
// 	entityMock := entities.New(cfg, log, repo, s, nil, nil, nil, nil, nil, nil, nil)

// 	type fields struct {
// 		entities *entities.Entity
// 		log      logger.Logger
// 	}
// 	tests := []struct {
// 		name             string
// 		fields           fields
// 		wantCode         int
// 		wantErr          error
// 		wantResponsePath string
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name: "get succesfully",
// 			fields: fields{
// 				entities: entityMock,
// 				log:      log,
// 			},
// 			wantCode:         200,
// 			wantErr:          nil,
// 			wantResponsePath: "testdata/get_nft_recent/200.json",
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			h := &Handler{
// 				entities: tt.fields.entities,
// 				log:      tt.fields.log,
// 			}
// 			w := httptest.NewRecorder()
// 			ctx, _ := gin.CreateTestContext(w)
// 			ctx.Request = httptest.NewRequest("GET", "/api/v1/nfts/new-listed", nil)

// 			h.GetNewListedNFTCollection(ctx)
// 			require.Equal(t, tt.wantCode, w.Code)
// 			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
// 			require.NoError(t, err)
// 			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.GetChains] response mismatched")
// 		})
// 	}
// }

func TestHandler_GetNFTDetail(t *testing.T) {
	cfg := config.LoadTestConfig()
	db := testhelper.LoadTestDB("../../migrations/test_seed")
	repo := pg.NewRepo(db)
	log := logger.NewLogrusLogger()
	s := pg.NewPostgresStore(&cfg)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	indexerMock := mock_indexer.NewMockService(ctrl)
	entityMock := entities.New(cfg, log, repo, s, nil, nil, nil, nil, indexerMock, nil, nil)

	tests := []struct {
		name             string
		querySymbol      string
		queryTokenId     string
		queryGuildId     string
		expectedAddress  string
		wantIndexerResp  *response.IndexerGetNFTTokenDetailResponse
		wantCode         int
		wantErr          error
		wantResponsePath string
	}{
		// TODO: Add test cases.
		{
			name:            "query match single record",
			querySymbol:     "PH",
			queryTokenId:    "1",
			queryGuildId:    "",
			expectedAddress: "0xb54FF1EBc9950fce19Ee9E055A382B1219f862f0",
			wantIndexerResp: &response.IndexerGetNFTTokenDetailResponse{
				Data: response.IndexerNFTTokenDetailData{
					TokenID:           "1",
					CollectionAddress: "0xb54FF1EBc9950fce19Ee9E055A382B1219f862f0",
					Name:              "Portalhead #1",
					Description:       "A Modern People Project.",
					Amount:            "",
					Image:             "https://portalheads.mypinata.cloud/ipfs/QmNip3GfiuE2rFyMrSfPaafgLyLNUiMSmCti1N8LAF5Re9/0x12791c9a355a9097e9ef4b894cbf46579ca259d281a27d0f6e6b27a00538607c.jpg",
					ImageCDN:          "",
					ThumbnailCDN:      "",
					ImageContentType:  "",
					RarityRank:        0,
					RarityTier:        "",
					Attributes: []response.IndexerNFTTokenAttribute{
						{
							CollectionAddress: "",
							TokenId:           "",
							TraitType:         "eyes",
							Value:             "White POV",
							Count:             0,
							Rarity:            "",
							Frequency:         "5.220%",
						},
						{
							CollectionAddress: "",
							TokenId:           "",
							TraitType:         "face",
							Value:             "None",
							Count:             0,
							Rarity:            "",
							Frequency:         "60.140%",
						},
						{
							CollectionAddress: "",
							TokenId:           "",
							TraitType:         "head",
							Value:             "Woke Up Like Dis",
							Count:             0,
							Rarity:            "",
							Frequency:         "2.750%",
						},
					},
					Rarity: &response.IndexerNFTTokenRarity{
						Rank:  0,
						Score: "",
						Total: 6771,
					},
					Owner: response.IndexerNftTokenOwner{
						OwnerAddress:      "0x5417A03667AbB6A059b3F174c1F67b1E83753046",
						CollectionAddress: "",
						TokenId:           "1",
					},
					Marketplace: []response.NftListingMarketplace{},
					MetadataID:  "",
				},
			},
			wantCode:         200,
			wantErr:          nil,
			wantResponsePath: "testdata/get_nft_detail/200-match.json",
		},
		{
			name:             "query no record found",
			querySymbol:      "qweasd",
			queryTokenId:     "1",
			queryGuildId:     "",
			wantCode:         200,
			wantErr:          nil,
			wantResponsePath: "testdata/get_nft_detail/200-no-data.json",
		},
		{
			name:             "query match multiple record - no default",
			querySymbol:      "NEKO",
			queryTokenId:     "1",
			queryGuildId:     "",
			wantCode:         200,
			wantErr:          nil,
			wantResponsePath: "testdata/get_nft_detail/200-suggest.json",
		},
		{
			name:             "query match multiple record - with default",
			querySymbol:      "NEKO",
			queryTokenId:     "1",
			queryGuildId:     "863278424433229854",
			expectedAddress:  "",
			wantCode:         200,
			wantErr:          nil,
			wantResponsePath: "testdata/get_nft_detail/200-default.json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				entities: entityMock,
				log:      log,
			}
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest("GET", fmt.Sprintf("/api/v1/nfts/%s/%s?guild_id=%s", tt.querySymbol, tt.queryTokenId, tt.queryGuildId), nil)
			ctx.Params = []gin.Param{
				{
					Key:   "symbol",
					Value: tt.querySymbol,
				},
				{
					Key:   "id",
					Value: tt.queryTokenId,
				},
			}
			indexerMock.EXPECT().GetNFTDetail(tt.expectedAddress, tt.queryTokenId).Return(tt.wantIndexerResp, nil).AnyTimes()

			h.GetNFTDetail(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)
			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.GetNFTDetail] response mismatched")
		})
	}
}

func TestHandler_CreateNFTCollection(t *testing.T) {
	cfg := config.LoadTestConfig()
	db := testhelper.LoadTestDB("../../migrations/test_seed")
	log := logger.NewLogrusLogger()
	s := pg.NewPostgresStore(&cfg)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	NFTCollection := mock_nftcollection.NewMockStore(ctrl)
	repo := pg.NewRepo(db)
	svc, _ := service.NewService(cfg, log)
	repo.NFTCollection = NFTCollection

	indexerMock := mock_indexer.NewMockService(ctrl)
	marketplaceMock := mock_marketplace.NewMockService(ctrl)
	abiMock := mock_abi.NewMockService(ctrl)
	discordMock := mock_discord.NewMockService(ctrl)
	svc.Discord = discordMock
	entityMock := entities.New(cfg, log, repo, s, nil, nil, nil, svc, indexerMock, abiMock, marketplaceMock)

	marketplaceDataRabby := response.OpenseaAssetContractResponse{
		Address: "0x7D1070fdbF0eF8752a9627a79b00221b53F231fA",
		Collection: response.OpenseaAssetContract{
			Image:   "https://openseauserdata.com/files/061eb8949cff84d0be850fc9a566e4fe.png",
			UrlName: "cyber-rabby",
		},
	}
	tests := []struct {
		name                     string
		req                      request.CreateNFTCollectionRequest
		address                  string
		collectionExist          bool
		wantMarketplaceOpensea   *response.OpenseaAssetContractResponse
		wantMarketplacePaintswap *response.PaintswapCollectionResponse
		wantMarketplaceOptimisim *response.QuixoticCollectionResponse
		wantIndexerContract      *response.IndexerContract
		wantName                 string
		wantSymbol               string
		wantImage                string
		wantCode                 int
		wantError                error
		wantResponsePath         string
	}{
		{
			name: "create new collection successful - with address",
			req: request.CreateNFTCollectionRequest{
				Address: "0x7D1070fdbF0eF8752a9627a79b00221b53F231fA",
				ChainID: "eth",
				Chain:   "1",
				Author:  "319132138849173505",
				GuildID: "863278424433229854",
			},
			collectionExist:        false,
			address:                "0x7D1070fdbF0eF8752a9627a79b00221b53F231fA",
			wantMarketplaceOpensea: &marketplaceDataRabby,
			wantName:               "Cyber Rabby",
			wantSymbol:             "RABBY",
			wantImage:              "",
			wantCode:               200,
			wantError:              nil,
			wantResponsePath:       "testdata/create_nft_collection/200-success.json",
		},
		{
			name: "create new collection successful - with market link",
			req: request.CreateNFTCollectionRequest{
				Address: "https://opensea.io/collection/cyber-rabby",
				ChainID: "eth",
				Chain:   "1",
				Author:  "319132138849173505",
				GuildID: "863278424433229854",
			},
			collectionExist:        false,
			address:                "0x7D1070fdbF0eF8752a9627a79b00221b53F231fA",
			wantMarketplaceOpensea: &marketplaceDataRabby,
			wantName:               "Cyber Rabby",
			wantSymbol:             "RABBY",
			wantImage:              "https://openseauserdata.com/files/061eb8949cff84d0be850fc9a566e4fe.png",
			wantCode:               200,
			wantError:              nil,
			wantResponsePath:       "testdata/create_nft_collection/200-success.json",
		},
		{
			name: "fail to create - nft existed - not synced",
			req: request.CreateNFTCollectionRequest{
				Address: "0x7aCeE5D0acC520faB33b3Ea25D4FEEF1FfebDE73",
				ChainID: "eth",
				Chain:   "1",
				Author:  "319132138849173505",
				GuildID: "863278424433229854",
			},
			collectionExist: true,
			address:         "0x7aCeE5D0acC520faB33b3Ea25D4FEEF1FfebDE73",
			wantIndexerContract: &response.IndexerContract{
				IsSynced: false,
			},
			wantCode:         400,
			wantError:        nil,
			wantResponsePath: "testdata/create_nft_collection/400-not-sync.json",
		},
		{
			name: "fail to create - nft existed - already synced",
			req: request.CreateNFTCollectionRequest{
				Address: "0x09E0dF4aE51111CA27d6B85708CFB3f1F7cAE982",
				ChainID: "eth",
				Chain:   "1",
				Author:  "319132138849173505",
				GuildID: "863278424433229854",
			},
			collectionExist: true,
			address:         "0x09E0dF4aE51111CA27d6B85708CFB3f1F7cAE982",
			wantIndexerContract: &response.IndexerContract{
				IsSynced: true,
			},
			wantCode:         400,
			wantError:        nil,
			wantResponsePath: "testdata/create_nft_collection/400-insync.json",
		},
		// TODO(trkhoi): turn on this test once get new opensea api key
		// {
		// 	name: "fail to get contract",
		// 	req: request.CreateNFTCollectionRequest{
		// 		Address: "0x09e0dF4ae51111Ca27d6B85708Cfb3F1f7cAe983",
		// 		ChainID: "eth",
		// 		Chain:   "1",
		// 		Author:  "319132138849173505",
		// 		GuildID: "863278424433229854",
		// 	},
		// 	collectionExist:  false,
		// 	address:          "0x09e0dF4ae51111Ca27d6B85708Cfb3F1f7cAe983",
		// 	wantCode:         500,
		// 	wantError:        errors.New("GetNFTCollections - failed to get opensea asset contract with address=0x09e0dF4ae51111Ca27d6B85708Cfb3F1f7cAe983: {success:false}"),
		// 	wantResponsePath: "testdata/create_nft_collection/500-cannot-find-contract.json",
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chainId, _ := strconv.Atoi(tt.req.Chain)
			h := &Handler{
				entities: entityMock,
				log:      log,
			}
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest("POST", "/api/v1/nfts/collections", nil)
			util.SetRequestBody(ctx, tt.req)

			// marketplaceMock.EXPECT().GetOpenseaAssetContract(tt.address).Return(tt.wantMarketplaceOpensea, tt.wantError).AnyTimes()
			marketplaceMock.EXPECT().GetCollectionFromPaintswap(tt.address).Return(tt.wantMarketplacePaintswap, tt.wantError).AnyTimes()
			marketplaceMock.EXPECT().GetCollectionFromQuixotic(tt.address).Return(tt.wantMarketplaceOptimisim, tt.wantError).AnyTimes()
			marketplaceMock.EXPECT().HandleMarketplaceLink(tt.req.Address, tt.req.ChainID).Return(tt.address).AnyTimes()

			abiMock.EXPECT().GetNameAndSymbol(tt.address, int64(chainId)).Return(tt.wantName, tt.wantSymbol, nil).AnyTimes()

			indexerMock.EXPECT().GetNFTContract(tt.address).Return(tt.wantIndexerContract, nil).AnyTimes()
			indexerMock.EXPECT().CreateERC721Contract(indexer.CreateERC721ContractRequest{
				Address: tt.address,
				ChainID: chainId,
				Name:    "Cyber Rabby",
				Symbol:  "RABBY",
				GuildID: tt.req.GuildID,
			}).Return(nil).AnyTimes()

			discordMock.EXPECT().NotifyAddNewCollection(tt.req.GuildID, tt.wantName, tt.wantSymbol, tt.req.ChainID, tt.wantImage).AnyTimes()

			if tt.collectionExist {
				NFTCollection.EXPECT().GetByAddress(tt.address).Return(nil, nil).AnyTimes()
			} else {
				NFTCollection.EXPECT().GetByAddress(tt.address).Return(nil, errors.New("record not found")).AnyTimes()
			}
			NFTCollection.EXPECT().Create(model.NFTCollection{
				Address:    tt.address,
				Symbol:     tt.wantSymbol,
				Name:       tt.wantName,
				ChainID:    tt.req.Chain,
				ERCFormat:  "ERC721",
				IsVerified: true,
				Author:     tt.req.Author,
				Image:      tt.wantImage,
			}).Return(&model.NFTCollection{
				ID:         util.GetNullUUID("8aa72c1b-5dcd-467f-9486-e9826d9a18e0"),
				Address:    tt.address,
				Symbol:     tt.wantSymbol,
				Name:       tt.wantName,
				ChainID:    tt.req.Chain,
				ERCFormat:  "ERC721",
				IsVerified: true,
				CreatedAt:  time.Date(2022, 8, 29, 5, 4, 3, 2, time.UTC),
				Author:     tt.req.Author,
				Image:      tt.wantImage,
			}, nil).AnyTimes()

			h.CreateNFTCollection(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)
			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.CreateNFTCollection] response mismatched")
		})
	}
}
