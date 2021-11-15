package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/serhatmorkoc/go-realworld-example/handler/api"
	as "github.com/serhatmorkoc/go-realworld-example/store/article"
	cs "github.com/serhatmorkoc/go-realworld-example/store/comment"
	"github.com/serhatmorkoc/go-realworld-example/store/shared/db"
	"github.com/serhatmorkoc/go-realworld-example/store/shared/db/seed"
	us "github.com/serhatmorkoc/go-realworld-example/store/user"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	logo := os.Getenv("CONSOLE_L")
	fmt.Println(logo)
	driver := os.Getenv("DB_DRIVER")
	host := os.Getenv("DB_HOST")
	database := os.Getenv("DB_DATABASE")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	maxOpenConnections, _ := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNECTIONS"))

	db, err := db.Connection(driver, host, database, user, password, port, maxOpenConnections)
	if err != nil {
		log.Fatal(err)
	}

	articleStore := as.NewArticleStore(db)
	userStore := us.NewUserStore(db)
	commentStore := cs.NewCommentStore(db)

	r := api.New(userStore, commentStore, articleStore)
	h := r.Handler()

	sd, _ := strconv.ParseBool(os.Getenv("DB_SEED"))
	if sd {
		if err = seed.Seed(userStore); err != nil {
			logrus.Fatalf("error occurred on database seed: %s", err.Error())
		}
	}

	logrus.Info("application starting")

	log.Println("application starting")

	go func() {
		s := http.Server{
			Addr:           ":8080",
			Handler:        h,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20, //1mb
		}

		err := s.ListenAndServe()
		if err != nil {
			log.Println("application failed to start")
			panic(err)
		}
	}()
	log.Println("application started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Info("application shutting down")

	log.Println("database closing")
	if err := db.Close(); err != nil {
		panic(err)
	}
	log.Println("database closed")
}
