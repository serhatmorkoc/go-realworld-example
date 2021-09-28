package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/serhatmorkoc/go-realworld-example/db"
	"github.com/serhatmorkoc/go-realworld-example/model"
	"github.com/serhatmorkoc/go-realworld-example/store"
	"log"
	"os"
	"reflect"
	"strconv"
	"time"
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

	/*	exUser := model.User{
			Email:     fmt.Sprintf("email-%d", "email"),
			Token:     fmt.Sprintf("token-%d", "token"),
			UserName:  fmt.Sprintf("username-%d", "username"),
			Bio:       fmt.Sprintf("bio-%d", "bie"),
			Image:     fmt.Sprintf("image-%d", "image"),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		lastInsertedId, err := us.Create(&exUser)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("last inserted id: %d\n", lastInsertedId)*/

	exUser1 := model.User{
		UserId:    138193,
		Email:     fmt.Sprintf("email-%s", "email1122222"),
		Token:     fmt.Sprintf("token-%s", "token"),
		UserName:  fmt.Sprintf("username-%s", "username"),
		Bio:       fmt.Sprintf("bio-%s", "bie"),
		Image:     fmt.Sprintf("image-%s", "image"),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err = us.Update(&exUser1)
	if err != nil {
		fmt.Println(err)
	}

	/*	r :=api.New(us)
		h := r.Handler()

		if err := http.ListenAndServe(":3000", h); err != nil {
			panic(err)
		}*/

}
