package user

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/serhatmorkoc/go-realworld-example/model"
	"io"
	"net/http"
	"strconv"
)

func HandlerFind(us model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//w.Header().Set("Content-Type", "application/json")

		val := chi.URLParam(r, "id")

		v, _ := strconv.ParseInt(val, 10, 64)
		user, err := us.Find(v)
		if err != nil {
			render.JSON(w, r, map[string]interface{}{
				"success": false,
				"message": []string{err.Error()},
				"data":    []interface{}{},
				"code":    http.StatusBadRequest,
			})
			return
		}

		//TODO: user boş gelirse json da nil olarak gösteriliyor.
		render.JSON(w, r, map[string]interface{}{
			"success": true,
			"errors":  []interface{}{},
			"data":    []interface{}{user},
			"code":    http.StatusOK,
		})
	}
}

func HandlerGetByEmail(us model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		val := chi.URLParam(r, "email")

		user, err := us.GetByEmail(val)
		if err != nil {
			render.JSON(w, r, map[string]interface{}{
				"success": false,
				"errors":  []interface{}{err.Error()},
				"data":    []interface{}{},
				"code":    http.StatusOK,
			})
			return
		}

		//TODO: user boş gelirse json da nil olarak gösteriliyor.
		render.JSON(w, r, map[string]interface{}{
			"success": true,
			"errors":  []interface{}{},
			"data":    []interface{}{user},
			"code":    http.StatusOK,
		})
	}
}

func HandlerList(us model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		users, err := us.List()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, map[string]interface{}{
				"success": false,
				"message": []interface{}{err.Error()},
				"data":    []interface{}{},
				"code":    http.StatusBadRequest,
			})
			return
		}

		//TODO: user boş gelirse json da nil olarak gösteriliyor.
		render.JSON(w, r, map[string]interface{}{
			"success": true,
			"errors":  []interface{}{},
			"data":    users,
			"code":    http.StatusOK,
		})
	}
}

func HandlerListRange(us model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		users, err := us.ListRange(model.UserParams{
			Sort: true,
			Page: 10,
			Size: 5,
		})
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, map[string]interface{}{
				"success": false,
				"message": []interface{}{err.Error()},
				"data":    []interface{}{},
				"code":    http.StatusBadRequest,
			})
			return
		}

		//TODO: user boş gelirse json da nil olarak gösteriliyor.
		render.JSON(w, r, map[string]interface{}{
			"success": true,
			"errors":  []interface{}{},
			"data":    users,
			"code":    http.StatusOK,
		})
	}
}

func HandlerCreate(us model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		body, err := io.ReadAll(r.Body)
		defer r.Body.Close()

		var user model.User
		if err := json.Unmarshal(body, &user); err != nil {
			render.JSON(w, r, map[string]interface{}{
				"success": false,
				"errors":  []interface{}{err.Error()},
				"data":    []interface{}{},
				"code":    http.StatusBadRequest,
			})
			return
		}

		affected, err := us.Create(&user)
		if err != nil {
			render.JSON(w, r, map[string]interface{}{
				"success": false,
				"errors":  []interface{}{err.Error()},
				"data":    []interface{}{},
				"code":    http.StatusInternalServerError,
			})
			return
		}

		if affected == 0 {
			render.JSON(w, r, map[string]interface{}{
				"success": false,
				"errors":  []interface{}{"operation failed"},
				"data":    []interface{}{},
				"code":    http.StatusInternalServerError,
			})
			return
		}

		render.JSON(w, r, map[string]interface{}{
			"success": true,
			"errors":  []interface{}{},
			"data":    []interface{}{user},
			"code":    http.StatusCreated,
		})
	}
}
