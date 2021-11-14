package db

import (
	"database/sql"
	"sync"
)

type DB struct {
	conn *sql.DB
	lock *sync.Mutex
}
