package user

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/repo/pg"
	"github.com/defipod/mochi/pkg/service"
	mock_processor "github.com/defipod/mochi/pkg/service/processor/mocks"
	"github.com/defipod/mochi/pkg/util/testhelper"
)

func TestHandler_GetUserCurrentGMStreak(t *testing.T) {
	db := testhelper.LoadTestDB("../../../migrations/test_seed")
	repo := pg.NewRepo(db)
	cfg := config.LoadTestConfig()
	log := logger.NewLogrusLogger()
	entity := entities.New(cfg, log, repo, nil, nil, nil, nil, nil, nil, nil, nil, nil)

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
			name:             "User has Gm streak",
			query:            "guild_id=552427722551459840&discord_id=393034938028392449",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/users/200-exist-gm-streak.json",
		},
		{
			name:             "400_empty_guild_id",
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/400-missing-discordID-guildID.json",
		},
		{
			name:             "400_empty_discord_id",
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/400-missing-discordID-guildID.json",
		},
		{
			name:             "User does not have Gm streak",
			query:            "guild_id=552427722551459840&discord_id=not_have_gm_streak",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/users/200-not-exist-gm-streak.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/users/gmstreak?%s", tt.query), nil)

			h.GetUserCurrentGMStreak(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.GetUserCurrentGMStreak] response mismatched")
		})
	}
}

func TestHandler_GetUserProfile(t *testing.T) {
	cfg := config.LoadTestConfig()
	db := testhelper.LoadTestDB("../../../migrations/test_seed")
	repo := pg.NewRepo(db)
	log := logger.NewLogrusLogger()
	s := pg.NewPostgresStore(&cfg)
	svc, _ := service.NewService(cfg, log)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	processorMock := mock_processor.NewMockService(ctrl)
	svc.Processor = processorMock
	entity := entities.New(cfg, log, repo, s, nil, nil, nil, svc, nil, nil, nil, nil)

	tests := []struct {
		name             string
		query            string
		wantCode         int
		wantError        error
		wantResponsePath string
	}{
		{
			name:             "User has profiles",
			query:            "guild_id=552427722551459840&profile_id=393034938028392449",
			wantCode:         http.StatusOK,
			wantError:        nil,
			wantResponsePath: "testdata/users/200-user-profiles.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				entities: entity,
				log:      log,
			}
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/profiles?%s", tt.query), nil)

			h.GetUserProfile(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEqf(t, string(expRespRaw), w.Body.String(), "[Handler.GetUserProfile] response mismatched")
		})
	}
}

func TestHandler_GetTopUsers(t *testing.T) {
	db := testhelper.LoadTestDB("../../../migrations/test_seed")
	repo := pg.NewRepo(db)
	cfg := config.LoadTestConfig()
	log := logger.NewLogrusLogger()
	entity := entities.New(cfg, log, repo, nil, nil, nil, nil, nil, nil, nil, nil, nil)

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
			name:             "Top users list",
			query:            "guild_id=552427722551459840&profile_id=393034938028392449",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/users/200-top-users.json",
		},
		{
			name:             "400_empty_guild_id",
			query:            "profile_id=393034938028392449",
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/400-missing-guildID.json",
		},
		{
			name:             "400_empty_profile_id",
			query:            "guild_id=552427722551459840",
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/400-missing-profileID.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/users/top?%s", tt.query), nil)

			h.GetTopUsers(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEqf(t, string(expRespRaw), w.Body.String(), "[Handler.GetTopUsers] response mismatched")
		})
	}
}
