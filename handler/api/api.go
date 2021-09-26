package api

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/serhatmorkoc/golang-realworld-example/model"
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
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		res,err := s.Users.GetByEmail("email")
		if err != nil {
			fmt.Println(err)
		}
		w.Write([]byte(fmt.Sprintf("hi %v", res)))
	})

	return r
}
