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
	"github.com/defipod/mochi/pkg/util/testhelper"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestHandler_GetUserUpvoteLeaderboard(t *testing.T) {
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
			name:             "success - get by streak count",
			query:            "streak",
			wantCode:         200,
			wantResponsePath: "testdata/users/200-vote-top-streak.json",
		},
		{
			name:             "success - get by total count",
			query:            "total",
			wantCode:         200,
			wantResponsePath: "testdata/users/200-vote-top-total.json",
		},
		{
			name:             "success - get by default",
			query:            "",
			wantCode:         200,
			wantResponsePath: "testdata/users/200-vote-top-total.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest("GET", fmt.Sprintf("/api/v1/users/upvote-leaderboard?by=%s", tt.query), nil)

			h.GetUserUpvoteLeaderboard(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)
			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.GetUserUpvoteLeaderboard] response mismatched")
		})
	}
}
