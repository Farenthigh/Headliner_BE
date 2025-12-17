package main

import (
	"fmt"
	"headliner-be/config"
	Entities "headliner-be/entites"
	"headliner-be/routers"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DbHost, config.DbPort, config.DbUser, config.DbPassword, config.DbSchema)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	app := fiber.New()

	Entities.Init(db, app)

	routers.InitUsersRoute(app, db)
	app.Listen(fmt.Sprintf(":%s", config.Port))

}
