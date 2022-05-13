package handler

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http/httptest"
// 	"testing"
// 	"time"

// 	"github.com/defipod/mochi/pkg/config"
// 	discordWalletMocks "github.com/defipod/mochi/pkg/discordwallet/mocks"
// 	"github.com/defipod/mochi/pkg/entities"
// 	"github.com/defipod/mochi/pkg/logger"
// 	"github.com/defipod/mochi/pkg/request"
// 	"github.com/ethereum/go-ethereum/accounts"
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/mock"
// )

// func Test_HandleDiscordWebhook(t *testing.T) {
// 	discordWallet := discordWalletMocks.IDiscordWallet{}
// 	discordWallet.On("GetAccountByWalletNumber", mock.Anything).Return(accounts.Account{
// 		Address: common.HexToAddress("0x65c150B7eF3B1adbB9cB2b8041C892b15eDde05A"),
// 	}, nil)

// 	cfg := config.Config{
// 		DBUser: "postgres",
// 		DBPass: "postgres",
// 		DBHost: "localhost",
// 		DBPort: "5434",
// 		DBName: "mochi_local",

// 		InDiscordWalletMnemonic: "holiday frequent toy bachelor auto use style result recycle crumble glue blouse",
// 		FantomRPC:               "https://rpc.ftm.tools",
// 		FantomScan:              "https://api.ftmscan.com/api?",
// 		FantomScanAPIKey:        "XEKSVDF5VWQDY5VY6ZNT6AK9QPQRH483EF",

// 		EthereumRPC:        "https://mainnet.infura.io/v3/5b389eb75c514cf6b1711d70084b0114",
// 		EthereumScan:       "https://api.etherscan.io/api?",
// 		EthereumScanAPIKey: "SM5BHYSNIRZ1HEWJ1JPHVTMJS95HRA6DQF",

// 		BscRPC:        "https://bsc-dataseed.binance.org",
// 		BscScan:       "https://api.bscscan.com/api?",
// 		BscScanAPIKey: "VTKF4RG4HP6WXQ5QTAJ8MHDDIUFYD6VZHC",

// 		DiscordToken: "OTY1NTMxOTUxOTMzMDU1MDY2.Yl0jtA.iPVk82cHTnmmSdRUyK8ygt9h8P4",

// 		RedisURL: "redis://localhost:6379/0",
// 	}

// 	l := logger.NewLogrusLogger()

// 	entities.Init(cfg, l)

// 	h := Handler{
// 		entities: entities.Get(),
// 	}

// 	type author struct {
// 		ID string
// 	}

// 	type argsData struct {
// 		Author    author
// 		GuildID   string `json:"guild_id"`
// 		ChannelID string `json:"channel_id"`
// 		Timestamp time.Time
// 		Content   string
// 	}

// 	type args struct {
// 		Event string
// 		Data  argsData
// 	}

// 	type result struct {
// 		code   int
// 		err    error
// 		status string
// 	}

// 	gin.SetMode(gin.TestMode)

// 	tests := []struct {
// 		name string
// 		args args
// 		want result
// 	}{
// 		{
// 			name: "successfully handled",
// 			args: args{request.MESSAGE_CREATE, argsData{author{"760874365037314100"}, "878692765683298344", "895659000996200508", time.Now(), "hello"}},
// 			want: result{
// 				code:   200,
// 				err:    nil,
// 				status: "OK",
// 			},
// 		},
// 		// {
// 		// 	name: "successfully get user info by discord_id",
// 		// 	args: args{request.MESSAGE_CREATE, argsData{author{"760874365037314100"}, "", "895659000996200508", time.Now(), ""}},
// 		// 	want: result{
// 		// 		code: 200,
// 		// 		err:  nil,
// 		// 	},
// 		// },
// 		// {
// 		// 	name: "user not verified",
// 		// 	args: args{request.MESSAGE_CREATE, argsData{author{"760874365037314100"}, "", "895659000996200508", time.Now(), ""}},
// 		// 	want: result{
// 		// 		code: 400,
// 		// 		err:  fmt.Errorf("unverified user"),
// 		// 	},
// 		// },
// 	}

// 	gin.SetMode(gin.TestMode)

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {

// 			body, err := json.Marshal(tt.args)
// 			if err != nil {
// 				t.Error(err)
// 				return
// 			}

// 			w := httptest.NewRecorder()
// 			ctx, _ := gin.CreateTestContext(w)
// 			ctx.Request = httptest.NewRequest("POST", "/api/v1/webhook", bytes.NewBuffer(body))

// 			h.HandleDiscordWebhook(ctx)

// 			var got struct {
// 				Status string
// 			}
// 			if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
// 				t.Error(err)
// 				return
// 			}
// 			if got.Status != "OK" {
// 				t.Errorf("Handler.HandleDiscordWebhook() status = %v, want %v", got.Status, tt.want.status)
// 				return
// 			}

// 			if tt.want.code != w.Code {
// 				t.Errorf("Handler.HandleDiscordWebhook() code = %v, want %v, error: %s", w.Code, tt.want.code, "")
// 				return
// 			}

// 			// if tt.want.err != nil {
// 			// 	if tt.want.err.Error() != got.Error {
// 			// 		t.Errorf("Handler.HandleDiscordWebhook() error = %v, want %v", got.Error, tt.want.err)
// 			// 	}
// 			// 	return
// 			// }

// 			// if tt.want.user.Address != got.User.Address {
// 			// 	t.Errorf("Handler.GetUserInfo() address = %v, want %v", got.User.Address, tt.want.user.Address)
// 			// 	return
// 			// }

// 			// if !reflect.DeepEqual(*got.User.XPs, *tt.want.user.XPs) {
// 			// 	t.Errorf("Handler.GetUserInfo() XPs = %v, want %v", got.User.XPs, tt.want.user.XPs)
// 			// 	return
// 			// }
// 		})
// 	}
// }
