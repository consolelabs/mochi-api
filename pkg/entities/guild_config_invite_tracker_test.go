package entities

import (
	"reflect"
	"testing"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/repo"
	mock_guildconfiginvitetracker "github.com/defipod/mochi/pkg/repo/guild_config_invite_tracker/mocks"
	"github.com/defipod/mochi/pkg/request"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func TestEntity_GetInviteTrackerLogChannel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mock repo
	gcitRepo := mock_guildconfiginvitetracker.NewMockStore(ctrl)
	repo := &repo.Repo{
		GuildConfigInviteTracker: gcitRepo,
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
		data *model.GuildConfigInviteTracker
		err  error
	}

	// create response
	guildConfigInviteTracker := model.GuildConfigInviteTracker{
		ID: uuid.NullUUID{
			UUID:  uuid.New(),
			Valid: true,
		},
		GuildID:   "962589711841525780",
		ChannelID: "964773702476656701",
	}

	type args struct {
		guildID string
	}
	tests := []struct {
		name    string
		args    args
		res     res
		want    *model.GuildConfigInviteTracker
		wantErr bool
	}{
		{
			name: "happy_case",
			args: args{
				guildID: "962589711841525780",
			},
			res: res{
				data: &guildConfigInviteTracker,
				err:  nil,
			},
			want:    &guildConfigInviteTracker,
			wantErr: false,
		},
		{
			name: "record_not_found",
			args: args{
				guildID: "962589711841525780",
			},
			res: res{
				data: nil,
				err:  gorm.ErrRecordNotFound,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gcitRepo.EXPECT().GetOne(tt.args.guildID).Return(tt.res.data, tt.res.err).Times(1)

			got, err := e.GetInviteTrackerLogChannel(tt.args.guildID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.GetInviteTrackerLogChannel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && tt.want != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entity.GetInviteTrackerLogChannel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntity_CreateOrUpdateInviteTrackerLogChannel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mock repo
	gcitRepo := mock_guildconfiginvitetracker.NewMockStore(ctrl)
	repo := &repo.Repo{
		GuildConfigInviteTracker: gcitRepo,
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
		req request.ConfigureInviteRequest
	}
	tests := []struct {
		name    string
		args    args
		res     res
		wantErr bool
	}{
		{
			name: "happy_case",
			args: args{
				req: request.ConfigureInviteRequest{
					GuildID:    "962589711841525780",
					LogChannel: "964773702476656701",
				},
			},
			res: res{
				err: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gcitRepo.EXPECT().Upsert(&model.GuildConfigInviteTracker{
				GuildID:    tt.args.req.GuildID,
				ChannelID:  tt.args.req.LogChannel,
				WebhookURL: model.JSONNullString{},
			}).Return(tt.res.err).Times(1)

			if err := e.CreateOrUpdateInviteTrackerLogChannel(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("Entity.CreateOrUpdateInviteTrackerLogChannel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
