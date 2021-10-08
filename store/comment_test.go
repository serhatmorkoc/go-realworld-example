package store

import (
	"context"
	"github.com/serhatmorkoc/go-realworld-example/database/databasetest"
	"github.com/serhatmorkoc/go-realworld-example/model"
	"testing"
	"time"
)

var noContext = context.TODO()

func TestComment(t *testing.T) {

	db, err := databasetest.Connect()
	if err != nil {
		t.Error(err)
		return
	}

	defer func() {
		 databasetest.Reset(db)

		if err := databasetest.Disconnect(db); err != nil {
			t.Error(err)
		}
	}()

	store := NewCommentStore(db).(*commentStore)
	t.Run("Create", testCommentCreate(store))
	t.Run("GetById", testCommentGetById(store))
}

func testCommentCreate(store *commentStore) func(t *testing.T) {
	return func(t *testing.T) {

		comment := &model.Comment{
			ArticleId: 1,
			Body:      "body",
			Author:    "author",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		affectedRow, err := store.Create(comment)
		if err != nil {
			t.Error(err)
		}

		if affectedRow == 0 {
			t.Errorf("Want comment ID assigned, got %d", comment.CommentId)
		}
	}
}

func testCommentGetById(store *commentStore) func(t *testing.T) {
	return func(t *testing.T){
		comment, err := store.GetByID(1)
		if err != nil {
			t.Error(err)
			return
		}

		if comment == nil {
			t.Error("err")
		}
	}
}
