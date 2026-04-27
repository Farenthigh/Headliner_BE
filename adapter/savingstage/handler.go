package savingstage

import (
	"headliner-be/usecase/savingstage"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type SavingStageHandler struct {
	usecase *savingstage.SavingStageUsecase
}

func NewSavingStageHandler(u *savingstage.SavingStageUsecase) *SavingStageHandler {
	return &SavingStageHandler{usecase: u}
}

type SaveStageRequest struct {
	UserID uint `json:"user_id"`
	Stage  int  `json:"stage"`
	Stars  int  `json:"stars"`
}

// SaveStage godoc
// @Summary      บันทึกผลการเล่นด่านออมเงิน
// @Description  บันทึกคะแนน จำนวนดาว และด่านที่เล่นจบของผู้เล่น
// @Tags         SavingStage
// @Accept       json
// @Produce      json
// @Param        request body SaveStageRequest true "ข้อมูลการบันทึกด่าน"
// @Success      200  {object}  map[string]interface{} "บันทึกสำเร็จ"
// @Router       /saving-stage/save [post]
func (h *SavingStageHandler) SaveStage(c fiber.Ctx) error {

	var req SaveStageRequest

	if err := c.Bind().Body(&req); err != nil {
		return err
	}

	err := h.usecase.SaveStage(req.UserID, req.Stage, req.Stars)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "saved",
	})
}

// GetLeaderboard godoc
// @Summary      ดึงอันดับ Leaderboard ของ SavingGame
// @Description  ดึงรายการผู้เล่นที่มีคะแนนสูงสุดตามจำนวน limit ที่กำหนด
// @Tags         SavingStage
// @Produce      json
// @Param        limit  query     int  false  "จำนวนอันดับที่ต้องการ (เริ่มต้น 50)"
// @Success      200  {array}   map[string]interface{} "ดึงข้อมูลสำเร็จ"
// @Router       /saving-stage/leaderboard [get]
func (h *SavingStageHandler) GetLeaderboard(c fiber.Ctx) error {

	limitStr := c.Query("limit", "50")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 50
	}

	board, err := h.usecase.GetLeaderboard(limit)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON(board)
}

// GetMyRank godoc
// @Summary      ดึงอันดับปัจจุบันของผู้เล่นตาม User ID
// @Description  ดึงข้อมูลอันดับปัจจุบันของผู้เล่นตาม User ID
// @Tags         SavingStage
// @Produce      json
// @Param        user_id  query     int  true  "User ID ของผู้เล่น"
// @Success      200  {object}  map[string]interface{} "ดึงข้อมูลสำเร็จ"
// @Router       /saving-stage/leaderboard/me [get]
func (h *SavingStageHandler) GetMyRank(c fiber.Ctx) error {

	userIDStr := c.Query("user_id")

	userID64, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		return c.Status(400).JSON("invalid user_id")
	}

	board, err := h.usecase.GetPlayerRank(uint(userID64))
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON(board)
}

// GetUserStages godoc
// @Summary      ดึงความคืบหน้าการเล่นด่านทั้งหมด
// @Description  รายการด่านทั้งหมดที่ผู้เล่นเคยเล่นผ่านไปแล้ว
// @Tags         SavingStage
// @Produce      json
// @Param        user_id  query     int  true  "User ID ของผู้เล่น"
// @Success      200  {array}   map[string]interface{} "ดึงข้อมูลสำเร็จ"
// @Router       /saving-stage/progress [get]
func (h *SavingStageHandler) GetUserStages(c fiber.Ctx) error {

	userIDStr := c.Query("user_id")

	userID64, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		return c.Status(400).JSON("invalid user_id")
	}

	stages, err := h.usecase.GetUserStages(uint(userID64))
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON(stages)
}

// GetUnlockStage godoc
// @Summary      ดึงข้อมูลด่านที่ปลดล็อก
// @Description  เช็คว่าผู้เล่นเล่นถึงด่านไหนและด่านถัดไปคือด่านอะไร
// @Tags         SavingStage
// @Produce      json
// @Param        user_id  query     int  true  "User ID ของผู้เล่น"
// @Success      200  {object}  map[string]interface{} "ดึงข้อมูลสำเร็จ"
// @Router       /saving-stage/unlock [get]
func (h *SavingStageHandler) GetUnlockStage(c fiber.Ctx) error {

    userIDStr := c.Query("user_id")

    userID64, err := strconv.ParseUint(userIDStr, 10, 64)
    if err != nil {
        return c.Status(400).JSON("invalid user_id")
    }

    nowStage, unlockedStages, err := h.usecase.GetUnlockStage(uint(userID64))
    if err != nil {
        return c.Status(500).JSON(err.Error())
    }

    nextStage := nowStage + 1

    return c.JSON(fiber.Map{
        "now_stage": nowStage + 1,
        "next_stage": nextStage + 1,
        "unlocked": unlockedStages,
    })
}

// GetStageStars godoc
// @Summary      ดึงจำนวนดาวที่ได้ในแต่ละด่าน
// @Description  ดูว่าในแต่ละด่านที่ผ่านมา ผู้เล่นได้ไปกี่ดาว (0-3 ดาว)
// @Tags         SavingStage
// @Produce      json
// @Param        user_id  query     int  true  "User ID ของผู้เล่น"
// @Success      200  {array}   map[string]interface{} "ดึงข้อมูลสำเร็จ"
// @Router       /saving-stage/stars [get]
func (h *SavingStageHandler) GetStageStars(c fiber.Ctx) error {

	userIDStr := c.Query("user_id")

	userID64, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		return c.Status(400).JSON("invalid user_id")
	}

	stages, err := h.usecase.GetStageStars(uint(userID64))
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.JSON(stages)
}