package user

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"github.com/serhatmorkoc/go-realworld-example/handler/render"
	"github.com/serhatmorkoc/go-realworld-example/model"
	"github.com/serhatmorkoc/go-realworld-example/service/auth"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

func HandleAuthentication(us model.UserStore) http.HandlerFunc {
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

		if err := req.validateLoginUserRequest(); err != nil {
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

func HandleRegistration(us model.UserStore) http.HandlerFunc {
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

		result, err := us.Create(r.Context(), &user)
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

func HandleCurrentUser(us model.UserStore) http.HandlerFunc {
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

func HandleUpdate(us model.UserStore) http.HandlerFunc {
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

		user, err := us.GetById(userId)
		if err != nil {
			render.BadRequest(w,err)
			return
		}

		if len(req.User.Email) != 0 {
			user.Email = req.User.Email
		}

		if len(req.User.Username) != 0 {
			user.UserName = req.User.Username
		}

		if len(req.User.Password) != 0 {
			user.Password = req.User.Password
		}

		if len(req.User.Bio) != 0 {
			user.Bio = req.User.Bio
		}

		if len(req.User.Image) != 0 {
			user.Image = req.User.Image
		}

		if err := us.Update(r.Context(), user); err != nil {
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

		userName := chi.URLParam(r, "username")
		if len(userName) == 0 {
			render.BadRequest(w, errors.New(""))
			return
		}

		user, err := us.GetByUsername(userName)
		if err != nil {
			render.NotFound(w, err)
			return
		}

		var profile ProfileResponse
		profile.Profile.Bio = user.Bio
		profile.Profile.Username = user.UserName
		profile.Profile.Image = user.Image

		render.JSON(w, profile, http.StatusOK)

	}
}

func HandleFollowUser(us model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var userName = chi.URLParam(r,"username")

		followerID, err := auth.GetUserId(r.Context())
		if err != nil {
			render.Unauthorized(w, err)
			return
		}

		followerUser, err := us.GetById(followerID)
		if err != nil {
			render.BadRequest(w,err)
			return
		}

		user, err := us.GetByUsername(userName)
		if err != nil {
			render.NotFound(w,err)
			return
		}

		if followerUser.UserName == user.UserName {
			render.BadRequest(w,errors.New("cannot follow yourself"))
			return
		}

		if err = us.AddFollower(user,followerID); err != nil {
			render.NotFound(w,err)
			return
		}

		var profile ProfileResponse
		profile.Profile.Bio = user.Bio
		profile.Profile.Username = user.UserName
		profile.Profile.Image = user.Image
		profile.Profile.Following = true

		render.JSON(w, profile, http.StatusOK)
	}
}

func (u LoginUserRequest) validateLoginUserRequest() error {
	if len(u.User.Email) == 0 {
		return errors.New("Email required")
	}

	if len(u.User.Password) == 0 {
		return errors.New("Password required")
	}

	return nil
}
