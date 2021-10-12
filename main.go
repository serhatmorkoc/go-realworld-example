package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/serhatmorkoc/go-realworld-example/database"
	"github.com/serhatmorkoc/go-realworld-example/database/seed"
	"github.com/serhatmorkoc/go-realworld-example/handler/api"
	"github.com/serhatmorkoc/go-realworld-example/store"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func main() {

	//logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	driver := os.Getenv("DB_DRIVER")
	host := os.Getenv("DB_HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dbName := os.Getenv("DB_DATABASE")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	sd, _ := strconv.ParseBool(os.Getenv("DB_SEED"))
	logo := os.Getenv("CONSOLE_L")
	maxConnections, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))

	fmt.Println(logo)
	fmt.Printf("driver: %s\n", driver)
	fmt.Printf("host: %s\n", host)
	fmt.Printf("port: %d\n", port)
	fmt.Printf("database: %s\n", dbName)
	fmt.Printf("username: %s\n", username)
	fmt.Printf("password: %s\n", password)
	fmt.Printf("seed: %t\n", sd)
	fmt.Println("------------------------------------")

	db, err := database.Connect(driver, host, dbName, username, password, port, maxConnections)
	if err != nil {
		logrus.Fatalf("error occured on database connection: %s", err.Error())
	}

	us := store.NewUserStore(db)
	cs := store.NewCommentStore(db)
	as := store.NewArticleStore(db)

	if sd {
		if err = seed.Seed(us); err != nil {
			logrus.Fatalf("error occurred on database seed: %s", err.Error())
		}
	}

	r := api.New(us, cs, as)
	h := r.Handler()

	var srv Server
	go func() {
		if err = srv.Run("3000",h); err != nil{
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Info("application started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Info("application shutting down")

/*	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}*/

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on database connection close: %s", err.Error())
	}

}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}