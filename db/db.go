package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/serhatmorkoc/golang-realworld-example/db/migrate"
	"time"
)

type Locker interface {
	Lock()
	Unlock()
	RLock()
	RUnlock()
}

/*type DB struct {
Conn   *sqlx.DB
Lock   Locker

}
*/
func Connect(driver string, dsn string, maxOpenConnections int) (*sql.DB, error) {

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

	return db,nil
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

	return migrate.Migrate(db)

	/*switch driver {
	case "mysql":
		return mysql.Migrate(db)
	case "postgres":
		return postgres.Migrate(db)
	default:
		return sqlite.Migrate(db)
	}
	*/
}
