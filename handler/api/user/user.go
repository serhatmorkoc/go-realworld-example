package user

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/serhatmorkoc/go-realworld-example/model"
	"net/http"
	"strconv"
)

func HandlerFind(us model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		val := chi.URLParam(r, "id")

		v, _ := strconv.ParseInt(val, 10, 64)
		res, err := us.Find(v)
		if err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(res)
	}
}

func GetByEmail(us model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		val := chi.URLParam(r, "email")

		res, err := us.GetByEmail(val)
		if err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(res)
	}
}

func HandlerList(us model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		res, err := us.List()
		if err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(res)
	}
}
