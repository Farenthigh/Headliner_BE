package routers

import (
    LeaderboardAdapter "headliner-be/adapter/leaderboard" // สมมติชื่อโฟลเดอร์
    LeaderboardUsecase "headliner-be/usecase/leaderboard"
    "github.com/gofiber/fiber/v3"
    "gorm.io/gorm"
)

func InitLeaderboardRoute(app *fiber.App, db *gorm.DB) {
    lbRepo := LeaderboardAdapter.NewLeaderboardGorm(db)
    lbUsecase := LeaderboardUsecase.NewLeaderboardService(lbRepo)
    lbHandler := LeaderboardAdapter.NewLeaderboardAdapter(lbUsecase)

    lb := app.Group("/leaderboard")
    lb.Get("/", lbHandler.GetLeaderboard) 
}