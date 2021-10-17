package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/serhatmorkoc/go-realworld-example/config"
	"github.com/serhatmorkoc/go-realworld-example/database/migrate/postgres"
	"strconv"
	"time"
)

func Connect(cfg *config.Config) (*sql.DB, error) {

	port, _ := strconv.Atoi(cfg.Database.Port)
	dsn, err := parseDSN(cfg.Database.Driver,
		cfg.Database.Host,
		cfg.Database.Name,
		cfg.Database.User,
		cfg.Database.Password,
		port)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open(cfg.Database.Driver, dsn)
	if err != nil {
		return nil, err
	}

	if err := pingDatabase(db); err != nil {
		return nil, err
	}

	if err := setupDatabase(db, cfg.Database.Driver); err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(3)

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

	return errors.New("database connection failed")
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
