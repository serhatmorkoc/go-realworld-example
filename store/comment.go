package store

import (
	"database/sql"
	"github.com/serhatmorkoc/go-realworld-example/model"
)

type commentStore struct {
	db *sql.DB
}

func NewCommentStore(db *sql.DB) model.CommentStore {
	return &commentStore{
		db: db,
	}
}


func (c *commentStore) GetAllBySlug(s string) ([]*model.Comment, error) {
	panic("implement me")
}

func (c *commentStore) GetByID(u uint) (*model.Comment, error) {
	panic("implement me")
}

func (c *commentStore) Create(comment *model.Comment) (int, error) {
	panic("implement me")
}

func (c *commentStore) Delete(comment *model.Comment) (int, error) {
	panic("implement me")
}

