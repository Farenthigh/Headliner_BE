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

// Chat godoc
// @Summary      ส่งข้อความคุยกับแชทบอท
// @Description  ส่งคำถามหรือข้อความไปยังระบบ AI LLM (Mystery Man) เพื่อขอคำปรึกษาทางการเงิน
// @Tags         Chat
// @Accept       json
// @Produce      json
// @Param        request body model_chat.ChatRequest true "ข้อมูล Player ID และข้อความที่จะส่ง"
// @Success      200  {object}  map[string]interface{} "ได้รับคำตอบจาก AI สำเร็จ"
// @Failure      400  {object}  map[string]interface{} "ข้อมูลไม่ถูกต้อง (เช่น ขาด Player ID หรือ ข้อความ)"
// @Failure      500  {object}  map[string]interface{} "เซิร์ฟเวอร์หรือระบบ AI มีปัญหา"
// @Router       /api/chat [post]
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