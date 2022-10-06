package entities

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/repo"
	mock_guildconfigdefaultroles "github.com/defipod/mochi/pkg/repo/guild_config_default_roles/mocks"
	mock_guildconfiggroupnftrole "github.com/defipod/mochi/pkg/repo/guild_config_group_nft_role/mocks"
	mock_guildconfiglevelrole "github.com/defipod/mochi/pkg/repo/guild_config_level_role/mocks"
	mock_guildconfigreactionroles "github.com/defipod/mochi/pkg/repo/guild_config_reaction_roles/mocks"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func TestEntity_ListAllReactionRoles(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mock repo
	gcrrRepo := mock_guildconfigreactionroles.NewMockStore(ctrl)
	repo := &repo.Repo{
		GuildConfigReactionRole: gcrrRepo,
	}

	// create entity
	cfg := config.LoadTestConfig()
	log := logger.NewLogrusLogger()
	e := &Entity{
		cfg:  cfg,
		log:  log,
		repo: repo,
	}

	type res struct {
		data []model.GuildConfigReactionRole
		err  error
	}
	type args struct {
		guildID string
	}
	tests := []struct {
		name    string
		args    args
		res     res
		want    *response.ListRoleReactionResponse
		wantErr bool
	}{
		{
			name: "1_message_1_reaction",
			args: args{
				guildID: "962589711841525780",
			},
			res: res{
				data: []model.GuildConfigReactionRole{
					{
						ID: uuid.NullUUID{
							UUID:  uuid.New(),
							Valid: true,
						},
						GuildID:       "962589711841525780",
						MessageID:     "964773702476656701",
						ReactionRoles: string([]byte(`[{"id": "1008961210546401391", "reaction": "✅"}]`)),
					},
				},
				err: nil,
			},
			want: &response.ListRoleReactionResponse{
				GuildID: "962589711841525780",
				Configs: []response.RoleReactionByMessage{
					{
						MessageID: "964773702476656701",
						Roles: []response.Role{
							{
								ID:       "1008961210546401391",
								Reaction: "✅",
							},
						},
					},
				},
				Success: true,
			},
			wantErr: false,
		},
		{
			name: "1_message_2_reaction",
			args: args{
				guildID: "962589711841525780",
			},
			res: res{
				data: []model.GuildConfigReactionRole{
					{
						ID: uuid.NullUUID{
							UUID:  uuid.New(),
							Valid: true,
						},
						GuildID:       "962589711841525780",
						MessageID:     "964773702476656701",
						ReactionRoles: string([]byte(`[{"id": "1008961210546401391", "reaction": "✅"},{"id": "1008961210546401391", "reaction": "🌟"}]`)),
					},
				},
				err: nil,
			},
			want: &response.ListRoleReactionResponse{
				GuildID: "962589711841525780",
				Configs: []response.RoleReactionByMessage{
					{
						MessageID: "964773702476656701",
						Roles: []response.Role{
							{
								ID:       "1008961210546401391",
								Reaction: "✅",
							},
							{
								ID:       "1008961210546401391",
								Reaction: "🌟",
							},
						},
					},
				},
				Success: true,
			},
			wantErr: false,
		},
		{
			name: "empty_config",
			args: args{
				guildID: "962589711841525780",
			},
			res: res{
				data: []model.GuildConfigReactionRole{},
				err:  nil,
			},
			want: &response.ListRoleReactionResponse{
				GuildID: "962589711841525780",
				Configs: []response.RoleReactionByMessage{},
				Success: true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gcrrRepo.EXPECT().ListAllByGuildID(tt.args.guildID).Return(tt.res.data, tt.res.err).Times(1)

			got, err := e.ListAllReactionRoles(tt.args.guildID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.ListAllReactionRoles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entity.ListAllReactionRoles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntity_UpdateConfigByMessageID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mock repo
	gcdrRepo := mock_guildconfigdefaultroles.NewMockStore(ctrl)
	gcdrRepo.EXPECT().GetAllByGuildID(gomock.Any()).Return(model.GuildConfigDefaultRole{}, gorm.ErrRecordNotFound).AnyTimes()

	gclrRepo := mock_guildconfiglevelrole.NewMockStore(ctrl)
	gclrRepo.EXPECT().GetByRoleID(gomock.Any(), gomock.Any()).Return(nil, gorm.ErrRecordNotFound).AnyTimes()

	gcgnrRepo := mock_guildconfiggroupnftrole.NewMockStore(ctrl)
	gcgnrRepo.EXPECT().GetByRoleID(gomock.Any(), gomock.Any()).Return(nil, gorm.ErrRecordNotFound).AnyTimes()

	gcrrRepo := mock_guildconfigreactionroles.NewMockStore(ctrl)
	gcrrRepo.EXPECT().ListAllByGuildID(gomock.Any()).Return([]model.GuildConfigReactionRole{}, nil).AnyTimes()

	repo := &repo.Repo{
		GuildConfigReactionRole: gcrrRepo,
		GuildConfigDefaultRole:  gcdrRepo,
		GuildConfigLevelRole:    gclrRepo,
		GuildConfigGroupNFTRole: gcgnrRepo,
	}

	// create entity
	cfg := config.LoadTestConfig()
	log := logger.NewLogrusLogger()
	e := &Entity{
		cfg:  cfg,
		log:  log,
		repo: repo,
	}

	type res struct {
		data model.GuildConfigReactionRole
		err  error
	}
	type args struct {
		req request.RoleReactionUpdateRequest
	}
	tests := []struct {
		name    string
		args    args
		res     []res
		want    *response.RoleReactionConfigResponse
		wantErr bool
	}{
		{
			name: "message_have_1_config",
			args: args{
				req: request.RoleReactionUpdateRequest{
					GuildID:   "962589711841525780",
					MessageID: "964773702476656701",
					Reaction:  "🌟",
					RoleID:    "974737489388519444",
				},
			},
			res: []res{
				{
					data: model.GuildConfigReactionRole{
						ID: uuid.NullUUID{
							UUID:  uuid.New(),
							Valid: true,
						},
						GuildID:       "962589711841525780",
						MessageID:     "964773702476656701",
						ReactionRoles: string([]byte(`[{"id": "1008961210546401391", "reaction": "✅"}]`)),
					},
					err: nil,
				},
				{
					err: nil,
				},
				{
					err: nil,
				},
			},
			want: &response.RoleReactionConfigResponse{
				GuildID:   "962589711841525780",
				MessageID: "964773702476656701",
				Roles: []response.Role{
					{
						ID:       "1008961210546401391",
						Reaction: "✅",
					},
					{
						ID:       "974737489388519444",
						Reaction: "🌟",
					},
				},
				Success: true,
			},
			wantErr: false,
		},
		{
			name: "no_config_for_message",
			args: args{
				req: request.RoleReactionUpdateRequest{
					GuildID:   "962589711841525780",
					MessageID: "964773702476656701",
					Reaction:  "🌟",
					RoleID:    "974737489388519444",
				},
			},
			res: []res{
				{
					data: model.GuildConfigReactionRole{},
					err:  gorm.ErrRecordNotFound,
				},
				{
					err: nil,
				},
				{
					err: nil,
				},
			},
			want: &response.RoleReactionConfigResponse{
				GuildID:   "962589711841525780",
				MessageID: "964773702476656701",
				Roles: []response.Role{
					{
						ID:       "974737489388519444",
						Reaction: "🌟",
					},
				},
				Success: true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gcrrRepo.EXPECT().GetByMessageID(tt.args.req.GuildID, tt.args.req.MessageID).Return(tt.res[0].data, tt.res[0].err).Times(1)

			roles := []response.Role{
				{
					ID:       tt.args.req.RoleID,
					Reaction: tt.args.req.Reaction,
				},
			}
			data, err := json.Marshal(roles)
			if err != nil {
				t.Errorf("Entity.UpdateConfigByMessageID() cannot marshal roles: %v", err)
				return
			}

			if tt.res[0].err == gorm.ErrRecordNotFound {
				gcrrRepo.EXPECT().CreateRoleConfig(tt.args.req, string(data)).Return(tt.res[1].err).Times(1)
			} else {
				resRoles := []response.Role{}
				err = json.Unmarshal([]byte(tt.res[0].data.ReactionRoles), &resRoles)
				if err != nil {
					t.Errorf("Entity.UpdateConfigByMessageID() cannot unmarshal roles: %v", err)
					return
				}

				roles = append(resRoles, roles...)
				data, err := json.Marshal(roles)
				if err != nil {
					t.Errorf("Entity.UpdateConfigByMessageID() cannot marshal roles: %v", err)
					return
				}
				gcrrRepo.EXPECT().UpdateRoleConfig(tt.args.req, string(data)).Return(tt.res[2].err).Times(1)
			}

			got, err := e.UpdateConfigByMessageID(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.UpdateConfigByMessageID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entity.UpdateConfigByMessageID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntity_RemoveSpecificRoleReaction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mock repo
	gcrrRepo := mock_guildconfigreactionroles.NewMockStore(ctrl)
	repo := &repo.Repo{
		GuildConfigReactionRole: gcrrRepo,
	}

	// create entity
	cfg := config.LoadTestConfig()
	log := logger.NewLogrusLogger()
	e := &Entity{
		cfg:  cfg,
		log:  log,
		repo: repo,
	}

	type res struct {
		data model.GuildConfigReactionRole
		err  error
	}
	type args struct {
		req request.RoleReactionUpdateRequest
		res []res
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "message_has_been_configured",
			args: args{
				req: request.RoleReactionUpdateRequest{
					GuildID:   "962589711841525780",
					MessageID: "964773702476656701",
					Reaction:  "🌟",
					RoleID:    "974737489388519444",
				},
				res: []res{
					{
						data: model.GuildConfigReactionRole{
							ID: uuid.NullUUID{
								UUID:  uuid.New(),
								Valid: true,
							},
							GuildID:       "962589711841525780",
							MessageID:     "964773702476656701",
							ReactionRoles: string([]byte(`[{"id": "1008961210546401391", "reaction": "✅"},{"id": "974737489388519444", "reaction": "🌟"}]`)),
						},
						err: nil,
					},
					{
						err: nil,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "role_does_not_config",
			args: args{
				req: request.RoleReactionUpdateRequest{
					GuildID:   "962589711841525780",
					MessageID: "964773702476656701",
					Reaction:  "✅",
					RoleID:    "1008961210546401391",
				},
				res: []res{
					{
						data: model.GuildConfigReactionRole{
							ID: uuid.NullUUID{
								UUID:  uuid.New(),
								Valid: true,
							},
							GuildID:       "962589711841525780",
							MessageID:     "964773702476656701",
							ReactionRoles: string([]byte(`[{"id": "974737489388519444", "reaction": "🌟"}]`)),
						},
						err: nil,
					},
					{
						err: nil,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "message_does_not_config",
			args: args{
				req: request.RoleReactionUpdateRequest{
					GuildID:   "962589711841525780",
					MessageID: "964773702476656701",
					Reaction:  "🌟",
					RoleID:    "974737489388519444",
				},
				res: []res{
					{
						data: model.GuildConfigReactionRole{},
						err:  gorm.ErrRecordNotFound,
					},
					{
						err: nil,
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gcrrRepo.EXPECT().
				GetByMessageID(tt.args.req.GuildID, tt.args.req.MessageID).
				Return(tt.args.res[0].data, tt.args.res[0].err).Times(1)

			if tt.args.res[0].err == nil {
				roles := []response.Role{}
				err := json.Unmarshal([]byte(tt.args.res[0].data.ReactionRoles), &roles)
				if err != nil {
					t.Errorf("Entity.UpdateConfigByMessageID() cannot unmarshal roles: %v", err)
					return
				}

				var updatedRoles []response.Role
				for _, r := range roles {
					if r.ID != tt.args.req.RoleID {
						updatedRoles = append(updatedRoles, r)
					}
				}

				data, err := json.Marshal(updatedRoles)
				if err != nil {
					t.Errorf("Entity.UpdateConfigByMessageID() cannot marshal roles: %v", err)
					return
				}

				gcrrRepo.EXPECT().
					UpdateRoleConfig(tt.args.req, string(data)).
					Return(tt.args.res[1].err).Times(1)
			}

			if err := e.RemoveSpecificRoleReaction(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("Entity.RemoveSpecificRoleReaction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEntity_ClearReactionMessageConfig(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mock repo
	gcrrRepo := mock_guildconfigreactionroles.NewMockStore(ctrl)
	repo := &repo.Repo{
		GuildConfigReactionRole: gcrrRepo,
	}

	// create entity
	cfg := config.LoadTestConfig()
	log := logger.NewLogrusLogger()
	e := &Entity{
		cfg:  cfg,
		log:  log,
		repo: repo,
	}

	type res struct {
		err error
	}
	type args struct {
		req request.RoleReactionUpdateRequest
		res res
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				req: request.RoleReactionUpdateRequest{
					GuildID:   "962589711841525780",
					MessageID: "964773702476656701",
					Reaction:  "🌟",
					RoleID:    "974737489388519444",
				},
				res: res{
					err: nil,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gcrrRepo.EXPECT().ClearMessageConfig(tt.args.req.GuildID, tt.args.req.MessageID).Return(tt.args.res.err).Times(1)
			if err := e.ClearReactionMessageConfig(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("Entity.ClearReactionMessageConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEntity_GetReactionRoleByMessageID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mock repo
	gcrrRepo := mock_guildconfigreactionroles.NewMockStore(ctrl)
	repo := &repo.Repo{
		GuildConfigReactionRole: gcrrRepo,
	}

	// create entity
	cfg := config.LoadTestConfig()
	log := logger.NewLogrusLogger()
	e := &Entity{
		cfg:  cfg,
		log:  log,
		repo: repo,
	}

	type res struct {
		data model.GuildConfigReactionRole
		err  error
	}
	type req struct {
		GuildID   string
		MessageID string
		Reaction  string
	}
	type args struct {
		req req
		res res
	}
	tests := []struct {
		name    string
		args    args
		want    *response.RoleReactionResponse
		wantErr bool
	}{
		{
			name: "role_configured",
			args: args{
				req: req{
					GuildID:   "962589711841525780",
					MessageID: "964773702476656701",
					Reaction:  "🌟",
				},
				res: res{
					data: model.GuildConfigReactionRole{
						ID: uuid.NullUUID{
							UUID:  uuid.New(),
							Valid: true,
						},
						GuildID:       "962589711841525780",
						MessageID:     "964773702476656701",
						ReactionRoles: string([]byte(`[{"id": "974737489388519444", "reaction": "🌟"}]`)),
					},
					err: nil,
				},
			},
			want: &response.RoleReactionResponse{
				GuildID:   "962589711841525780",
				MessageID: "964773702476656701",
				Role: response.Role{
					ID:       "974737489388519444",
					Reaction: "🌟",
				},
			},
			wantErr: false,
		},
		{
			name: "role_is_not_configured",
			args: args{
				req: req{
					GuildID:   "962589711841525780",
					MessageID: "964773702476656701",
					Reaction:  "✅",
				},
				res: res{
					data: model.GuildConfigReactionRole{
						ID: uuid.NullUUID{
							UUID:  uuid.New(),
							Valid: true,
						},
						GuildID:       "962589711841525780",
						MessageID:     "964773702476656701",
						ReactionRoles: string([]byte(`[{"id": "974737489388519444", "reaction": "🌟"}]`)),
					},
					err: nil,
				},
			},
			want: &response.RoleReactionResponse{
				GuildID:   "962589711841525780",
				MessageID: "964773702476656701",
				Role:      response.Role{},
			},
			wantErr: false,
		},
		{
			name: "message_is_not_configured",
			args: args{
				req: req{
					GuildID:   "962589711841525780",
					MessageID: "964773702476656123",
					Reaction:  "🌟",
				},
				res: res{
					data: model.GuildConfigReactionRole{},
					err:  gorm.ErrRecordNotFound,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gcrrRepo.EXPECT().GetByMessageID(tt.args.req.GuildID, tt.args.req.MessageID).Return(tt.args.res.data, tt.args.res.err).Times(1)

			got, err := e.GetReactionRoleByMessageID(tt.args.req.GuildID, tt.args.req.MessageID, tt.args.req.Reaction)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.GetReactionRoleByMessageID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entity.GetReactionRoleByMessageID() = %v, want %v", got, tt.want)
			}
		})
	}
}
