package model

import "time"

type UserStore interface {
	Find(id int64) ([]*User, error)
	GetByEmail(string) (*User, error)
	GetByUsername(string) (*User, error)
	Create(*User) error
	Update(*User) error
	Delete(*User) error
	List(*User) error
	ListRange(*User) error
	AddFollower(user *User, followerID uint) error
	RemoveFollower(user *User, followerID uint) error
	IsFollower(userID, followerID uint) (bool, error)
}

type User struct {
	UserId    int64  `json:"user_id"`
	Email     string `json:"email"`
	Token     string
	UserName  string
	Bio       *string
	Image     *string
	CreatedAt time.Time
	UpdatedAt time.Time
}
