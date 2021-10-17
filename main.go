package main

import (
	"context"
	"fmt"
	"github.com/serhatmorkoc/go-realworld-example/config"
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
	Config     *config.Config
}

func main() {

	//logrus.SetFormatter(new(logrus.JSONFormatter))

	cfg := config.NewConfig()

	logo := os.Getenv("CONSOLE_L")
	fmt.Println(logo)
	fmt.Printf("driver: %s\n", cfg.Database.Driver)
	fmt.Printf("host: %s\n", cfg.Database.Host)
	fmt.Printf("port: %s\n", cfg.Database.Port)
	fmt.Printf("database: %s\n", cfg.Database.Name)
	fmt.Printf("username: %s\n", cfg.Database.User)
	fmt.Printf("password: %s\n", cfg.Database.Password)
	fmt.Println("------------------------------------")

	db, err := database.Connect(cfg)
	if err != nil {
		logrus.Fatalf("error occured on database connection: %s", err.Error())
	}

	us := store.NewUserStore(db)
	cs := store.NewCommentStore(db)
	as := store.NewArticleStore(db)

	sd, _ := strconv.ParseBool(os.Getenv("DB_SEED"))
	if sd {
		if err = seed.Seed(us); err != nil {
			logrus.Fatalf("error occurred on database seed: %s", err.Error())
		}
	}

	r := api.New(us, cs, as)
	h := r.Handler()

	logrus.Info("application starting")

	var srv Server
	go func() {
		if err = srv.Run(h); err != nil {
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

	logrus.Info("application shut down")

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on database connection close: %s", err.Error())
	}
}

func (s *Server) Run(handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + "3000",
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
