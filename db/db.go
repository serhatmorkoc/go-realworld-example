package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/serhatmorkoc/go-realworld-example/db/migrate/postgres"
	"time"
)

func Connect(driver, dsn string, maxOpenConnections int) (*sql.DB, error) {

	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}

	if err := pingDatabase(db); err != nil {
		return nil, err
	}

	if err := setupDatabase(db, driver); err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(maxOpenConnections)

	return db, nil
}

func pingDatabase(db *sql.DB) (err error) {
	for i := 0; i < 5; i++ {
		err = db.Ping()
		if err == nil {
			return nil
		}
		time.Sleep(1 * time.Second)
	}

	return errors.New("")
}

func setupDatabase(db *sql.DB, driver string) error {

	return postgres.Migrate(db)

	switch driver {
	case "postgres":
		return postgres.Migrate(db)
	case "mysql":
		return errors.New("mysql is not supported")
	default:
		return errors.New("driver is not supported")
	}

}
