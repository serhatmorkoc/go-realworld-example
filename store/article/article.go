package article

import (
	"fmt"
	"github.com/lib/pq"
	"github.com/pkg/errors"
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

func (as *articleStore) GetAll(tag, author, favorited string, limit, offset int) ([]*model.Article, error) {

	query := "SELECT * FROM articles WHERE 1=1 "

	if len(tag) != 0 {
		query = fmt.Sprintf(query + "AND slug='%s' " , tag)
	}
	if len(author) != 0 {
		query = fmt.Sprintf(query + "AND WHERE author='%s' ", author)
	}

	query = fmt.Sprintf(query + "LIMIT %d ",limit)
	query = fmt.Sprintf(query + "OFFSET %d ", offset)

	var articles []*model.Article
	err := as.db.Read(func(execer db.Execer) error {
		rows, err := execer.Query(query)
		if err != nil {
			return err
		}

		var article model.Article
		for rows.Next() {
			 err = rows.Scan(&article.ArticleId,
				 &article.UserId,
				 &article.Title,
				 &article.Description,
				 &article.Body,
				 pq.Array(&article.TagList),
				 &article.CreatedAt,
				 &article.UpdatedAt)
			 if err != nil {
				 return err
			 }
			 articles = append(articles,&article)
		}

		return err
	})

	if len(articles) == 0 {
		return nil, errors.New("sql: no rows in result set")
	}
	return articles, err

}

func (as *articleStore) GetById(u uint64) (*model.Article, error) {
	panic("implement me")
}

func (as *articleStore) Create(article *model.Article) error {
	panic("implement me")
}

func (as *articleStore) Update(article *model.Article) (int64, error) {
	panic("implement me")
}

func (as *articleStore) Delete(u uint64) (int64, error) {
	panic("implement me")
}
