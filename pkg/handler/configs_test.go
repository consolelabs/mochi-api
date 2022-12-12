package handler

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
	"github.com/defipod/mochi/pkg/util"
	"github.com/defipod/mochi/pkg/util/testhelper"
)

// func TestHandler_ConfigLevelRole(t *testing.T) {
// 	cfg := config.LoadTestConfig()
// 	db := testhelper.LoadTestDB("../../migrations/test_seed")
// 	log := logger.NewLogrusLogger()
// 	s := pg.NewPostgresStore(&cfg)

// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	repo := pg.NewRepo(db)
// 	entityMock := entities.New(cfg, log, repo, s, nil, nil, nil, nil, nil, nil, nil)

// 	tests := []struct {
// 		name             string
// 		req              request.ConfigLevelRoleRequest
// 		wantCode         int
// 		wantError        error
// 		wantResponsePath string
// 	}{
// 		{
// 			name: "success config level role",
// 			req: request.ConfigLevelRoleRequest{
// 				GuildID: "895659000996200508",
// 				RoleID:  "1003867749841383425",
// 				Level:   5,
// 			},
// 			wantCode:         200,
// 			wantError:        nil,
// 			wantResponsePath: "testdata/200-message-ok-uppercase.json",
// 		},
// 		{
// 			name: "success upsert level role 3",
// 			req: request.ConfigLevelRoleRequest{
// 				GuildID: "895659000996200508",
// 				RoleID:  "1003867749841383426",
// 				Level:   3,
// 			},
// 			wantCode:         200,
// 			wantError:        nil,
// 			wantResponsePath: "testdata/200-message-ok-uppercase.json",
// 		},
// 		{
// 			name: "fail to config - lack of guildID",
// 			req: request.ConfigLevelRoleRequest{
// 				RoleID: "1003867749841383426",
// 				Level:  3,
// 			},
// 			wantCode:         400,
// 			wantError:        nil,
// 			wantResponsePath: "testdata/400-missing-guildID.json",
// 		},
// 		{
// 			name: "fail to config - lack of roleID",
// 			req: request.ConfigLevelRoleRequest{
// 				GuildID: "895659000996200508",
// 				Level:   3,
// 			},
// 			wantCode:         400,
// 			wantError:        nil,
// 			wantResponsePath: "testdata/400-missing-roleID.json",
// 		},
// 		{
// 			name: "fail to config - lack of level",
// 			req: request.ConfigLevelRoleRequest{
// 				GuildID: "895659000996200508",
// 				RoleID:  "1003867749841383426",
// 			},
// 			wantCode:         400,
// 			wantError:        nil,
// 			wantResponsePath: "testdata/config_level_role/400-missing-level.json",
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			h := &Handler{
// 				entities: entityMock,
// 				log:      log,
// 			}
// 			w := httptest.NewRecorder()
// 			ctx, _ := gin.CreateTestContext(w)
// 			ctx.Request = httptest.NewRequest("POST", "/api/v1/configs/level-roles", nil)
// 			util.SetRequestBody(ctx, tt.req)
// 			h.ConfigLevelRole(ctx)
// 			require.Equal(t, tt.wantCode, w.Code)
// 			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
// 			require.NoError(t, err)
// 			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.ConfigLevelRole] response mismatched")
// 		})
// 	}
// }

// func TestHandler_GetLevelRoleConfigs(t *testing.T) {
// 	cfg := config.LoadTestConfig()
// 	db := testhelper.LoadTestDB("../../migrations/test_seed")
// 	log := logger.NewLogrusLogger()
// 	s := pg.NewPostgresStore(&cfg)

// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	repo := pg.NewRepo(db)
// 	entityMock := entities.New(cfg, log, repo, s, nil, nil, nil, nil, nil, nil, nil)

// 	type args struct {
// 		guildID string
// 	}
// 	tests := []struct {
// 		name             string
// 		req              args
// 		wantCode         int
// 		wantError        error
// 		wantResponsePath string
// 	}{
// 		{
// 			name: "success get level role",
// 			req: args{
// 				guildID: "895659000996200508",
// 			},
// 			wantCode:         200,
// 			wantError:        nil,
// 			wantResponsePath: "testdata/get_level_role_configs/200-success.json",
// 		},
// 		{
// 			name:             "fail to get - lack of guildID",
// 			req:              args{},
// 			wantCode:         400,
// 			wantError:        nil,
// 			wantResponsePath: "testdata/400-missing-guildID.json",
// 		},
// 		{
// 			name: "no data to get",
// 			req: args{
// 				guildID: "abc",
// 			},
// 			wantCode:         200,
// 			wantError:        nil,
// 			wantResponsePath: "testdata/get_level_role_configs/200-no-data.json",
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			h := &Handler{
// 				entities: entityMock,
// 				log:      log,
// 			}
// 			w := httptest.NewRecorder()
// 			ctx, _ := gin.CreateTestContext(w)
// 			ctx.Params = []gin.Param{
// 				{
// 					Key:   "guild_id",
// 					Value: tt.req.guildID,
// 				},
// 			}
// 			ctx.Request = httptest.NewRequest("GET", "/api/v1/configs/level-roles/", nil)
// 			h.GetLevelRoleConfigs(ctx)
// 			require.Equal(t, tt.wantCode, w.Code)
// 			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
// 			require.NoError(t, err)
// 			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.GetLevelRoleConfigs] response mismatched")
// 		})
// 	}
// }

// func TestHandler_RemoveLevelRoleConfig(t *testing.T) {
// 	cfg := config.LoadTestConfig()
// 	db := testhelper.LoadTestDB("../../migrations/test_seed")
// 	log := logger.NewLogrusLogger()
// 	s := pg.NewPostgresStore(&cfg)

// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	repo := pg.NewRepo(db)
// 	entityMock := entities.New(cfg, log, repo, s, nil, nil, nil, nil, nil, nil, nil)

// 	type args struct {
// 		guildID string
// 		level   string
// 	}

// 	tests := []struct {
// 		name             string
// 		req              args
// 		wantCode         int
// 		wantError        error
// 		wantResponsePath string
// 	}{
// 		{
// 			name: "success delete config level role",
// 			req: args{
// 				guildID: "895659000996200508",
// 				level:   "2",
// 			},
// 			wantCode:         200,
// 			wantError:        nil,
// 			wantResponsePath: "testdata/200-message-ok-uppercase.json",
// 		},
// 		{
// 			name:             "fail to delete level role - lack of guildid",
// 			req:              args{},
// 			wantCode:         400,
// 			wantError:        nil,
// 			wantResponsePath: "testdata/400-missing-guildID.json",
// 		},
// 		{
// 			name: "fail to delete level role - lack of level",
// 			req: args{
// 				guildID: "895659000996200508",
// 			},
// 			wantCode:         400,
// 			wantError:        nil,
// 			wantResponsePath: "testdata/remove_level_role_config/400-missing-level.json",
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			h := &Handler{
// 				entities: entityMock,
// 				log:      log,
// 			}
// 			w := httptest.NewRecorder()
// 			ctx, _ := gin.CreateTestContext(w)
// 			ctx.Params = []gin.Param{
// 				{
// 					Key:   "guild_id",
// 					Value: tt.req.guildID,
// 				},
// 			}
// 			q := url.Values{}
// 			q.Add("level", tt.req.level)
// 			ctx.Request = httptest.NewRequest("DELETE", "/api/v1/configs/level-roles/", nil)
// 			ctx.Request.URL.RawQuery = q.Encode()
// 			h.RemoveLevelRoleConfig(ctx)
// 			require.Equal(t, tt.wantCode, w.Code)
// 			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
// 			require.NoError(t, err)
// 			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.ConfigLevelRole] response mismatched")
// 		})
// 	}
// }

// func newState() (*discordgo.State, error) {
// 	role := mockrole.New(
// 		mockrole.WithID("1012578894550937620"),
// 		mockrole.WithName("nftrole1"),
// 		mockrole.WithPermissions(discordgo.PermissionViewChannel),
// 	)

// 	rolenotNFT := mockrole.New(
// 		mockrole.WithID("1012578894550937621"),
// 		mockrole.WithName("notnftrole"),
// 		mockrole.WithPermissions(discordgo.PermissionViewChannel),
// 	)

// 	botUser := mockuser.New(
// 		mockuser.WithID(mockconstants.TestUser+"Bot"),
// 		mockuser.WithUsername(mockconstants.TestUser+"Bot"),
// 		mockuser.WithBotFlag(true),
// 	)

// 	botMember := mockmember.New(
// 		mockmember.WithUser(botUser),
// 		mockmember.WithGuildID("895659000996200508"),
// 		mockmember.WithRoles(role),
// 	)

// 	userMember := mockmember.New(
// 		mockmember.WithUser(mockuser.New(
// 			mockuser.WithID(mockconstants.TestUser),
// 			mockuser.WithUsername(mockconstants.TestUser),
// 		)),
// 		mockmember.WithGuildID("895659000996200508"),
// 		mockmember.WithRoles(role, rolenotNFT),
// 	)

// 	return mockstate.New(
// 		mockstate.WithUser(botUser),
// 		mockstate.WithGuilds(
// 			mockguild.New(
// 				mockguild.WithID("895659000996200508"),
// 				mockguild.WithRoles(role, rolenotNFT),
// 				mockguild.WithMembers(botMember, userMember),
// 			),
// 		),
// 	)
// }

// func TestHandler_ListGuildNFTRoles(t *testing.T) {
// 	cfg := config.LoadTestConfig()
// 	db := testhelper.LoadTestDB("../../migrations/test_seed")
// 	log := logger.NewLogrusLogger()
// 	s := pg.NewPostgresStore(&cfg)

// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	repo := pg.NewRepo(db)
// 	state, _ := newState()
// 	session, _ := mocksession.New(
// 		mocksession.WithState(state),
// 		mocksession.WithClient(&http.Client{
// 			Transport: mockrest.NewTransport(state),
// 		}),
// 	)

// 	entityMock := entities.New(cfg, log, repo, s, nil, session, nil, nil, nil, nil, nil)

// 	type args struct {
// 		guildID string
// 	}

// 	tests := []struct {
// 		name             string
// 		req              args
// 		wantCode         int
// 		wantError        error
// 		wantResponsePath string
// 	}{
// 		{
// 			name: "success get nftrole",
// 			req: args{
// 				guildID: "895659000996200508",
// 			},
// 			wantCode:         200,
// 			wantError:        nil,
// 			wantResponsePath: "testdata/list_nft_roles/200-success.json",
// 		},
// 		{
// 			name:             "fail to get - lack of guildID",
// 			req:              args{},
// 			wantCode:         400,
// 			wantError:        nil,
// 			wantResponsePath: "testdata/400-missing-guildID.json",
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			h := &Handler{
// 				entities: entityMock,
// 				log:      log,
// 			}
// 			w := httptest.NewRecorder()
// 			ctx, _ := gin.CreateTestContext(w)
// 			q := url.Values{}
// 			q.Add("guild_id", tt.req.guildID)
// 			ctx.Request = httptest.NewRequest("GET", "/api/v1/configs/nft-roles", nil)
// 			ctx.Request.URL.RawQuery = q.Encode()
// 			h.ListGuildGroupNFTRoles(ctx)
// 			require.Equal(t, tt.wantCode, w.Code)
// 			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
// 			require.NoError(t, err)
// 			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.ConfigLevelRole] response mismatched")
// 		})
// 	}
// }

// func TestHandler_NewGuildNFTRole(t *testing.T) {
// 	cfg := config.LoadTestConfig()
// 	db := testhelper.LoadTestDB("../../migrations/test_seed")
// 	log := logger.NewLogrusLogger()
// 	s := pg.NewPostgresStore(&cfg)

// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	repo := pg.NewRepo(db)
// 	entityMock := entities.New(cfg, log, repo, s, nil, nil, nil, nil, nil, nil, nil)

// 	tests := []struct {
// 		name             string
// 		req              request.ConfigGroupNFTRoleRequest
// 		wantCode         int
// 		wantError        error
// 		wantResponsePath string
// 	}{
// 		{
// 			name: "success config nftrole",
// 			req: request.ConfigGroupNFTRoleRequest{
// 				CollectionAddress: []string{"1a42432c-b1a8-4874-b7cc-875a5086742a"},
// 				GuildID:           "863278424433229854",
// 				RoleID:            "1012578894550937621",
// 				NumberOfTokens:    1,
// 				GroupName:         "test",
// 			},
// 			wantCode:         201,
// 			wantError:        nil,
// 			wantResponsePath: "testdata/new_guild_nft_role/201-success.json",
// 		},
// 		{
// 			name: "fail to config - lack of guildID",
// 			req: request.ConfigNFTRoleRequest{
// 				GuildConfigNFTRole: model.GuildConfigNFTRole{
// 					NFTCollectionID: util.GetNullUUID("1a42432c-b1a8-4874-b7cc-875a5086742a"),
// 					RoleID:          "1012578894550937621",
// 					NumberOfTokens:  1,
// 				},
// 			},
// 			wantCode:         400,
// 			wantError:        nil,
// 			wantResponsePath: "testdata/400-missing-guildID.json",
// 		},
// 		{
// 			name: "fail to config - lack of nft collection id",
// 			req: request.ConfigNFTRoleRequest{
// 				GuildConfigNFTRole: model.GuildConfigNFTRole{
// 					GuildID:        "863278424433229854",
// 					RoleID:         "1012578894550937621",
// 					NumberOfTokens: 1,
// 				},
// 			},
// 			wantCode:         400,
// 			wantError:        nil,
// 			wantResponsePath: "testdata/new_guild_nft_role/400-missing-nft.json",
// 		},
// 		{
// 			name: "fail to config - lack of roleID",
// 			req: request.ConfigNFTRoleRequest{
// 				GuildConfigNFTRole: model.GuildConfigNFTRole{
// 					NFTCollectionID: util.GetNullUUID("1a42432c-b1a8-4874-b7cc-875a5086742a"),
// 					GuildID:         "863278424433229854",
// 					NumberOfTokens:  1,
// 				},
// 			},
// 			wantCode:         400,
// 			wantError:        nil,
// 			wantResponsePath: "testdata/400-missing-roleID.json",
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			h := &Handler{
// 				entities: entityMock,
// 				log:      log,
// 			}
// 			w := httptest.NewRecorder()
// 			ctx, _ := gin.CreateTestContext(w)
// 			ctx.Request = httptest.NewRequest("POST", "/api/v1/configs/nft-roles", nil)
// 			util.SetRequestBody(ctx, tt.req)
// 			h.NewGuildGroupNFTRole(ctx)
// 			require.Equal(t, tt.wantCode, w.Code)
// 			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
// 			require.NoError(t, err)
// 			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.ConfigLevelRole] response mismatched")
// 		})
// 	}
// }

// func TestHandler_RemoveGuildNFTRole(t *testing.T) {
// 	cfg := config.LoadTestConfig()
// 	db := testhelper.LoadTestDB("../../migrations/test_seed")
// 	log := logger.NewLogrusLogger()
// 	s := pg.NewPostgresStore(&cfg)

// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()
// 	repo := pg.NewRepo(db)
// 	entityMock := entities.New(cfg, log, repo, s, nil, nil, nil, nil, nil, nil, nil)

// 	type args struct {
// 		configID string
// 	}

// 	tests := []struct {
// 		name             string
// 		req              args
// 		wantCode         int
// 		wantError        error
// 		wantResponsePath string
// 	}{
// 		{
// 			name: "success delete nftrole",
// 			req: args{
// 				configID: "ee531d57-8c2e-45f8-b328-5f8bd57470b6",
// 			},
// 			wantCode:         200,
// 			wantError:        nil,
// 			wantResponsePath: "testdata/200-message-ok-uppercase.json",
// 		},
// 		{
// 			name: "fail to delete - lack of guildID",
// 			req: args{
// 				configID: "",
// 			},
// 			wantCode:         400,
// 			wantError:        nil,
// 			wantResponsePath: "testdata/remove_guild_nft_role/400-missing-configID.json",
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			h := &Handler{
// 				entities: entityMock,
// 				log:      log,
// 			}
// 			w := httptest.NewRecorder()
// 			ctx, _ := gin.CreateTestContext(w)
// 			ctx.Params = []gin.Param{
// 				{
// 					Key:   "config_id",
// 					Value: tt.req.configID,
// 				},
// 			}
// 			ctx.Request = httptest.NewRequest("DELETE", "/api/v1/configs/nft-roles", nil)
// 			h.RemoveGuildNFTRole(ctx)
// 			require.Equal(t, tt.wantCode, w.Code)
// 			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
// 			require.NoError(t, err)
// 			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.ConfigLevelRole] response mismatched")
// 		})
// 	}
// }

func TestHandler_GetWelcomeChannelConfig(t *testing.T) {
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
			name:             "guild_is_configured",
			query:            "guild_id=895659000996200508",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/guild_config_welcome_channels/200-get-ok.json",
		},
		{
			name:             "guild_id_is_empty",
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/400-missing-guildID.json",
		},
		{
			name:             "guild_is_not_configured",
			query:            "guild_id=863278424433229854",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/200-data-null.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/configs/welcome?%s", tt.query), nil)

			h.GetWelcomeChannelConfig(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.GetWelcomeChannelConfig] response mismatched")
		})
	}
}

func TestHandler_UpsertWelcomeChannelConfig(t *testing.T) {
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
		args             request.UpsertWelcomeConfigRequest
		query            string
		wantCode         int
		wantResponsePath string
	}{
		{
			name: "update_new_msg",
			args: request.UpsertWelcomeConfigRequest{
				GuildID:    "895659000996200508",
				ChannelID:  "1016919074221064256",
				WelcomeMsg: "Welcome to the guild!",
			},
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/guild_config_welcome_channels/200-upsert-new-msg.json",
		},
		{
			name: "update_with_empty_msg",
			args: request.UpsertWelcomeConfigRequest{
				GuildID:   "895659000996200508",
				ChannelID: "1016919074221064256",
			},
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/guild_config_welcome_channels/200-upsert-new-msg.json",
		},
		{
			name: "update_with_new_channel",
			args: request.UpsertWelcomeConfigRequest{
				GuildID:    "895659000996200508",
				ChannelID:  "1016919074221064123",
				WelcomeMsg: "Welcome to the guild!",
			},
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/guild_config_welcome_channels/200-upsert-new-channel.json",
		},
		{
			name: "no_channel_id_and_guild_id",
			args: request.UpsertWelcomeConfigRequest{
				WelcomeMsg: "Welcome to the guild!",
			},
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/400-missing-guildID.json",
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
			ctx.Request = httptest.NewRequest(http.MethodPost, "/api/v1/configs/welcome", bytes.NewBuffer(body))

			h.UpsertWelcomeChannelConfig(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.UpsertWelcomeChannelConfig] response mismatched")
		})
	}
}

func TestHandler_DeleteWelcomeChannelConfig(t *testing.T) {
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
		args             request.DeleteWelcomeConfigRequest
		wantCode         int
		wantResponsePath string
	}{
		{
			name: "configuration_exist",
			args: request.DeleteWelcomeConfigRequest{
				GuildID: "895659000996200508",
			},
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/200-message-ok.json",
		},
		{
			name: "configuration_does_not_exist",
			args: request.DeleteWelcomeConfigRequest{
				GuildID: "863278424433229854",
			},
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/200-message-ok.json",
		},
		{
			name:             "empty_input",
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/400-missing-guildID.json",
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
			ctx.Request = httptest.NewRequest(http.MethodDelete, "/api/v1/configs/welcome", bytes.NewBuffer(body))

			h.DeleteWelcomeChannelConfig(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.DeleteWelcomeChannelConfig] response mismatched")
		})
	}
}

func TestHandler_GetVoteChannelConfig(t *testing.T) {
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
			name:             "guild_is_configured",
			query:            "guild_id=895659000996200508",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/guild_config_vote_channels/200-get-ok.json",
		},
		{
			name:             "guild_id_is_empty",
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/400-missing-guildID.json",
		},
		{
			name:             "guild_is_not_configured",
			query:            "guild_id=863278424433229854",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/200-data-null.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/configs/upvote?%s", tt.query), nil)

			h.GetVoteChannelConfig(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.GetVoteChannelConfig] response mismatched")
		})
	}
}

func TestHandler_UpsertVoteChannelConfig(t *testing.T) {
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
		args             request.UpsertVoteChannelConfigRequest
		query            string
		wantCode         int
		wantResponsePath string
	}{
		{
			name: "update_new_channel",
			args: request.UpsertVoteChannelConfigRequest{
				GuildID:   "895659000996200508",
				ChannelID: "1016919074221064123",
			},
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/guild_config_vote_channels/200-upsert-new-channel.json",
		},
		{
			name:             "no_channel_id_and_guild_id",
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/400-missing-guildID.json",
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
			ctx.Request = httptest.NewRequest(http.MethodPost, "/api/v1/configs/upvote", bytes.NewBuffer(body))

			h.UpsertVoteChannelConfig(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.UpsertVoteChannelConfig] response mismatched")
		})
	}
}

func TestHandler_DeleteVoteChannelConfig(t *testing.T) {
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
		args             request.DeleteWelcomeConfigRequest
		query            string
		wantCode         int
		wantResponsePath string
	}{
		{
			name: "configuration_exist",
			args: request.DeleteWelcomeConfigRequest{
				GuildID: "895659000996200508",
			},
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/200-message-ok.json",
		},
		{
			name: "configuration_does_not_exist",
			args: request.DeleteWelcomeConfigRequest{
				GuildID: "863278424433229854",
			},
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/200-message-ok.json",
		},
		{
			name:             "empty_input",
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/400-missing-guildID.json",
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
			ctx.Request = httptest.NewRequest(http.MethodDelete, "/api/v1/configs/upvote", bytes.NewBuffer(body))

			h.DeleteVoteChannelConfig(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.DeleteVoteChannelConfig] response mismatched")
		})
	}
}

func TestHandler_GetGuildTokens(t *testing.T) {
	cfg := config.LoadTestConfig()
	db := testhelper.LoadTestDB("../../migrations/test_seed")
	repo := pg.NewRepo(db)
	log := logger.NewLogrusLogger()
	s := pg.NewPostgresStore(&cfg)

	entityMock := entities.New(cfg, log, repo, s, nil, nil, nil, nil, nil, nil, nil)
	h := &Handler{
		entities: entityMock,
		log:      log,
	}
	tests := []struct {
		name             string
		guildId          string
		wantError        error
		wantCode         int
		wantResponsePath string
	}{
		// TODO: Add test cases.
		{
			name:             "success - with guild id",
			guildId:          "863278424433229854",
			wantCode:         200,
			wantResponsePath: "testdata/guild_config_tokens/200-with-id.json",
		},
		{
			name:             "success - without guild id",
			guildId:          "",
			wantCode:         200,
			wantResponsePath: "testdata/guild_config_tokens/200-without-id.json",
		},
		{
			name:             "success - with wrong guild id",
			guildId:          "123123123",
			wantCode:         200,
			wantResponsePath: "testdata/guild_config_tokens/200-without-id.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest("GET", fmt.Sprintf("/api/v1/configs/tokens?guild_id=%s", tt.guildId), nil)

			h.GetGuildTokens(ctx)

			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)
			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.GetGuildTokens] response mismatched")
		})
	}
}

func TestHandler_UpsertGuildTokenConfig(t *testing.T) {
	cfg := config.LoadTestConfig()
	db := testhelper.LoadTestDB("../../migrations/test_seed")
	repo := pg.NewRepo(db)
	log := logger.NewLogrusLogger()
	s := pg.NewPostgresStore(&cfg)

	entityMock := entities.New(cfg, log, repo, s, nil, nil, nil, nil, nil, nil, nil)
	h := &Handler{
		entities: entityMock,
		log:      log,
	}
	tests := []struct {
		name             string
		req              request.UpsertGuildTokenConfigRequest
		wantError        error
		wantCode         int
		wantResponsePath string
	}{
		{
			name: "success - add - new token",
			req: request.UpsertGuildTokenConfigRequest{
				GuildID: "863278424433229854",
				Symbol:  "TOMB",
				Active:  true,
			},
			wantError:        nil,
			wantCode:         200,
			wantResponsePath: "testdata/200-message-ok.json",
		},
		{
			name: "success - add - same token",
			req: request.UpsertGuildTokenConfigRequest{
				GuildID: "863278424433229854",
				Symbol:  "FTM",
				Active:  true,
			},
			wantError:        nil,
			wantCode:         200,
			wantResponsePath: "testdata/200-message-ok.json",
		},
		{
			name: "success - remove - existed token",
			req: request.UpsertGuildTokenConfigRequest{
				GuildID: "863278424433229854",
				Symbol:  "SHIB",
				Active:  false,
			},
			wantError:        nil,
			wantCode:         200,
			wantResponsePath: "testdata/200-message-ok.json",
		},
		{
			name: "failed - remove - non existed token",
			req: request.UpsertGuildTokenConfigRequest{
				GuildID: "863278424433229854",
				Symbol:  "ABCC",
				Active:  false,
			},
			wantError:        nil,
			wantCode:         200,
			wantResponsePath: "testdata/404_record_not_found.json",
		},
		{
			name: "failed - missing guild",
			req: request.UpsertGuildTokenConfigRequest{
				GuildID: "",
				Symbol:  "ABCC",
				Active:  false,
			},
			wantError:        nil,
			wantCode:         200,
			wantResponsePath: "testdata/400-missing-guildID.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest("POST", "/api/v1/configs/token", nil)
			util.SetRequestBody(ctx, tt.req)

			h.UpsertGuildTokenConfig(ctx)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)
			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.UpsertGuildTokenConfig] response mismatched")
		})
	}
}
