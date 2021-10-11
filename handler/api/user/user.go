package user

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/serhatmorkoc/go-realworld-example/handler/render"
	"github.com/serhatmorkoc/go-realworld-example/model"
	"io"
	"net/http"
	"strconv"
	"time"
)

func HandlerCreate(us model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req model.UserRequest
		body, err := io.ReadAll(r.Body)
		defer r.Body.Close()

		if err := json.Unmarshal(body, &req); err != nil {
			render.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}

		user := model.User{
			UserName:  req.Username,
			Email:     req.Email,
			Password:  req.Password,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		affected, err := us.Create(&user)
		if err != nil {
			render.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}

		if affected == 0 {
			render.ErrorJSON(w, model.ErrOperationFailed, http.StatusBadRequest)
			return
		}

		res := model.Response{
			User: model.UserResponse{
				Username: user.UserName,
				Email:    user.Email,
				Image:    user.Image,
				Bio:      user.Bio,
				Token:    "token",
			}}

		render.SingleSuccessJSON(w, res)

	}
}

func HandlerFind(us model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		val := chi.URLParam(r, "id")

		v, _ := strconv.ParseInt(val, 10, 64)
		user, err := us.GetById(v)
		if err != nil {
			render.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}

		render.SingleSuccessJSON(w, user)
	}
}

func HandlerList(us model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		users, err := us.GetAll()
		if err != nil {
			render.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}

		render.MultipleSuccessJSON(w, users)
	}
}

func HandlerListRange(us model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		users, err := us.GetAllRange(
			model.UserParams{
				Page: 10,
				Size: 5,
			})
		if err != nil {
			render.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}

		render.MultipleSuccessJSON(w, users)
	}
}
