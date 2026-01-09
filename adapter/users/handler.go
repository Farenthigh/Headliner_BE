package UsersAdapter

import (
	UsersModels "headliner-be/model/users"
	UsersUsecase "headliner-be/usecase/users"
	"headliner-be/utils"

	"github.com/gofiber/fiber/v3"
)

type UsersHandler struct {
	UsersUsecase UsersUsecase.UsersUsecase
}

func NewUsersAdapter(usersUsecase UsersUsecase.UsersUsecase) *UsersHandler {
	return &UsersHandler{
		UsersUsecase: usersUsecase,
	}
}

func (a *UsersHandler) Register(c fiber.Ctx) error {
	var users UsersModels.RegisterInput
	if err := c.Bind().Body(&users); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid request body", err.Error(), nil)
	}
	if !utils.IsValidEmail(users.Email) {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid email format", "", nil)
	}
	message, err := a.UsersUsecase.Register(&users)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, message, err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, message, "", nil)
}

func (a *UsersHandler) Login(c fiber.Ctx) error {
	var users UsersModels.LoginInput
	if err := c.Bind().Body(&users); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid request body", err.Error(), nil)
	}
	data, err := a.UsersUsecase.Login(&users)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, data, err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Login successful", "", fiber.Map{
		"token": data,
	})
}

func (a *UsersHandler) CreateCharacter(c fiber.Ctx) error {
	var users UsersModels.CreateCharacterInput
	if err := c.Bind().Body(&users); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid request body", err.Error(), nil)
	}
	userID := c.Locals("userID")

	message, err := a.UsersUsecase.CreateCharacter(uint(userID.(float64)), &users)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, message, err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, message, "", nil)
	// return nil
}

func (a *UsersHandler) GetUserData(c fiber.Ctx) error {
	userID := c.Locals("userID")

	user, err := a.UsersUsecase.GetUserData(uint(userID.(float64)))
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to get user data", err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "User data retrieved successfully", "", fiber.Map{
		"id":        user.ID,
		"username":  user.Username,
		"email":     user.Email,
		"character": user.Character,
	})
}
