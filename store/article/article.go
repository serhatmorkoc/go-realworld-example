package article

import (
	"github.com/serhatmorkoc/go-realworld-example/model"
	"github.com/serhatmorkoc/go-realworld-example/store/shared/db"
)

type articleStore struct {
	db *db.DB
}

func NewArticleStore(db *db.DB) model.ArticleStore {
	return &articleStore{
		db: db,
	}
}

func (as *articleStore) GetAll(u uint64) ([]*model.Article, error) {
	panic("implement me")
}

func (as *articleStore) GetById(u uint64) (*model.Article, error) {
	panic("implement me")
}

func (as *articleStore) Create(article *model.Article) (int64, error) {
	panic("implement me")

	//tx, err := as.db.Begin()
	//if err != nil {
	//	return 0, err
	//}
	//
	//article.CreatedAt = time.Now()
	//article.UpdatedAt = time.Now()
	//
	//query := "INSERT INTO articles(slug, title, description,body,tag_list,favorited,favorites_count,created_at,updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)"
	//result, execErr := tx.Exec(query, article.Slug, article.Title, article.Description, article.Body, pq.Array(article.TagList), article.Favorited, article.FavoritesCount, article.CreatedAt, article.UpdatedAt)
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

func (as *articleStore) Update(article *model.Article) (int64, error) {
	panic("implement me")
}

func (as *articleStore) Delete(u uint64) (int64, error) {
	panic("implement me")
}
