package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/serhatmorkoc/go-realworld-example/database/migrate/postgres"
	"time"
)

func Connect(driver, host, database, username, password string, port, maxOpenConnections int) (*sql.DB, error) {

	dsn, err := parseDSN(driver, host, database, username, password, port)
	if err != nil {
		return nil, err
	}

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

func parseDSN(driver, host, database, username, password string, port int) (string, error) {

	switch driver {
	case "postgres":
		return postgresParseDSN(host, database, username, password, port), nil
	case "mysql":
		return "", errors.New("mysql is not supported")
	default:
		return "", errors.New("driver is not supported")
	}
}

func postgresParseDSN(host, database, username, password string, port int) string {

	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, username, password, database)
}
