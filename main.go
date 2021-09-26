package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/serhatmorkoc/golang-realworld-example/db"
	"github.com/serhatmorkoc/golang-realworld-example/handler/api"
	"github.com/serhatmorkoc/golang-realworld-example/store"
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

	if err := godotenv.Load("env/local.env"); err != nil {
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

	result, err := us.Find(11)
	if err != nil {
		log.Fatal(err)
	}

	if reflect.ValueOf(result).IsNil() {
		fmt.Println("IsNil")
	}

	gbe, err := us.GetByEmail("tt")
	if err != nil {
		log.Fatal(err)
	}

	if reflect.ValueOf(gbe).IsNil() {
		fmt.Println("IsNil:gbe")
	}

	if gbe == nil {
		fmt.Println("IsNil1")
	}

	r :=api.New(us)
	h := r.Handler()
	http.ListenAndServe(":3000", h)
}
