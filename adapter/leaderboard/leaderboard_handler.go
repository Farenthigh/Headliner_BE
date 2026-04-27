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

// GetLeaderboard godoc
// @Summary      ดึงข้อมูลก Leaderboard
// @Description  ดึงข้อมูลอันดับและคะแนนรวมของผู้เล่นทั้งหมดในระบบเพื่อแสดงผลบน Leaderboard
// @Tags         Leaderboard
// @Produce      json
// @Success      200  {object}  map[string]interface{} "ดึงข้อมูลสำเร็จ"
// @Failure      500  {object}  map[string]interface{} "เซิร์ฟเวอร์มีปัญหา"
// @Router       /leaderboard/ [get]
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

// SaveScore godoc
// @Summary      บันทึกคะแนนลง Leaderboard
// @Description  บันทึกหรืออัปเดตคะแนนรวมของผู้เล่นเพื่อจัดอันดับบน Leaderboard
// @Tags         Leaderboard
// @Accept       json
// @Produce      json
// @Param        request body Entities.Leaderboard true "ข้อมูลผู้เล่นและคะแนนที่ต้องการบันทึก"
// @Success      200  {object}  map[string]interface{} "บันทึกคะแนนสำเร็จ"
// @Failure      400  {object}  map[string]interface{} "รูปแบบข้อมูลไม่ถูกต้อง"
// @Failure      500  {object}  map[string]interface{} "เซิร์ฟเวอร์มีปัญหา"
// @Router       /leaderboard/ [post]
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