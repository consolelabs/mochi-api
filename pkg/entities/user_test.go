package entities

import (
	"errors"
	"reflect"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/defipod/mochi/pkg/cache"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/discordwallet"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/repo"
	mock_config_xp_level "github.com/defipod/mochi/pkg/repo/config_xp_level/mocks"
	mock_discord_guilds "github.com/defipod/mochi/pkg/repo/discord_guilds/mocks"
	mock_discord_user_gm_streak "github.com/defipod/mochi/pkg/repo/discord_user_gm_streak/mocks"
	mock_guild_user_activity_log "github.com/defipod/mochi/pkg/repo/guild_user_activity_log/mocks"
	mock_guild_user_xp "github.com/defipod/mochi/pkg/repo/guild_user_xp/mocks"
	"github.com/defipod/mochi/pkg/repo/pg"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/service"
	"github.com/defipod/mochi/pkg/service/abi"
	mock_discord "github.com/defipod/mochi/pkg/service/discord/mocks"
	"github.com/defipod/mochi/pkg/service/indexer"
	"github.com/defipod/mochi/pkg/service/marketplace"
	mock_processor "github.com/defipod/mochi/pkg/service/processor/mocks"
	"github.com/golang/mock/gomock"
)

func TestEntity_SendGiftXp(t *testing.T) {
	type fields struct {
		repo     *repo.Repo
		store    repo.Store
		log      logger.Logger
		dcwallet discordwallet.IDiscordWallet
		discord  *discordgo.Session
		cache    cache.Cache
		svc      *service.Service
		cfg      config.Config
	}
	type args struct {
		request.GiftXPRequest
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
		FantomRPC:               "https://rpc.ftm.tools",
		FantomScan:              "https://api.ftmscan.com/api?",
		FantomScanAPIKey:        "XEKSVDF5VWQDY5VY6ZNT6AK9QPQRH483EF",

		EthereumRPC:        "https://mainnet.infura.io/v3/5b389eb75c514cf6b1711d70084b0114",
		EthereumScan:       "https://api.etherscan.io/api?",
		EthereumScanAPIKey: "SM5BHYSNIRZ1HEWJ1JPHVTMJS95HRA6DQF",

		BscRPC:        "https://bsc-dataseed.binance.org",
		BscScan:       "https://api.bscscan.com/api?",
		BscScanAPIKey: "VTKF4RG4HP6WXQ5QTAJ8MHDDIUFYD6VZHC",

		DiscordToken: "OTcxNjMyNDMzMjk0MzQ4Mjg5.G5BEgF.rv-16ZuTzzqOv2W76OljymFxxnNpjVjCnOkn98",

		RedisURL: "redis://localhost:6379/0",
	}

	s := pg.NewPostgresStore(&cfg)
	r := pg.NewRepo(s.DB())
	uXp := mock_guild_user_xp.NewMockStore(ctrl)
	uLog := mock_guild_user_activity_log.NewMockStore(ctrl)
	discordMock := mock_discord.NewMockService(ctrl)

	r.GuildUserXP = uXp
	r.GuildUserActivityLog = uLog
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.HandleUserActivityResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test gift xp successfully",
			fields: fields{
				repo: r,
				log:  logger.NewLogrusLogger(),
				svc: &service.Service{
					Discord: discordMock,
				},
			},
			args: args{
				GiftXPRequest: request.GiftXPRequest{
					GuildID:       "552427722551459840",
					UserDiscordID: "973069332034752522",
					XPAmount:      10,
				},
			},
			want: &response.HandleUserActivityResponse{
				GuildID:      "552427722551459840",
				UserID:       "973069332034752522",
				Action:       "gifted",
				CurrentXP:    20,
				CurrentLevel: 0,
				LevelUp:      false,
				AddedXP:      10,
			},
			wantErr: false,
		},
		{
			name: "case user not exist in server, cannot gift xp to this user",
			fields: fields{
				repo: r,
				log:  logger.NewLogrusLogger(),
			},
			args: args{
				GiftXPRequest: request.GiftXPRequest{
					GuildID:       "abc",
					UserDiscordID: "abc",
					XPAmount:      10,
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	guildUserXpParam := model.GuildUserXP{
		GuildID: "552427722551459840",
		UserID:  "973069332034752522",
		TotalXP: 20,
		Level:   0,
		Guild:   &model.DiscordGuild{},
	}
	guildUserActivityLog := model.GuildUserActivityLog{
		GuildID:      "552427722551459840",
		UserID:       "973069332034752522",
		EarnedXP:     10,
		ActivityName: "gifted",
	}

	// case sucess send gift xp
	uXp.EXPECT().GetOne("552427722551459840", "973069332034752522").Return(&guildUserXpParam, nil).AnyTimes()
	uLog.EXPECT().CreateOne(guildUserActivityLog).Return(nil).AnyTimes()
	// case cannot find user, user not exist in server discord
	uXp.EXPECT().GetOne("abc", "abc").Return(nil, errors.New("Error cannot find user in server")).AnyTimes()
	discordMock.EXPECT().SendLevelUpMessage("", "", &response.HandleUserActivityResponse{
		GuildID:   "552427722551459840",
		UserID:    "973069332034752522",
		Action:    "gifted",
		AddedXP:   10,
		CurrentXP: 20,
	}).AnyTimes()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Entity{
				repo:     tt.fields.repo,
				store:    tt.fields.store,
				log:      tt.fields.log,
				dcwallet: tt.fields.dcwallet,
				discord:  tt.fields.discord,
				cache:    tt.fields.cache,
				svc:      tt.fields.svc,
				cfg:      tt.fields.cfg,
			}
			got, err := e.SendGiftXp(tt.args.GiftXPRequest)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.SendGiftXp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entity.SendGiftXp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntity_GetUserProfile(t *testing.T) {
	type fields struct {
		repo     *repo.Repo
		store    repo.Store
		log      logger.Logger
		dcwallet discordwallet.IDiscordWallet
		discord  *discordgo.Session
		cache    cache.Cache
		svc      *service.Service
		cfg      config.Config
	}
	type args struct {
		guildID string
		userID  string
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := config.LoadTestConfig()

	s := pg.NewPostgresStore(&cfg)
	r := pg.NewRepo(s.DB())
	log := logger.NewLogrusLogger()
	svc, _ := service.NewService(cfg, log)
	uXp := mock_guild_user_xp.NewMockStore(ctrl)
	cXp := mock_config_xp_level.NewMockStore(ctrl)
	dcG := mock_discord_guilds.NewMockStore(ctrl)
	processor := mock_processor.NewMockService(ctrl)
	svc.Processor = processor

	r.GuildUserXP = uXp
	r.ConfigXPLevel = cXp
	r.DiscordGuilds = dcG

	cXpValue := model.ConfigXpLevel{}

	cXp.EXPECT().GetNextLevel(gomock.Any(), gomock.Any()).Return(&cXpValue, nil).AnyTimes()

	dcGValue := model.DiscordGuild{}

	dcG.EXPECT().GetByID("981128899280908299").Return(&dcGValue, nil).AnyTimes()

	userXP := model.GuildUserXP{
		GuildID: "981128899280908299",
		UserID:  "963641551416881183",
	}

	uXp.EXPECT().GetOne("981128899280908299", "963641551416881183").Return(&userXP, nil).AnyTimes()

	uXp.EXPECT().GetOne("abc", "abc").Return(nil, errors.New("cannot find user")).AnyTimes()
	uXp.EXPECT().GetOne("abc", "963641551416881183").Return(nil, errors.New("cannot find guild")).AnyTimes()
	processor.EXPECT().GetUserFactionXp("963641551416881183").Return(&model.GetUserFactionXpsResponse{}, nil).AnyTimes()

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.GetUserProfileResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test get user profile successfully",
			fields: fields{
				repo: r,
				svc:  svc,
			},
			args: args{
				guildID: "981128899280908299",
				userID:  "963641551416881183",
			},
			want: &response.GetUserProfileResponse{
				ID:             "963641551416881183",
				CurrentLevel:   &cXpValue,
				NextLevel:      &cXpValue,
				GuildXP:        0,
				NrOfActions:    0,
				Progress:       1,
				Guild:          &dcGValue,
				UserWallet:     &model.UserWallet{},
				UserFactionXps: &model.UserFactionXpsMapping{},
			},
			wantErr: false,
		},
		{
			name: "case user not exist in server, cannot gift xp to this user",
			fields: fields{
				repo: r,
				svc:  svc,
			},
			args: args{
				guildID: "abc",
				userID:  "abc",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "case guild ID not exist",
			fields: fields{
				repo: r,
				svc:  svc,
			},
			args: args{
				guildID: "abc",
				userID:  "963641551416881183",
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Entity{
				repo:     tt.fields.repo,
				store:    tt.fields.store,
				log:      tt.fields.log,
				dcwallet: tt.fields.dcwallet,
				discord:  tt.fields.discord,
				cache:    tt.fields.cache,
				svc:      tt.fields.svc,
				cfg:      tt.fields.cfg,
			}
			got, err := e.GetUserProfile(tt.args.guildID, tt.args.userID)

			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.GetUserProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entity.GetUserProfile() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestEntity_GetTopUsers(t *testing.T) {
	type fields struct {
		repo     *repo.Repo
		store    repo.Store
		log      logger.Logger
		dcwallet discordwallet.IDiscordWallet
		discord  *discordgo.Session
		cache    cache.Cache
		svc      *service.Service
		cfg      config.Config
	}

	type args struct {
		guildID string
		userID  string
		limit   int
		page    int
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := config.LoadTestConfig()

	s := pg.NewPostgresStore(&cfg)
	r := pg.NewRepo(s.DB())

	uXp := mock_guild_user_xp.NewMockStore(ctrl)
	dcG := mock_discord_guilds.NewMockStore(ctrl)

	r.GuildUserXP = uXp
	r.DiscordGuilds = dcG

	userXP := model.GuildUserXP{
		GuildID: "981128899280908299",
		UserID:  "963641551416881183",
	}

	leaderboard := []model.GuildUserXP{}

	uXp.EXPECT().GetTopUsers(gomock.Any(), gomock.Any(), gomock.Any()).Return(leaderboard, nil).AnyTimes()

	dcGValue := model.DiscordGuild{}

	dcG.EXPECT().GetByID("981128899280908299").Return(&dcGValue, nil).AnyTimes()

	uXp.EXPECT().GetOne("981128899280908299", "963641551416881183").Return(&userXP, nil).AnyTimes()

	uXp.EXPECT().GetOne("abc", "abc").Return(nil, errors.New("cannot find user")).AnyTimes()
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *response.TopUser
		wantErr bool
	}{
		{
			name: "test get successfully",
			fields: fields{
				repo: r,
			},
			args: args{
				guildID: "981128899280908299",
				userID:  "963641551416881183",
				limit:   5,
				page:    0,
			},
			want: &response.TopUser{
				Author:      &userXP,
				Leaderboard: leaderboard,
			},
			wantErr: false,
		},
		{
			name: "test user does not exist",
			fields: fields{
				repo: r,
			},
			args: args{
				guildID: "abc",
				userID:  "abc",
				limit:   5,
				page:    0,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Entity{
				repo:     tt.fields.repo,
				store:    tt.fields.store,
				log:      tt.fields.log,
				dcwallet: tt.fields.dcwallet,
				discord:  tt.fields.discord,
				cache:    tt.fields.cache,
				svc:      tt.fields.svc,
				cfg:      tt.fields.cfg,
			}
			got, err := e.GetTopUsers(tt.args.guildID, tt.args.userID, tt.args.limit, tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.GetTopUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entity.GetTopUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntity_GetUserCurrentGMStreak(t *testing.T) {
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
	discordUserGmStreak := mock_discord_user_gm_streak.NewMockStore(ctrl)

	r.DiscordUserGMStreak = discordUserGmStreak

	type args struct {
		discordID string
		guildID   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.DiscordUserGMStreak
		want1   int
		wantErr bool
	}{
		{
			name: "User has Gm streak",
			fields: fields{
				repo: r,
			},
			args: args{
				guildID:   "552427722551459840",
				discordID: "393034938028392449",
			},
			want:    &model.DiscordUserGMStreak{GuildID: "552427722551459840", DiscordID: "393034938028392449", StreakCount: 1, TotalCount: 1},
			want1:   200,
			wantErr: false,
		},
		{
			name: "User does not have Gm streak",
			fields: fields{
				repo: r,
			},
			args: args{
				guildID:   "552427722551459840",
				discordID: "not_have_gm_streak",
			},
			want:    &model.DiscordUserGMStreak{},
			want1:   200,
			wantErr: false,
		},
	}

	discordUserGmStreak.EXPECT().GetByDiscordIDGuildID("393034938028392449", "552427722551459840").Return(&model.DiscordUserGMStreak{
		GuildID:     "552427722551459840",
		DiscordID:   "393034938028392449",
		StreakCount: 1,
		TotalCount:  1,
	}, nil).AnyTimes()
	discordUserGmStreak.EXPECT().GetByDiscordIDGuildID("not_have_gm_streak", "552427722551459840").Return(&model.DiscordUserGMStreak{}, nil).AnyTimes()

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
			got, got1, err := e.GetUserCurrentGMStreak(tt.args.discordID, tt.args.guildID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.GetUserCurrentGMStreak() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entity.GetUserCurrentGMStreak() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Entity.GetUserCurrentGMStreak() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
