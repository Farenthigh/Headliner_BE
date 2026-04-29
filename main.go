// package main

// import (
// 	"fmt"
// 	"headliner-be/config"
// 	Entities "headliner-be/entities"
// 	"headliner-be/utils"
// 	"log"

// 	"headliner-be/routers"

// 	"github.com/gofiber/fiber/v3"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"

// 	_ "headliner-be/docs"

// 	swagger "github.com/Flussen/swagger-fiber-v3"
// )

// // @title           Headliner API
// // @version         1.0
// // @description     API Documentation สำหรับเกม Headliner
// // @host            localhost:8080
// // @BasePath        /

// func main() {

// 	fmt.Println("------- Starting Services -------")

// 	dsn := config.DbURL
// 	if dsn == "" {
// 	    dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
// 	    config.DbUser, config.DbPassword, config.DbHost, config.DbPort, config.DbSchema)
// 	}

// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// if err != nil {
//     // ใส่ log เพื่อดูว่า dsn ที่สร้างออกมาหน้าตาเป็นยังไง (ช่วย debug ได้ดีมาก)
//     log.Printf("Current DSN: %s", dsn)
//     panic(fmt.Sprintf("failed to connect database: %v", err))
// }
// 	db.AutoMigrate(&Entities.Users{}, &Entities.StageLog{}, &Entities.Leaderboard{})//สร้าง table อัตโนมัติ
// 	app := fiber.New()
// 	utils.InitFirebase()
// 	app.Get("/swagger/*", swagger.HandlerDefault)

// 	Entities.Init(db, app)

// 	routers.InitUsersRoute(app, db)

// 	routers.InitChatRoute(app)

// 	routers.InitAchievementRoute(app, db)

// 	routers.InitStageRoute(app, db)
// 	routers.InitSavingStageRoute(app, db)
// 	routers.InitLeaderboardRoute(app, db)

// 	port := config.Port
//     if port == "" {
//         port = "8080"
//     }

//     // สั่งให้แอปฟังที่พอร์ตนั้น
//     log.Fatal(app.Listen(":" + port))

// }

package main

import (
	"headliner-be/config"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	// พยายามโหลด .env ถ้าไม่มี (เช่นบน Cloud) จะข้ามไปอ่านค่าจากระบบแทน
	_ = godotenv.Load()

	// ดึงค่า PORT จาก Env ถ้าไม่ได้ตั้งไว้ให้ใช้ 8080
	port := config.Port
	if port == "" {
		port = "8080"
	}

	app := fiber.New()

	// ตัวอย่าง Route พื้นฐาน
	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Backend is running!",
			"db_host": config.DbHost, // ลองดึงค่า Env มาโชว์
		})
	})

	log.Printf("Server is starting on port %s", port)
	
	// สั่งรันแอป
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}