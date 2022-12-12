package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestHandler_GetInviteTrackerConfig(t *testing.T) {
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
			query:            "962589711841525780",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/guild_config_invite_trackers/200.json",
		},
		{
			name:             "400_missing_guild_id",
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/guild_config_invite_trackers/400_missing_guild_id.json",
		},
		{
			name:             "404_guild_id_does_not_exist",
			query:            "962589711841525123",
			wantCode:         http.StatusNotFound,
			wantResponsePath: "testdata/404_record_not_found.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/community/invites/config?guild_id=%s", tt.query), nil)

			h.GetInviteTrackerConfig(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.GetInviteTrackerConfig] response mismatched")
		})
	}
}

func TestHandler_ConfigureInvites(t *testing.T) {
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
			args: request.ConfigureInviteRequest{
				LogChannel: "964773702476656701",
				GuildID:    "962589711841525780",
			},
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/200-message-ok.json",
		},
		{
			name: "400_invalid_json",
			args: map[string]interface{}{
				"log_channel": 964773702476656701,
				"guild_id":    962589711841525780,
			},
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/guild_config_invite_trackers/400_invalid_json.json",
		},
		{
			name: "400_missing_guild_id",
			args: request.ConfigureInviteRequest{
				LogChannel: "964773702476656701",
			},
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/400-missing-guildID.json",
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
			ctx.Request = httptest.NewRequest(http.MethodPost, "/api/v1/community/invites/config", bytes.NewBuffer(body))

			h.ConfigureInvites(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.GetInviteTrackerConfig] response mismatched")
		})
	}
}

func TestHandler_GetInvitesLeaderboard(t *testing.T) {
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
		args             string
		wantCode         int
		wantResponsePath string
	}{
		{
			name:             "200_ok",
			args:             "962589711841525780",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/response/user_invites_aggregation/200_get_invites_leaderboard_response.json",
		},
		{
			name:             "400_missing_guild_id",
			args:             "",
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/400-missing-guildID.json",
		},
		{
			name:             "200_guild_does_not_have_leaderboard",
			args:             "962589711841525123",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/200-data-null.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/community/invites/leaderboard/%s", tt.args), nil)
			ctx.AddParam("id", tt.args)

			h.GetInvitesLeaderboard(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.GetInvitesLeaderboard] response mismatched")
		})
	}
}

func TestHandler_InvitesAggregation(t *testing.T) {
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
		args             []string
		wantCode         int
		wantResponsePath string
	}{
		{
			name: "200_ok",
			args: []string{
				"guild_id=962589711841525780",
				"inviter_id=962592086849376266",
			},
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/response/user_invites_aggregation/200.json",
		},
		{
			name: "200_invalid_guild_id",
			args: []string{
				"guild_id=962589711841525123",
				"inviter_id=962592086849376266",
			},
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/response/user_invites_aggregation/200_zero_invite.json",
		},
		{
			name: "400_missing_guild_id",
			args: []string{
				"inviter_id=962592086849376266",
			},
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/400-missing-guildID.json",
		},
		{
			name: "400_missing_inviter_id",
			args: []string{
				"guild_id=962589711841525780",
			},
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/400_missing_inviter_id.json",
		},
		{
			name:             "400_missing_query",
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/400-missing-guildID.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/community/invites/aggregation?%s", strings.Join(tt.args, "&")), nil)

			h.InvitesAggregation(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.InvitesAggregation] response mismatched")
		})
	}
}
