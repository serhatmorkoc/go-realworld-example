package store

import (
	"database/sql"
	"github.com/pkg/errors"
	"github.com/serhatmorkoc/go-realworld-example/model"
)

type userStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) model.UserStore {
	return &userStore{
		db: db,
	}
}

func (us *userStore) GetById(id int64) (*model.User, error) {

	var user model.User
	err := us.db.QueryRow("SELECT * FROM users where user_id = $1 LIMIT 1", id).Scan(
		&user.UserId,
		&user.Email,
		&user.Password,
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

func (us *userStore) GetByEmail(s string) (*model.User, error) {

	var user model.User
	err := us.db.QueryRow("SELECT * FROM users where email = $1 LIMIT 1", s).Scan(
		&user.UserId,
		&user.Email,
		&user.Password,
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

	var user model.User
	err := us.db.QueryRow("SELECT * FROM users where username = $1 LIMIT 1", s).Scan(
		&user.UserId,
		&user.Email,
		&user.Password,
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

func (us *userStore) Create(user *model.User) (*model.User, error) {

	if err := user.Validate(); err != nil {
		return nil, err
	}

	tx, err := us.db.Begin()
	if err != nil {
		return nil, err
	}

	query := "INSERT INTO users (email, password, username, bio, image, created_at, updated_at) VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING user_id"
	row := tx.QueryRow(query, user.Email, user.Password, user.UserName, user.Bio, user.Image, user.CreatedAt, user.UpdatedAt)
	err = row.Scan(&user.UserId)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return nil, rollbackErr
		}
		return nil, err
	}

	if user.UserId == 0 {
		return nil, errors.New("")
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return user, nil
}

func (us *userStore) Update(user *model.User) (int64, error) {

	tx, err := us.db.Begin()
	if err != nil {
		return 0, err
	}

	query := "UPDATE users SET email=:email, password=:password, username=:username, bio=:bio, image=:image, created_at=:created_at, updated_at=:updated_at  WHERE user_id=:user_id RETURNING user_id"

	result, execErr := tx.Exec(query,
		sql.Named("email", user.Email),
		sql.Named("password", user.Password),
		sql.Named("username", user.UserName),
		sql.Named("bio", user.Bio),
		sql.Named("image", user.Image),
		sql.Named("created_at", user.CreatedAt),
		sql.Named("updated_at", user.UpdatedAt),
		sql.Named("user_id", user.UserId))

	if execErr != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return 0, rollbackErr
		}

		return 0, execErr
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 1, nil
	}

	return rowsAffected, nil
}

func (us *userStore) Delete(user *model.User) error {
	panic("implement me")
}

func (us *userStore) GetAll() ([]*model.User, error) {

	var users []*model.User

	rows, err := us.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User

		err = rows.Scan(
			&user.UserId,
			&user.Email,
			&user.Password,
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

func (us *userStore) GetAllRange(params model.UserParams) ([]*model.User, error) {

	query := "SELECT * FROM users LIMIT %d OFFSET %d"

	var users []*model.User

	rows, err := us.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User

		err = rows.Scan(
			&user.UserId,
			&user.Email,
			&user.Password,
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

func (us *userStore) AddFollower(user *model.User, followerID uint) error {
	panic("implement me")
}

func (us *userStore) RemoveFollower(user *model.User, followerID uint) error {
	panic("implement me")
}

func (us *userStore) IsFollower(userID, followerID uint) (bool, error) {
	panic("implement me")
}
