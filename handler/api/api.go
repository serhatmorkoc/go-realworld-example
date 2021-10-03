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

	//log := logrus.New()
	//r.Use(logger.Logger("router", log))

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.NoCache)

	r.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// documentation for developers
/*	opts := sw.SwaggerUIOpts{SpecURL: "/swagger.yaml"}
	sh := sw.SwaggerUI(opts, nil)
	r.Handle("/docs", sh)

	opts1 := sw.RedocOpts{SpecURL: "/swagger.yaml", Path: "docs1"}
	sh1 := sw.Redoc(opts1, nil)
	r.Handle("/docs1", sh1)*/

	r.Route("/user", func(r chi.Router) {

		r.Get("/list", user.HandlerList(s.Users))
		r.Get("/list/range", user.HandlerListRange(s.Users))
		r.Get("/id/{id}", user.HandlerFind(s.Users))
		r.Post("/create", user.HandlerCreate(s.Users))

	})

	return r
}

