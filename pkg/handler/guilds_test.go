package handler

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/defipod/mochi/pkg/config"
// 	"github.com/defipod/mochi/pkg/entities"
// 	"github.com/defipod/mochi/pkg/logger"
// 	"github.com/defipod/mochi/pkg/model"
// 	"github.com/defipod/mochi/pkg/repo/pg"
// 	"github.com/defipod/mochi/pkg/request"
// 	"github.com/defipod/mochi/pkg/util/testhelper"
// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/require"
// )

// func TestHandler_GetGuild(t *testing.T) {
// 	db := testhelper.LoadTestDB("../../migrations/test_seed")
// 	repo := pg.NewRepo(db)
// 	cfg := config.LoadTestConfig()
// 	log := logger.NewLogrusLogger()
// 	entity := entities.New(cfg, log, repo, nil, nil, nil, nil, nil, nil, nil, nil)

// 	h := &Handler{
// 		entities: entity,
// 		log:      log,
// 	}

// 	tests := []struct {
// 		name             string
// 		args             string
// 		wantCode         int
// 		wantResponsePath string
// 	}{
// 		{
// 			name:             "Has record guild",
// 			args:             "552427722551459840",
// 			wantCode:         http.StatusOK,
// 			wantResponsePath: "testdata/guilds/200-exist-discord-guild.json",
// 		},
// 		{
// 			name:             "400_empty_guild_id",
// 			wantCode:         http.StatusBadRequest,
// 			args:             "",
// 			wantResponsePath: "testdata/400-missing-guildID.json",
// 		},
// 		{
// 			name:             "Not have record guild",
// 			args:             "not_have_record",
// 			wantCode:         http.StatusOK,
// 			wantResponsePath: "testdata/200-data-null.json",
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			w := httptest.NewRecorder()
// 			ctx, _ := gin.CreateTestContext(w)
// 			ctx.Request = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/guilds/%s", tt.args), nil)
// 			ctx.AddParam("guild_id", tt.args)

// 			h.GetGuild(ctx)
// 			require.Equal(t, tt.wantCode, w.Code)
// 			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
// 			require.NoError(t, err)

// 			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.GetGuild] response mismatched")
// 		})
// 	}
// }

// func TestHandler_UpdateGuild(t *testing.T) {
// 	db := testhelper.LoadTestDB("../../migrations/test_seed")
// 	repo := pg.NewRepo(db)
// 	cfg := config.LoadTestConfig()
// 	log := logger.NewLogrusLogger()
// 	entity := entities.New(cfg, log, repo, nil, nil, nil, nil, nil, nil, nil, nil)

// 	h := &Handler{
// 		entities: entity,
// 		log:      log,
// 	}

// 	modelTest := model.DiscordGuild{
// 		GlobalXP:   true,
// 		LogChannel: "test",
// 		Active:     true,
// 	}
// 	tests := []struct {
// 		name             string
// 		param            string
// 		args             interface{}
// 		wantCode         int
// 		wantResponsePath string
// 	}{
// 		{
// 			name:  "Has record guild and update successfully",
// 			param: "552427722551459840",
// 			args: request.UpdateGuildRequest{
// 				GlobalXP:   &modelTest.GlobalXP,
// 				LogChannel: &modelTest.LogChannel,
// 				Active:     &modelTest.Active,
// 			},
// 			wantCode:         http.StatusOK,
// 			wantResponsePath: "testdata/200-message-ok.json",
// 		},
// 		{
// 			name:  "Not have record guild",
// 			param: "not_have_record",
// 			args: request.UpdateGuildRequest{
// 				GlobalXP:   &modelTest.GlobalXP,
// 				LogChannel: &modelTest.LogChannel,
// 				Active:     &modelTest.Active,
// 			},
// 			wantCode:         http.StatusOK,
// 			wantResponsePath: "testdata/200-data-null.json",
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			body, err := json.Marshal(tt.args)
// 			if err != nil {
// 				t.Error(err)
// 				return
// 			}

// 			w := httptest.NewRecorder()
// 			ctx, _ := gin.CreateTestContext(w)
// 			ctx.Request = httptest.NewRequest(http.MethodPost, "/api/v1/guilds/"+tt.param, bytes.NewBuffer(body))
// 			ctx.AddParam("guild_id", tt.param)

// 			h.UpdateGuild(ctx)
// 			require.Equal(t, tt.wantCode, w.Code)
// 			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
// 			require.NoError(t, err)

// 			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.UpdateGuild] response mismatched")
// 		})
// 	}
// }
