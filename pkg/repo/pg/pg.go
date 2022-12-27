package pg

import (
	"context"
	"database/sql"
	errs "errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"gorm.io/plugin/dbresolver"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/model/errors"
	"github.com/defipod/mochi/pkg/repo"
	"github.com/defipod/mochi/pkg/util"
)

// store is implimentation of repository
type store struct {
	database     *gorm.DB
	shutdownFunc func() error
}

// Shutdown close database connection
func (s *store) Shutdown() error {
	if s.shutdownFunc != nil {
		return s.shutdownFunc()
	}
	return nil
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
func NewPostgresStore(cfg *config.Config) repo.Store {
	ds := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPass,
		cfg.DBHost, cfg.DBPort, cfg.DBName,
	)
	conn, err := sql.Open("postgres", ds)
	if err != nil {
		panic(err)
	}

	l := logrus.StandardLogger()
	if cfg.Debug {
		l.SetLevel(logrus.DebugLevel)
	}

	db, err := gorm.Open(postgres.New(
		postgres.Config{Conn: conn}),
		&gorm.Config{
			Logger: newLogrusLogger(l),
		})
	if err != nil {
		panic(err)
	}

	var readDialectors []gorm.Dialector
	for _, r := range cfg.DBReadHosts {
		if r == "" {
			continue
		}

		port := cfg.DBPort
		if strings.Contains(r, ":") {
			p := strings.Split(r, ":")
			r = p[0]
			port = p[1]
		}
		ds := fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=disable",
			cfg.DBUser, cfg.DBPass,
			r, port, cfg.DBName,
		)
		conn, err := sql.Open("postgres", ds)
		if err != nil {
			l.Errorf("cannot init replica %v:%v", r, port)
			continue
		}
		readDialectors = append(readDialectors, postgres.New(postgres.Config{Conn: conn}))
	}

	if len(readDialectors) > 0 {
		err := db.Use(dbresolver.Register(dbresolver.Config{
			Replicas: readDialectors,
			// sources/replicas load balancing policy
			Policy: dbresolver.RandomPolicy{},
		}))
		if err != nil {
			panic(err)
		}
	}

	return &store{
		database:     db,
		shutdownFunc: conn.Close,
	}
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
