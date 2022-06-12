package entities

import (
	//"errors"
	"errors"
	"testing"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/defipod/mochi/pkg/cache"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/discordwallet"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/repo"
	"github.com/defipod/mochi/pkg/repo/pg"
	"github.com/defipod/mochi/pkg/service"
	"github.com/defipod/mochi/pkg/util"
	"github.com/golang/mock/gomock"

	mock_discord_guild_stats "github.com/defipod/mochi/pkg/repo/discord_guild_stats/mocks"
)

func TestEntity_UpdateGuildStats(t *testing.T) {
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
		stat model.DiscordGuildStat
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

		DiscordToken: "OTgzNjE5MTMyNTMwNDkxNDMz.Gi9ci5.SiEdAaDilYktiABZGunihCr-488lgnwNN7Wf10",

		RedisURL: "redis://localhost:6379/0",
	}
	s := pg.NewPostgresStore(&cfg)
	r := pg.NewRepo(s.DB())

	disGuStat := mock_discord_guild_stats.NewMockStore(ctrl)
	r.DiscordGuildStats = disGuStat

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "update guild stats successfully",
			fields: fields{
				repo: r,
			},
			args: args{stat: model.DiscordGuildStat{
				ID:                       util.GetNullUUID("45fad0a6-7f71-48d1-b343-ae51614980ed"),
				GuildID:                  "863278424433229854",
				NrOfMembers:              4,
				NrOfUsers:                1,
				NrOfBots:                 3,
				NrOfChannels:             5,
				NrOfTextChannels:         1,
				NrOfVoiceChannels:        4,
				NrOfStageChannels:        0,
				NrOfCategories:           2,
				NrOfAnnouncementChannels: 0,
				NrOfEmojis:               0,
				NrOfStaticEmojis:         0,
				NrOfAnimatedEmojis:       0,
				NrOfStickers:             0,
				NrOfCustomStickers:       0,
				NrOfServerStickers:       0,
				NrOfRoles:                4,
				CreatedAt:                time.Date(2022, 06, 8, 13, 51, 38, 6763, time.UTC)}},
			wantErr: false,
		},
		{
			name: "guild id incorrect",
			fields: fields{
				repo: r,
			},
			args: args{stat: model.DiscordGuildStat{
				ID:                       util.GetNullUUID("45fad0a6-7f71-48d1-b343-ae51614980ed"),
				GuildID:                  "abc",
				NrOfMembers:              4,
				NrOfUsers:                1,
				NrOfBots:                 3,
				NrOfChannels:             5,
				NrOfTextChannels:         1,
				NrOfVoiceChannels:        4,
				NrOfStageChannels:        0,
				NrOfCategories:           2,
				NrOfAnnouncementChannels: 0,
				NrOfEmojis:               0,
				NrOfStaticEmojis:         0,
				NrOfAnimatedEmojis:       0,
				NrOfStickers:             0,
				NrOfCustomStickers:       0,
				NrOfServerStickers:       0,
				NrOfRoles:                4,
				CreatedAt:                time.Date(2022, 06, 8, 13, 51, 38, 6763, time.UTC)}},
			wantErr: true,
		},
	}

	guildStatReturned := model.DiscordGuildStat{
		ID:                       util.GetNullUUID("45fad0a6-7f71-48d1-b343-ae51614980ed"),
		GuildID:                  "863278424433229854",
		NrOfMembers:              4,
		NrOfUsers:                1,
		NrOfBots:                 3,
		NrOfChannels:             5,
		NrOfTextChannels:         1,
		NrOfVoiceChannels:        4,
		NrOfStageChannels:        0,
		NrOfCategories:           2,
		NrOfAnnouncementChannels: 0,
		NrOfEmojis:               0,
		NrOfStaticEmojis:         0,
		NrOfAnimatedEmojis:       0,
		NrOfStickers:             0,
		NrOfCustomStickers:       0,
		NrOfServerStickers:       0,
		NrOfRoles:                4,
		CreatedAt:                time.Date(2022, 06, 8, 13, 51, 38, 6763, time.UTC),
	}
	guildStatWrong := model.DiscordGuildStat{
		ID:                       util.GetNullUUID("45fad0a6-7f71-48d1-b343-ae51614980ed"),
		GuildID:                  "abc",
		NrOfMembers:              4,
		NrOfUsers:                1,
		NrOfBots:                 3,
		NrOfChannels:             5,
		NrOfTextChannels:         1,
		NrOfVoiceChannels:        4,
		NrOfStageChannels:        0,
		NrOfCategories:           2,
		NrOfAnnouncementChannels: 0,
		NrOfEmojis:               0,
		NrOfStaticEmojis:         0,
		NrOfAnimatedEmojis:       0,
		NrOfStickers:             0,
		NrOfCustomStickers:       0,
		NrOfServerStickers:       0,
		NrOfRoles:                4,
		CreatedAt:                time.Date(2022, 06, 8, 13, 51, 38, 6763, time.UTC),
	}
	disGuStat.EXPECT().UpsertOne(guildStatReturned).Return(nil).AnyTimes()
	disGuStat.EXPECT().UpsertOne(guildStatWrong).Return(errors.New("Error cannot find guild")).AnyTimes()
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
			if err := e.UpdateGuildStats(tt.args.stat); (err != nil) != tt.wantErr {
				t.Errorf("Entity.UpdateGuildStats() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
