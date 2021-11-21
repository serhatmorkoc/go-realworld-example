package user

type loginUserRequest struct {
	User struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password"`
	}
}

type createUserRequest struct {
	User struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	} `json:"user"`
}

type updateUserRequest struct {
	User struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
		Bio      string `json:"bio"`
		Image    string `json:"image"`
	} `json:"user"`
}

type profileRequest struct {
	Username string `json:"username"`
}
