package handler

import (
	"net/http/httptest"
	"testing"

	"github.com/defipod/mochi/pkg/config"
	discordWalletMocks "github.com/defipod/mochi/pkg/discordwallet/mocks"
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/repo/pg"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

func TestHandler_Healthz(t *testing.T) {

	discordWallet := discordWalletMocks.IDiscordWallet{}
	discordWallet.On("GetAccountByWalletNumber", mock.Anything).Return(accounts.Account{
		Address: common.HexToAddress("0x65c150B7eF3B1adbB9cB2b8041C892b15eDde05A"),
	}, nil)

	cfg := config.Config{
		DBUser: "postgres",
		DBPass: "postgres",
		DBHost: "localhost",
		DBPort: "5434",
		DBName: "mochi_local",
	}

	s, _ := pg.NewPostgresStore(&cfg)

	l := logger.NewJSONLogger(
		logger.WithServiceName("test"),
		logger.WithHostName(""),
	)

	h := Handler{
		log:      newHandlerLog(l, ""),
		repo:     pg.NewRepo(s.DB().Debug()),
		dcwallet: &discordWallet,
		entities: entities.New(cfg, l, s, &discordWallet, nil),
	}

	type args struct {
		addr      string
		discordID string
		guildID   string
	}

	type result struct {
		code int
		err  error
	}

	gin.SetMode(gin.TestMode)

	tests := []struct {
		name string
		args args
		want result
	}{
		{
			name: "heath check success",
			want: result{
				code: 200,
				err:  nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest("GET", "/healthz", nil)
			h.Healthz(ctx)

			if tt.want.code != w.Code {
				t.Errorf("Handler.GetUserInfo() code = %v, want %v", w.Code, tt.want.code)
				return
			}
		})
	}
}
