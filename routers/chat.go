package routers

import (
    "github.com/gofiber/fiber/v3"
    "headliner-be/adapter/chat"
    "headliner-be/usecase/chat"
)

func InitChatRoute(app *fiber.App) {
    chatUsecase := usecase_chat.NewChatUsecase()

    chatHandler := adapter_chat.NewChatHandler(chatUsecase)

    api := app.Group("/api")
    api.Post("/chat", chatHandler.Chat)
}