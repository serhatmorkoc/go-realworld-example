package model

import "time"

type CommentStore interface {
	GetAllBySlug(string) ([]*Comment, error)
	GetByID(uint64) (*Comment, error)
	Create(*Comment) (int64, error)
	Delete(*Comment) (int64, error)
}

type Comment struct {
	CommentId uint64    `json:"comment_id"`
	ArticleId uint64    `json:"article_id"`
	Body      string    `json:"body"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (comment *Comment) Validate() error {

	return nil
}
