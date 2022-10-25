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

func TestHandler_GetAllRoleReactionConfigs(t *testing.T) {
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
			name:             "200_ok",
			query:            "guild_id=895659000996200508",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/response/role_reaction_config/200.json",
		},
		{
			name:             "400_empty_guild_id",
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/400-missing-guildID.json",
		},
		{
			name:             "200_guild_does_not_have_config",
			query:            "guild_id=895659000996200123",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/response/role_reaction_config/200_empty_config.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/configs/reaction-roles?%s", tt.query), nil)

			h.GetAllRoleReactionConfigs(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.GetAllRoleReactionConfigs] response mismatched")
		})
	}
}

func TestHandler_AddReactionRoleConfig(t *testing.T) {
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
		args             interface{}
		wantCode         int
		wantResponsePath string
	}{
		{
			name: "200_ok",
			args: request.RoleReactionUpdateRequest{
				GuildID:   "895659000996200508",
				MessageID: "1012576367268872234",
				ChannelID: "",
				Reaction:  "🍐",
				RoleID:    "1012192448300208191",
			},
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/response/role_reaction_config/200_add_reaction_role_config.json",
		},
		{
			name: "200_guild_does_not_have_config",
			args: request.RoleReactionUpdateRequest{
				GuildID:   "863278424433229854",
				MessageID: "1012576367268872234",
				ChannelID: "",
				Reaction:  "🍐",
				RoleID:    "1012192448300208191",
			},
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/response/role_reaction_config/200_guild_does_not_have_config.json",
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
			ctx.Request = httptest.NewRequest(http.MethodPost, "/api/v1/configs/reaction-roles", bytes.NewBuffer(body))

			h.AddReactionRoleConfig(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.AddReactionRoleConfig] response mismatched")
		})
	}
}

func TestHandler_RemoveReactionRoleConfig(t *testing.T) {
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
		args             interface{}
		wantCode         int
		wantResponsePath string
	}{
		{
			name: "200_with_role_id_and_reaction",
			args: request.RoleReactionUpdateRequest{
				GuildID:   "895659000996200508",
				MessageID: "1012576367268872234",
				Reaction:  "🌟",
				RoleID:    "1007248566957396018",
			},
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/200-message-ok.json",
		},
		{
			name: "200_with_empty_role_id_and_reaction",
			args: request.RoleReactionUpdateRequest{
				GuildID:   "895659000996200508",
				MessageID: "1012576367268872234",
			},
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/200-message-ok.json",
		},
		{
			name:             "200_with_empty_body",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/200-message-ok.json",
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
			ctx.Request = httptest.NewRequest(http.MethodDelete, "/api/v1/configs/reaction-roles", bytes.NewBuffer(body))

			h.RemoveReactionRoleConfig(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.RemoveReactionRoleConfig] response mismatched")
		})
	}
}

func TestHandler_FilterConfigByReaction(t *testing.T) {
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
		args             interface{}
		wantCode         int
		wantResponsePath string
	}{
		{
			name: "200_data_already_exists",
			args: request.RoleReactionRequest{
				GuildID:   "863278424433229854",
				MessageID: "1012576367268872234",
				Reaction:  "🍐",
			},
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/response/role_reaction_config/200_filter_config.json",
		},
		{
			name: "200_message_does_not_exist",
			args: request.RoleReactionRequest{
				GuildID:   "863278424433229854",
				MessageID: "1012576367268872123",
				Reaction:  "🍐",
			},
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/200-data-null.json",
		},
		{
			name: "200_reaction_does_not_exist",
			args: request.RoleReactionRequest{
				GuildID:   "863278424433229854",
				MessageID: "1012576367268872234",
				Reaction:  "🌟",
			},
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/response/role_reaction_config/200_filter_config_reaction_does_not_exist.json",
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
			ctx.Request = httptest.NewRequest(http.MethodPost, "/api/v1/configs/reaction-roles/filter", bytes.NewBuffer(body))

			h.FilterConfigByReaction(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.FilterConfigByReaction] response mismatched")
		})
	}
}
