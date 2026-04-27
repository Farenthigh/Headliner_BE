package stage

import (
	stageUsecase "headliner-be/usecase/stage"

	"strconv"

	"github.com/gofiber/fiber/v3"
)

type StageHandler struct {
	usecase *stageUsecase.StageUsecase
}

func NewStageHandler(u *stageUsecase.StageUsecase) *StageHandler {
	return &StageHandler{usecase: u}
}

type SaveStageRequest struct {
	UserID uint `json:"user_id"`
	Stage  int  `json:"stage"`
	Stars  int  `json:"stars"`
}

// SaveStage godoc
// @Summary      บันทึกผลการเล่นด่านปกติ
// @Description  บันทึกข้อมูลการเล่นจบด่าน คะแนน และดาวของผู้เล่นในโหมดปกติ
// @Tags         Stage
// @Accept       json
// @Produce      json
// @Param        request body SaveStageRequest true "ข้อมูลการบันทึกด่าน"
// @Success      200  {object}  map[string]interface{} "บันทึกสำเร็จ"
// @Router       /stage/save [post]
func (h *StageHandler) SaveStage(c fiber.Ctx) error {

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
// @Summary      ดึงอันดับ Leaderboard ของด่านปกติ
// @Description  ดึงรายการผู้เล่นที่มีคะแนนสูงสุดในโหมดปกติ
// @Tags         Stage
// @Produce      json
// @Param        limit  query     int  false  "จำนวนอันดับที่ต้องการ (เริ่มต้น 50)"
// @Success      200  {array}   map[string]interface{} "ดึงข้อมูลสำเร็จ"
// @Router       /stage/leaderboard [get]
func (h *StageHandler) GetLeaderboard(c fiber.Ctx) error {

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
// @Summary      เช็คอันดับของผู้เล่นในด่านปกติ
// @Description  ดึงข้อมูลอันดับปัจจุบันของผู้เล่นในโหมดปกติ
// @Tags         Stage
// @Produce      json
// @Param        user_id  query     int  true  "User ID ของผู้เล่น"
// @Success      200  {object}  map[string]interface{} "ดึงข้อมูลสำเร็จ"
// @Router       /stage/leaderboard/me [get]
func (h *StageHandler) GetMyRank(c fiber.Ctx) error {

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
// @Summary      ดึงความคืบหน้าการเล่นด่านปกติทั้งหมด
// @Description  รายการด่านปกติทั้งหมดที่ผู้เล่นเคยเล่นผ่านไปแล้ว
// @Tags         Stage
// @Produce      json
// @Param        user_id  query     int  true  "User ID ของผู้เล่น"
// @Success      200  {array}   map[string]interface{} "ดึงข้อมูลสำเร็จ"
// @Router       /stage/progress [get]
func (h *StageHandler) GetUserStages(c fiber.Ctx) error {

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
// @Summary      ดึงข้อมูลด่านปกติที่ปลดล็อกล่าสุด
// @Description  เช็คว่าในโหมดปกติผู้เล่นเล่นถึงด่านไหน
// @Tags         Stage
// @Produce      json
// @Param        user_id  query     int  true  "User ID ของผู้เล่น"
// @Success      200  {object}  map[string]interface{} "ดึงข้อมูลสำเร็จ"
// @Router       /stage/unlock [get]
func (h *StageHandler) GetUnlockStage(c fiber.Ctx) error {

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
        "now_stage": nowStage,
        "next_stage": nextStage,
        "unlocked": unlockedStages,
    })
}

// GetStageStars godoc
// @Summary      ดึงจำนวนดาวที่ได้ในแต่ละด่านปกติ
// @Description  ดูจำนวนดาวที่ทำได้ในแต่ละด่านของโหมดปกติ
// @Tags         Stage
// @Produce      json
// @Param        user_id  query     int  true  "User ID ของผู้เล่น"
// @Success      200  {array}   map[string]interface{} "ดึงข้อมูลสำเร็จ"
// @Router       /stage/stars [get]
func (h *StageHandler) GetStageStars(c fiber.Ctx) error {

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