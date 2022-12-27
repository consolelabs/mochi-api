package configcommunity

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

func TestHandler_LinkUserTelegramWithDiscord(t *testing.T) {
	db := testhelper.LoadTestDB("../../../migrations/test_seed")
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
			name: "add new association",
			args: request.LinkUserTelegramWithDiscordRequest{
				DiscordID:        "963641551416881183",
				TelegramUsername: "hmh",
			},
			wantCode:         201,
			wantResponsePath: "testdata/200-data-null.json",
		},
		{
			name: "upsert association successfully",
			args: request.LinkUserTelegramWithDiscordRequest{
				DiscordID:        "463379262620041226",
				TelegramUsername: "anhnh",
			},
			wantCode:         201,
			wantResponsePath: "testdata/200-data-null.json",
		},
		{
			name: "upsert association falied by duplicate telegram username",
			args: request.LinkUserTelegramWithDiscordRequest{
				DiscordID:        "963641551416881183",
				TelegramUsername: "trkhoi",
			},
			wantCode:         500,
			wantResponsePath: "testdata/telegram/500-duplicate-username.json",
		},
		{
			name: "empty discord id",
			args: request.LinkUserTelegramWithDiscordRequest{
				TelegramUsername: "hmhabc",
			},
			wantCode:         400,
			wantResponsePath: "testdata/telegram/400-missing-discordID-telegramusername.json",
		},
		{
			name: "empty telegram username",
			args: request.LinkUserTelegramWithDiscordRequest{
				DiscordID: "963641551416881183",
			},
			wantCode:         400,
			wantResponsePath: "testdata/telegram/400-missing-discordID-telegramusername.json",
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
			ctx.Request = httptest.NewRequest(http.MethodPost, "/api/v1/configs/telegram", bytes.NewBuffer(body))
			ctx.Request.Header.Set("Content-Type", "application/json")

			h.LinkUserTelegramWithDiscord(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.LinkUserTelegramWithDiscord] response mismatched")
		})
	}
}

func TestHandler_GetLinkedTelegram(t *testing.T) {
	db := testhelper.LoadTestDB("../../../migrations/test_seed")
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
			name:             "400_telegram_username_is_required",
			query:            "",
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/telegram/400-teleram-username-required.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/config-community/telegram?%s", tt.query), nil)

			h.GetLinkedTelegram(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.GetLinkedTelegram] response mismatched")
		})
	}
}

func TestHandler_GetAllTwitterConfig(t *testing.T) {
	db := testhelper.LoadTestDB("../../../migrations/test_seed")
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
			name:             "200_message_ok",
			query:            "?guild_id=testt",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/200_data_empty_slice.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/config-community/twitter%s", tt.query), nil)

			h.GetAllTwitterConfig(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.GetAllTwitterConfig] response mismatched")
		})
	}
}

func TestHandler_GetTwitterHashtagConfig(t *testing.T) {
	db := testhelper.LoadTestDB("../../../migrations/test_seed")
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
			name:             "200_message_ok",
			param:            "testt",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/200-data-null.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/config-community/twitter/hashtag/%s", tt.param), nil)

			h.GetTwitterHashtagConfig(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.GetTwitterHashtagConfig] response mismatched")
		})
	}
}

func TestHandler_GetAllTwitterHashtagConfig(t *testing.T) {
	db := testhelper.LoadTestDB("../../../migrations/test_seed")
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
			name:             "200_data_empty_slice",
			param:            "",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/200_data_empty_slice.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodGet, "/api/v1/config-community/twitter/hashtag", nil)

			h.GetAllTwitterHashtagConfig(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.GetAllTwitterHashtagConfig] response mismatched")
		})
	}
}

func TestHandler_GetTwitterBlackList(t *testing.T) {
	db := testhelper.LoadTestDB("../../../migrations/test_seed")
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
			name:             "200_message_ok",
			query:            "?guild_id=testt",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/200_data_empty_slice.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/config-community/twitter/blacklist%s", tt.query), nil)

			h.GetTwitterBlackList(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.GetTwitterBlackList] response mismatched")
		})
	}
}

func TestHandler_DeleteTwitterHashtagConfig(t *testing.T) {
	db := testhelper.LoadTestDB("../../../migrations/test_seed")
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
			name:             "200_ok",
			param:            "testt",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/200-message-ok.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/config-community/twitter/hashtag/%s", tt.param), nil)

			h.DeleteTwitterHashtagConfig(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.DeleteTwitterHashtagConfig] response mismatched")
		})
	}
}

func TestHandler_DeleteFromTwitterBlackList(t *testing.T) {
	db := testhelper.LoadTestDB("../../../migrations/test_seed")
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
		// {
		// 	name: "200_ok",
		// 	args: request.DeleteFromTwitterBlackListRequest{
		// 		GuildID:   "testt",
		// 		TwitterID: "testt",
		// 	},
		// 	wantCode:         http.StatusOK,
		// 	wantResponsePath: "testdata/200-message-ok.json",
		// },
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
			ctx.Request = httptest.NewRequest(http.MethodDelete, "/api/v1/config-community/twitter/blacklist", bytes.NewBuffer(body))

			h.DeleteFromTwitterBlackList(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.DeleteFromTwitterBlackList] response mismatched")
		})
	}
}
