package entities

import (
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	baseerrs "github.com/defipod/mochi/pkg/model/errors"
	"github.com/defipod/mochi/pkg/repo"
	mock_telegram_association "github.com/defipod/mochi/pkg/repo/user_telegram_discord_association/mocks"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
)

func TestEntity_GetByTelegramUsername(t *testing.T) {
	type fields struct {
		repo *repo.Repo
	}
	type args struct {
		telegramUsername string
	}
	type want struct {
		res *response.GetLinkedTelegramResponse
		err error
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	teleRepo := mock_telegram_association.NewMockStore(ctrl)
	r := &repo.Repo{
		UserTelegramDiscordAssociation: teleRepo,
	}
	cfg := config.LoadTestConfig()
	log := logger.NewLogrusLogger()
	e := &Entity{
		cfg:  cfg,
		log:  log,
		repo: r,
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "get successfully telegramusername",
			fields: fields{
				repo: r,
			},
			args: args{
				telegramUsername: "anhnh",
			},
			want: want{
				res: &response.GetLinkedTelegramResponse{
					Data: &model.UserTelegramDiscordAssociation{
						TelegramUsername: "anhnh",
						DiscordID:        "12345",
					},
				},
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "not found telegramusername",
			fields: fields{
				repo: r,
			},
			args: args{
				telegramUsername: "abc",
			},
			want: want{
				res: nil,
				err: baseerrs.ErrRecordNotFound,
			},
			wantErr: true,
		},
	}

	teleRepo.EXPECT().GetOneByTelegramUsername("anhnh").Return(&model.UserTelegramDiscordAssociation{
		TelegramUsername: "anhnh",
		DiscordID:        "12345",
	}, nil).Times(1)
	teleRepo.EXPECT().GetOneByTelegramUsername("abc").Return(nil, gorm.ErrRecordNotFound).Times(1)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := e.GetByTelegramUsername(tt.args.telegramUsername)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.LinkUserTelegramWithDiscord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want.res) {
				t.Errorf("Entity.LinkUserTelegramWithDiscord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntity_LinkUserTelegramWithDiscord(t *testing.T) {

	type fields struct {
		repo *repo.Repo
	}
	type args struct {
		association request.LinkUserTelegramWithDiscordRequest
	}
	type want struct {
		res *response.LinkUserTelegramWithDiscordResponse
		err error
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	teleRepo := mock_telegram_association.NewMockStore(ctrl)
	r := &repo.Repo{
		UserTelegramDiscordAssociation: teleRepo,
	}
	cfg := config.LoadTestConfig()
	log := logger.NewLogrusLogger()
	e := &Entity{
		cfg:  cfg,
		log:  log,
		repo: r,
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "upsert telegram discord user association",
			fields: fields{
				repo: r,
			},
			args: args{
				association: request.LinkUserTelegramWithDiscordRequest{
					DiscordID:        "123456",
					TelegramUsername: "abc",
				},
			},
			want: want{
				res: &response.LinkUserTelegramWithDiscordResponse{
					Data: nil,
				},
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "duplicate telegram username",
			fields: fields{
				repo: r,
			},
			args: args{
				association: request.LinkUserTelegramWithDiscordRequest{
					DiscordID:        "123467",
					TelegramUsername: "abc",
				},
			},
			want: want{
				res: nil,
				err: errors.New("conflict username"),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			teleRepo.EXPECT().Upsert(&model.UserTelegramDiscordAssociation{
				DiscordID:        tt.args.association.DiscordID,
				TelegramUsername: tt.args.association.TelegramUsername,
			}).Return(tt.want.err).Times(1)
			got, err := e.LinkUserTelegramWithDiscord(tt.args.association)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.LinkUserTelegramWithDiscord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want.res) {
				t.Errorf("Entity.LinkUserTelegramWithDiscord() = %v, want %v", got, tt.want)
			}
		})
	}
}
