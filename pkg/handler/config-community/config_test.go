package configcommunity

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
