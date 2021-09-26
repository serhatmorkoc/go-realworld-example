package store

import (
	"database/sql"
	"fmt"
	"github.com/serhatmorkoc/go-realworld-example/model"
	"runtime/debug"
)

func NewUserStore(db *sql.DB) model.UserStore {
	return &userStore{
		db: db,
	}
}

type userStore struct {
	db *sql.DB
	//ctx context.Context
}

func (us *userStore) Find(id int64) ([]model.User, error) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	rows, err := us.db.Query("SELECT * FROM users where user_id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	//users := []*model.User{}

	for rows.Next() {
		var user model.User
		//var user = new(model.User)

		err = rows.Scan(
			&user.UserId,
			&user.Email,
			&user.Token,
			&user.UserName,
			&user.Bio,
			&user.Image,
			&user.CreatedAt,
			&user.UpdatedAt)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (us *userStore) GetByEmail(s string) (*model.User, error) {

	defer func() {
		if err  := recover(); err  != nil {
			fmt.Println(err )
			debug.PrintStack()
		}
	}()

	var user model.User
	err := us.db.QueryRow("SELECT * FROM users where email = $1 LIMIT 1", s).Scan(
		&user.UserId,
		&user.Email,
		&user.Token,
		&user.UserName,
		&user.Bio,
		&user.Image,
		&user.CreatedAt,
		&user.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (us *userStore) GetByUsername(s string) (*model.User, error) {
	panic("implement me")
}

func (us *userStore) Create(user *model.User) error {
	panic("implement me")
}

func (us *userStore) Update(user *model.User) error {
	panic("implement me")
}

func (us *userStore) Delete(user *model.User) error {
	panic("implement me")
}

func (us *userStore) List(user *model.User) error {
	panic("implement me")
}

func (us *userStore) ListRange(user *model.User) error {
	panic("implement me")
}

func (us *userStore) AddFollower(user *model.User, followerID uint) error {
	panic("implement me")
}

func (us *userStore) RemoveFollower(user *model.User, followerID uint) error {
	panic("implement me")
}

func (us *userStore) IsFollower(userID, followerID uint) (bool, error) {
	panic("implement me")
}
