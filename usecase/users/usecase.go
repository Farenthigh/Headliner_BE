package UsersUsecase

import (
	"errors"
	Entities "headliner-be/entites"
	UsersModels "headliner-be/model/users"
	"headliner-be/utils"

	"golang.org/x/crypto/bcrypt"
)

type UsersUsecase interface {
	Register(*UsersModels.RegisterInput) (string, error)
	Login(*UsersModels.LoginInput) (string, error)
}

type UsersService struct {
	usersRepo UsersRepository
}

func NewUsersService(usersrepo UsersRepository) UsersUsecase {
	return &UsersService{
		usersRepo: usersrepo,
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (service *UsersService) Register(users *UsersModels.RegisterInput) (string, error) {
	if users.Password != users.ConfirmPassword {
		return "Password and Confirm Password must be the same", errors.New("Password and Confirm Password must be the same")
	}
	HashPassword, err := HashPassword(users.Password)
	if err != nil {
		return "Failed to hash password", err
	}
	var EntitiesUsers = &Entities.Users{
		Username:  users.Username,
		Email:     users.Email,
		Password:  HashPassword,
		Character: users.Character,
	}
	if err := service.usersRepo.Register(EntitiesUsers); err != nil {
		return "Failed to register user", err
	}
	return "User registered successfully", nil
}

func (service *UsersService) Login(user *UsersModels.LoginInput) (string, error) {
	existingUser, err := service.usersRepo.GetUserByEmail(user.Email)
	if err != nil {
		return "Internal server error", err
	}
	if existingUser == nil {
		return "User not found", errors.New("User not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)); err != nil {
		return "Invalid password", err
	}
	token, err := utils.CreateToken(existingUser.ID, existingUser.Email, existingUser.Username)
	if err != nil {
		return "Failed to create token", err
	}
	return token, nil
}
