package model

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const MinPasswordLength = 6
const PasswordKeyLength = 64

type UserStore interface {
	GetById(id int64) (*User, error)
	GetByEmail(string) (*User, error)
	GetByUsername(string) (*User, error)
	Create(*User) (int64, error)
	Update(*User) (int64, error)
	Delete(*User) error


	GetAll() ([]*User, error)
	GetAllRange(params UserParams) ([]*User, error)
	AddFollower(user *User, followerID uint) error
	RemoveFollower(user *User, followerID uint) error
	IsFollower(userID, followerID uint) (bool, error)
}

type Request struct {
	User UserRequest `json:"user"`
}

type UserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Response struct {
	User UserResponse `json:"user"`
}

type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Image    string `json:"image"`
	Bio      string `json:"bio"`
	Token    string `json:"token"`
}

type UserParams struct {
	Page int64
	Size int64
}

type User struct {
	UserId    int64     `json:"user_id"`
	UserName  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Bio       string    `json:"bio"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (user *User) Validate() error {

	if len(user.UserName) == 0 {
		return errors.New("username cannot be empty")
	}

	if len(user.Email) == 0 {
		return errors.New("email cannot be empty")
	}

	if len(user.Password) == 0 {
		return errors.New("email cannot be empty")
	}

	return nil
}

func (user *User) HashPassword(password string) (string, error) {
	if len(password) == 0 {
		return "", errors.New("password should not be empty")
	}
	h, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(h), err
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}