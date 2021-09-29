package model

import "time"

type UserStore interface {
	Find(id int64) ([]*User, error)
	GetByEmail(string) (*User, error)
	GetByUsername(string) (*User, error)
	Create(*User) (int64, error)
	Update(*User) (int64, error)
	Delete(*User) error
	List() ([]*User,error)
	ListRange(*User) error
	AddFollower(user *User, followerID uint) error
	RemoveFollower(user *User, followerID uint) error
	IsFollower(userID, followerID uint) (bool, error)
}

type User struct {
	UserId    int64  `json:"user_id"`
	Email     string `json:"email"`
	Token     string `json:"token"`
	UserName  string `json:"userName"`
	Bio       string `json:"bio"`
	Image     string `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
