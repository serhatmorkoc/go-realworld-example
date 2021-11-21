package article

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/serhatmorkoc/go-realworld-example/handler/render"
	"github.com/serhatmorkoc/go-realworld-example/model"
	"github.com/serhatmorkoc/go-realworld-example/service/auth"
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

		slug := r.URL.Query().Get("slug")

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

		render.JSON(w, articles, http.StatusOK)
		return

	}
}

func HandlerCreate(as model.ArticleStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		userId, err := auth.GetUserId(r.Context())
		if err != nil {
			render.Unauthorized(w, err)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			render.BadRequest(w, err)
			return
		}
		r.Body.Close()

		var request createArticleRequest
		if err := json.Unmarshal(body, &request); err != nil {
			render.BadRequest(w, err)
			return
		}

		a := request.createRequestToModel()
		a.UserId = int64(userId)
		if err := as.Create(a); err != nil {
			render.BadRequest(w, err)
			return
		}

		render.JSON(w, a, http.StatusCreated)
	}
}

func (req createArticleRequest) createRequestToModel() *model.Article {

	return &model.Article{
		Title: req.Article.Title,
	}

}
