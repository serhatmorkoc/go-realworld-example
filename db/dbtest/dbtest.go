package dbtest

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/serhatmorkoc/golang-realworld-example/db"
	"os"
)

func Connect() (*sqlx.DB, error) {

	var (
		driver         = "sqlite3"
		dsn         = ""
		maxConnections = 0
	)

	if os.Getenv("DB_NAME") != "" {
		driver = os.Getenv("DB_USER")
	}

	return db.Connect(driver, dsn,maxConnections)
}
