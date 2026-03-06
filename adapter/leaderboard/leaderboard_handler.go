package leaderboard

import (
    leaderboard_usecase "headliner-be/usecase/leaderboard"  // ← import usecase
    "github.com/gofiber/fiber/v3"
    // ลบ import Entities ออก ไม่ได้ใช้
)

type LeaderboardAdapter struct {
    usecase leaderboard_usecase.LeaderboardUsecase  // ← ใช้ type จาก usecase package
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