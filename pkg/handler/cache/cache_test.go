package cache

import (
	"bytes"
	"encoding/json"
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

func TestHandler_SetUpvoteMessageCache(t *testing.T) {
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
			name: "400_missing_guild_id",
			args: request.SetUpvoteMessageCacheRequest{
				GuildID:   "",
				MessageID: "test",
				ChannelID: "test",
				UserID:    "test",
			},
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/400-missing-required-fields.json",
		},
		{
			name: "400_missing_user_id",
			args: request.SetUpvoteMessageCacheRequest{
				GuildID:   "test",
				MessageID: "test",
				ChannelID: "test",
				UserID:    "",
			},
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/400-missing-required-fields.json",
		},
		// {
		// 	name: "Set upvote message cache succesfully",
		// 	args: request.SetUpvoteMessageCacheRequest{
		// 		GuildID:   "test",
		// 		MessageID: "test",
		// 		ChannelID: "test",
		// 		UserID:    "test",
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
			ctx.Request = httptest.NewRequest(http.MethodPost, "/api/v1/cache/upvote", bytes.NewBuffer(body))

			h.SetUpvoteMessageCache(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.SetUpvoteMessageCache] response mismatched")
		})
	}
}
