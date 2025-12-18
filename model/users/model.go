package UsersModels

type RegisterInput struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Character       int    `json:"character"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
