package model

type Article struct {

	ArticleId      int64
	Slug           string
	Title          string
	Description    string
	Body           string
	TagList        []string
	CreatedAt      int64
	UpdatedAt      int64
	FavoritesCount int64
	Author         string
	Dummy          byte
}
