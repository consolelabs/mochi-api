package community

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

func TestHandler_UpdateUserFeedback(t *testing.T) {
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
		args             interface{}
		wantCode         int
		wantResponsePath string
	}{
		{
			name: "Missing guild_id",
			args: request.UpdateUserFeedbackRequest{
				ID:     "test",
				Status: "test",
			},
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/feedback/400-invalid-status.json",
		},
		// {
		// 	name: "Config repost reaction conversation succesfully",
		// 	args: request.UpdateUserFeedbackRequest{
		// 		ID:     "test",
		// 		Status: "none",
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
			ctx.Request = httptest.NewRequest(http.MethodPut, "/api/v1/community/feedback", bytes.NewBuffer(body))

			h.UpdateUserFeedback(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.UpdateUserFeedback] response mismatched")
		})
	}
}

func TestHandler_HandleUserFeedback(t *testing.T) {
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
		args             interface{}
		wantCode         int
		wantResponsePath string
	}{
		// {
		// 	name: "missing body",
		// 	args: request.UserFeedbackRequest{
		// 		DiscordID: "",
		// 		Username:  "",
		// 		MessageID: "",
		// 		Feedback:  "",
		// 		Avatar:    "",
		// 		Command:   "",
		// 	},
		// 	wantCode:         http.StatusBadRequest,
		// 	wantResponsePath: "testdata/400-missing-guildID.json",
		// },
		// {
		// 	name: "Config repost reaction conversation succesfully",
		// 	args: request.UserFeedbackRequest{
		// 		DiscordID: "test",
		// 		Username:  "test",
		// 		MessageID: "test",
		// 		Feedback:  "test",
		// 		Avatar:    "test",
		// 		Command:   "test",
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
			ctx.Request = httptest.NewRequest(http.MethodPost, "/api/v1/community/feedback", bytes.NewBuffer(body))

			h.HandleUserFeedback(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.CreateConfigRepostReactionConversation] response mismatched")
		})
	}
}

func TestHandler_GetAllUserFeedback(t *testing.T) {
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
			name:             "Filter by command",
			query:            "filter=command",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/feedback/200-get-all-user-feedback.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/community/feedback?%s", tt.query), nil)

			h.GetAllUserFeedback(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.GetAllUserFeedback] response mismatched")
		})
	}
}

func TestHandler_GetUserQuestList(t *testing.T) {
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
		// {
		// 	name:             "With user id",
		// 	query:            "user_id=test",
		// 	wantCode:         http.StatusOK,
		// 	wantResponsePath: "testdata/quests/200-get-by-user-id-test.json",
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/community/quests?%s", tt.query), nil)

			h.GetUserQuestList(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.GetUserQuestList] response mismatched")
		})
	}
}
