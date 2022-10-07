package entities

import (
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
	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
)

func TestEntity_ConfigLevelRole(t *testing.T) {
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
		err error
	}
	type args struct {
		req request.ConfigLevelRoleRequest
	}
	tests := []struct {
		name    string
		args    args
		res     res
		wantErr bool
	}{
		{
			name: "guild_has_been_configured",
			args: args{
				req: request.ConfigLevelRoleRequest{
					GuildID: "895659000996200508",
					RoleID:  "1003867749841383425",
					Level:   2,
				},
			},
		},
		{
			name: "guild_is_not_configured",
			args: args{
				req: request.ConfigLevelRoleRequest{
					GuildID: "895659000996200123",
					RoleID:  "1003867749841383425",
					Level:   2,
				},
			},
		},
		{
			name: "internal_server_error",
			args: args{
				req: request.ConfigLevelRoleRequest{
					GuildID: "895659000996200456",
					RoleID:  "1003867749841383425",
					Level:   2,
				},
			},
			res: res{
				err: gorm.ErrInvalidTransaction,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gclrRepo.EXPECT().UpsertOne(model.GuildConfigLevelRole{
				GuildID: tt.args.req.GuildID,
				RoleID:  tt.args.req.RoleID,
				Level:   tt.args.req.Level,
			}).Return(tt.res.err).Times(1)

			if err := e.ConfigLevelRole(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("Entity.ConfigLevelRole() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEntity_GetGuildLevelRoleConfigs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mock repo
	gclrRepo := mock_guildconfiglevelrole.NewMockStore(ctrl)
	repo := &repo.Repo{
		GuildConfigLevelRole: gclrRepo,
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
		data []model.GuildConfigLevelRole
		err  error
	}
	type args struct {
		guildID string
	}
	tests := []struct {
		name    string
		args    args
		res     res
		want    []model.GuildConfigLevelRole
		wantErr bool
	}{
		{
			name: "guild_has_been_configured",
			args: args{
				guildID: "895659000996200508",
			},
			res: res{
				data: []model.GuildConfigLevelRole{
					{
						GuildID: "895659000996200508",
						RoleID:  "1003867749841383425",
						Level:   2,
					},
					{
						GuildID: "895659000996200508",
						RoleID:  "1003867749841383123",
						Level:   3,
					},
				},
			},
			want: []model.GuildConfigLevelRole{
				{
					GuildID: "895659000996200508",
					RoleID:  "1003867749841383425",
					Level:   2,
				},
				{
					GuildID: "895659000996200508",
					RoleID:  "1003867749841383123",
					Level:   3,
				},
			},
		},
		{
			name: "guild_is_not_configured",
			args: args{
				guildID: "895659000996200123",
			},
			res: res{
				data: []model.GuildConfigLevelRole{},
			},
			want: []model.GuildConfigLevelRole{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gclrRepo.EXPECT().GetByGuildID(tt.args.guildID).Return(tt.res.data, tt.res.err).Times(1)

			got, err := e.GetGuildLevelRoleConfigs(tt.args.guildID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.GetGuildLevelRoleConfigs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entity.GetGuildLevelRoleConfigs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntity_RemoveGuildLevelRoleConfig(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mock repo
	gclrRepo := mock_guildconfiglevelrole.NewMockStore(ctrl)
	repo := &repo.Repo{
		GuildConfigLevelRole: gclrRepo,
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
		guildID string
		level   int
	}
	tests := []struct {
		name    string
		args    args
		res     res
		wantErr bool
	}{
		{
			name: "guild_has_been_configured",
			args: args{
				guildID: "895659000996200508",
				level:   2,
			},
		},
		{
			name: "guild_is_not_configured",
			args: args{
				guildID: "895659000996200123",
				level:   2,
			},
		},
		{
			name: "internal_server_error",
			args: args{
				guildID: "895659000996200508",
				level:   2,
			},
			res: res{
				err: gorm.ErrInvalidTransaction,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gclrRepo.EXPECT().DeleteOne(tt.args.guildID, tt.args.level).Return(tt.res.err).Times(1)

			if err := e.RemoveGuildLevelRoleConfig(tt.args.guildID, tt.args.level); (err != nil) != tt.wantErr {
				t.Errorf("Entity.RemoveGuildLevelRoleConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
