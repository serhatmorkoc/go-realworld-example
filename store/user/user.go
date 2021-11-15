package user

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/serhatmorkoc/go-realworld-example/model"
	"github.com/serhatmorkoc/go-realworld-example/store/shared/db"
	"time"
)

type userStore struct {
	db *db.DB
}

func NewUserStore(db *db.DB) model.UserStore {
	return &userStore{
		db: db,
	}
}

func (us *userStore) GetById(id uint) (*model.User, error) {

	var user model.User
	err := us.db.Read(func(execer db.Execer) error {

		err := execer.QueryRow("SELECT * FROM users where user_id = $1 LIMIT 1", id).Scan(
			&user.UserId,
			&user.UserName,
			&user.Email,
			&user.Password,
			&user.Bio,
			&user.Image,
			&user.CreatedAt,
			&user.UpdatedAt)

		return err
	})

	return &user, err
}

func (us *userStore) GetByEmail(email string) (*model.User, error) {

	var user model.User
	err := us.db.Read(func(execer db.Execer) error {
		query := "SELECT * FROM users where email = $1 LIMIT 1"
		err := execer.QueryRow(query, email).Scan(
			&user.UserId,
			&user.Email,
			&user.Password,
			&user.UserName,
			&user.Bio,
			&user.Image,
			&user.CreatedAt,
			&user.UpdatedAt)

		return err
	})

	return &user, err
}

func (us *userStore) GetByUsername(username string) (*model.User, error) {

	var user model.User
	err := us.db.Read(func(execer db.Execer) error {
		query := "SELECT * FROM users where username = $1 LIMIT 1"
		err := execer.QueryRow(query, username).Scan(
			&user.UserId,
			&user.Email,
			&user.Password,
			&user.UserName,
			&user.Bio,
			&user.Image,
			&user.CreatedAt,
			&user.UpdatedAt)

		return err
	})

	return &user, err
}

func (us *userStore) Create(ctx context.Context, user *model.User) (*model.User, error) {

	if err := user.Validate(); err != nil {
		return nil, err
	}

	err := us.db.Update(func(execer db.Execer) error {
		query := "INSERT INTO users (email, password, username, bio, image, created_at, updated_at) VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING user_id"
		_, err := execer.ExecContext(ctx, query, user.Email, user.Password, user.UserName, user.Bio, user.Image, user.CreatedAt, user.UpdatedAt)
		return err
	})

	if user.UserId == 0 {
		return nil, errors.New("user not found")
	}

	return user, err
}

func (us *userStore) Update(ctx context.Context, user *model.User) error {

	err := us.db.Update(func(execer db.Execer) error {
		user.UpdatedAt = time.Now()
		query := "UPDATE users SET email=$1, password=$2, username=$3, bio=$4, image=$5, updated_at=$6 WHERE user_id=$7"
		a, err := execer.ExecContext(ctx, query, user.Email, user.Password, user.UserName, user.Bio, user.Image, user.UpdatedAt, user.UserId)
		fmt.Println(a)
		return err
	})

	return err
}

func (us *userStore) Delete(user *model.User) error {
	panic("implement me")
}

func (us *userStore) GetAll() ([]*model.User, error) {
	panic("implement me")

	//var users []*model.User
	//
	//rows, err := us.db.Query("SELECT * FROM users")
	//if err != nil {
	//	return nil, err
	//}
	//defer rows.Close()
	//
	//for rows.Next() {
	//	var user model.User
	//
	//	err = rows.Scan(
	//		&user.UserId,
	//		&user.Email,
	//		&user.Password,
	//		&user.UserName,
	//		&user.Bio,
	//		&user.Image,
	//		&user.CreatedAt,
	//		&user.UpdatedAt)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	users = append(users, &user)
	//}
	//
	//if err := rows.Err(); err != nil {
	//	return nil, err
	//}
	//
	//return users, nil
}

func (us *userStore) GetAllRange(params model.UserParams) ([]*model.User, error) {

	panic("implement me")

	//query := "SELECT * FROM users LIMIT %d OFFSET %d"
	//
	//var users []*model.User
	//
	//rows, err := us.db.Query(query)
	//if err != nil {
	//	return nil, err
	//}
	//defer rows.Close()
	//
	//for rows.Next() {
	//	var user model.User
	//
	//	err = rows.Scan(
	//		&user.UserId,
	//		&user.Email,
	//		&user.Password,
	//		&user.UserName,
	//		&user.Bio,
	//		&user.Image,
	//		&user.CreatedAt,
	//		&user.UpdatedAt)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	users = append(users, &user)
	//}
	//
	//if err := rows.Err(); err != nil {
	//	return nil, err
	//}
	//
	//return users, nil
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
