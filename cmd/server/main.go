package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	mdwgin "github.com/consolelabs/mochi-toolkit/http/middleware/gin"
	typesetservice "github.com/consolelabs/mochi-typeset/common/service/typeset"
	typesetqueue "github.com/consolelabs/mochi-typeset/queue"
	"github.com/getsentry/sentry-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	_ "github.com/defipod/mochi/docs"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/handler"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/routes"
	"github.com/defipod/mochi/pkg/util"
)

// @title          Swagger API
// @version        1.0
// @description    This is a swagger for mochi api.
// @termsOfService http://swagger.io/terms/

// @contact.name  API Support
// @contact.url   http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name access_token
func main() {
	cfg := config.LoadConfig(config.DefaultConfigLoaders())
	log := logger.NewLogrusLogger()

	err := sentry.Init(sentry.ClientOptions{
		Dsn:        "https://b632003ad0874c5182ee572bb7ad3c6c@sentry.daf.ug/4",
		Debug:      true,
		ServerName: "mochi-api",
	})
	if err != nil {
		log.Fatal(err, "can't init sentry service")
	}
	defer sentry.Flush(2 * time.Second)

	// *** entities ***
	err = entities.Init(cfg, log)
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

	// implement fetch discord stats
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
	r.Use(func(ctx *gin.Context) {
		cr := mdwgin.CaptureRequest(ctx, &mdwgin.CaptureRequestOptions{
			ExcludePaths: []string{"/healthz", "/webhook/discord"},
		})
		if cr == nil {
			ctx.Next()
			return
		}
		b, err := json.Marshal(cr)
		if err != nil {
			l.Error(err, "cannot marshal capture request")
			ctx.Next()
			return
		}
		go func() {
			kfkMsg := typesetqueue.KafkaMessage{
				Type:   typesetqueue.KAFKA_MESSAGE_TYPE_AUDIT,
				Data:   b,
				Sender: typesetservice.SERVICE_MOCHI_API,
			}
			body, err := json.Marshal(kfkMsg)
			if err != nil {
				return
			}
			util.SendRequest(util.SendRequestQuery{
				Method: http.MethodPost,
				URL:    fmt.Sprintf("%v/api/v1/audit", cfg.MochiAuditServerHost),
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				Body: bytes.NewReader(body),
			})
		}()
	})

	h := handler.New(e, l)

	r.Use(func(c *gin.Context) {
		allowOrigins := []string{"*"}

		cors.New(
			cors.Config{
				AllowOrigins: allowOrigins,
				AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"},
				AllowHeaders: []string{"Origin", "Host",
					"Content-Type", "Content-Length",
					"Accept-Encoding", "Accept-Language", "Accept",
					"X-CSRF-Token", "Authorization", "X-Requested-With", "X-Access-Token", "X-Request-Id"},
				ExposeHeaders:    []string{"MeAllowMethodsntent-Length"},
				AllowCredentials: true,
			},
		)(c)
	})

	// handlers
	r.GET("/healthz", h.Healthcheck.Healthz)

	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// load API here
	routes.NewRoutes(r, h, cfg)

	return r
}
