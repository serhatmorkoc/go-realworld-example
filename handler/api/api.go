package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/serhatmorkoc/go-realworld-example/handler/api/user"
	"github.com/serhatmorkoc/go-realworld-example/model"
	"net/http"
)

type Server struct {
	Users model.UserStore
}

func New(users model.UserStore) Server {
	return Server{
		Users: users,
	}
}

func (s Server) Handler() http.Handler {
	r := chi.NewRouter()

	/*"github.com/goware/cors"
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	})

	r.Use(cors.Handler)*/

	r.Use(middleware.Recoverer)
	r.Use(middleware.NoCache)
	r.Use(middleware.Logger)

	r.Route("/user", func(r chi.Router) {

		r.Get("/list", user.HandlerList(s.Users))
		r.Get("/list/range", user.HandlerListRange(s.Users))
		r.Get("/id/{id}", user.HandlerFind(s.Users))
		r.Post("/create", user.HandlerCreate(s.Users))


	})

	return r
}
