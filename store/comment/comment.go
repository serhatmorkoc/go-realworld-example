package comment

import (
	"github.com/serhatmorkoc/go-realworld-example/model"
	"github.com/serhatmorkoc/go-realworld-example/store/shared/db"
)

type commentStore struct {
	db *db.DB
}

func NewCommentStore(db *db.DB) model.CommentStore {
	return &commentStore{
		db: db,
	}
}

func (cs *commentStore) GetAllBySlug(slug string) ([]*model.Comment, error) {
	panic("implement me")

	//var comments []*model.Comment
	//
	//rows, err := cs.db.Query("SELECT * FROM comments WHERE")
	//if err != nil {
	//	return nil, err
	//}
	//defer rows.Close()
	//
	//for rows.Next() {
	//	var comment model.Comment
	//
	//	err = rows.Scan(
	//		&comment.CommentId,
	//		&comment.ArticleId,
	//		&comment.Body,
	//		&comment.Author,
	//		&comment.CreatedAt,
	//		&comment.UpdatedAt)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	comments = append(comments, &comment)
	//}
	//
	//if err := rows.Err(); err != nil {
	//	return nil, err
	//}
	//
	//return comments, nil
}

func (cs *commentStore) GetByID(id int64) (*model.Comment, error) {

	panic("implement me")

	//var comment model.Comment
	//err := cs.db.QueryRow("SELECT * FROM comments WHERE comment_id=$1 LIMIT 1", id).Scan(
	//	&comment.CommentId,
	//	&comment.ArticleId,
	//	&comment.Body,
	//	&comment.Author,
	//	&comment.CreatedAt,
	//	&comment.UpdatedAt,
	//)
	//if err == sql.ErrNoRows {
	//	return nil, nil
	//}
	//
	//if err != nil {
	//	return nil, err
	//}
	//
	//return &comment, nil
}

func (cs *commentStore) Create(comment *model.Comment) (int64, error) {

	panic("implement me")

	//tx, err := cs.db.Begin()
	//if err != nil {
	//	return 0, err
	//}
	//
	//query := "INSERT INTO comments (article_id, body, author, created_at, updated_at) VALUES($1,$2,$3,$4,$5)"
	//
	//result, execErr := tx.Exec(query, comment.ArticleId, comment.Body, comment.Author, comment.CreatedAt, comment.UpdatedAt)
	//if execErr != nil {
	//	rollbackErr := tx.Rollback()
	//	if rollbackErr != nil {
	//		return 0, rollbackErr
	//	}
	//	return 0, execErr
	//}
	//
	//if err := tx.Commit(); err != nil {
	//	return 0, err
	//}
	//
	//rowsAffected, err := result.RowsAffected()
	//if err != nil {
	//	return 0, nil
	//}
	//
	//return rowsAffected, nil
}

func (cs *commentStore) Delete(id int64) (int64, error) {

	panic("implement me")

	//tx, err := cs.db.Begin()
	//if err != nil {
	//	return 0, err
	//}
	//
	//query := "DELETE FROM comments where comment_id = $1"
	//
	//result, execErr := tx.Exec(query, id)
	//if execErr != nil {
	//	rollbackErr := tx.Rollback()
	//	if rollbackErr != nil {
	//		return 0, rollbackErr
	//	}
	//	return 0, execErr
	//}
	//
	//if err := tx.Commit(); err != nil {
	//	fmt.Println(err)
	//	return 0, err
	//}
	//
	//rowsAffected, err := result.RowsAffected()
	//if err != nil {
	//	return 0, nil
	//}
	//
	//return rowsAffected, nil
}
