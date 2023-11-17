package sentrygo

import (
	"errors"
	"time"

	"github.com/getsentry/sentry-go"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
)

type sentrygo struct {
	client *sentry.Client
	config *config.Config
	logger logger.Logger
}

func New(cfg *config.Config, l logger.Logger) Service {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:   cfg.SentryDsn,
		Debug: true,
	})
	if err != nil {
		l.Fatal(err, "can't init sentry service")
	}

	return &sentrygo{
		client: sentry.CurrentHub().Client(),
		config: cfg,
		logger: l,
	}
}

func (s *sentrygo) Flush(timeout time.Duration) bool {
	return s.client.Flush(timeout)
}

func (s *sentrygo) CaptureErrorEvent(msg string, data map[string]interface{}) {
	scope := sentry.NewScope()
	scope.SetLevel(sentry.LevelError)
	event := sentry.NewEvent()
	event.Level = sentry.LevelError
	event.Message = msg
	event.Extra = data
	s.client.CaptureEvent(event, &sentry.EventHint{
		Data:              data,
		EventID:           msg,
		OriginalException: errors.New(msg),
	}, scope)
}
