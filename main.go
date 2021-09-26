package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/serhatmorkoc/go-realworld-example/db"
	"github.com/serhatmorkoc/go-realworld-example/handler/api"
	"github.com/serhatmorkoc/go-realworld-example/store"
	"log"
	"net/http"
	"os"
	"reflect"
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
	maxConnectionsString := os.Getenv("DATABASE_MAX_CONNECTIONS")
	maxConnections, _ := strconv.Atoi(maxConnectionsString)

	fmt.Printf("data source name: %s\n", dsn)
	fmt.Printf("driver: %s\n", driver)
	fmt.Printf("max connections: %d\n", maxConnections)

	db, err := db.Connect(driver, dsn, maxConnections)
	if err != nil {
		panic(err)
	}
	us := store.NewUserStore(db)

	user1, err := us.Find(11)
	if err != nil {
		log.Fatal(err)
	}

	if reflect.ValueOf(user1).IsNil() {
		fmt.Println("IsNil:user1")
	}

	user2, err := us.GetByEmail("email")
	if err != nil {
		log.Fatal(err)
	}

	if reflect.ValueOf(user2).IsNil() {
		fmt.Println("IsNil:user2")
	}


	r :=api.New(us)
	h := r.Handler()

	if err := http.ListenAndServe(":3000", h); err != nil {
		panic(err)
	}

}
