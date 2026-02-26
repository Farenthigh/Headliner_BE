package AchivementAdapter

import (
	"fmt"
	Entities "headliner-be/entites"
	AchivementUsecase "headliner-be/usecase/achivement"
	"headliner-be/utils"

	"github.com/gofiber/fiber/v3"
)


type AchivementHandler struct {
	achivementUsecase AchivementUsecase.AchivementUsecase
}

func NewAchivementAdapter(achivementUsecase AchivementUsecase.AchivementUsecase) *AchivementHandler {
	return &AchivementHandler{
		achivementUsecase: achivementUsecase,
	}
}

func (a *AchivementHandler) CreateAchivement(c fiber.Ctx) error {
	var achivement Entities.Achivement
	if err := c.Bind().Body(&achivement); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid request body", err.Error(), nil)
	}
	err := a.achivementUsecase.CreateAchivement(&achivement)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to create achivement", err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Achivement created successfully", "", nil)
}

func (a *AchivementHandler) GetAllAchivements(c fiber.Ctx) error {
	achivements, err := a.achivementUsecase.GetAllAchivements()
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to retrieve achivements", err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Achivements retrieved successfully", "", achivements)
}

func (a *AchivementHandler) GetAchivementByID(c fiber.Ctx) error {
	idParam := c.Params("id")
	if idParam == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "ID parameter is required", "", nil)
	}
	var achivementID uint
	_, err := fmt.Sscanf(idParam, "%d", &achivementID)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid ID format", err.Error(), nil)
	}
	achivement, err := a.achivementUsecase.GetAchivementByID(achivementID)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to retrieve achivement", err.Error(), nil)
	}
	if achivement == nil {
		return utils.ResponseJSON(c, fiber.StatusNotFound, "Achivement not found", "", nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Achivement retrieved successfully", "", achivement)
}