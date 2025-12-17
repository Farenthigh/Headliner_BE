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
	message, err := a.UsersUsecase.Register(&users)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, message, err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, message, "", nil)
}
