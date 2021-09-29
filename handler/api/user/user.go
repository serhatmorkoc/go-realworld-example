package user

import (
	"encoding/json"
	"fmt"
	"github.com/serhatmorkoc/go-realworld-example/handler/api"
	"net/http"
)

func HandlerList(s api.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		res,err := s.Users.List()
		if err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(res)
	}
}
