package pg

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	errs "errors"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/model/errors"
	"github.com/defipod/mochi/pkg/repo"
	"github.com/defipod/mochi/pkg/util"
	"gorm.io/gorm"
)

// store is implimentation of repository
type store struct {
	database *gorm.DB
}

// DB database connection
func (s *store) DB() *gorm.DB {
	return s.database
}

// NewTransaction for database connection
func (s *store) NewTransaction() (newRepo repo.Store, finallyFn repo.FinallyFunc) {
	newDB := s.database.Begin()

	finallyFn = func(err error) error {
		if err != nil {
			nErr := newDB.Rollback().Error
			if nErr != nil {
				return errors.NewStringError(nErr.Error(), http.StatusInternalServerError)
			}
			return errors.NewStringError(err.Error(), util.ParseErrorCode(err))
		}

		cErr := newDB.Commit().Error
		if cErr != nil {
			return errors.NewStringError(cErr.Error(), http.StatusInternalServerError)
		}
		return nil
	}

	return &store{database: newDB}, finallyFn
}

// NewPostgresStore postgres init by gorm
func NewPostgresStore(cfg *config.Config) (repo.Store, func() error) {
	ds := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPass,
		cfg.DBHost, cfg.DBPort, cfg.DBName,
	)
	conn, err := sql.Open("postgres", ds)
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.New(
		postgres.Config{Conn: conn}),
		&gorm.Config{
			Logger: newLogrusLogger(logrus.StandardLogger()),
		})
	if err != nil {
		panic(err)
	}

	return &store{database: db}, conn.Close
}

// NewStore postgres init by gorm
func NewStore(db *gorm.DB) repo.Store {
	return &store{database: db}
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
