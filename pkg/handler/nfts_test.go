package handler

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/repo/pg"
	"github.com/defipod/mochi/pkg/response"
	mock_indexer "github.com/defipod/mochi/pkg/service/indexer/mocks"
	"github.com/defipod/mochi/pkg/util/testhelper"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestHandler_GetNewListedNFTCollection(t *testing.T) {
	cfg := config.LoadTestConfig()
	db := testhelper.LoadTestDB("../../migrations/test_seed")
	repo := pg.NewRepo(db)
	log := logger.NewLogrusLogger()
	s := pg.NewPostgresStore(&cfg)
	entityMock := entities.New(cfg, log, repo, s, nil, nil, nil, nil, nil, nil, nil)

	type fields struct {
		entities *entities.Entity
		log      logger.Logger
	}
	tests := []struct {
		name             string
		fields           fields
		wantCode         int
		wantErr          error
		wantResponsePath string
	}{
		// TODO: Add test cases.
		{
			name: "get succesfully",
			fields: fields{
				entities: entityMock,
				log:      log,
			},
			wantCode:         200,
			wantErr:          nil,
			wantResponsePath: "testdata/get_nft_recent/200.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				entities: tt.fields.entities,
				log:      tt.fields.log,
			}
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest("GET", "/api/v1/nfts/new-listed", nil)

			h.GetNewListedNFTCollection(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)
			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.GetChains] response mismatched")
		})
	}
}

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
					MetadataID: "",
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
