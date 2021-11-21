package model

import "time"

type ArticleStore interface {
	GetAll(tag, author, favorited string, limit, offset int) ([]*Article, error)
	GetById(uint64) (*Article, error)
	Create(*Article) (int64, error)
	Update(*Article) (int64, error)
	Delete(uint64) (int64, error)
}

type Article struct {
	ArticleId      int64
	UserId         int64
	Slug           string
	Title          string
	Description    string
	Body           string
	TagList        []string
	Favorited      bool
	FavoritesCount int64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
