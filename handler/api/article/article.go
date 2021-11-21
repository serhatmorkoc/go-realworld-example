package article

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

type singleArticleResponse struct {
	Article struct {
		Slug           string    `json:"slug"`
		Title          string    `json:"title"`
		Description    string    `json:"description"`
		Body           string    `json:"body"`
		TagList        []string  `json:"tagList"`
		CreatedAt      time.Time `json:"createdAt"`
		UpdatedAt      time.Time `json:"updatedAt"`
		Favorited      bool      `json:"favorited"`
		FavoritesCount int       `json:"favoritesCount"`
		Author         struct {
			Username  string `json:"username"`
			Bio       string `json:"bio"`
			Image     string `json:"image"`
			Following bool   `json:"following"`
		} `json:"author"`
	} `json:"article"`
}

var multipleArticleResponse []singleArticleResponse

func HandleArticleList(as model.ArticleStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		slug :=r.URL.Query().Get("slug")


		favorited := chi.URLParam(r, "favortied")
		limit, err := strconv.Atoi(chi.URLParam(r, "limit"))
		if err != nil {
			limit = 20
		}

		if limit == 0 {
			limit = 20
		}

		offset, err := strconv.Atoi(chi.URLParam(r, "offset"))
		if err != nil {
			offset = 0
		}

		articles, err := as.GetAll(slug, "", favorited, limit, offset)
		if err != nil {
			render.NotFound(w, err)
			return
		}

		render.JSON(w,articles, http.StatusOK)
		return

	}
}

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
