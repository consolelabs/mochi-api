package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"

	"github.com/defipod/mochi/pkg/cache"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/discordwallet"
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/repo/pg"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/service/abi"
	"github.com/defipod/mochi/pkg/service/indexer"
	"github.com/defipod/mochi/pkg/service/marketplace"
)

func Test_HandleDiscordWebhook(t *testing.T) {
	cfg := config.Config{
		DBUser: "postgres",
		DBPass: "postgres",
		DBHost: "localhost",
		DBPort: "5434",
		DBName: "mochi_local",

		RedisURL: "redis://localhost:6379/0",
	}

	redisOpt, err := redis.ParseURL(cfg.RedisURL)
	if err != nil {
		t.Error(err)
		return
	}

	cache, err := cache.NewRedisCache(redisOpt)
	if err != nil {
		t.Error(err)
		return
	}

	s := pg.NewPostgresStore(&cfg)
	repo := pg.NewRepo(s.DB())
	l := logger.NewLogrusLogger()
	indexer := indexer.NewIndexer(cfg, l)
	abi := abi.NewAbi(&cfg)
	marketplace := marketplace.NewMarketplace(&cfg)

	e := entities.New(cfg, l, repo, s, &discordwallet.DiscordWallet{}, nil, cache, nil, indexer, abi, marketplace)

	h := Handler{
		entities: e,
		log:      l,
	}

	// upsert guild
	if err := repo.DiscordGuilds.CreateOrReactivate(model.DiscordGuild{
		ID:        "878692765683298344",
		Name:      "test-server",
		BotScopes: model.JSONArrayString{"*"},
		GlobalXP:  false,
	}); err != nil {
		t.Error(err)
		return
	}

	type author struct {
		ID string
	}

	type argsData struct {
		Author    author
		GuildID   string `json:"guild_id"`
		ChannelID string `json:"channel_id"`
		Timestamp time.Time
		Content   string
	}

	type args struct {
		Event string
		Data  argsData
	}

	type result struct {
		code   int
		err    error
		status string
		resp   *response.HandleUserActivityResponse
		Type   string
	}

	gin.SetMode(gin.TestMode)

	tests := []struct {
		name string
		args args
		want result
	}{
		{
			name: "bad request - no event specified",
			args: args{"", argsData{author{"760874365037314100"}, "", "895659000996200508", time.Now(), ""}},
			want: result{
				code: 400,
				err:  nil,
				resp: nil,
			},
		},
		{
			name: "internal server error - invalid guild ID",
			args: args{request.MESSAGE_CREATE, argsData{author{"760874365037314100"}, "", "895659000996200508", time.Now(), "abc"}},
			want: result{
				code: 500,
				err:  nil,
				resp: nil,
			},
		},
		// {
		// 	name: "successfully handled user chat and add xp",
		// 	args: args{request.MESSAGE_CREATE, argsData{author{"760874365037314100"}, "878692765683298344", "895659000996200508", time.Now(), "hello"}},
		// 	want: result{
		// 		code:   200,
		// 		err:    nil,
		// 		status: "OK",
		// 		Type:   "level_up",
		// 		resp: &response.HandleUserActivityResponse{
		// 			UserID: "760874365037314100",
		// 			Action: "chat",
		// 		},
		// 	},
		// },
	}

	gin.SetMode(gin.TestMode)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			xpID := fmt.Sprintf(`%s_%s_chat_xp_cooldown`, tt.args.Data.Author.ID, tt.args.Data.GuildID)
			exists, err := cache.GetBool(xpID)
			if err != nil {
				t.Error(err)
				return
			}

			body, err := json.Marshal(tt.args)
			if err != nil {
				t.Error(err)
				return
			}

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest("POST", "/api/v1/webhook", bytes.NewBuffer(body))

			h.HandleDiscordWebhook(ctx)

			var got struct {
				Status string
				Type   string
				Data   *response.HandleUserActivityResponse
			}
			if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
				t.Error(err)
				return
			}

			if tt.want.code != w.Code {
				t.Errorf("Handler.HandleDiscordWebhook() code = %v, want %v, error: %s", w.Code, tt.want.code, "")
				return
			}

			if tt.want.resp == nil && got.Data != nil {
				t.Errorf("Handler.HandleDiscordWebhook() resp = %v, want %v, error: %s", got.Data, tt.want.resp, "")
				return
			}

			if !exists && tt.want.resp != nil {
				if tt.want.Type != got.Type {
					t.Errorf("Handler.HandleDiscordWebhook() type = %v, want %v, error: %s", got.Type, tt.want.Type, "")
					return
				}

				if tt.want.resp.UserID != got.Data.UserID {
					t.Errorf("Handler.HandleDiscordWebhook() resp.UserID = %v, want %v, error: %s", got.Data.UserID, tt.want.resp.UserID, "")
					return
				}

				if tt.want.resp.Action != got.Data.Action {
					t.Errorf("Handler.HandleDiscordWebhook() resp.Action = %v, want %v, error: %s", got.Data.Action, tt.want.resp.Action, "")
					return
				}
			} else if exists {
				if got.Data != nil {
					t.Errorf("Handler.HandleDiscordWebhook() resp = %v, want %v, error: %s", got.Data, nil, "")
					return
				}
			}
		})
	}
}
