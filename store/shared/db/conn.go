package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/serhatmorkoc/go-realworld-example/store/shared/db/migrate/postgres"
	"sync"
	"time"
)

func Connection(driver, host, database, username, password string, port, maxOpenConnections int) (*DB, error) {
	dsn, err := parseDSN(driver, host, database, username, password, port)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open(driver, dsn)
	if err != nil {
		//debug.PrintStack()
		return nil, err
	}

	if err := pingDatabase(db); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConnections)
	db.SetMaxIdleConns(5)
	db.SetConnMaxIdleTime(1 * time.Second)
	db.SetConnMaxLifetime(30 * time.Second)

	if err := setupDatabase(db, driver); err != nil {
		return nil, err
	}

	return &DB{
		Conn: db,
		lock: new(sync.Mutex),
	}, nil
}

func pingDatabase(db *sql.DB) error {
	r := 3
	for i := 0; i < r; i++ {
		err := db.Ping()
		if err == nil {
			return nil
		}
		time.Sleep(1 * time.Second)
	}

	//debug.PrintStack()
	return errPingDatabase
}

func setupDatabase(db *sql.DB, driver string) error {

	return postgres.Migrate(db)

	switch driver {
	case "postgres":
		return postgres.Migrate(db)
	default:
		return errUnSupportedDriver
	}

}

func parseDSN(driver, host, database, username, password string, port int) (string, error) {

	switch driver {
	case "postgres":
		return postgreParseDSN(host, database, username, password, port), nil
	default:
		return "", errUnSupportedDriver
	}
}

func postgreParseDSN(host, database, username, password string, port int) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, username, password, database)
}
