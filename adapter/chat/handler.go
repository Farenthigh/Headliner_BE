package adapter_chat

import (
    "net/http"
    "github.com/gofiber/fiber/v3"
    "headliner-be/model/chat"
    "headliner-be/usecase/chat"
)

type ChatHandler struct {
    usecase usecase_chat.ChatUsecase
}

func NewChatHandler(u usecase_chat.ChatUsecase) *ChatHandler {
    return &ChatHandler{usecase: u}
}

func (h *ChatHandler) Chat(c fiber.Ctx) error {
    var req model_chat.ChatRequest

    if err := c.Bind().Body(&req); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid JSON format",
        })
    }

	if req.PlayerID == "" || req.Message == "" {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "error": "player_id and message are required",
        })
    }

    res, err := h.usecase.AskAI(req)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    return c.Status(http.StatusOK).JSON(res)
}