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

type SentryCapturePayload struct {
	Message string
	Tags    map[string]string
	Extra   map[string]interface{}
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

func (s *sentrygo) CaptureErrorEvent(p SentryCapturePayload) {
	scope := sentry.NewScope()
	scope.SetLevel(sentry.LevelError)
	scope.SetTags(p.Tags)
	event := sentry.NewEvent()
	event.Level = sentry.LevelError
	event.Message = p.Message
	event.Extra = p.Extra
	s.client.CaptureEvent(event, &sentry.EventHint{
		Data:              p.Extra,
		EventID:           p.Message,
		OriginalException: errors.New(p.Message),
	}, scope)
}
