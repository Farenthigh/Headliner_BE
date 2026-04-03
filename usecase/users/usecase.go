package UsersUsecase

import (
	"errors"

	Entities "headliner-be/entities"
	UsersModels "headliner-be/model/users"
	"headliner-be/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UsersUsecase interface {
	Register(*UsersModels.RegisterInput) (string, error)
	Login(*UsersModels.LoginInput) (string, error)
	CreateCharacter(uint, *UsersModels.CreateCharacterInput) (string, error)
	GetUserData(uint) (*Entities.Users, error)
	UpdateUsername(uint, *UsersModels.UpdateUsernameInput) (string, error)
	UpdatePassword(uint, *UsersModels.UpdatePasswordInput) (string, error)

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
	hashPassword, err := HashPassword(users.Password)
	if err != nil {
		return "Failed to hash password", err
	}
	var EntitiesUsers = &Entities.Users{
		Email:    users.Email,
		Password: hashPassword,
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
	var EntitiesUsers = &Entities.Users{
		ID:        userID,
		Username:  users.Username,
		Character: users.Character,
	}
	user, err := service.usersRepo.GetUserByUsername(users.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return "Internal server error", err
	}
	if user != nil {
		return "Username already exists", errors.New("Username already exists")
	}
	if err := service.usersRepo.CreateCharacter(EntitiesUsers); err != nil {
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
func (service *UsersService) UpdateUsername(userID uint, input *UsersModels.UpdateUsernameInput) (string, error) {

	if input.Username == "" {
		return "Username cannot be empty", errors.New("Username cannot be empty")
	}

	user, err := service.usersRepo.GetUserByUsername(input.Username)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return "Internal server error", err
	}

	if user != nil {
		return "Username already exists", errors.New("Username already exists")
	}

	err = service.usersRepo.UpdateUsername(userID, input.Username)

	if err != nil {
		return "Failed to update username", err
	}

	return "Username updated successfully", nil
}

func (service *UsersService) UpdatePassword(userID uint, input *UsersModels.UpdatePasswordInput) (string, error) {

	user, err := service.usersRepo.GetUserByID(userID)

	if err != nil {
		return "User not found", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.CurrentPassword)); err != nil {
		return "Incorrect current password", err
	}

	hashedPassword, err := HashPassword(input.NewPassword)

	if err != nil {
		return "Failed to hash password", err
	}

	err = service.usersRepo.UpdatePassword(userID, hashedPassword)

	if err != nil {
		return "Failed to update password", err
	}

	return "Password updated successfully", nil
}

func (u *userUseCase) UpdateChatbotName(userID uint, newName string) error {
    result := u.db.Model(&entities.User{}).Where("id = ?", userID).Update("chatbot_name", newName)
    if result.Error != nil {
        return result.Error
    }
    return nil
}