package model

import (
	"errors"
	"time"
)

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

type UserParams struct {
	Page int64
	Size int64
}

type User struct {
	UserId    int64     `json:"user_id"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	UserName  string    `json:"userName"`
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

	return nil
}
