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
	"github.com/defipod/mochi/pkg/repo"
	"github.com/defipod/mochi/pkg/repo/pg"
	"github.com/defipod/mochi/pkg/routes"
	"github.com/defipod/mochi/pkg/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func main() {
	cls := config.DefaultConfigLoaders()
	cfg := config.LoadConfig(cls)

	log := initLog(cfg)

	s, close := pg.NewPostgresStore(&cfg)
	defer func() {
		err := close()
		if err != nil {
			log.Fatalf("Error closing postgres store:", err)
		}
	}()

	repo := pg.NewRepo(s.DB())

	dcwallet, err := discordwallet.New(cfg, log, s)
	if err != nil {
		log.Fatalf("failed to init discord wallet: %v", err)
	}

	discord, err := discordgo.New("Bot " + cfg.DiscordToken)
	if err != nil {
		log.Fatalf("failed to init discord: %v", err)
	}

	redisOpt, err := redis.ParseURL(cfg.RedisURL)
	if err != nil {
		log.Fatalf("failed to init redis: %v", err)
	}

	redisCache, err := cache.NewRedisCache(redisOpt)
	if err != nil {
		log.Fatalf("failed to init redis cache: %v", err)
	}

	entities := entities.New(
		log,
		repo,
		dcwallet,
		discord,
		redisCache,
	)

	service := service.NewService()

	router := setupRouter(cfg, log, s, dcwallet, entities, service)

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

func setupRouter(cfg config.Config, l logger.Log, s repo.Store, dcwallet *discordwallet.DiscordWallet, entities *entities.Entity, service *service.Service) *gin.Engine {
	r := gin.New()
	pprof.Register(r)
	r.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/healthz"),
		gin.Recovery(),
	)

	h, err := handler.New(cfg, l, s, dcwallet, entities, service)
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

	routes.NewRoutes(r, h, cfg, s)

	return r
}
