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

type UnlockAchievementRequest struct {
	UserID        uint `json:"user_id"`
	AchievementID uint `json:"achievement_id"`
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

func (a *AchievementHandler) UnlockAchievement(c fiber.Ctx) error {
	var req UnlockAchievementRequest
	if err := c.Bind().Body(&req); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid request body", err.Error(), nil)
	}

	userAch := &Entities.UserAchievement{
		UserID:        req.UserID,
		AchievementID: req.AchievementID,
	}

	if err := a.achievementUsecase.SaveUserAchievement(userAch); err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to unlock achievement", err.Error(), nil)
	}

	return utils.ResponseJSON(c, fiber.StatusOK, "Achievement unlocked successfully", "", userAch)
}

func (a *AchievementHandler) GetUserAchievements(c fiber.Ctx) error {
	userIDParam := c.Params("userId")
	if userIDParam == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "User ID parameter is required", "", nil)
	}

	var userID uint
	if _, err := fmt.Sscanf(userIDParam, "%d", &userID); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid User ID format", err.Error(), nil)
	}

	userAchievements, err := a.achievementUsecase.GetUserAchievements(userID)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to retrieve user achievements", err.Error(), nil)
	}

	return utils.ResponseJSON(c, fiber.StatusOK, "User achievements retrieved successfully", "", userAchievements)
}