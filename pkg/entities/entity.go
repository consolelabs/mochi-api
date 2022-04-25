package entities

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/go-redis/redis/v8"

	"github.com/defipod/mochi/pkg/cache"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/consts"
	"github.com/defipod/mochi/pkg/discordwallet"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/repo"
	"github.com/defipod/mochi/pkg/repo/pg"
	"github.com/defipod/mochi/pkg/service"
)

var (
	ErrRecordNotFound = errors.New("not found")
)

type Entity struct {
	repo     *repo.Repo
	store    repo.Store
	log      logger.Logger
	dcwallet discordwallet.IDiscordWallet
	discord  *discordgo.Session
	cache    cache.Cache
	svc      *service.Service
}

var e *Entity

func Init(cfg config.Config, log logger.Logger) error {
	// *** db ***
	s := pg.NewPostgresStore(&cfg)
	repo := pg.NewRepo(s.DB())

	// *** dcwallet ***
	dcwallet, err := discordwallet.New(cfg, log, s)
	if err != nil {
		log.Fatal(err, "failed to init discord wallet")
	}

	// *** discord **
	discord, err := discordgo.New("Bot " + cfg.DiscordToken)
	if err != nil {
		log.Fatal(err, "failed to init discord")
	}
	setDiscordIntents(discord)
	log.Infof("discord intents: %d", discord.Identify.Intents)

	u, err := discord.User("@me")
	if err != nil {
		log.Fatal(err, "failed to get discord bot user")
	}
	log.Infof("Connected to discord: %s", u.Username)

	// *** cache ***
	redisOpt, err := redis.ParseURL(cfg.RedisURL)
	if err != nil {
		log.Fatal(err, "failed to init redis")
	}

	cache, err := cache.NewRedisCache(redisOpt)
	if err != nil {
		log.Fatal(err, "failed to init redis cache")
	}

	service := service.NewService()

	// *** init entity ***
	e = &Entity{
		repo:     repo,
		store:    s,
		log:      log,
		dcwallet: dcwallet,
		discord:  discord,
		cache:    cache,
		svc:      service,
	}

	if e.discord != nil && e.cache != nil {
		if err := e.initInviteTrackerCache(); err != nil {
			log.Error(err, "failed to init invite tracker cache")
		}
	}

	return nil
}

func Get() *Entity {
	return e
}

func Shutdown() error {
	e.store.Shutdown()
	return nil
}

func setDiscordIntents(discord *discordgo.Session) {
	discord.Identify.Intents = discordgo.MakeIntent(discordgo.IntentGuilds)
	discord.Identify.Intents = discordgo.MakeIntent(discordgo.IntentGuildMessages)
	discord.Identify.Intents = discordgo.MakeIntent(discordgo.IntentGuildMessageReactions)
	discord.Identify.Intents = discordgo.MakeIntent(discordgo.IntentGuildMembers)
	discord.Identify.Intents = discordgo.MakeIntent(discordgo.IntentDirectMessages)
	discord.Identify.Intents = discordgo.MakeIntent(discordgo.IntentGuildInvites)
}

func (e *Entity) initInviteTrackerCache() error {
	guilds, err := e.GetGuilds()
	if err != nil {
		return fmt.Errorf("failed to get guilds: %w", err)
	}

	for _, guild := range guilds.Data {
		invites, err := e.discord.GuildInvites(guild.ID)
		if err != nil {
			return fmt.Errorf("failed to get invites for guild %s: %w", guild.ID, err)
		}

		invitesUses := make(map[string]string)
		for _, invite := range invites {
			invitesUses[invite.Code] = strconv.Itoa(invite.Uses)
		}

		if len(invitesUses) > 0 {
			if err := e.cache.HashSet(consts.CachePrefixInviteTracker+guild.ID, invitesUses, 0); err != nil {
				return fmt.Errorf("failed to cache invites for guild %s: %w", guild.ID, err)
			}
		}

		e.log.Fields(logger.Fields{guild.ID: invites}).Debug("cache guild invites")
	}

	return nil
}
