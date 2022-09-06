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
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/util"
	"github.com/defipod/mochi/pkg/util/testhelper"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestHandler_GetGuildTokens(t *testing.T) {
	cfg := config.LoadTestConfig()
	db := testhelper.LoadTestDB("../../migrations/test_seed")
	repo := pg.NewRepo(db)
	log := logger.NewLogrusLogger()
	s := pg.NewPostgresStore(&cfg)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	entityMock := entities.New(cfg, log, repo, s, nil, nil, nil, nil, nil, nil, nil)

	tests := []struct {
		name             string
		guildId          string
		wantError        error
		wantCode         int
		wantResponsePath string
	}{
		// TODO: Add test cases.
		{
			name:             "success - with guild id",
			guildId:          "863278424433229854",
			wantCode:         200,
			wantResponsePath: "testdata/guild_config_tokens/200-with-id.json",
		},
		{
			name:             "success - without guild id",
			guildId:          "",
			wantCode:         200,
			wantResponsePath: "testdata/guild_config_tokens/200-without-id.json",
		},
		{
			name:             "success - with wrong guild id",
			guildId:          "123123123",
			wantCode:         200,
			wantResponsePath: "testdata/guild_config_tokens/200-without-id.json",
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
			ctx.Request = httptest.NewRequest("GET", fmt.Sprintf("/api/v1/configs/tokens?guild_id=%s", tt.guildId), nil)

			h.GetGuildTokens(ctx)

			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)
			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.GetGuildTokens] response mismatched")
		})
	}
}

func TestHandler_UpsertGuildTokenConfig(t *testing.T) {
	cfg := config.LoadTestConfig()
	db := testhelper.LoadTestDB("../../migrations/test_seed")
	repo := pg.NewRepo(db)
	log := logger.NewLogrusLogger()
	s := pg.NewPostgresStore(&cfg)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	//guildConfigToken := mock_guild_config_token.NewMockStore(ctrl)
	//repo.GuildConfigToken = guildConfigToken
	entityMock := entities.New(cfg, log, repo, s, nil, nil, nil, nil, nil, nil, nil)
	tests := []struct {
		name             string
		req              request.UpsertGuildTokenConfigRequest
		wantError        error
		wantCode         int
		wantResponsePath string
	}{
		// TODO: Add test cases.
		{
			name: "success",
			req: request.UpsertGuildTokenConfigRequest{
				GuildID: "863278424433229854",
				Symbol:  "TOMB",
				Active:  true,
			},
			wantCode:         200,
			wantResponsePath: "testdata/guild_config_tokens/200-create-ok.json",
		},
		{
			name: "fail - invalid request",
			req: request.UpsertGuildTokenConfigRequest{
				GuildID: "",
				Symbol:  "TOMB",
				Active:  true,
			},
			wantCode:         400,
			wantResponsePath: "testdata/guild_config_tokens/400-create-invalid-req.json",
		},
		{
			name: "fail - cannot find symbol",
			req: request.UpsertGuildTokenConfigRequest{
				GuildID: "863278424433229854",
				Symbol:  "TOMBA",
				Active:  true,
			},
			wantCode:         500,
			wantResponsePath: "testdata/guild_config_tokens/500-create-symbol-not-found.json",
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
			ctx.Request = httptest.NewRequest("POST", "/api/v1/configs/tokens", nil)
			util.SetRequestBody(ctx, tt.req)
			// guildConfigToken.EXPECT().UpsertMany([]model.GuildConfigToken{{GuildID: tt.req.GuildID, TokenID: tt.req.Symbol}}).Return()

			h.UpsertGuildTokenConfig(ctx)

			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)
			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.GetGuildTokens] response mismatched")

		})
	}
}
