package databasetest

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
	"strconv"
)

func Connect() (*sql.DB, error) {

	if err := godotenv.Load("../.env"); err != nil {
		panic("Error loading .env file")
	}

	driver := os.Getenv("DB_DRIVER")
	host := os.Getenv("DB_HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dbName := os.Getenv("DB_DATABASE")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, username, password, dbName)

	return sql.Open(driver, dsn)
}

func Reset(sql *sql.DB) {
	sql.Exec("truncate table comments; ALTER SEQUENCE comments_seq RESTART WITH 1;")


	//sql.Exec("DELETE FROM users")
	//sql.Exec("DELETE FROM reset")
}

func Disconnect(sql *sql.DB) error {
	return sql.Close()
}
