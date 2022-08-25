package handler

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/repo/pg"
	"github.com/stretchr/testify/require"

	"github.com/defipod/mochi/pkg/util/testhelper"
	"github.com/gin-gonic/gin"
)

func TestHandler_GetNewListedNFTCollection(t *testing.T) {
	cfg := config.Config{
		DBUser: "postgres",
		DBPass: "postgres",
		DBHost: "localhost",
		DBPort: "25434",
		DBName: "mochi_local_test",

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
	db := testhelper.LoadTestDB()
	repo := pg.NewRepo(db)
	log := logger.NewLogrusLogger()
	s := pg.NewPostgresStore(&cfg)
	entityMock := entities.New(cfg, log, repo, s, nil, nil, nil, nil, nil, nil, nil)

	wr := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(wr)
	type fields struct {
		entities *entities.Entity
		log      logger.Logger
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
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
			args: args{
				c: context,
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
			h.GetNewListedNFTCollection(tt.args.c)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.GetChains] response mismatched")
		})
	}
}
