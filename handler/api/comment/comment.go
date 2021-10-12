package comment

import (
	"github.com/go-chi/chi"
	"github.com/serhatmorkoc/go-realworld-example/handler/render"
	"github.com/serhatmorkoc/go-realworld-example/model"
	"net/http"
	"strconv"
)

func HandlerDelete(store model.CommentStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		val := chi.URLParam(r, "id")

		v, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			render.BadRequest(w, err)
			return
		}

		result, err := store.Delete(v)
		if err != nil {
			render.BadRequest(w, err)
			return
		}

		render.JSON(w, result, http.StatusOK)
	}
}