package leaderboard

import (
    leaderboard_usecase "headliner-be/usecase/leaderboard"  
    "github.com/gofiber/fiber/v3"
    Entities "headliner-be/entities"
)

type LeaderboardAdapter struct {
    usecase leaderboard_usecase.LeaderboardUsecase  
}

func NewLeaderboardAdapter(usecase leaderboard_usecase.LeaderboardUsecase) *LeaderboardAdapter {
    return &LeaderboardAdapter{usecase: usecase}
}

func (h *LeaderboardAdapter) GetLeaderboard(c fiber.Ctx) error {
    lbData, err := h.usecase.GetLeaderboard()
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(fiber.Map{
        "message": "success",
        "data":    fiber.Map{"leaderBoard": lbData},
    })
}

func (h *LeaderboardAdapter) SaveScore(c fiber.Ctx) error {
	var input Entities.Leaderboard

	if err := c.Bind().Body(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	err := h.usecase.SaveScore(input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "Score saved successfully",
	})
}