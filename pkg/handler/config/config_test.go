package config

import (
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

func TestHandler_ToggleActivityConfig(t *testing.T) {
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
		param            string
		query            string
		wantCode         int
		wantResponsePath string
	}{
		// {
		// 	name:             "Toggle activiy status",
		// 	param:            "chat",
		// 	query:            "guild_id=863278424433229854",
		// 	wantCode:         http.StatusOK,
		// 	wantResponsePath: "testdata/activity-config/200-toggle-activity-success.json",
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/v1/data/activities/%s?%s", tt.param, tt.query), nil)

			h.ToggleActivityConfig(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.ToggleActivityConfig] response mismatched")
		})
	}
}
