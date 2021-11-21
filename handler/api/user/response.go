package user

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
