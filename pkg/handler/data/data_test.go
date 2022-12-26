package data

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
	"github.com/defipod/mochi/pkg/util/testhelper"
)

func TestHandler_MetricByProperties(t *testing.T) {
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
			name:             "Get metrics by property success",
			query:            "q=nft_collections",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/metrics/200-metrics-by-properties.json",
		},
	}

	for _, tt := range tests {
		t.Run(t.Name(), func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/data/metrics?%s", tt.query), nil)

			h.MetricByProperties(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.MetricByProperties] response mismatched")
		})
	}
}

func TestHandler_AddGitbookClick(t *testing.T) {
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
		// {
		// 	name:             "Missing url",
		// 	query:            "command=test",
		// 	wantCode:         http.StatusBadRequest,
		// 	wantResponsePath: "testdata/usage-stats/400-missing-url-and-command.json",
		// },
		// {
		// 	name:             "Missing command",
		// 	query:            "url=test",
		// 	wantCode:         http.StatusOK,
		// 	wantResponsePath: "testdata/metrics/200-metrics-by-properties.json",
		// },
		// {
		// 	name:             "Get metrics by property success",
		// 	query:            "url=test&command=test",
		// 	wantCode:         http.StatusOK,
		// 	wantResponsePath: "testdata/metrics/200-metrics-by-properties.json",
		// },
	}

	for _, tt := range tests {
		t.Run(t.Name(), func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/data/usage-stats/gitbook?%s", tt.query), nil)

			h.MetricByProperties(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.MetricByProperties] response mismatched")
		})
	}
}

func TestHandler_AddServersUsageStat(t *testing.T) {
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
		// 	name: "Add server usage stats successfully",
		// 	args: request.UsageInformation{
		// 		UserID:          "test",
		// 		GuildID:         "test",
		// 		Command:         "test",
		// 		Args:            "test",
		// 		Success:         true,
		// 		ExecutionTimeMs: 1000,
		// 	},
		// 	wantCode:         http.StatusOK,
		// 	wantResponsePath: "testdata/usage-stats/200-message-ok.json",
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
			ctx.Request = httptest.NewRequest(http.MethodPost, "/api/v1/data/usage-stats", bytes.NewBuffer(body))

			h.AddServersUsageStat(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.AddServersUsageStat] response mismatched")
		})
	}
}
