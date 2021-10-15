package user

//noinspection
import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/serhatmorkoc/go-realworld-example/handler/render"
	"github.com/serhatmorkoc/go-realworld-example/model"
	"github.com/serhatmorkoc/go-realworld-example/service/auth"
	"io"
	"net/http"
	"time"
)

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

type CreateUserResponse struct {
	User struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Image    string `json:"image"`
		Bio      string `json:"bio"`
		Token    string `json:"token"`
	} `json:"user"`
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

type ProfileResponse struct {
	Profile struct {
		Username  string `json:"username"`
		Bio       string `json:"bio"`
		Image     string `json:"image"`
		Following bool   `json:"following"`
	} `json:"profile"`
}

func HandlerCreate(us model.UserStore) http.HandlerFunc {
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

		result, err := us.Create(&user)
		if err != nil {
			render.BadRequest(w, err)
			return
		}

		if result == nil {
			render.BadRequest(w, err)
			return
		}

		token, err := auth.GenerateToken(result.UserId)
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

func HandlerCurrentUser(us model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		props, _ := r.Context().Value("props").(jwt.MapClaims)

		str := fmt.Sprintf("%v", props["user_id"])
		w.Write([]byte(str))
	}
}

func HandlerUpdate(us model.UserStore) http.HandlerFunc {
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

func HandlerProfile(us model.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
