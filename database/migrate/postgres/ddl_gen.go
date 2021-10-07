package postgres

import (
	"database/sql"
)

var migrations = []struct {
	name string
	stmt string
}{
	{
		name: "create-table-users",
		stmt: createTableUsers,
	},
	{
		name: "create-table-comments",
		stmt: createTableComments,
	},
}

func Migrate(db *sql.DB) error {
	if err := createTable(db); err != nil {
		return err
	}

	completed, err := selectCompleted(db)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	for _, migration := range migrations {
		if _, ok := completed[migration.name]; ok {
			continue
		}

		if _, err := db.Exec(migration.stmt); err != nil {
			return err
		}
		if err := insertMigration(db, migration.name); err != nil {
			return err
		}

	}
	return nil
}

func selectCompleted(db *sql.DB)(map[string]struct{}, error) {
	migrations := map[string]struct{}{}
	rows, err := db.Query(migrationSelect)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		migrations[name] = struct{}{}
	}
	return migrations,nil
}

func createTable(db *sql.DB) error {
	_, err := db.Exec(createTableMigrations)
	return err
}

func insertMigration(db *sql.DB, name string) error {
	_, err := db.Exec(migrationInsert, name)
	return err
}

var createTableMigrations = `
CREATE TABLE IF NOT EXISTS migrations (
 name VARCHAR(255)
,UNIQUE(name)
)
`

var migrationInsert = `
INSERT INTO migrations (name) VALUES ($1)
`

var migrationSelect = `
SELECT name FROM migrations
`

var createTableUsers = `
CREATE SEQUENCE IF NOT EXISTS users_seq;

CREATE TABLE IF NOT EXISTS users
(
    user_id int DEFAULT NEXTVAL('users_seq') NOT NULL,
    email VARCHAR NOT NULL,
    token VARCHAR NOT NULL,
    username VARCHAR NOT NULL,
    bio TEXT,
    image VARCHAR,
    created_at Timestamp NOT NULL,
    updated_at Timestamp NOT NULL,
    CONSTRAINT PK_users PRIMARY KEY
    (user_id)
);
`

var createTableComments = `
CREATE SEQUENCE IF NOT EXISTS comments_seq;

CREATE TABLE IF NOT EXISTS comments
(
    comment_id int DEFAULT NEXTVAL ('comments_seq') NOT NULL,
    article_id int NOT NULL,
    body TEXT NOT NULL,
    author VARCHAR NOT NULL,
    created_at Timestamp NOT NULL,
    updated_at Timestamp NOT NULL,
    CONSTRAINT PK_comments PRIMARY KEY
    (comment_id)
);
`
