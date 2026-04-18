package savingstage

import (
	"headliner-be/usecase/savingstage"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type SavingStageHandler struct {
	usecase *savingstage.SavingStageUsecase
}

func NewSavingStageHandler(u *savingstage.SavingStageUsecase) *SavingStageHandler {
	return &SavingStageHandler{usecase: u}
}

type SaveStageRequest struct {
	UserID uint `json:"user_id"`
	Stage  int  `json:"stage"`
	Stars  int  `json:"stars"`
}

func (h *SavingStageHandler) SaveStage(c fiber.Ctx) error {

	var req SaveStageRequest

	if err := c.Bind().Body(&req); err != nil {
		return err
	}

	err := h.usecase.SaveStage(req.UserID, req.Stage, req.Stars)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "saved",
	})
}

func (h *SavingStageHandler) GetLeaderboard(c fiber.Ctx) error {

	limitStr := c.Query("limit", "50")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 50
	}

	board, err := h.usecase.GetLeaderboard(limit)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON(board)
}

func (h *SavingStageHandler) GetMyRank(c fiber.Ctx) error {

	userIDStr := c.Query("user_id")

	userID64, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		return c.Status(400).JSON("invalid user_id")
	}

	board, err := h.usecase.GetPlayerRank(uint(userID64))
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON(board)
}

func (h *SavingStageHandler) GetUserStages(c fiber.Ctx) error {

	userIDStr := c.Query("user_id")

	userID64, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		return c.Status(400).JSON("invalid user_id")
	}

	stages, err := h.usecase.GetUserStages(uint(userID64))
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON(stages)
}

func (h *SavingStageHandler) GetUnlockStage(c fiber.Ctx) error {

    userIDStr := c.Query("user_id")

    userID64, err := strconv.ParseUint(userIDStr, 10, 64)
    if err != nil {
        return c.Status(400).JSON("invalid user_id")
    }

    nowStage, unlockedStages, err := h.usecase.GetUnlockStage(uint(userID64))
    if err != nil {
        return c.Status(500).JSON(err.Error())
    }

    nextStage := nowStage + 1

    return c.JSON(fiber.Map{
        "now_stage": nowStage + 1,
        "next_stage": nextStage + 1,
        "unlocked": unlockedStages,
    })
}

func (h *SavingStageHandler) GetStageStars(c fiber.Ctx) error {

	userIDStr := c.Query("user_id")

	userID64, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		return c.Status(400).JSON("invalid user_id")
	}

	stages, err := h.usecase.GetStageStars(uint(userID64))
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON(stages)
}