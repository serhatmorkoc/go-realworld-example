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
	{
		name: "create-table-articles",
		stmt: createTableArticles,
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
CREATE TABLE IF NOT EXISTS users(
    user_id SERIAL PRIMARY KEY,
    username VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    bio TEXT,
    image VARCHAR,
    created_at Timestamp NOT NULL,
    updated_at Timestamp NOT NULL,
    UNIQUE(username),
    UNIQUE(email)
);
`

var createTableComments = `
CREATE TABLE IF NOT EXISTS comments(
    comment_id SERIAL PRIMARY KEY,
    article_id INTEGER NOT NULL,
    body TEXT NOT NULL,
    author VARCHAR NOT NULL,
    created_at Timestamp NOT NULL,
    updated_at Timestamp NOT NULL
);
`

var createTableArticles = `
CREATE TABLE IF NOT EXISTS articles(
    article_id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    slug VARCHAR NOT NULL,
    title VARCHAR NOT NULL,
    description VARCHAR NOT NULL,
    body TEXT NOT NULL,
    tag_list VARCHAR NOT NULL,
    favorited BOOLEAN NOT NULL,
    favorites_count INTEGER NOT NULL,
    created_at Timestamp NOT NULL,
    updated_at Timestamp NOT NULL
);
`
