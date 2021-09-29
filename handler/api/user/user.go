package user

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"github.com/serhatmorkoc/go-realworld-example/model"
	"io"
	"net/http"
	"strconv"
)

type BaseResponse struct {
	Status  int
	Message string
	Data    interface{}
}

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

func HandlerList(us model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		users, err := us.List()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		var br BaseResponse
		br.Data = users
		br.Status = 200
		br.Message = "test"

		usersJson, err := json.Marshal(br)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}



		w.WriteHeader(http.StatusOK)
		w.Write(usersJson)
	}
}

func HandlerCreate(us model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)


		body, err := io.ReadAll(r.Body)
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				//TODO:
			}
		}(r.Body)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response := make(map[string]string)
		var user model.User
		if err := json.Unmarshal(body, &user); err != nil {

			w.WriteHeader(http.StatusBadRequest)
			response["message"] = err.Error()
			jsonResponse, _ := json.Marshal(response)
			w.Write(jsonResponse)
			return
		}

		affected, err := us.Create(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if affected == 0 {
			http.Error(w, errors.New("unsuccessful").Error(), http.StatusBadRequest)
			return
		}

		response["message"] = "successful"
		jsonResponse, err := json.Marshal(response)
		w.Write(jsonResponse)
	}
}
