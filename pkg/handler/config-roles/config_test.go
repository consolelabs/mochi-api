package configrole

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

func TestHandler_GetDefaultRolesByGuildID(t *testing.T) {
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
		query            string
		wantCode         int
		wantResponsePath string
	}{
		{
			name:             "guild_has_been_configured",
			query:            "guild_id=895659000996200508",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/response/default_role/200_guild_has_been_configured.json",
		},
		{
			name:             "guild_is_not_configured",
			query:            "guild_id=863278424433229854",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/response/default_role/200_guild_is_not_configured.json",
		},
		{
			name:             "empty_query",
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/400-missing-guildID.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/configs/default-roles?%s", tt.query), nil)

			h.GetDefaultRolesByGuildID(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.GetDefaultRolesByGuildID] response mismatched")
		})
	}
}

func TestHandler_CreateDefaultRole(t *testing.T) {
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
		args             request.CreateDefaultRoleRequest
		wantCode         int
		wantResponsePath string
	}{
		{
			name: "update_new_role",
			args: request.CreateDefaultRoleRequest{
				GuildID: "895659000996200508",
				RoleID:  "1012576984783671123",
			},
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/response/default_role/200_update_new_role.json",
		},
		{
			name: "create_new_config_for_new_guild",
			args: request.CreateDefaultRoleRequest{
				GuildID: "863278424433229854",
				RoleID:  "1011659315729403958",
			},
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/response/default_role/200_create_new_config_for_new_guild.json",
		},
		{
			name:             "empty_request",
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/response/default_role/400_empty_request.json",
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
			ctx.Request = httptest.NewRequest(http.MethodPost, "/api/v1/configs/default-roles", bytes.NewBuffer(body))

			h.CreateDefaultRole(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.CreateDefaultRole] response mismatched")
		})
	}
}

func TestHandler_DeleteDefaultRoleByGuildID(t *testing.T) {
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
		query            string
		wantCode         int
		wantResponsePath string
	}{
		{
			name:             "guild_has_been_configured",
			query:            "guild_id=895659000996200508",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/200-message-ok.json",
		},
		{
			name:             "guild_is_not_configured",
			query:            "guild_id=863278424433229854",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/200-message-ok.json",
		},
		{
			name:             "empty_query",
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/400-missing-guildID.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/configs/default-roles?%s", tt.query), nil)

			h.DeleteDefaultRoleByGuildID(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.DeleteDefaultRoleByGuildID] response mismatched")
		})
	}
}

func TestHandler_GetAllRoleReactionConfigs(t *testing.T) {
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
		query            string
		wantCode         int
		wantResponsePath string
	}{
		{
			name:             "200_ok",
			query:            "guild_id=895659000996200508",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/response/role_reaction_config/200.json",
		},
		{
			name:             "400_empty_guild_id",
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/400-missing-guildID.json",
		},
		{
			name:             "200_guild_does_not_have_config",
			query:            "guild_id=895659000996200123",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/response/role_reaction_config/200_empty_config.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/configs/reaction-roles?%s", tt.query), nil)

			h.GetAllRoleReactionConfigs(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.GetAllRoleReactionConfigs] response mismatched")
		})
	}
}

func TestHandler_AddReactionRoleConfig(t *testing.T) {
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
			name: "200_ok",
			args: request.RoleReactionUpdateRequest{
				GuildID:   "895659000996200508",
				MessageID: "1012576367268872234",
				ChannelID: "",
				Reaction:  "🍐",
				RoleID:    "1012192448300208191",
			},
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/response/role_reaction_config/200_add_reaction_role_config.json",
		},
		{
			name: "200_guild_does_not_have_config",
			args: request.RoleReactionUpdateRequest{
				GuildID:   "863278424433229854",
				MessageID: "1012576367268872234",
				ChannelID: "",
				Reaction:  "🍐",
				RoleID:    "1012192448300208191",
			},
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/response/role_reaction_config/200_guild_does_not_have_config.json",
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
			ctx.Request = httptest.NewRequest(http.MethodPost, "/api/v1/configs/reaction-roles", bytes.NewBuffer(body))

			h.AddReactionRoleConfig(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.AddReactionRoleConfig] response mismatched")
		})
	}
}

func TestHandler_RemoveReactionRoleConfig(t *testing.T) {
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
			name: "200_with_role_id_and_reaction",
			args: request.RoleReactionUpdateRequest{
				GuildID:   "895659000996200508",
				MessageID: "1012576367268872234",
				Reaction:  "🌟",
				RoleID:    "1007248566957396018",
			},
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/200-message-ok.json",
		},
		{
			name: "200_with_empty_role_id_and_reaction",
			args: request.RoleReactionUpdateRequest{
				GuildID:   "895659000996200508",
				MessageID: "1012576367268872234",
			},
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/200-message-ok.json",
		},
		{
			name:             "200_with_empty_body",
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
			ctx.Request = httptest.NewRequest(http.MethodDelete, "/api/v1/configs/reaction-roles", bytes.NewBuffer(body))

			h.RemoveReactionRoleConfig(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.RemoveReactionRoleConfig] response mismatched")
		})
	}
}

func TestHandler_FilterConfigByReaction(t *testing.T) {
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
			name: "200_data_already_exists",
			args: request.RoleReactionRequest{
				GuildID:   "863278424433229854",
				MessageID: "1012576367268872234",
				Reaction:  "🍐",
			},
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/response/role_reaction_config/200_filter_config.json",
		},
		{
			name: "200_message_does_not_exist",
			args: request.RoleReactionRequest{
				GuildID:   "863278424433229854",
				MessageID: "1012576367268872123",
				Reaction:  "🍐",
			},
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/200-data-null.json",
		},
		{
			name: "200_reaction_does_not_exist",
			args: request.RoleReactionRequest{
				GuildID:   "863278424433229854",
				MessageID: "1012576367268872234",
				Reaction:  "🌟",
			},
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/response/role_reaction_config/200_filter_config_reaction_does_not_exist.json",
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
			ctx.Request = httptest.NewRequest(http.MethodPost, "/api/v1/configs/reaction-roles/filter", bytes.NewBuffer(body))

			h.FilterConfigByReaction(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.FilterConfigByReaction] response mismatched")
		})
	}
}

func TestHandler_ConfigLevelRole(t *testing.T) {
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
		args             request.ConfigLevelRoleRequest
		wantCode         int
		wantResponsePath string
	}{
		{
			name: "create_new_config",
			args: request.ConfigLevelRoleRequest{
				GuildID: "895659000996200508",
				RoleID:  "1003867749841383425",
				Level:   5,
			},
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/200-message-ok.json",
		},
		{
			name: "update_new_level_for_old_role",
			args: request.ConfigLevelRoleRequest{
				GuildID: "895659000996200508",
				RoleID:  "1003862842707017729",
				Level:   6,
			},
			wantCode:         http.StatusInternalServerError,
			wantResponsePath: "testdata/config/500_role_has_been_used_for_level_role.json",
		},
		{
			name:             "empty_body",
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/400-missing-guildID.json",
		},
		{
			name: "update_level_with_zero_value",
			args: request.ConfigLevelRoleRequest{
				GuildID: "895659000996200508",
				RoleID:  "1003862842707017729",
			},
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/config/400_invalid_level.json",
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
			ctx.Request = httptest.NewRequest(http.MethodPost, "/api/v1/configs/level-roles", bytes.NewBuffer(body))

			h.ConfigLevelRole(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.ConfigLevelRole] response mismatched")
		})
	}
}

func TestHandler_GetLevelRoleConfigs(t *testing.T) {
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
		param            string
		wantCode         int
		wantResponsePath string
	}{
		{
			name:             "guild_has_been_configured",
			param:            "895659000996200508",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/guild_config_level_roles/200_get_guild_has_been_configured.json",
		},
		{
			name:             "guild_is_not_configured",
			param:            "863278424433229854",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/200_data_empty_slice.json",
		},
		{
			name:             "empty_param",
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/400-missing-guildID.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/configs/level-roles/%s", tt.param), nil)
			ctx.AddParam("guild_id", tt.param)

			h.GetLevelRoleConfigs(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.GetLevelRoleConfigs] response mismatched")
		})
	}
}

func TestHandler_RemoveLevelRoleConfig(t *testing.T) {
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
		param            string
		query            string
		wantCode         int
		wantResponsePath string
	}{
		{
			name:             "level_has_been_configured",
			param:            "895659000996200508",
			query:            "level=2",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/200-message-ok.json",
		},
		{
			name:             "level_is_not_configured",
			param:            "895659000996200508",
			query:            "level=5",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/200-message-ok.json",
		},
		{
			name:             "guild_is_not_configured",
			param:            "863278424433229854",
			query:            "level=2",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/200-message-ok.json",
		},
		{
			name:             "empty_query",
			param:            "895659000996200508",
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/guild_config_level_roles/400_level_is_required.json",
		},
		{
			name:             "empty_param",
			query:            "level=5",
			wantCode:         http.StatusBadRequest,
			wantResponsePath: "testdata/400-missing-guildID.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/configs/level-roles/%s?%s", tt.param, tt.query), nil)
			ctx.AddParam("guild_id", tt.param)

			h.RemoveLevelRoleConfig(ctx)
			require.Equal(t, tt.wantCode, w.Code)
			expRespRaw, err := ioutil.ReadFile(tt.wantResponsePath)
			require.NoError(t, err)

			require.JSONEq(t, string(expRespRaw), w.Body.String(), "[Handler.RemoveLevelRoleConfig] response mismatched")
		})
	}
}

// func TestHandler_ListGuildNFTRoles(t *testing.T) {
// 	cfg := config.LoadTestConfig()
// 	db := testhelper.LoadTestDB("../../../migrations/test_seed")
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
// 	db := testhelper.LoadTestDB("../../../migrations/test_seed")
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
// 	db := testhelper.LoadTestDB("../../../migrations/test_seed")
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