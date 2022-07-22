package entities

import (
	"errors"
	"reflect"
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
	"github.com/defipod/mochi/pkg/service/abi"
	"github.com/defipod/mochi/pkg/service/indexer"
	"github.com/defipod/mochi/pkg/service/marketplace"
	"github.com/golang/mock/gomock"

	mock_discord_guilds "github.com/defipod/mochi/pkg/repo/discord_guilds/mocks"
	mock_config_verify "github.com/defipod/mochi/pkg/repo/guild_config_wallet_verification_message/mocks"
)

func TestEntity_GetGuildConfigWalletVerificationMessage(t *testing.T) {
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
	log := logger.NewLogrusLogger()

	mockDiscordGuilds := mock_discord_guilds.NewMockStore(ctrl)
	mockGuildConfigVerify := mock_config_verify.NewMockStore(ctrl)

	r.DiscordGuilds = mockDiscordGuilds
	r.GuildConfigWalletVerificationMessage = mockGuildConfigVerify

	type args struct {
		guildId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.GuildConfigWalletVerificationMessage
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "successful",
			fields: fields{
				repo: r,
				log:  log,
			},
			args: args{
				guildId: "863278424433229854",
			},
			want: &model.GuildConfigWalletVerificationMessage{
				GuildID:          "863278424433229854",
				VerifyChannelID:  "986854719999864863",
				Content:          "",
				EmbeddedMessage:  nil,
				CreatedAt:        time.Date(2022, 1, 2, 3, 4, 5, 6, time.UTC),
				DiscordMessageID: "999891586395685004",
			},
			wantErr: false,
		},
		{
			name: "invalid guild id",
			fields: fields{
				repo: r,
				log:  log,
			},
			args: args{
				guildId: "123456",
			},
			want:    nil,
			wantErr: true,
		},
	}
	guild := &model.DiscordGuild{
		Name: "test",
	}
	config := &model.GuildConfigWalletVerificationMessage{
		GuildID:          "863278424433229854",
		VerifyChannelID:  "986854719999864863",
		Content:          "",
		EmbeddedMessage:  nil,
		CreatedAt:        time.Date(2022, 1, 2, 3, 4, 5, 6, time.UTC),
		DiscordMessageID: "999891586395685004",
	}
	// case success
	mockDiscordGuilds.EXPECT().GetByID("863278424433229854").Return(guild, nil).AnyTimes()
	mockGuildConfigVerify.EXPECT().GetOne("863278424433229854").Return(config, nil).AnyTimes()

	// case fail - invalid guild id
	mockDiscordGuilds.EXPECT().GetByID("123456").Return(nil, errors.New("invalid guild id")).AnyTimes()

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
			got, err := e.GetGuildConfigWalletVerificationMessage(tt.args.guildId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Entity.GetGuildConfigWalletVerificationMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entity.GetGuildConfigWalletVerificationMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
