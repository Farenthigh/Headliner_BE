package AchievementAdapter

import (
	"fmt"
	Entities "headliner-be/entities"
	AchievementUsecase "headliner-be/usecase/achievement"
	"headliner-be/utils"

	"github.com/gofiber/fiber/v3"
)


type AchievementHandler struct {
	achievementUsecase AchievementUsecase.AchievementUsecase
}

type UnlockAchievementRequest struct {
	UserID        uint `json:"user_id"`
	AchievementID uint `json:"achievement_id"`
}

func NewAchievementHandler(achievementUsecase AchievementUsecase.AchievementUsecase) *AchievementHandler {
	return &AchievementHandler{
		achievementUsecase: achievementUsecase,
	}
}

// CreateAchievement godoc
// @Summary      สร้าง Achievement ใหม่
// @Description  เพิ่มข้อมูลความสำเร็จ (Achievement) ใหม่เข้าสู่ระบบ
// @Tags         Achievement
// @Accept       json
// @Produce      json
// @Param        request body Entities.Achievement true "ข้อมูล Achievement ที่ต้องการสร้าง"
// @Success      200  {object}  map[string]interface{} "สร้างสำเร็จ"
// @Failure      400  {object}  map[string]interface{} "ข้อมูลไม่ถูกต้อง"
// @Failure      500  {object}  map[string]interface{} "เซิร์ฟเวอร์มีปัญหา"
// @Router       /achievements/ [post]
func (a *AchievementHandler) CreateAchievement(c fiber.Ctx) error {
	var achievement Entities.Achievement
	if err := c.Bind().Body(&achievement); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid request body", err.Error(), nil)
	}
	err := a.achievementUsecase.CreateAchievement(&achievement)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to create achievement", err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Achievement created successfully", "", nil)
}

// GetAllAchievements godoc
// @Summary      ดึงข้อมูล Achievement ทั้งหมด
// @Description  ดึงรายการเงื่อนไขความสำเร็จ (Achievement) ทั้งหมดที่มีในระบบ
// @Tags         Achievement
// @Produce      json
// @Success      200  {object}  map[string]interface{} "ดึงข้อมูลสำเร็จ"
// @Failure      500  {object}  map[string]interface{} "เซิร์ฟเวอร์มีปัญหา"
// @Router       /achievements/ [get]
func (a *AchievementHandler) GetAllAchievements(c fiber.Ctx) error {
	achievements, err := a.achievementUsecase.GetAllAchievements()
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to retrieve achievements", err.Error(), nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Achievements retrieved successfully", "", achievements)
}

// GetAchievementByID godoc
// @Summary      ดึงข้อมูล Achievement ตาม ID
// @Description  ค้นหาและดึงรายละเอียดของ Achievement แบบเจาะจงด้วย ID
// @Tags         Achievement
// @Produce      json
// @Param        id   path      int  true  "รหัส Achievement ID"
// @Success      200  {object}  map[string]interface{} "ดึงข้อมูลสำเร็จ"
// @Failure      400  {object}  map[string]interface{} "รูปแบบ ID ไม่ถูกต้อง"
// @Failure      404  {object}  map[string]interface{} "ไม่พบข้อมูล"
// @Failure      500  {object}  map[string]interface{} "เซิร์ฟเวอร์มีปัญหา"
// @Router       /achievements/{id} [get]
func (a *AchievementHandler) GetAchievementByID(c fiber.Ctx) error {
	idParam := c.Params("id")
	if idParam == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "ID parameter is required", "", nil)
	}
	var achievementID uint
	_, err := fmt.Sscanf(idParam, "%d", &achievementID)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid ID format", err.Error(), nil)
	}
	achievement, err := a.achievementUsecase.GetAchievementByID(achievementID)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to retrieve achievement", err.Error(), nil)
	}
	if achievement == nil {
		return utils.ResponseJSON(c, fiber.StatusNotFound, "Achievement not found", "", nil)
	}
	return utils.ResponseJSON(c, fiber.StatusOK, "Achievement retrieved successfully", "", achievement)
}

// UnlockAchievement godoc
// @Summary      ปลดล็อก Achievement ให้ผู้เล่น
// @Description  บันทึกข้อมูลว่าผู้เล่นได้รับ (Unlock) Achievement นี้เรียบร้อยแล้ว
// @Tags         Achievement
// @Accept       json
// @Produce      json
// @Param        request body UnlockAchievementRequest true "ข้อมูล User ID และ Achievement ID"
// @Success      200  {object}  map[string]interface{} "ปลดล็อกสำเร็จ"
// @Failure      400  {object}  map[string]interface{} "ข้อมูลไม่ถูกต้อง"
// @Failure      500  {object}  map[string]interface{} "เซิร์ฟเวอร์มีปัญหา"
// @Router       /achievements/unlock [post]
func (a *AchievementHandler) UnlockAchievement(c fiber.Ctx) error {
	var req UnlockAchievementRequest
	if err := c.Bind().Body(&req); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid request body", err.Error(), nil)
	}

	userAch := &Entities.UserAchievement{
		UserID:        req.UserID,
		AchievementID: req.AchievementID,
	}

	if err := a.achievementUsecase.SaveUserAchievement(userAch); err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to unlock achievement", err.Error(), nil)
	}

	return utils.ResponseJSON(c, fiber.StatusOK, "Achievement unlocked successfully", "", userAch)
}

// GetUserAchievements godoc
// @Summary      ดึงรายการ Achievement ที่ผู้เล่นปลดล็อกแล้ว
// @Description  ดูรายการว่าผู้เล่นคนนี้ทำเงื่อนไข Achievement อะไรสำเร็จไปแล้วบ้าง
// @Tags         Achievement
// @Produce      json
// @Param        userId   path      int  true  "รหัส User ID ของผู้เล่น"
// @Success      200  {object}  map[string]interface{} "ดึงข้อมูลสำเร็จ"
// @Failure      400  {object}  map[string]interface{} "รูปแบบ User ID ไม่ถูกต้อง"
// @Failure      500  {object}  map[string]interface{} "เซิร์ฟเวอร์มีปัญหา"
// @Router       /achievements/user/{userId} [get]
func (a *AchievementHandler) GetUserAchievements(c fiber.Ctx) error {
	userIDParam := c.Params("userId")
	if userIDParam == "" {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "User ID parameter is required", "", nil)
	}

	var userID uint
	if _, err := fmt.Sscanf(userIDParam, "%d", &userID); err != nil {
		return utils.ResponseJSON(c, fiber.StatusBadRequest, "Invalid User ID format", err.Error(), nil)
	}

	userAchievements, err := a.achievementUsecase.GetUserAchievements(userID)
	if err != nil {
		return utils.ResponseJSON(c, fiber.StatusInternalServerError, "Failed to retrieve user achievements", err.Error(), nil)
	}

	return utils.ResponseJSON(c, fiber.StatusOK, "User achievements retrieved successfully", "", userAchievements)
}