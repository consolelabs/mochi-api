package service

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	errs "errors"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/service/coingecko"
	"github.com/defipod/mochi/pkg/service/discord"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

// import "github.com/defipod/api/pkg/service/binance"

type Service struct {
	CoinGecko coingecko.Service
	Discord   discord.Service
}

func NewService(
	cfg config.Config,
) (*Service, error) {

	discordSvc, err := discord.NewService(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to init discord: %w", err)
	}

	ds := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPass,
		cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	conn, err := sql.Open("postgres", ds)
	if err != nil {
		panic(err)
	}
	_, err = gorm.Open(postgres.New(
		postgres.Config{Conn: conn}),
		&gorm.Config{
			Logger: newLogrusLogger(logrus.StandardLogger()),
		})
	if err != nil {
		panic(err)
	}

	return &Service{
		CoinGecko: coingecko.NewService(),
		Discord:   discordSvc,
	}, nil
}

type logger struct {
	SlowThreshold         time.Duration
	SourceField           string
	SkipErrRecordNotFound bool
}

func newLogrusLogger(*logrus.Logger) *logger {
	return &logger{
		SlowThreshold:         time.Second,
		SkipErrRecordNotFound: true,
	}
}

func (l *logger) LogMode(gormlogger.LogLevel) gormlogger.Interface {
	return l
}

func (l *logger) Info(ctx context.Context, s string, args ...interface{}) {
	logrus.WithContext(ctx).Infof(s, args...)
}

func (l *logger) Warn(ctx context.Context, s string, args ...interface{}) {
	logrus.WithContext(ctx).Warnf(s, args...)
}

func (l *logger) Error(ctx context.Context, s string, args ...interface{}) {
	logrus.WithContext(ctx).Errorf(s, args...)
}

func (l *logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, _ := fc()
	fields := logrus.Fields{}
	if l.SourceField != "" {
		fields[l.SourceField] = utils.FileWithLineNum()
	}
	if err != nil && !(errs.Is(err, gorm.ErrRecordNotFound) && l.SkipErrRecordNotFound) {
		fields[logrus.ErrorKey] = err
		logrus.WithContext(ctx).WithFields(fields).Errorf("%s [%s]", sql, elapsed)
		return
	}

	if l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
		logrus.WithContext(ctx).WithFields(fields).Warnf("%s [%s]", sql, elapsed)
		return
	}

	logrus.WithContext(ctx).WithFields(fields).Debugf("%s [%s]", sql, elapsed)
}
