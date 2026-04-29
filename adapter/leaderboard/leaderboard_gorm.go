package leaderboard

import (
	Entities "headliner-be/entities"
	leaderboard_usecase "headliner-be/usecase/leaderboard" 

	"gorm.io/gorm"
)

type LeaderboardGorm struct {
	db *gorm.DB
}

func NewLeaderboardGorm(db *gorm.DB) leaderboard_usecase.LeaderboardRepository {
	return &LeaderboardGorm{db: db}
}

func (r *LeaderboardGorm) GetAllLeaderboard() ([]Entities.LeaderboardEntry, error) {
	var results []Entities.LeaderboardEntry
	err := r.db.Table("leaderboard").
		Select("users.username, leaderboard.saving_game_score, leaderboard.tax_game_score, leaderboard.updated_at").
		Joins("join users on users.id = leaderboard.user_id").
		Scan(&results).Error
	return results, err
}


func (r *LeaderboardGorm) SaveScore(data Entities.Leaderboard) error {
	var existingRecord Entities.Leaderboard
	result := r.db.Where("user_id = ?", data.UserID).First(&existingRecord)

	// 1. สั่ง SUM รวมดาวที่ดีที่สุดของ "เกมออมเงิน" จากตาราง saving_stage_logs ที่คุณมีอยู่แล้ว
	var totalSavingStars int
	r.db.Table("saving_stage_logs").Where("user_id = ?", data.UserID).Select("COALESCE(SUM(stars), 0)").Scan(&totalSavingStars)

	// 2. สั่ง SUM รวมดาวที่ดีที่สุดของ "เกมภาษี" จากตาราง stage_logs ที่คุณมีอยู่แล้ว
	var totalTaxStars int
	r.db.Table("stage_logs").Where("user_id = ?", data.UserID).Select("COALESCE(SUM(stars), 0)").Scan(&totalTaxStars)

	if result.Error == nil {
		// ถ้ามีประวัติในลีดเดอร์บอร์ดอยู่แล้ว ให้อัปเดตคะแนนเป็น "ผลรวมของดาวล่าสุด" ทันที (แทนการใช้ +=)
		existingRecord.SavingGameScore = totalSavingStars
		existingRecord.TaxGameScore = totalTaxStars

		// อัปเดตเวลาที่ใช้เล่นด้วย
		if data.SavingGameTime > 0 {
			existingRecord.SavingGameTime = data.SavingGameTime
		}
		if data.TaxGameTime > 0 {
			existingRecord.TaxGameTime = data.TaxGameTime
		}

		return r.db.Save(&existingRecord).Error
	}

	// ถ้าเพิ่งเล่นจบด่านแรกและยังไม่มีชื่อในลีดเดอร์บอร์ด ให้สร้างแถวใหม่เลย
	data.SavingGameScore = totalSavingStars
	data.TaxGameScore = totalTaxStars
	return r.db.Create(&data).Error
}