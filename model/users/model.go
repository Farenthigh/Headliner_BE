package UsersModels

type RegisterInput struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type CreateCharacterInput struct {
	Username  string `json:"username"`
	Character int    `json:"character"`
}
type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UpdateUsernameInput struct {
	Username string `json:"username"`
}

type UpdatePasswordInput struct {
	CurrentPassword string `json:"currentPassword"`
	NewPassword     string `json:"newPassword"`
}
