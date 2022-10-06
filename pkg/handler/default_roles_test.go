package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/repo/pg"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/util/testhelper"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestHandler_GetDefaultRolesByGuildID(t *testing.T) {
	db := testhelper.LoadTestDB("../../migrations/test_seed")
	repo := pg.NewRepo(db)
	cfg := config.LoadTestConfig()
	log := logger.NewLogrusLogger()
	entity := entities.New(cfg, log, repo, nil, nil, nil, nil, nil, nil, nil, nil)

	h := &Handler{
		entities: entity,
		log:      log,
	}

	tests := []struct {
		name             string
		query            string
		wantCode         int
		wantResponsePath string
	}{
		{
			name:             "guild_has_been_configured",
			query:            "guild_id=895659000996200508",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/response/default_role/200_guild_has_been_configured.json",
		},
		{
			name:             "guild_is_not_configured",
			query:            "guild_id=863278424433229854",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/response/default_role/200_guild_is_not_configured.json",
		},
		{
			name:             "empty_query",
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/400-missing-guildID.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/configs/default-roles?%s", tt.query), nil)

			h.GetDefaultRolesByGuildID(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.GetDefaultRolesByGuildID] response mismatched")
		})
	}
}

func TestHandler_CreateDefaultRole(t *testing.T) {
	db := testhelper.LoadTestDB("../../migrations/test_seed")
	repo := pg.NewRepo(db)
	cfg := config.LoadTestConfig()
	log := logger.NewLogrusLogger()
	entity := entities.New(cfg, log, repo, nil, nil, nil, nil, nil, nil, nil, nil)

	h := &Handler{
		entities: entity,
		log:      log,
	}

	tests := []struct {
		name             string
		args             request.CreateDefaultRoleRequest
		wantCode         int
		wantResponsePath string
	}{
		{
			name: "update_new_role",
			args: request.CreateDefaultRoleRequest{
				GuildID: "895659000996200508",
				RoleID:  "1012576984783671123",
			},
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/response/default_role/200_update_new_role.json",
		},
		{
			name: "create_new_config_for_new_guild",
			args: request.CreateDefaultRoleRequest{
				GuildID: "863278424433229854",
				RoleID:  "1011659315729403958",
			},
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/response/default_role/200_create_new_config_for_new_guild.json",
		},
		{
			name:             "empty_request",
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/response/default_role/400_empty_request.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, err := json.Marshal(tt.args)
			if err != nil {
				t.Error(err)
				return
			}

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodPost, "/api/v1/configs/default-roles", bytes.NewBuffer(body))

			h.CreateDefaultRole(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.CreateDefaultRole] response mismatched")
		})
	}
}

func TestHandler_DeleteDefaultRoleByGuildID(t *testing.T) {
	db := testhelper.LoadTestDB("../../migrations/test_seed")
	repo := pg.NewRepo(db)
	cfg := config.LoadTestConfig()
	log := logger.NewLogrusLogger()
	entity := entities.New(cfg, log, repo, nil, nil, nil, nil, nil, nil, nil, nil)

	h := &Handler{
		entities: entity,
		log:      log,
	}

	tests := []struct {
		name             string
		query            string
		wantCode         int
		wantResponsePath string
	}{
		{
			name:             "guild_has_been_configured",
			query:            "guild_id=895659000996200508",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/200-message-ok.json",
		},
		{
			name:             "guild_is_not_configured",
			query:            "guild_id=863278424433229854",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/200-message-ok.json",
		},
		{
			name:             "empty_query",
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/400-missing-guildID.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/configs/default-roles?%s", tt.query), nil)

			h.DeleteDefaultRoleByGuildID(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.DeleteDefaultRoleByGuildID] response mismatched")
		})
	}
}
