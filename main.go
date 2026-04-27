package main

import (
	"fmt"
	"headliner-be/config"
	Entities "headliner-be/entities"
	"headliner-be/utils"

	"headliner-be/routers"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	swagger "github.com/Flussen/swagger-fiber-v3"
	_ "headliner-be/docs"
)

// @title           Headliner API
// @version         1.0
// @description     API Documentation สำหรับเกม Headliner
// @host            localhost:8080
// @BasePath        /

func main() {
	godotenv.Load()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DbHost, config.DbPort, config.DbUser, config.DbPassword, config.DbSchema)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Entities.Users{}, &Entities.StageLog{}, &Entities.Leaderboard{})//สร้าง table อัตโนมัติ
	app := fiber.New()
	utils.InitFirebase()
	app.Get("/swagger/*", swagger.HandlerDefault)

	Entities.Init(db, app)

	routers.InitUsersRoute(app, db)

	routers.InitChatRoute(app)

	routers.InitAchievementRoute(app, db)

	routers.InitStageRoute(app, db)
	routers.InitSavingStageRoute(app, db)
	routers.InitLeaderboardRoute(app, db)

	app.Listen(fmt.Sprintf(":%s", config.Port))

}
