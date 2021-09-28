package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/serhatmorkoc/go-realworld-example/db"
	"github.com/serhatmorkoc/go-realworld-example/db/seed"
	"github.com/serhatmorkoc/go-realworld-example/store"
	"os"
	"strconv"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234563"
	dbname   = "conduit"
)

func main() {

	if err := godotenv.Load("local.env"); err != nil {
		panic("Error loading .env file")
	}

	driver := os.Getenv("DATABASE_DRIVER")
	dsn := os.Getenv("DATABASE_DATASOURCE")
	sd, _ := strconv.ParseBool(os.Getenv("DATABASE_SEED"))
	maxConnections, _ := strconv.Atoi(os.Getenv("DATABASE_MAX_CONNECTIONS"))

	fmt.Println("\n   __________  __    ___    _   ________\n  / ____/ __ \\/ /   /   |  / | / / ____/\n / / __/ / / / /   / /| | /  |/ / / __  \n/ /_/ / /_/ / /___/ ___ |/ /|  / /_/ /  \n\\____/\\____/_____/_/  |_/_/ |_/\\____/   \n                                        ")
	fmt.Printf("data source name: %s\n", dsn)
	fmt.Printf("driver: %s\n", driver)
	fmt.Printf("max connections: %d\n", maxConnections)
	fmt.Printf("seed: %t\n", sd)
	fmt.Println("------------------------------------")

	db, err := db.Connect(driver, dsn, maxConnections)
	if err != nil {
		panic(err)
	}

	us := store.NewUserStore(db)

	if sd {
		seed.Seed(us)
	}

	/*	r :=api.New(us)
		h := r.Handler()

		if err := http.ListenAndServe(":3000", h); err != nil {
			panic(err)
		}*/

}
