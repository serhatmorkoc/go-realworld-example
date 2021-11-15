package db

import (
	"context"
	"database/sql"
	"runtime/debug"
	"sync"
)

type DB struct {
	conn *sql.DB
	lock *sync.Mutex
}

type Execer interface {
	Queryer
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

type Queryer interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

func (db *DB) Read(fn func(Execer) error) error {
	db.lock.Lock()
	defer db.lock.Unlock()

	err := fn(db.conn)
	return err
}

func (db *DB) Update(fn func(Execer) error) error {
	db.lock.Lock()
	defer db.lock.Unlock()

	tx, err := db.conn.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			err = tx.Rollback()
			debug.PrintStack()

		} else if err != nil {
			tx.Rollback()

		} else {
			err = tx.Commit()
		}
	}()

	err = fn(tx)
	return err
}

func (db *DB) Close() error {
	return db.conn.Close()
}
