package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/defipod/mochi/pkg/cache"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/discordwallet"
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/handler"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/repo/pg"
	"github.com/defipod/mochi/pkg/routes"
	"github.com/defipod/mochi/pkg/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func main() {
	// *** config ***
	cls := config.DefaultConfigLoaders()
	cfg := config.LoadConfig(cls)
	log := initLog(cfg)

	// *** db ***
	s, close := pg.NewPostgresStore(&cfg)
	defer func() {
		err := close()
		if err != nil {
			log.Fatalf("Error closing postgres store:", err)
		}
	}()

	repo := pg.NewRepo(s.DB())

	// *** dcwallet ***
	dcwallet, err := discordwallet.New(cfg, log, s)
	if err != nil {
		log.Fatalf("failed to init discord wallet: %v", err)
	}

	// *** discord **
	discord, err := discordgo.New("Bot " + cfg.DiscordToken)
	if err != nil {
		log.Fatalf("failed to init discord: %v", err)
	}
	setDiscordIntents(discord)
	log.Info("discord intents:", discord.Identify.Intents)

	u, err := discord.User("@me")
	if err != nil {
		log.Fatalf("failed to get discord bot user: %v", err)
	}
	log.Infof("Connected to discord: %s", u.Username)


	// *** cache ***
	redisOpt, err := redis.ParseURL(cfg.RedisURL)
	if err != nil {
		log.Fatalf("failed to init redis: %v", err)
	}

	redisCache, err := cache.NewRedisCache(redisOpt)
	if err != nil {
		log.Fatalf("failed to init redis cache: %v", err)
	}

	service := service.NewService()

	// *** entities ***
	entities, err := entities.New(
		log,
		repo,
		dcwallet,
		discord,
		redisCache,
		service,
	)
	if err != nil {
		log.Fatalf("failed to init entities: %v", err)
	}

	router := setupRouter(cfg, log, entities)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	shutdownServer(srv, cfg.GetShutdownTimeout(), log)

}

func shutdownServer(srv *http.Server, timeout time.Duration, l logger.Log) {
	l.Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		l.Error("Server Shutdown:", err)
	}
	l.Info("Server exiting")
}

func initLog(cfg config.Config) logger.Log {
	return logger.NewJSONLogger(
		logger.WithServiceName(cfg.ServiceName),
		logger.WithHostName(cfg.BaseURL),
	)
}

func setDiscordIntents(discord *discordgo.Session) {
	discord.Identify.Intents = discordgo.MakeIntent(discordgo.IntentGuilds)
	discord.Identify.Intents = discordgo.MakeIntent(discordgo.IntentGuildMessages)
	discord.Identify.Intents = discordgo.MakeIntent(discordgo.IntentGuildMessageReactions)
	discord.Identify.Intents = discordgo.MakeIntent(discordgo.IntentGuildMembers)
	discord.Identify.Intents = discordgo.MakeIntent(discordgo.IntentDirectMessages)
	discord.Identify.Intents = discordgo.MakeIntent(discordgo.IntentGuildInvites)
}

func setupRouter(cfg config.Config, l logger.Log, entities *entities.Entity) *gin.Engine {
	r := gin.New()
	pprof.Register(r)
	r.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/healthz"),
		gin.Recovery(),
	)

	h, err := handler.New(cfg, l, entities)
	if err != nil {
		l.Fatal(err)
	}

	corsOrigins := cfg.GetCORS()
	r.Use(func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		allowOrigins := corsOrigins

		// allow all localhosts and all GET method
		if origin != "" && (strings.Contains(origin, "http://localhost") || c.Request.Method == "GET") {
			allowOrigins = []string{origin}
		} else {
			// suport wildcard cors: https://*.domain.com
			for _, url := range allowOrigins {
				if strings.Contains(origin, strings.Replace(url, "https://*", "", 1)) {
					allowOrigins = []string{origin}
					break
				}
			}
		}

		cors.New(
			cors.Config{
				AllowOrigins: allowOrigins,
				AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"},
				AllowHeaders: []string{"Origin", "Host",
					"Content-Type", "Content-Length",
					"Accept-Encoding", "Accept-Language", "Accept",
					"X-CSRF-Token", "Authorization", "X-Requested-With", "X-Access-Token"},
				ExposeHeaders:    []string{"MeAllowMethodsntent-Length"},
				AllowCredentials: true,
			},
		)(c)
	})

	// handlers
	r.GET("/healthz", h.Healthz)

	routes.NewRoutes(r, h, cfg)

	return r
}
