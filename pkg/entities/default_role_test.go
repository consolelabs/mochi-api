package entities

import (
	"reflect"
	"testing"
	"time"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/repo"
	mock_guildconfigdefaultroles "github.com/defipod/mochi/pkg/repo/guild_config_default_roles/mocks"
	"github.com/defipod/mochi/pkg/response"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func TestEntity_GetDefaultRoleByGuildID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mock repo
	gcdrRepo := mock_guildconfigdefaultroles.NewMockStore(ctrl)
	repo := &repo.Repo{
		GuildConfigDefaultRole: gcdrRepo,
	}

	// create entity
	cfg := config.LoadTestConfig()
	log := logger.NewLogrusLogger()
	e := &Entity{
		cfg:  cfg,
		log:  log,
		repo: repo,
	}

	createdAt, _ := time.Parse(time.RFC3339, "2022-08-26T04:19:40.192695+00")
	type res struct {
		data model.GuildConfigDefaultRole
		err  error
	}
	type args struct {
		guildID string
	}
	tests := []struct {
		name    string
		args    args
		res     res
		want    *response.DefaultRole
		wantErr bool
	}{
		{
			name: "guild_has_been_configured",
			args: args{
				guildID: "895659000996200508",
			},
			res: res{
				data: model.GuildConfigDefaultRole{
					ID: uuid.NullUUID{
						UUID:  uuid.New(),
						Valid: true,
					},
					GuildID:   "895659000996200508",
					RoleID:    "1003862842707017729",
					CreatedAt: createdAt,
				},
			},
			want: &response.DefaultRole{
				GuildID: "895659000996200508",
				RoleID:  "1003862842707017729",
			},
		},
		{
			name: "guild_is_not_configured",
			args: args{
				guildID: "895659000996200508",
			},
			res: res{
				err: gorm.ErrRecordNotFound,
			},
			want: &response.DefaultRole{
				GuildID: "895659000996200508",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gcdrRepo.EXPECT().GetAllByGuildID(tt.args.guildID).Return(tt.res.data, tt.res.err).Times(1)

			got, err := e.GetDefaultRoleByGuildID(tt.args.guildID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.GetDefaultRoleByGuildID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entity.GetDefaultRoleByGuildID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntity_CreateDefaultRoleConfig(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mock repo
	gcdrRepo := mock_guildconfigdefaultroles.NewMockStore(ctrl)
	repo := &repo.Repo{
		GuildConfigDefaultRole: gcdrRepo,
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
		roleID  string
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
				roleID:  "1003862842707017729",
			},
		},
		{
			name: "guild_is_not_configured",
			args: args{
				guildID: "895659000996200123",
				roleID:  "1003862842707017729",
			},
		},
		{
			name: "internal_server_error",
			args: args{
				guildID: "895659000996200123",
				roleID:  "1003862842707017729",
			},
			res: res{
				err: gorm.ErrInvalidTransaction,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gcdrRepo.EXPECT().CreateDefaultRoleIfNotExist(model.GuildConfigDefaultRole{
				GuildID: tt.args.guildID,
				RoleID:  tt.args.roleID,
			}).Return(tt.res.err).Times(1)

			if err := e.CreateDefaultRoleConfig(tt.args.guildID, tt.args.roleID); (err != nil) != tt.wantErr {
				t.Errorf("Entity.CreateDefaultRoleConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEntity_DeleteDefaultRoleConfig(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mock repo
	gcdrRepo := mock_guildconfigdefaultroles.NewMockStore(ctrl)
	repo := &repo.Repo{
		GuildConfigDefaultRole: gcdrRepo,
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
			},
		},
		{
			name: "guild_is_not_configured",
			args: args{
				guildID: "895659000996200123",
			},
		},
		{
			name: "guild_is_not_configured",
			args: args{
				guildID: "895659000996200123",
			},
			res: res{
				err: gorm.ErrInvalidTransaction,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gcdrRepo.EXPECT().DeleteByGuildID(tt.args.guildID).Return(tt.res.err).Times(1)

			if err := e.DeleteDefaultRoleConfig(tt.args.guildID); (err != nil) != tt.wantErr {
				t.Errorf("Entity.DeleteDefaultRoleConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
