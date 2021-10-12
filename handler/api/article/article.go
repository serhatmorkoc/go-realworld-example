package article

import (
	"encoding/json"
	"github.com/serhatmorkoc/go-realworld-example/handler/render"
	"github.com/serhatmorkoc/go-realworld-example/model"
	"io"
	"net/http"
)

func HandlerCreate(as model.ArticleStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		body, err := io.ReadAll(r.Body)
		if err != nil {
			//
		}
		r.Body.Close()

		var article model.Article
		if err := json.Unmarshal(body, &article); err != nil {
			render.BadRequest(w, err)
			return
		}

		affected, err := as.Create(&article)
		if err != nil {
			render.BadRequest(w, err)
			return
		}
		if affected == 0 {
			//render.ErrorJSON(w, model.ErrOperationFailed, http.StatusBadRequest)
			render.BadRequest(w, err)
			return
		}

		render.JSON(w, article, http.StatusCreated)
	}
}