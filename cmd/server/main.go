package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/handler"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig(config.DefaultConfigLoaders())
	log := logger.NewLogrusLogger()
	// *** entities ***
	err := entities.Init(cfg, log)
	if err != nil {
		log.Fatal(err, "failed to init entities")
	}

	router := setupRouter(cfg, log, entities.Get())

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err, "failed to listen and serve")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	shutdownServer(srv, cfg.GetShutdownTimeout(), log)
	if err := entities.Shutdown(); err != nil {
		log.Fatal(err, "failed to shutdown entities")
	}

}

func shutdownServer(srv *http.Server, timeout time.Duration, l logger.Logger) {
	l.Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		l.Error(err, "failed to shutdown server")
	}

	l.Info("Server exiting")
}

func setupRouter(cfg config.Config, l logger.Logger, e *entities.Entity) *gin.Engine {
	r := gin.New()
	pprof.Register(r)
	r.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/healthz"),
		gin.Recovery(),
	)

	h, err := handler.New(e)
	if err != nil {
		l.Fatal(err, "failed to init handler")
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

	routes.NewRoutes(r, h)

	return r
}
