package UsersUsecase

import (
	"errors"

	Entities "headliner-be/entities"
	UsersModels "headliner-be/model/users"
	"headliner-be/utils"

	"golang.org/x/crypto/bcrypt"
)

type UsersUsecase interface {
	Register(*UsersModels.RegisterInput) (string, error)
	Login(*UsersModels.LoginInput) (string, error)
	CreateCharacter(uint, *UsersModels.CreateCharacterInput) (string, error)
	GetUserData(uint) (*Entities.Users, error)

	SetChatbotName(uint, *UsersModels.SetChatbotNameInput) (string, error)
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
		Email:    users.Email,
		Password: HashPassword,
	}
	if err := service.usersRepo.Register(EntitiesUsers); err != nil {
		return "Failed to register user", err
	}
	return "User registered successfully", nil
}

func (service *UsersService) Login(users *UsersModels.LoginInput) (string, error) {
	existingUser, err := service.usersRepo.GetUserByEmail(users.Email)
	if err != nil {
		return "Internal server error", err
	}
	if existingUser == nil {
		return "User not found", errors.New("User not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(users.Password)); err != nil {
		return "Invalid password", err
	}
	token, err := utils.CreateToken(existingUser.ID, existingUser.Email, existingUser.Username)
	if err != nil {
		return "Failed to create token", err
	}
	return token, nil
}

func (service *UsersService) CreateCharacter(userID uint, users *UsersModels.CreateCharacterInput) (string, error) {
	var entitiesUsers = &Entities.Users{
		ID:        userID,
		Username:  users.Username,
		Character: users.Character,
	}
	if err := service.usersRepo.CreateCharacter(entitiesUsers); err != nil {
		return "Failed to create character", err
	}
	return "Character created successfully", nil
}

func (service *UsersService) GetUserData(userID uint) (*Entities.Users, error) {
	user, err := service.usersRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *UsersService) SetChatbotName(userID uint, input *UsersModels.SetChatbotNameInput) (string, error) {
	if input.ChatbotName == "" {
		return "Chatbot name is required", errors.New("chatbot name is required")
	}
	if err := service.usersRepo.SetChatbotName(userID, input.ChatbotName); err != nil {
		return "Failed to set chatbot name", err
	}
	return "Chatbot name set successfully", nil
}