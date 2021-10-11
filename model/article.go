package model

import "time"

type ArticleStore interface {
	GetAll(uint64) ([]*Article, error)
	GetById(uint64) (*Article, error)
	Create(*Article) (int64, error)
	Update(*Article) (int64, error)
	Delete(uint64) (int64, error)

}

type Article struct {
	ArticleId      int64
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
