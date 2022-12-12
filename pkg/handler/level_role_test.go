package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/repo/pg"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/util/testhelper"
)

func TestHandler_ConfigLevelRole(t *testing.T) {
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
		args             request.ConfigLevelRoleRequest
		wantCode         int
		wantResponsePath string
	}{
		{
			name: "create_new_config",
			args: request.ConfigLevelRoleRequest{
				GuildID: "895659000996200508",
				RoleID:  "1003867749841383425",
				Level:   5,
			},
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/200-message-ok.json",
		},
		{
			name: "update_new_level_for_old_role",
			args: request.ConfigLevelRoleRequest{
				GuildID: "895659000996200508",
				RoleID:  "1003862842707017729",
				Level:   6,
			},
			wantCode:         http.StatusInternalServerError,
			wantResponsePath: "testdata/config/500_role_has_been_used_for_level_role.json",
		},
		{
			name:             "empty_body",
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/400-missing-guildID.json",
		},
		{
			name: "update_level_with_zero_value",
			args: request.ConfigLevelRoleRequest{
				GuildID: "895659000996200508",
				RoleID:  "1003862842707017729",
			},
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/config/400_invalid_level.json",
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
			ctx.Request = httptest.NewRequest(http.MethodPost, "/api/v1/configs/level-roles", bytes.NewBuffer(body))

			h.ConfigLevelRole(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.ConfigLevelRole] response mismatched")
		})
	}
}

func TestHandler_GetLevelRoleConfigs(t *testing.T) {
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
		param            string
		wantCode         int
		wantResponsePath string
	}{
		{
			name:             "guild_has_been_configured",
			param:            "895659000996200508",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/guild_config_level_roles/200_get_guild_has_been_configured.json",
		},
		{
			name:             "guild_is_not_configured",
			param:            "863278424433229854",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/200_data_empty_slice.json",
		},
		{
			name:             "empty_param",
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/400-missing-guildID.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/configs/level-roles/%s", tt.param), nil)
			ctx.AddParam("guild_id", tt.param)

			h.GetLevelRoleConfigs(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.GetLevelRoleConfigs] response mismatched")
		})
	}
}

func TestHandler_RemoveLevelRoleConfig(t *testing.T) {
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
		param            string
		query            string
		wantCode         int
		wantResponsePath string
	}{
		{
			name:             "level_has_been_configured",
			param:            "895659000996200508",
			query:            "level=2",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/200-message-ok.json",
		},
		{
			name:             "level_is_not_configured",
			param:            "895659000996200508",
			query:            "level=5",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/200-message-ok.json",
		},
		{
			name:             "guild_is_not_configured",
			param:            "863278424433229854",
			query:            "level=2",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/200-message-ok.json",
		},
		{
			name:             "empty_query",
			param:            "895659000996200508",
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/guild_config_level_roles/400_level_is_required.json",
		},
		{
			name:             "empty_param",
			query:            "level=5",
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/400-missing-guildID.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/configs/level-roles/%s?%s", tt.param, tt.query), nil)
			ctx.AddParam("guild_id", tt.param)

			h.RemoveLevelRoleConfig(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.RemoveLevelRoleConfig] response mismatched")
		})
	}
}
