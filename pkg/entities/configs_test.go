package entities

import (
	"errors"
	"reflect"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/defipod/mochi/pkg/cache"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/consts"
	"github.com/defipod/mochi/pkg/discordwallet"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/repo"
	mock_guild_config_gm_gn "github.com/defipod/mochi/pkg/repo/guild_config_gm_gn/mocks"
	mock_guild_config_level_role "github.com/defipod/mochi/pkg/repo/guild_config_level_role/mocks"
	mock_guild_config_repost_reaction "github.com/defipod/mochi/pkg/repo/guild_config_repost_reaction/mocks"
	mock_guild_config_token "github.com/defipod/mochi/pkg/repo/guild_config_token/mocks"
	mock_guildconfigvotechannel "github.com/defipod/mochi/pkg/repo/guild_config_vote_channel/mocks"
	mock_guildconfigwelcomechannel "github.com/defipod/mochi/pkg/repo/guild_config_welcome_channel/mocks"
	mock_message_repost_history "github.com/defipod/mochi/pkg/repo/message_repost_history/mocks"
	"github.com/defipod/mochi/pkg/repo/pg"
	mock_token "github.com/defipod/mochi/pkg/repo/token/mocks"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/service"
	"github.com/defipod/mochi/pkg/service/abi"
	"github.com/defipod/mochi/pkg/service/indexer"
	"github.com/defipod/mochi/pkg/service/marketplace"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func TestEntity_GetUserRoleByLevel(t *testing.T) {
	type fields struct {
		repo        *repo.Repo
		store       repo.Store
		log         logger.Logger
		dcwallet    discordwallet.IDiscordWallet
		discord     *discordgo.Session
		cache       cache.Cache
		svc         *service.Service
		cfg         config.Config
		indexer     indexer.Service
		abi         abi.Service
		marketplace marketplace.Service
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := config.Config{
		DBUser: "postgres",
		DBPass: "postgres",
		DBHost: "localhost",
		DBPort: "5434",
		DBName: "mochi_local",

		InDiscordWalletMnemonic: "holiday frequent toy bachelor auto use style result recycle crumble glue blouse",
		FantomRPC:               "sample",
		FantomScan:              "sample",
		FantomScanAPIKey:        "sample",

		EthereumRPC:        "sample",
		EthereumScan:       "sample",
		EthereumScanAPIKey: "sample",

		BscRPC:        "sample",
		BscScan:       "sample",
		BscScanAPIKey: "sample",

		DiscordToken: "sample",

		RedisURL: "redis://localhost:6379/0",
	}

	s := pg.NewPostgresStore(&cfg)
	r := pg.NewRepo(s.DB())
	guildConfigLR := mock_guild_config_level_role.NewMockStore(ctrl)

	r.GuildConfigLevelRole = guildConfigLR

	type args struct {
		guildID string
		level   int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "successful",
			fields: fields{
				repo: r,
			},
			args: args{
				guildID: "863278424433229854",
				level:   3,
			},
			want:    "abc",
			wantErr: false,
		},
		{
			name: "invalid guild id",
			fields: fields{
				repo: r,
			},
			args: args{
				guildID: "863278424433229abc",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "not found lr with highest level",
			fields: fields{
				repo: r,
			},
			args: args{
				guildID: "863278424433229854",
				level:   1,
			},
			want:    "",
			wantErr: true,
		},
	}
	config := model.GuildConfigLevelRole{
		GuildID: "863278424433229854",
		Level:   5,
		RoleID:  "abc",
		LevelConfig: &model.ConfigXpLevel{
			Level: 5,
			MinXP: 100,
		},
	}
	guildConfigLR.EXPECT().GetHighest("863278424433229854", 3).Return(&config, nil).AnyTimes()
	guildConfigLR.EXPECT().GetHighest("863278424433229abc", gomock.Any()).Return(nil, errors.New("invalid guild id")).AnyTimes()
	guildConfigLR.EXPECT().GetHighest(gomock.Any(), 1).Return(nil, gorm.ErrRecordNotFound)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Entity{
				repo:        tt.fields.repo,
				store:       tt.fields.store,
				log:         tt.fields.log,
				dcwallet:    tt.fields.dcwallet,
				discord:     tt.fields.discord,
				cache:       tt.fields.cache,
				svc:         tt.fields.svc,
				cfg:         tt.fields.cfg,
				indexer:     tt.fields.indexer,
				abi:         tt.fields.abi,
				marketplace: tt.fields.marketplace,
			}
			got, err := e.GetUserRoleByLevel(tt.args.guildID, tt.args.level)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.GetUserRoleByLevel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entity.GetUserRoleByLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntity_GetGmConfig(t *testing.T) {
	type fields struct {
		repo        *repo.Repo
		store       repo.Store
		log         logger.Logger
		dcwallet    discordwallet.IDiscordWallet
		discord     *discordgo.Session
		cache       cache.Cache
		svc         *service.Service
		cfg         config.Config
		indexer     indexer.Service
		abi         abi.Service
		marketplace marketplace.Service
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cfg := config.LoadTestConfig()
	s := pg.NewPostgresStore(&cfg)
	r := pg.NewRepo(s.DB())
	guildConfigGmGn := mock_guild_config_gm_gn.NewMockStore(ctrl)

	r.GuildConfigGmGn = guildConfigGmGn

	type args struct {
		guildID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.GuildConfigGmGn
		wantErr bool
	}{
		{
			name: "Guild has config GM channel",
			fields: fields{
				repo: r,
			},
			args: args{
				guildID: "552427722551459840",
			},
			want:    &model.GuildConfigGmGn{GuildID: "552427722551459840", ChannelID: "701029345795375114"},
			wantErr: false,
		},
		{
			name: "Guild does not have config GM channel",
			fields: fields{
				repo: r,
			},
			args: args{
				guildID: "not_have_config_gm_channel",
			},
			want:    nil,
			wantErr: false,
		},
	}

	guildConfigGmGn.EXPECT().GetByGuildID("not_have_config_gm_channel").Return(nil, nil).AnyTimes()

	configGmGn := model.GuildConfigGmGn{
		GuildID:   "552427722551459840",
		ChannelID: "701029345795375114",
	}
	guildConfigGmGn.EXPECT().GetByGuildID("552427722551459840").Return(&configGmGn, nil).AnyTimes()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Entity{
				repo:        tt.fields.repo,
				store:       tt.fields.store,
				log:         tt.fields.log,
				dcwallet:    tt.fields.dcwallet,
				discord:     tt.fields.discord,
				cache:       tt.fields.cache,
				svc:         tt.fields.svc,
				cfg:         tt.fields.cfg,
				indexer:     tt.fields.indexer,
				abi:         tt.fields.abi,
				marketplace: tt.fields.marketplace,
			}
			got, err := e.GetGmConfig(tt.args.guildID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.GetGmConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entity.GetGmConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntity_UpsertGmConfig(t *testing.T) {
	type fields struct {
		repo        *repo.Repo
		store       repo.Store
		log         logger.Logger
		dcwallet    discordwallet.IDiscordWallet
		discord     *discordgo.Session
		cache       cache.Cache
		svc         *service.Service
		cfg         config.Config
		indexer     indexer.Service
		abi         abi.Service
		marketplace marketplace.Service
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cfg := config.LoadTestConfig()
	s := pg.NewPostgresStore(&cfg)
	r := pg.NewRepo(s.DB())
	guildConfigGmGn := mock_guild_config_gm_gn.NewMockStore(ctrl)

	r.GuildConfigGmGn = guildConfigGmGn

	type args struct {
		req request.UpsertGmConfigRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Upsert config gm successfully",
			fields: fields{
				repo: r,
			},
			args: args{
				req: request.UpsertGmConfigRequest{
					GuildID:   "552427722551459840",
					ChannelID: "701029345795375114",
				},
			},
			wantErr: false,
		},
	}

	guildConfigGmGn.EXPECT().UpsertOne(&model.GuildConfigGmGn{
		GuildID:   "552427722551459840",
		ChannelID: "701029345795375114",
	}).Return(nil).AnyTimes()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Entity{
				repo:        tt.fields.repo,
				store:       tt.fields.store,
				log:         tt.fields.log,
				dcwallet:    tt.fields.dcwallet,
				discord:     tt.fields.discord,
				cache:       tt.fields.cache,
				svc:         tt.fields.svc,
				cfg:         tt.fields.cfg,
				indexer:     tt.fields.indexer,
				abi:         tt.fields.abi,
				marketplace: tt.fields.marketplace,
			}
			if err := e.UpsertGmConfig(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("Entity.UpsertGmConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEntity_GetWelcomeChannelConfig(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mock repo
	gcwcRepo := mock_guildconfigwelcomechannel.NewMockStore(ctrl)
	repo := &repo.Repo{
		GuildConfigWelcomeChannel: gcwcRepo,
	}

	// create entity
	cfg := config.LoadTestConfig()
	log := logger.NewLogrusLogger()
	e := &Entity{
		cfg:  cfg,
		log:  log,
		repo: repo,
	}

	guildConfigWelcomeChannel := model.GuildConfigWelcomeChannel{
		ID: uuid.NullUUID{
			UUID:  uuid.New(),
			Valid: true,
		},
		GuildID:        "895659000996200508",
		ChannelID:      "1016919074221064256",
		WelcomeMessage: "Welcome to the guild!",
	}

	type res struct {
		data *model.GuildConfigWelcomeChannel
		err  error
	}
	type args struct {
		guildID string
	}
	tests := []struct {
		name    string
		args    args
		res     res
		want    *model.GuildConfigWelcomeChannel
		wantErr bool
	}{
		{
			name: "guild_is_configured",
			args: args{
				guildID: "895659000996200508",
			},
			res: res{
				data: &guildConfigWelcomeChannel,
			},
			want:    &guildConfigWelcomeChannel,
			wantErr: false,
		},
		{
			name: "guild_is_not_configured",
			args: args{
				guildID: "895659000996200123",
			},
			res: res{
				data: nil,
				err:  gorm.ErrRecordNotFound,
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gcwcRepo.EXPECT().GetByGuildID(tt.args.guildID).Return(tt.res.data, tt.res.err).Times(1)

			got, err := e.GetWelcomeChannelConfig(tt.args.guildID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.GetWelcomeChannelConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entity.GetWelcomeChannelConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntity_UpsertWelcomeChannelConfig(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mock repo
	gcwcRepo := mock_guildconfigwelcomechannel.NewMockStore(ctrl)
	repo := &repo.Repo{
		GuildConfigWelcomeChannel: gcwcRepo,
	}

	// create entity
	cfg := config.LoadTestConfig()
	log := logger.NewLogrusLogger()
	e := &Entity{
		cfg:  cfg,
		log:  log,
		repo: repo,
	}

	guildConfigWelcomeChannel := model.GuildConfigWelcomeChannel{
		ID: uuid.NullUUID{
			UUID:  uuid.New(),
			Valid: true,
		},
		GuildID:        "895659000996200508",
		ChannelID:      "1016919074221064256",
		WelcomeMessage: "Welcome to the guild!",
	}
	guildConfigWelcomeChannelForNewConfig := model.GuildConfigWelcomeChannel{
		ID: uuid.NullUUID{
			UUID:  uuid.New(),
			Valid: true,
		},
		GuildID:        "895659000996200508",
		ChannelID:      "1016919074221064256",
		WelcomeMessage: "Greetings $name :wave: Welcome to the guild! Hope you enjoy your stay.",
	}

	type res struct {
		data *model.GuildConfigWelcomeChannel
		err  error
	}
	type args struct {
		req request.UpsertWelcomeConfigRequest
	}
	tests := []struct {
		name    string
		args    args
		res     []res
		want    *model.GuildConfigWelcomeChannel
		wantErr bool
	}{
		{
			name: "update_new_msg",
			args: args{
				req: request.UpsertWelcomeConfigRequest{
					GuildID:    "895659000996200508",
					ChannelID:  "1016919074221064256",
					WelcomeMsg: "Welcome to the guild!",
				},
			},
			res: []res{
				{
					data: &guildConfigWelcomeChannel,
				},
				{},
			},
			want:    &guildConfigWelcomeChannel,
			wantErr: false,
		},
		{
			name: "update_with_empty_msg",
			args: args{
				req: request.UpsertWelcomeConfigRequest{
					GuildID:   "895659000996200508",
					ChannelID: "1016919074221064256",
				},
			},
			res: []res{
				{
					data: &guildConfigWelcomeChannel,
					err:  nil,
				},
				{
					data: &guildConfigWelcomeChannel,
					err:  nil,
				},
			},
			want:    &guildConfigWelcomeChannel,
			wantErr: false,
		},
		{
			name: "create_new_config_with_empty_msg",
			args: args{
				req: request.UpsertWelcomeConfigRequest{
					GuildID:   "895659000996200123",
					ChannelID: "1016919074221064256",
				},
			},
			res: []res{
				{
					data: &guildConfigWelcomeChannelForNewConfig,
				},
				{
					err: gorm.ErrRecordNotFound,
				},
			},
			want:    &guildConfigWelcomeChannelForNewConfig,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			upsertData := &model.GuildConfigWelcomeChannel{
				GuildID:        tt.args.req.GuildID,
				ChannelID:      tt.args.req.ChannelID,
				WelcomeMessage: tt.args.req.WelcomeMsg,
			}
			if tt.args.req.WelcomeMsg == "" {
				gcwcRepo.EXPECT().GetByGuildID(tt.args.req.GuildID).Return(tt.res[1].data, tt.res[1].err).Times(1)
				upsertData.WelcomeMessage = "Greetings $name :wave: Welcome to the guild! Hope you enjoy your stay."
				if tt.res[1].data != nil {
					upsertData.WelcomeMessage = tt.res[1].data.WelcomeMessage
				}
			}
			gcwcRepo.EXPECT().UpsertOne(upsertData).Return(tt.res[0].data, tt.res[0].err).Times(1)

			got, err := e.UpsertWelcomeChannelConfig(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.UpsertWelcomeChannelConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entity.UpsertWelcomeChannelConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntity_DeleteWelcomeChannelConfig(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mock repo
	gcwcRepo := mock_guildconfigwelcomechannel.NewMockStore(ctrl)
	repo := &repo.Repo{
		GuildConfigWelcomeChannel: gcwcRepo,
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
		req request.DeleteWelcomeConfigRequest
	}
	tests := []struct {
		name    string
		args    args
		res     res
		wantErr bool
	}{
		{
			name: "guild_is_configured",
			args: args{
				req: request.DeleteWelcomeConfigRequest{
					GuildID: "895659000996200508",
				},
			},
			wantErr: false,
		},
		{
			name: "guild_is_not_configured",
			args: args{
				req: request.DeleteWelcomeConfigRequest{
					GuildID: "895659000996200123",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gcwcRepo.EXPECT().DeleteOne(&model.GuildConfigWelcomeChannel{
				GuildID: tt.args.req.GuildID,
			}).Return(tt.res.err).Times(1)

			if err := e.DeleteWelcomeChannelConfig(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("Entity.DeleteWelcomeChannelConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEntity_GetVoteChannelConfig(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mock repo
	gcvcRepo := mock_guildconfigvotechannel.NewMockStore(ctrl)
	repo := &repo.Repo{
		GuildConfigVoteChannel: gcvcRepo,
	}

	// create entity
	cfg := config.LoadTestConfig()
	log := logger.NewLogrusLogger()
	e := &Entity{
		cfg:  cfg,
		log:  log,
		repo: repo,
	}

	guildConfigVoteChannel := model.GuildConfigVoteChannel{
		ID: uuid.NullUUID{
			UUID:  uuid.New(),
			Valid: true,
		},
		GuildID:   "895659000996200508",
		ChannelID: "1016919074221064256",
	}

	type res struct {
		data *model.GuildConfigVoteChannel
		err  error
	}
	type args struct {
		guildID string
	}
	tests := []struct {
		name    string
		args    args
		res     res
		want    *model.GuildConfigVoteChannel
		wantErr bool
	}{
		{
			name: "guild_is_configured",
			args: args{
				guildID: "895659000996200508",
			},
			res: res{
				data: &guildConfigVoteChannel,
			},
			want:    &guildConfigVoteChannel,
			wantErr: false,
		},
		{
			name: "guild_is_not_configured",
			args: args{
				guildID: "895659000996200123",
			},
			res: res{
				data: nil,
				err:  gorm.ErrRecordNotFound,
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gcvcRepo.EXPECT().GetByGuildID(tt.args.guildID).Return(tt.res.data, tt.res.err).Times(1)

			got, err := e.GetVoteChannelConfig(tt.args.guildID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.GetVoteChannelConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entity.GetVoteChannelConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntity_UpsertVoteChannelConfig(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mock repo
	gcvcRepo := mock_guildconfigvotechannel.NewMockStore(ctrl)
	repo := &repo.Repo{
		GuildConfigVoteChannel: gcvcRepo,
	}

	// create entity
	cfg := config.LoadTestConfig()
	log := logger.NewLogrusLogger()
	e := &Entity{
		cfg:  cfg,
		log:  log,
		repo: repo,
	}

	guildConfigVoteChannel := model.GuildConfigVoteChannel{
		ID: uuid.NullUUID{
			UUID:  uuid.New(),
			Valid: true,
		},
		GuildID:   "895659000996200508",
		ChannelID: "1016919074221064256",
	}

	type res struct {
		data *model.GuildConfigVoteChannel
		err  error
	}
	type args struct {
		req request.UpsertVoteChannelConfigRequest
	}
	tests := []struct {
		name    string
		args    args
		res     res
		want    *model.GuildConfigVoteChannel
		wantErr bool
	}{
		{
			name: "guild_is_configured",
			args: args{
				req: request.UpsertVoteChannelConfigRequest{
					GuildID:   "895659000996200508",
					ChannelID: "1016919074221064256",
				},
			},
			res: res{
				data: &guildConfigVoteChannel,
			},
			want:    &guildConfigVoteChannel,
			wantErr: false,
		},
		{
			name: "guild_is_not_configured",
			args: args{
				req: request.UpsertVoteChannelConfigRequest{
					GuildID:   "895659000996200123",
					ChannelID: "1016919074221064256",
				},
			},
			res: res{
				data: &guildConfigVoteChannel,
			},
			want:    &guildConfigVoteChannel,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gcvcRepo.EXPECT().UpsertOne(&model.GuildConfigVoteChannel{
				GuildID:   tt.args.req.GuildID,
				ChannelID: tt.args.req.ChannelID,
			}).Return(tt.res.data, tt.res.err).Times(1)

			got, err := e.UpsertVoteChannelConfig(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.UpsertVoteChannelConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entity.UpsertVoteChannelConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntity_DeleteVoteChannelConfig(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mock repo
	gcvcRepo := mock_guildconfigvotechannel.NewMockStore(ctrl)
	repo := &repo.Repo{
		GuildConfigVoteChannel: gcvcRepo,
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
		req request.DeleteVoteChannelConfigRequest
	}
	tests := []struct {
		name    string
		res     res
		args    args
		wantErr bool
	}{
		{
			name: "guild_is_configured",
			args: args{
				req: request.DeleteVoteChannelConfigRequest{
					GuildID: "895659000996200508",
				},
			},
			wantErr: false,
		},
		{
			name: "guild_is_not_configured",
			args: args{
				req: request.DeleteVoteChannelConfigRequest{
					GuildID: "895659000996200123",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gcvcRepo.EXPECT().DeleteOne(&model.GuildConfigVoteChannel{
				GuildID: tt.args.req.GuildID,
			}).Return(tt.res.err).Times(1)

			if err := e.DeleteVoteChannelConfig(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("Entity.DeleteVoteChannelConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEntity_GetGuildRepostReactionConfigs(t *testing.T) {
	type fields struct {
		repo        *repo.Repo
		store       repo.Store
		log         logger.Logger
		dcwallet    discordwallet.IDiscordWallet
		discord     *discordgo.Session
		cache       cache.Cache
		svc         *service.Service
		cfg         config.Config
		indexer     indexer.Service
		abi         abi.Service
		marketplace marketplace.Service
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cfg := config.LoadTestConfig()
	s := pg.NewPostgresStore(&cfg)
	r := pg.NewRepo(s.DB())
	cfgRepost := mock_guild_config_repost_reaction.NewMockStore(ctrl)
	r.GuildConfigRepostReaction = cfgRepost

	type args struct {
		guildID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.GuildConfigRepostReaction
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Get guild repost reaction config successfully",
			fields: fields{
				repo: r,
			},
			args: args{
				guildID: "552427722551459840",
			},
			want: []model.GuildConfigRepostReaction{
				{
					GuildID:         "552427722551459840",
					Emoji:           "test",
					Quantity:        1,
					RepostChannelID: "test",
				},
			},
			wantErr: false,
		},
		{
			name: "Get guild repost reaction not have config",
			fields: fields{
				repo: r,
			},
			args: args{
				guildID: "not_exist",
			},
			want: []model.GuildConfigRepostReaction{
				{
					GuildID: "not_exist",
				},
			},
			wantErr: false,
		},
	}

	cfgRepost.EXPECT().GetByGuildIDAndReactionType("552427722551459840", consts.ReactionTypeMessage).Return([]model.GuildConfigRepostReaction{{
		GuildID:         "552427722551459840",
		Emoji:           "test",
		Quantity:        1,
		RepostChannelID: "test",
	},
	}, nil).AnyTimes()

	cfgRepost.EXPECT().GetByGuildIDAndReactionType("not_exist", consts.ReactionTypeMessage).Return([]model.GuildConfigRepostReaction{{
		GuildID: "not_exist",
	},
	}, nil).AnyTimes()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Entity{
				repo:        tt.fields.repo,
				store:       tt.fields.store,
				log:         tt.fields.log,
				dcwallet:    tt.fields.dcwallet,
				discord:     tt.fields.discord,
				cache:       tt.fields.cache,
				svc:         tt.fields.svc,
				cfg:         tt.fields.cfg,
				indexer:     tt.fields.indexer,
				abi:         tt.fields.abi,
				marketplace: tt.fields.marketplace,
			}
			got, err := e.GetGuildRepostReactionConfigs(tt.args.guildID, consts.ReactionTypeMessage)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.GetGuildRepostReactionConfigs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entity.GetGuildRepostReactionConfigs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntity_ConfigRepostReaction(t *testing.T) {
	type fields struct {
		repo        *repo.Repo
		store       repo.Store
		log         logger.Logger
		dcwallet    discordwallet.IDiscordWallet
		discord     *discordgo.Session
		cache       cache.Cache
		svc         *service.Service
		cfg         config.Config
		indexer     indexer.Service
		abi         abi.Service
		marketplace marketplace.Service
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cfg := config.LoadTestConfig()
	s := pg.NewPostgresStore(&cfg)
	r := pg.NewRepo(s.DB())
	cfgRepost := mock_guild_config_repost_reaction.NewMockStore(ctrl)
	r.GuildConfigRepostReaction = cfgRepost

	type args struct {
		req request.ConfigRepostRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Upsert repost reaction successfully",
			fields: fields{
				repo: r,
			},
			args: args{
				req: request.ConfigRepostRequest{
					GuildID:         "552427722551459840",
					Emoji:           "test",
					Quantity:        1,
					RepostChannelID: "test",
				},
			},
		},
	}

	cfgRepost.EXPECT().UpsertOne(model.GuildConfigRepostReaction{
		GuildID:         "552427722551459840",
		Emoji:           "test",
		Quantity:        1,
		RepostChannelID: "test",
		ReactionType:    "message"}).Return(nil).AnyTimes()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Entity{
				repo:        tt.fields.repo,
				store:       tt.fields.store,
				log:         tt.fields.log,
				dcwallet:    tt.fields.dcwallet,
				discord:     tt.fields.discord,
				cache:       tt.fields.cache,
				svc:         tt.fields.svc,
				cfg:         tt.fields.cfg,
				indexer:     tt.fields.indexer,
				abi:         tt.fields.abi,
				marketplace: tt.fields.marketplace,
			}
			if err := e.ConfigRepostReaction(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("Entity.ConfigRepostReaction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEntity_RemoveGuildRepostReactionConfig(t *testing.T) {
	type fields struct {
		repo        *repo.Repo
		store       repo.Store
		log         logger.Logger
		dcwallet    discordwallet.IDiscordWallet
		discord     *discordgo.Session
		cache       cache.Cache
		svc         *service.Service
		cfg         config.Config
		indexer     indexer.Service
		abi         abi.Service
		marketplace marketplace.Service
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cfg := config.LoadTestConfig()
	s := pg.NewPostgresStore(&cfg)
	r := pg.NewRepo(s.DB())
	cfgRepost := mock_guild_config_repost_reaction.NewMockStore(ctrl)
	r.GuildConfigRepostReaction = cfgRepost

	type args struct {
		guildID string
		emoji   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Remove repost reaction successfully",
			fields: fields{
				repo: r,
			},
			args: args{
				guildID: "552427722551459840",
				emoji:   "test",
			},
			wantErr: false,
		},
	}
	cfgRepost.EXPECT().DeleteOne("552427722551459840", "test").Return(nil).AnyTimes()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Entity{
				repo:        tt.fields.repo,
				store:       tt.fields.store,
				log:         tt.fields.log,
				dcwallet:    tt.fields.dcwallet,
				discord:     tt.fields.discord,
				cache:       tt.fields.cache,
				svc:         tt.fields.svc,
				cfg:         tt.fields.cfg,
				indexer:     tt.fields.indexer,
				abi:         tt.fields.abi,
				marketplace: tt.fields.marketplace,
			}
			if err := e.RemoveGuildRepostReactionConfig(tt.args.guildID, tt.args.emoji); (err != nil) != tt.wantErr {
				t.Errorf("Entity.RemoveGuildRepostReactionConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEntity_EditMessageRepost(t *testing.T) {
	type fields struct {
		repo        *repo.Repo
		store       repo.Store
		log         logger.Logger
		dcwallet    discordwallet.IDiscordWallet
		discord     *discordgo.Session
		cache       cache.Cache
		svc         *service.Service
		cfg         config.Config
		indexer     indexer.Service
		abi         abi.Service
		marketplace marketplace.Service
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cfg := config.LoadTestConfig()
	s := pg.NewPostgresStore(&cfg)
	r := pg.NewRepo(s.DB())
	log := logger.NewLogrusLogger()
	msgRepostHistory := mock_message_repost_history.NewMockStore(ctrl)
	r.MessageRepostHistory = msgRepostHistory

	type args struct {
		req *request.EditMessageRepostRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Edit message repost successfully",
			fields: fields{
				repo: r,
			},
			args: args{
				req: &request.EditMessageRepostRequest{
					GuildID:         "552427722551459840",
					OriginMessageID: "origin_msg",
					RepostMessageID: "repost_msg",
					OriginChannelID: "origin_channel",
					RepostChannelID: "repost_channel",
				},
			},
			wantErr: false,
		},
		{
			name: "Edit message repost failed",
			fields: fields{
				repo: r,
				log:  log,
			},
			args: args{
				req: &request.EditMessageRepostRequest{
					GuildID:         "fail_guild_id",
					OriginMessageID: "origin_msg",
					RepostMessageID: "repost_msg",
					OriginChannelID: "origin_channel",
					RepostChannelID: "repost_channel",
				},
			},
			wantErr: true,
		},
	}
	msgRepostHistory.EXPECT().EditMessageRepost(&request.EditMessageRepostRequest{
		GuildID:         "552427722551459840",
		OriginMessageID: "origin_msg",
		RepostMessageID: "repost_msg",
		OriginChannelID: "origin_channel",
		RepostChannelID: "repost_channel",
	}).Return(nil).AnyTimes()
	msgRepostHistory.EXPECT().EditMessageRepost(&request.EditMessageRepostRequest{
		GuildID:         "fail_guild_id",
		OriginMessageID: "origin_msg",
		RepostMessageID: "repost_msg",
		OriginChannelID: "origin_channel",
		RepostChannelID: "repost_channel",
	}).Return(errors.New("edit repost msg failed")).AnyTimes()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Entity{
				repo:        tt.fields.repo,
				store:       tt.fields.store,
				log:         tt.fields.log,
				dcwallet:    tt.fields.dcwallet,
				discord:     tt.fields.discord,
				cache:       tt.fields.cache,
				svc:         tt.fields.svc,
				cfg:         tt.fields.cfg,
				indexer:     tt.fields.indexer,
				abi:         tt.fields.abi,
				marketplace: tt.fields.marketplace,
			}
			if err := e.EditMessageRepost(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("Entity.EditMessageRepost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEntity_GetGuildTokens(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := config.LoadTestConfig()

	log := logger.NewLogrusLogger()
	svc, _ := service.NewService(cfg, log)
	mockToken := mock_token.NewMockStore(ctrl)
	mockGuildConfigToken := mock_guild_config_token.NewMockStore(ctrl)

	s := pg.NewPostgresStore(&cfg)
	r := pg.NewRepo(s.DB())
	r.Token = mockToken
	r.GuildConfigToken = mockGuildConfigToken

	e := &Entity{
		repo: r,
		log:  log,
		svc:  svc,
		cfg:  cfg,
	}

	defaultTokens := []model.Token{
		{
			ID:      1,
			Address: "0x7aCeE5D0acC520faB33b3Ea25D4FEEF1FfebDE73",
			Symbol:  "btc",
			ChainID: 1,
			Name:    "Bitcoin",
		},
		{
			ID:      2,
			Address: "0x7aCeE5D0acC520faB33b3Ea25D4FEEF1FfebDEA4",
			Symbol:  "ftm",
			ChainID: 1,
			Name:    "Fantom",
		},
	}

	type args struct {
		guildID string
	}
	tests := []struct {
		name        string
		args        args
		guildTokens []model.GuildConfigToken
		want        []model.Token
		wantErr     bool
	}{
		{
			name: "success - default tokens",
			args: args{
				guildID: "",
			},
			want:    defaultTokens,
			wantErr: false,
		},
		{
			name: "success - guild tokens",
			args: args{
				guildID: "863278424433229854",
			},
			guildTokens: []model.GuildConfigToken{
				{
					Token: &model.Token{
						ID:      3,
						Address: "0x7aCeE5D0acC520faB33b3Ea25D4FEEF1FfebD2AA",
						Symbol:  "Cake",
						ChainID: 1,
						Name:    "Pancake",
					},
				},
			},
			want: []model.Token{
				{
					ID:      3,
					Address: "0x7aCeE5D0acC520faB33b3Ea25D4FEEF1FfebD2AA",
					Symbol:  "Cake",
					ChainID: 1,
					Name:    "Pancake",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockToken.EXPECT().GetDefaultTokens().Return(defaultTokens, nil).AnyTimes()
			mockGuildConfigToken.EXPECT().GetByGuildID(tt.args.guildID).Return(tt.guildTokens, nil).AnyTimes()

			got, err := e.GetGuildTokens(tt.args.guildID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.GetGuildTokens() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entity.GetGuildTokens() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntity_UpsertGuildTokenConfig(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := config.LoadTestConfig()

	log := logger.NewLogrusLogger()
	svc, _ := service.NewService(cfg, log)
	mockToken := mock_token.NewMockStore(ctrl)
	mockGuildConfigToken := mock_guild_config_token.NewMockStore(ctrl)

	s := pg.NewPostgresStore(&cfg)
	r := pg.NewRepo(s.DB())
	r.Token = mockToken
	r.GuildConfigToken = mockGuildConfigToken

	e := &Entity{
		repo: r,
		log:  log,
		svc:  svc,
		cfg:  cfg,
	}
	type args struct {
		req request.UpsertGuildTokenConfigRequest
	}
	tests := []struct {
		name       string
		args       args
		tokenFound model.Token
		tokenErr   error
		wantErr    bool
	}{
		{
			name: "success - add token",
			args: args{
				req: request.UpsertGuildTokenConfigRequest{
					GuildID: "863278424433229854",
					Symbol:  "btc",
					Active:  true,
				},
			},
			tokenFound: model.Token{
				ID:      3,
				Address: "0x7aCeE5D0acC520faB33acwwqa25D4FEEF1FfebD2AA",
				Symbol:  "Bitcoin",
				ChainID: 1,
				Name:    "Bitcoin",
			},
			tokenErr: nil,
			wantErr:  false,
		},
		{
			name: "success - remove existed token",
			args: args{
				req: request.UpsertGuildTokenConfigRequest{
					GuildID: "863278424433229854",
					Symbol:  "btc",
					Active:  false,
				},
			},
			tokenFound: model.Token{
				ID:      3,
				Address: "0x7aCeE5D0acC520faB33acwwqa25D4FEEF1FfebD2AA",
				Symbol:  "Bitcoin",
				ChainID: 1,
				Name:    "Bitcoin",
			},
			tokenErr: nil,
			wantErr:  false,
		},
		{
			name: "failed - remove non existed token",
			args: args{
				req: request.UpsertGuildTokenConfigRequest{
					GuildID: "863278424433229854",
					Symbol:  "abcd",
					Active:  true,
				},
			},
			tokenErr: errors.New("record not found"),
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockToken.EXPECT().GetBySymbol(tt.args.req.Symbol, true).Return(tt.tokenFound, tt.tokenErr).AnyTimes()
			mockGuildConfigToken.EXPECT().UpsertMany([]model.GuildConfigToken{{
				GuildID: tt.args.req.GuildID,
				TokenID: tt.tokenFound.ID,
				Active:  tt.args.req.Active,
			}}).Return(nil).AnyTimes()

			if err := e.UpsertGuildTokenConfig(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("Entity.UpsertGuildTokenConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
