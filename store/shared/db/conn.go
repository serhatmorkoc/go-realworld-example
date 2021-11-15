package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
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
		return nil, err
	}

	if err := pingDatabase(db); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConnections)
	db.SetMaxIdleConns(5)
	db.SetConnMaxIdleTime(1 * time.Second)
	db.SetConnMaxLifetime(30 * time.Second)

	return &DB{
		conn: db,
		lock: new(sync.Mutex),
	}, nil
}

func pingDatabase(db *sql.DB) error {
	r := 5
	for i := 0; i < r; i++ {
		err := db.Ping()
		if err == nil {
			return nil
		}
		time.Sleep(1 * time.Second)
	}

	return errPingDatabase
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
