package AchievementAdapter

import (
	"fmt"
	Entities "headliner-be/entites"
	AchievementUsecase "headliner-be/usecase/achievement"
	"headliner-be/utils"

	"github.com/gofiber/fiber/v3"
)


type AchievementHandler struct {
	achievementUsecase AchievementUsecase.AchievementUsecase
}

func NewAchievementHandler(achievementUsecase AchievementUsecase.AchievementUsecase) *AchievementHandler {
	return &AchievementHandler{
		achievementUsecase: achievementUsecase,
	}
}

func (a *AchievementHandler) CreateAchievement(c fiber.Ctx) error {
	var achievement Entities.Achievement
	if err := c.Bind().Body(&achievement); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid request body", err.Error(), nil)
	}
	err := a.achievementUsecase.CreateAchievement(&achievement)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to create achievement", err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Achievement created successfully", "", nil)
}

func (a *AchievementHandler) GetAllAchievements(c fiber.Ctx) error {
	achievements, err := a.achievementUsecase.GetAllAchievements()
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to retrieve achievements", err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Achievements retrieved successfully", "", achievements)
}

func (a *AchievementHandler) GetAchievementByID(c fiber.Ctx) error {
	idParam := c.Params("id")
	if idParam == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "ID parameter is required", "", nil)
	}
	var achievementID uint
	_, err := fmt.Sscanf(idParam, "%d", &achievementID)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid ID format", err.Error(), nil)
	}
	achievement, err := a.achievementUsecase.GetAchievementByID(achievementID)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to retrieve achievement", err.Error(), nil)
	}
	if achievement == nil {
		return utils.ResponseJSON(c, fiber.StatusNotFound, "Achievement not found", "", nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Achievement retrieved successfully", "", achievement)
}