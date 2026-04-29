package main

import (
	"fmt"
	"headliner-be/config"
	Entities "headliner-be/entities"
	"headliner-be/utils"
	"log"
	"os"

	"github.com/joho/godotenv"

	"headliner-be/routers"

	"github.com/gofiber/fiber/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "headliner-be/docs"

	swagger "github.com/Flussen/swagger-fiber-v3"
)

// @title           Headliner API
// @version         1.0
// @description     API Documentation สำหรับเกม Headliner
// @host            localhost:8080
// @BasePath        /

func main() {
	_ = godotenv.Load() // โหลด .env ไฟล์ก่อน เพื่อให้ config สามารถอ่านค่าต่างๆ ได้
	fmt.Println("------- Starting Services -------")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbSchema := os.Getenv("DB_SCHEMA")

	port := os.Getenv("PORT")

	dsn := config.DbURL
	if dsn == "" {
		dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
			dbUser, dbPassword, dbHost, dbPort, dbSchema)
	}

db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
if err != nil {
    // ใส่ log เพื่อดูว่า dsn ที่สร้างออกมาหน้าตาเป็นยังไง (ช่วย debug ได้ดีมาก)
    log.Printf("Current DSN: %s", dsn)
    panic(fmt.Sprintf("failed to connect database: %v", err))
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

    // สั่งให้แอปฟังที่พอร์ตนั้น
    log.Fatal(app.Listen(":" + port))

}