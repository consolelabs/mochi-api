package entities

import (
	"errors"

	"github.com/bwmarrin/discordgo"
	"github.com/go-redis/redis/v8"
	"github.com/go-rod/rod"

	"github.com/defipod/mochi/pkg/cache"
	"github.com/defipod/mochi/pkg/chain"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/discordwallet"
	"github.com/defipod/mochi/pkg/kafka"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/repo"
	"github.com/defipod/mochi/pkg/repo/pg"
	"github.com/defipod/mochi/pkg/service"
	"github.com/defipod/mochi/pkg/service/abi"
	"github.com/defipod/mochi/pkg/service/indexer"
	"github.com/defipod/mochi/pkg/service/marketplace"
	"github.com/defipod/mochi/pkg/vaultwallet"
)

var (
	ErrRecordNotFound = errors.New("not found")
)

type Entity struct {
	repo        *repo.Repo
	store       repo.Store
	log         logger.Logger
	dcwallet    discordwallet.IDiscordWallet
	vaultwallet vaultwallet.IVaultWallet
	discord     *discordgo.Session
	cache       cache.Cache
	svc         *service.Service
	cfg         config.Config
	indexer     indexer.Service
	abi         abi.Service
	marketplace marketplace.Service
	solana      chain.Solana
	kafka       kafka.Kafka
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

	// *** vaultwallet ***
	vaultwallet, err := vaultwallet.New(cfg, log, s)
	if err != nil {
		log.Fatal(err, "failed to init vault wallet")
	}

	// *** discord **
	discord, err := discordgo.New("Bot " + cfg.DiscordToken)
	if err != nil {
		log.Fatal(err, "failed to init discord")
	}
	setDiscordIntents(discord)
	log.Infof("discord intents: %d", discord.Identify.Intents)

	// u, err := discord.User("@me")
	// if err != nil {
	// 	log.Fatal(err, "failed to get discord bot user")
	// }
	// log.Infof("Connected to discord: %s", u.Username)

	// *** cache ***
	redisOpt, err := redis.ParseURL(cfg.RedisURL)
	if err != nil {
		log.Fatal(err, "failed to init redis")
	}

	cache, err := cache.NewRedisCache(redisOpt)
	if err != nil {
		log.Fatal(err, "failed to init redis cache")
	}

	service, err := service.NewService(cfg, log)
	if err != nil {
		log.Fatal(err, "failed to init service")
	}

	kafka := kafka.New(cfg.Kafka.Brokers)

	errCh := make(chan error)
	go func(ch chan error) {
		err := kafka.RunProducer()
		if err != nil {
			errCh <- err
		}
	}(errCh)

	// *** init entity ***
	e = &Entity{
		repo:        repo,
		store:       s,
		log:         log,
		dcwallet:    dcwallet,
		vaultwallet: vaultwallet,
		discord:     discord,
		cache:       cache,
		svc:         service,
		cfg:         cfg,
		indexer:     indexer.NewIndexer(cfg, log),
		abi:         abi.NewAbi(&cfg),
		marketplace: marketplace.NewMarketplace(&cfg),
		solana:      *chain.NewSolanaClient(&cfg, log, cache),
		kafka:       *kafka,
	}

	return nil
}

func Get() *Entity {
	return e
}
func (e *Entity) GetSvc() *service.Service {
	return e.svc
}
func (e *Entity) GetLogger() logger.Logger {
	return e.log
}
func (e *Entity) GetKafka() *kafka.Kafka {
	return &e.kafka
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

func New(cfg config.Config, log logger.Logger, repo *repo.Repo, store repo.Store, dcwallet discordwallet.IDiscordWallet, discord *discordgo.Session, cache cache.Cache, svc *service.Service, indexer indexer.Service, abi abi.Service, marketplace marketplace.Service, page *rod.Page) *Entity {
	return &Entity{
		repo:        repo,
		store:       store,
		log:         log,
		dcwallet:    dcwallet,
		discord:     discord,
		cache:       cache,
		svc:         svc,
		cfg:         cfg,
		indexer:     indexer,
		abi:         abi,
		marketplace: marketplace,
	}
}
