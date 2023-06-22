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

func TestHandler_GetRepostReactionConfigs(t *testing.T) {
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
		args             string
		wantCode         int
		wantResponsePath string
	}{
		{
			name:             "Guild has repost reaction configs",
			args:             "552427722551459840",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/guild_config_repost_reactions/200-exist-guild-config-repost-reaction.json",
		},
		{
			name:             "400_empty_guild_id",
			wantCode:         http.StatusBadRequest,
			args:             "",
			wantResponsePath: "testdata/400-missing-guildID.json",
		},
		{
			name:             "Guild does not have repost reaction configs",
			args:             "not_have_record",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/200_data_empty_slice.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/configs/repost-reactions/%s?reaction_type=message", tt.args), nil)
			ctx.AddParam("guild_id", tt.args)

			h.GetRepostReactionConfigs(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.GetRepostReactionConfigs] response mismatched")
		})
	}
}

func TestHandler_RemoveRepostReactionConfig(t *testing.T) {
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
			name: "Remove but missing guild_id",
			args: request.ConfigRepostRequest{
				Emoji:           "test",
				Quantity:        1,
				RepostChannelID: "test",
			},
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/400-missing-guildID.json",
		},
		{
			name: "Remove but missing emoji",
			args: request.ConfigRepostRequest{
				GuildID:         "552427722551459840",
				Quantity:        1,
				RepostChannelID: "test",
			},
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/guild_config_repost_reactions/400-missing-emoji.json",
		},
		{
			name: "Remove successfully",
			args: request.ConfigRepostRequest{
				GuildID:         "552427722551459840",
				Emoji:           "<:approve:1013775827051237486>",
				Quantity:        1,
				RepostChannelID: "test",
			},
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
			ctx.Request = httptest.NewRequest(http.MethodDelete, "/api/v1/configs/repost-reactions", bytes.NewBuffer(body))

			h.RemoveRepostReactionConfig(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.RemoveRepostReactionConfig] response mismatched")
		})
	}
}

func TestHandler_ConfigRepostReaction(t *testing.T) {
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
			name: "Config repost reaction but missing guild_id",
			args: request.ConfigRepostRequest{
				Emoji:           "test",
				Quantity:        1,
				RepostChannelID: "test",
			},
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/400-missing-guildID.json",
		},
		{
			name: "Config repost reaction but missing emoji",
			args: request.ConfigRepostRequest{
				GuildID:         "test",
				Quantity:        1,
				RepostChannelID: "test",
			},
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/guild_config_repost_reactions/400-missing-emoji.json",
		},
		{
			name: "Config repost reaction but missing quantity",
			args: request.ConfigRepostRequest{
				GuildID:         "test",
				Emoji:           "test",
				Quantity:        0,
				RepostChannelID: "test",
			},
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/guild_config_repost_reactions/400-missing-quantity.json",
		},
		{
			name: "Config repost reaction but missing repost channel id",
			args: request.ConfigRepostRequest{
				GuildID:  "test",
				Emoji:    "test",
				Quantity: 1,
			},
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/guild_config_repost_reactions/400-missing-repost-channel-id.json",
		},
		{
			name: "Config repost reaction succesfully",
			args: request.ConfigRepostRequest{
				GuildID:         "552427722551459840",
				Emoji:           "test",
				Quantity:        1,
				RepostChannelID: "test",
			},
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
			ctx.Request = httptest.NewRequest(http.MethodPost, "/api/v1/configs/repost-reactions", bytes.NewBuffer(body))

			h.ConfigRepostReaction(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.ConfigRepostReaction] response mismatched")
		})
	}
}

func TestHandler_EditMessageRepost(t *testing.T) {
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
			name: "edit msg repost reaction succesfully",
			args: request.EditMessageRepostRequest{
				GuildID:         "552427722551459840",
				OriginMessageID: "origin_msg",
				OriginChannelID: "origin_channel",
				RepostMessageID: "repost_msg",
				RepostChannelID: "repost_channel",
			},
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
			ctx.Request = httptest.NewRequest(http.MethodPost, "/api/v1/configs/repost-reactions/message-repost", bytes.NewBuffer(body))

			h.EditMessageRepost(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.EditMessageRepost] response mismatched")
		})
	}
}

func TestHandler_CreateConfigRepostReactionConversation(t *testing.T) {
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
			name: "Config repost reaction conversation but missing guild_id",
			args: request.ConfigRepostReactionStartStop{
				GuildID:         "",
				EmojiStart:      "test",
				EmojiStop:       "test",
				RepostChannelID: "test",
			},
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/400-missing-guildID.json",
		},
		{
			name: "Config repost reaction conversation but missing start emoji",
			args: request.ConfigRepostReactionStartStop{
				GuildID:         "test",
				EmojiStart:      "",
				EmojiStop:       "test",
				RepostChannelID: "test",
			},
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/guild_config_repost_reactions/400-missing-emoji.json",
		},
		{
			name: "Config repost reaction conversation but missing stop emoji",
			args: request.ConfigRepostReactionStartStop{
				GuildID:         "test",
				EmojiStart:      "test",
				EmojiStop:       "",
				RepostChannelID: "test",
			},
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/guild_config_repost_reactions/400-missing-emoji.json",
		},
		{
			name: "Config repost reaction conversation but missing repost channel id",
			args: request.ConfigRepostReactionStartStop{
				GuildID:         "test",
				EmojiStart:      "test",
				EmojiStop:       "test",
				RepostChannelID: "",
			},
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/guild_config_repost_reactions/400-missing-repost-channel-id.json",
		},
		{
			name: "Config repost reaction conversation succesfully",
			args: request.ConfigRepostReactionStartStop{
				GuildID:         "552427722551459840",
				EmojiStart:      "test",
				EmojiStop:       "test",
				RepostChannelID: "test",
			},
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
			ctx.Request = httptest.NewRequest(http.MethodPost, "/api/v1/community/repost-reactions/conversation", bytes.NewBuffer(body))

			h.CreateConfigRepostReactionConversation(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.CreateConfigRepostReactionConversation] response mismatched")
		})
	}
}

func TestHandler_RemoveConfigRepostReactionConversation(t *testing.T) {
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
			name: "Remove config repost reaction conversation but missing guild_id",
			args: request.ConfigRepostReactionStartStop{
				GuildID:    "",
				EmojiStart: "test",
				EmojiStop:  "test",
			},
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/400-missing-guildID.json",
		},
		{
			name: "Remove config repost reaction conversation but missing start emoji",
			args: request.ConfigRepostReactionStartStop{
				GuildID:    "test",
				EmojiStart: "",
				EmojiStop:  "test",
			},
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/guild_config_repost_reactions/400-missing-emoji.json",
		},
		{
			name: "Remove config repost reaction conversation but missing stop emoji",
			args: request.ConfigRepostReactionStartStop{
				GuildID:    "test",
				EmojiStart: "test",
				EmojiStop:  "",
			},
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/guild_config_repost_reactions/400-missing-emoji.json",
		},
		{
			name: "Remove config repost reaction conversation succesfully",
			args: request.ConfigRepostReactionStartStop{
				GuildID:    "552427722551459840",
				EmojiStart: "test",
				EmojiStop:  "test",
			},
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
			ctx.Request = httptest.NewRequest(http.MethodDelete, "/api/v1/community/repost-reactions/conversation", bytes.NewBuffer(body))

			h.RemoveConfigRepostReactionConversation(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.RemoveConfigRepostReactionConversation] response mismatched")
		})
	}
}

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

func TestHandler_CreateTwitterPost(t *testing.T) {
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
			name: "Twitter post successfully",
			args: request.TwitterPost{
				TwitterID:     "test",
				TwitterHandle: "test",
				TweetID:       "test",
				GuildID:       "test",
				Content:       "test",
			},
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/twitter/200-message-ok.json",
		},
		// {
		// 	name:             "Twitter post without body",
		// 	args:             request.TwitterPost{},
		// 	wantCode:         http.StatusBadRequest,
		// 	wantResponsePath: "testdata/twitter/400-error-eol.json",
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
			ctx.Request = httptest.NewRequest(http.MethodPost, "/api/v1/community/twitter", bytes.NewBuffer(body))

			h.CreateTwitterPost(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.CreateTwitterPost] response mismatched")
		})
	}
}
