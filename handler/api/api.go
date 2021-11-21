package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/serhatmorkoc/go-realworld-example/handler/api/article"
	"github.com/serhatmorkoc/go-realworld-example/handler/api/user"
	middleware1 "github.com/serhatmorkoc/go-realworld-example/middleware"
	"github.com/serhatmorkoc/go-realworld-example/model"
	"net/http"
)

type Server struct {
	Users    model.UserStore
	Comments model.CommentStore
	Articles model.ArticleStore
}

func New(users model.UserStore, comments model.CommentStore, articles model.ArticleStore) Server {
	return Server{
		Users:    users,
		Comments: comments,
		Articles: articles,
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

	/*
		r.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

			// documentation for developers
			opts := sw.SwaggerUIOpts{SpecURL: "/swagger.yaml"}
			sh := sw.SwaggerUI(opts, nil)
			r.Handle("/docs", sh)

			opts1 := sw.RedocOpts{SpecURL: "/swagger.yaml", Path: "docs1"}
			sh1 := sw.Redoc(opts1, nil)
			r.Handle("/docs1", sh1)*/

	r.Route("/api/users", func(r chi.Router) {
		r.Post("/login", user.HandleAuthentication(s.Users))
		r.Post("/", user.HandleRegistration(s.Users))
	})

	r.Route("/api/user", func(r chi.Router) {
		r.Use(middleware1.ValidateJWT)
		r.Get("/", user.HandleCurrentUser(s.Users))
		r.Put("/", user.HandleUpdate(s.Users))
	})

	r.Route("/api/profiles", func(r chi.Router) {
		r.Get("/{username}", user.HandleProfile(s.Users))
		r.With(middleware1.ValidateJWT).Post("/{username}/follow", user.HandleFollowUser(s.Users))
	})

	r.Route("/api/articles", func(r chi.Router) {
		r.Get("/", article.HandleArticleList(s.Articles))
	})

	return r
}
