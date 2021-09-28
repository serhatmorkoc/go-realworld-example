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

func (us *userStore) Find(id int64) ([]*model.User, error) {

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

	var users []*model.User
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

		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (us *userStore) GetByEmail(s string) (*model.User, error) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
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

func (us *userStore) Create(user *model.User) (int64, error) {

	tx, err := us.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	stmt, err := us.db.Prepare("INSERT INTO public.users (email, token, username, bio, image, created_at, updated_at) " +
		"VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING user_id")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var lastInsertId int64
	if err := stmt.QueryRow(user.Email, user.Token, user.UserName, user.Bio, user.Image, user.CreatedAt, user.UpdatedAt).
		Scan(&lastInsertId); err != nil {
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return lastInsertId, nil
}

func (us *userStore) Update(user *model.User) (int64, error) {
	tx, err := us.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	stmt, err := us.db.Prepare("UPDATE users SET email=$2, token=$3, username=$4, bio=$5, image=$6, created_at=$7, updated_at=$8 " +
		"WHERE user_id=$1 RETURNING user_id")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.UserId, user.Email, user.Token, user.UserName, user.Bio, user.Image, user.CreatedAt, user.UpdatedAt)

	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	if rows != 0 {
		return rows, nil
	}

	return 0,nil
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
