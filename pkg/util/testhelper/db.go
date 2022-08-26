package testhelper

import (
	"database/sql"
	"sync"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/go-testfixtures/testfixtures/v3"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db              *gorm.DB
	fixtures        *testfixtures.Loader
	singletonTestDB sync.Once
)

func LoadTestDB(seedPath string) *gorm.DB {
	var err error
	var conn *sql.DB

	singletonTestDB.Do(func() {
		// initiate logger
		l := logger.NewLogrusLogger()

		conn, err = sql.Open("postgres", "host=localhost port=25432 user=postgres password=postgres dbname=mochi_local_test sslmode=disable")
		if err != nil {
			l.Fatalf(err, "failed to open database connection")
			return
		}

		// load fixture and restore db
		fixtures, err = testfixtures.New(
			testfixtures.Database(conn),
			testfixtures.Dialect("postgres"),
			testfixtures.Directory(seedPath),
		)
		if err != nil {
			l.Fatalf(err, "failed to load fixture")
			return
		}

		if err = fixtures.Load(); err != nil {
			l.Fatalf(err, "failed to load fixture")
			return
		}

		db, err = gorm.Open(postgres.New(
			postgres.Config{Conn: conn}))
		if err != nil {
			l.Fatalf(err, "gorm: failed to open database connection")
		}
	})

	return db
}
