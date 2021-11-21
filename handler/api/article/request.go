package article

type createArticleRequest struct {
	Article struct{
		Title string `json:"title"`
		Description string `json:"description"`
		Body string `json:"body"`
		TagList []string `json:"tagList"`
	} `json:"article"`
}
