package user

//noinspection
import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
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
	User struct{
		Email string `json:"email" validate:"required"`
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
	User UserResponse
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

type UserResponse struct {
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

func HandleCurrentUser(us model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		props, _ := r.Context().Value("userAuthCtx").(jwt.MapClaims)

		str := fmt.Sprintf("user id:%v - user name:%v", props["user_id"], props["user_name"])
		w.Write([]byte(str))
	}
}

func HandleUpdate(us model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

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

		user := model.User{
			Email:     req.User.Email,
			UserName:  req.User.Username,
			Password:  req.User.Password,
			Bio:       req.User.Bio,
			Image:     req.User.Image,
			UpdatedAt: time.Now(),
		}

		affected, err := us.Update(&user)
		if err != nil {
			render.BadRequest(w, err)
			return
		}

		if affected == 0 {
			render.BadRequest(w, model.ErrOperationFailed)
			return
		}

		var res UserResponse
		res.User.Email = user.Email
		res.User.Username = user.UserName
		res.User.Image = user.Image
		res.User.Bio = user.Bio

		render.JSON(w, res, http.StatusOK)
	}
}

func HandleProfile(us model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func HandleFind(us model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req LoginUserRequest

		body ,err := io.ReadAll(r.Body)
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
			render.BadRequest(w,err)
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

		var res UserResponse
		res.User.Email = user.Email
		res.User.Username = user.UserName
		res.User.Image = user.Image
		res.User.Bio = user.Bio
		res.User.Token = token

		render.JSON(w, res, http.StatusOK)
	}
}

/*func HandlerFind(us model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		val := chi.URLParam(r, "id")

		v, _ := strconv.ParseInt(val, 10, 64)
		user, err := us.GetById(v)
		if err != nil {
			render.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}

		render.SingleSuccessJSON(w, user)
	}
}

func HandlerList(us model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		users, err := us.GetAll()
		if err != nil {
			render.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}

		render.MultipleSuccessJSON(w, users)
	}
}

func HandlerListRange(us model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		users, err := us.GetAllRange(
			model.UserParams{
				Page: 10,
				Size: 5,
			})
		if err != nil {
			render.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}

		render.MultipleSuccessJSON(w, users)
	}
}*/

func (u LoginUserRequest) ValidateLoginUserRequest() error {
	if len(u.User.Email) == 0 {
		return errors.New("Email required")
	}

	if len(u.User.Password) == 0 {
		return errors.New("Password required")
	}

	return nil
}