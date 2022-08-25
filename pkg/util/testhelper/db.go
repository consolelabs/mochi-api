package testhelper

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/go-testfixtures/testfixtures/v3"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	db              *gorm.DB
	fixtures        *testfixtures.Loader
	singletonTestDB sync.Once
)

func LoadTestDB() *gorm.DB {
	var err error
	var conn *sql.DB

	singletonTestDB.Do(func() {
		// initiate logger
		l := logger.NewLogrusLogger()

		conn, err = sql.Open("postgres", "host=localhost port=25432 user=postgres password=postgres dbname=mochi_local_test sslmode=disable")
		if err != nil {
			fmt.Println(err)
			l.Fatalf(err, "failed to open database connection")
			return
		}

		path, err := os.Getwd()
		if err != nil {
			l.Error(err, "unable to get dir")
		}
		fmt.Println(path) // for example /home/user

		// load fixture and restore db
		fixtures, err = testfixtures.New(
			testfixtures.Database(conn),
			testfixtures.Dialect("postgres"),
			testfixtures.Directory("../../../migrations/test_seed"),
		)
		if err != nil {
			fmt.Print(err)
			l.Fatalf(err, "failed to load fixture")
			return
		}

		if err = fixtures.Load(); err != nil {
			l.Fatalf(err, "failed to load fixture")
			return
		}

		db, err = gorm.Open(postgres.New(
			postgres.Config{Conn: conn}),
			&gorm.Config{
				NamingStrategy: schema.NamingStrategy{
					SingularTable: true,
				},
			})
		if err != nil {
			l.Fatalf(err, "gorm: failed to open database connection")
		}
	})

	return db
}
