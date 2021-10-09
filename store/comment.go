package store

import (
	"database/sql"
	"fmt"
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

	var comments []*model.Comment

	rows, err := cs.db.Query("SELECT * FROM comments WHERE")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment model.Comment

		err = rows.Scan(
			&comment.CommentId,
			&comment.ArticleId,
			&comment.Body,
			&comment.Author,
			&comment.CreatedAt,
			&comment.UpdatedAt)
		if err != nil {
			return nil, err
		}

		comments = append(comments, &comment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
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

func (cs *commentStore) Create(comment *model.Comment) (int64, error) {

	tx, err := cs.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	query := "INSERT INTO comments (article_id, body, author, created_at, updated_at) VALUES($1,$2,$3,$4,$5)"

	result, execErr := tx.Exec(query, comment.ArticleId, comment.Body, comment.Author, comment.CreatedAt, comment.UpdatedAt)
	if execErr != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return 0, rollbackErr
		}

		fmt.Printf("insert failed: %v", execErr)
		return 0, execErr
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, nil
	}

	return rowsAffected, nil
}

func (cs *commentStore) Delete(id int64) (int64, error) {
	tx, err := cs.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	query := "DELETE FROM comments where comment_id = $1"

	result, execErr := tx.Exec(query, id)
	if execErr != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return 0, rollbackErr
		}
		return 0, execErr
	}

	if err := tx.Commit(); err != nil {
		fmt.Println(err)
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, nil
	}

	return rowsAffected, nil
}
