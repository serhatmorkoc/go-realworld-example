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


func (cs *commentStore) GetAllBySlug(s string) ([]*model.Comment, error) {
	panic("implement me")
}

func (cs *commentStore) GetByID(id uint64) (*model.Comment, error) {

	var comment model.Comment
	err := cs.db.QueryRow("SELECT * FROM comments WHERE comment_id=$1 LIMIT 1", id).Scan(
		&comment.CommentId,
		&comment.ArticleId,
		&comment.Body,
		&comment.Author,
		&comment.CreatedAt,
		&comment.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &comment, nil
}

func (cs *commentStore) Create(comment *model.Comment) (int, error) {
	panic("implement me")
}

func (cs *commentStore) Delete(comment *model.Comment) (int, error) {
	panic("implement me")
}

