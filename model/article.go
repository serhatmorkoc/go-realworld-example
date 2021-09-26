package model

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
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
