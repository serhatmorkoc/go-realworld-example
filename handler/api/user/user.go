package user

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/serhatmorkoc/go-realworld-example/handler/render"
	"github.com/serhatmorkoc/go-realworld-example/model"
	"github.com/serhatmorkoc/go-realworld-example/service/auth"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

//Request

type LoginUserRequest struct {
	User struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password"`
	}
}

type CreateUserRequest struct {
	User struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	} `json:"user"`
}

type UpdateUserRequest struct {
	User struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
		Bio      string `json:"bio"`
		Image    string `json:"image"`
	} `json:"user"`
}

type ProfileRequest struct {
	Username string `json:"username"`
}

//Response

type LoginUserResponse struct {
	User BaseUserResponse
}

type CreateUserResponse struct {
	User struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Image    string `json:"image"`
		Bio      string `json:"bio"`
		Token    string `json:"token"`
	} `json:"user"`
}

type ProfileResponse struct {
	Profile struct {
		Username  string `json:"username"`
		Bio       string `json:"bio"`
		Image     string `json:"image"`
		Following bool   `json:"following"`
	} `json:"profile"`
}

type BaseUserResponse struct {
	User struct {
		Email    string `json:"email"`
		Token    string `json:"token"`
		Username string `json:"username"`
		Bio      string `json:"bio"`
		Image    string `json:"image"`
	} `json:"user"`
}

func HandleCreate(ctx context.Context, us model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req CreateUserRequest
		body, err := io.ReadAll(r.Body)
		if err != nil {
			render.BadRequest(w, err)
			return
		}
		defer r.Body.Close()

		if err := json.Unmarshal(body, &req); err != nil {
			render.BadRequest(w, err)
			return
		}

		user := model.User{
			UserName:  req.User.Username,
			Email:     req.User.Email,
			Password:  req.User.Password,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		result, err := us.Create(ctx, &user)
		if err != nil {
			render.BadRequest(w, err)
			return
		}

		if result == nil {
			render.BadRequest(w, err)
			return
		}

		token, err := auth.GenerateToken(result.UserId, result.UserName)
		if err != nil {
			render.BadRequest(w, err)
			return
		}

		var res CreateUserResponse
		res.User.Username = user.UserName
		res.User.Email = user.Email
		res.User.Image = user.Image
		res.User.Bio = user.Bio
		res.User.Token = token

		render.JSON(w, res, http.StatusCreated)
	}
}

func HandleCurrentUser(ctx context.Context, us model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		userId, err := auth.GetUserId(r.Context())
		if err != nil {
			render.Unauthorized(w, err)
			return
		}

		user, err := us.GetById(userId)
		if err != nil {
			render.NotFound(w, err)
			return
		}

		render.JSON(w, user, http.StatusOK)
	}
}

func HandleUpdate(ctx context.Context, us model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		userId, err := auth.GetUserId(r.Context())
		if err != nil {
			render.Unauthorized(w, err)
			return
		}

		var req UpdateUserRequest
		body, err := io.ReadAll(r.Body)
		if err != nil {
			render.BadRequest(w, err)
			return
		}
		defer r.Body.Close()

		if err := json.Unmarshal(body, &req); err != nil {
			render.BadRequest(w, err)
			return
		}

		var user model.User
		user.UserId = userId

		if len(req.User.Email) != 0 {
			user.Email = req.User.Email
		}

		if len(req.User.Username) != 0 {
			user.Password = req.User.Username
		}

		if len(req.User.Bio) != 0 {
			user.Bio = req.User.Bio
		}

		if len(req.User.Image) != 0 {
			user.Image = req.User.Image
		}

		if err := us.Update(ctx, &user); err != nil {
			render.BadRequest(w, err)
			return
		}

		result, err := us.GetById(userId)
		if err != nil {
			render.NotFound(w, err)
			return
		}

		//TODO: token set edilecek.

		render.JSON(w, result, http.StatusOK)
	}
}

func HandleProfile(us model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func HandleFind(us model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req LoginUserRequest

		body, err := io.ReadAll(r.Body)
		if err != nil {
			render.BadRequest(w, err)
			return
		}

		defer r.Body.Close()

		if err := json.Unmarshal(body, &req); err != nil {
			render.BadRequest(w, err)
			return
		}

		if err := req.ValidateLoginUserRequest(); err != nil {
			render.BadRequest(w, err)
			return
		}

		user, err := us.GetByEmail(req.User.Email)
		if err != nil {
			render.NotFound(w, err)
			logrus.Info("api: cannot find user")
			return
		}

		token, err := auth.GenerateToken(user.UserId, user.UserName)
		if err != nil {
			render.BadRequest(w, err)
			return
		}

		var res BaseUserResponse
		res.User.Email = user.Email
		res.User.Username = user.UserName
		res.User.Image = user.Image
		res.User.Bio = user.Bio
		res.User.Token = token

		render.JSON(w, res, http.StatusOK)
	}
}

func (u LoginUserRequest) ValidateLoginUserRequest() error {
	if len(u.User.Email) == 0 {
		return errors.New("Email required")
	}

	if len(u.User.Password) == 0 {
		return errors.New("Password required")
	}

	return nil
}
