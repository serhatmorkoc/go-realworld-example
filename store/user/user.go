package user

import (
	"context"
	"github.com/pkg/errors"
	"github.com/serhatmorkoc/go-realworld-example/model"
	"github.com/serhatmorkoc/go-realworld-example/store/shared/db"
)

type userStore struct {
	db *db.DB
}

func NewUserStore(db *db.DB) model.UserStore {
	return &userStore{
		db: db,
	}
}

func (us *userStore) GetById(id int64) (*model.User, error) {

	var user model.User
	err := us.db.Read(func(execer db.Execer) error {

		err := execer.QueryRow("SELECT * FROM users where user_id = $1 LIMIT 1", id).Scan(
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

func (us *userStore) Update(user *model.User) (int64, error) {

	panic("implement me")

	//tx, err := us.db.Begin()
	//if err != nil {
	//	return 0, err
	//}
	//
	//query := "UPDATE users SET email=:email, password=:password, username=:username, bio=:bio, image=:image, created_at=:created_at, updated_at=:updated_at  WHERE user_id=:user_id RETURNING user_id"
	//
	//result, execErr := tx.Exec(query,
	//	sql.Named("email", user.Email),
	//	sql.Named("password", user.Password),
	//	sql.Named("username", user.UserName),
	//	sql.Named("bio", user.Bio),
	//	sql.Named("image", user.Image),
	//	sql.Named("created_at", user.CreatedAt),
	//	sql.Named("updated_at", user.UpdatedAt),
	//	sql.Named("user_id", user.UserId))
	//
	//if execErr != nil {
	//	rollbackErr := tx.Rollback()
	//	if rollbackErr != nil {
	//		return 0, rollbackErr
	//	}
	//
	//	return 0, execErr
	//}
	//
	//if err := tx.Commit(); err != nil {
	//	return 0, err
	//}
	//
	//rowsAffected, err := result.RowsAffected()
	//if err != nil {
	//	return 1, nil
	//}
	//
	//return rowsAffected, nil
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
