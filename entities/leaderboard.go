package Entities

import (
	"time"

	"gorm.io/gorm"
)

// แก้ไข Struct นี้เพื่อใช้สร้างตาราง (AutoMigrate)
type Leaderboard struct {
	gorm.Model
	UserID          uint `json:"user_id" gorm:"uniqueIndex"` // บังคับให้ 1 คนมีแค่ 1 แถว
	SavingGameScore int  `json:"saving_game_score"`
	TaxGameScore    int  `json:"tax_game_score"`
	SavingGameTime  int  `json:"saving_game_time"`
	TaxGameTime     int  `json:"tax_game_time"`
}

// Struct ตัวนี้ของเดิม ถูกต้องแล้วครับ
type LeaderboardEntry struct {
	Username        string    `json:"username"`
	SavingGameScore int       `json:"saving_game_score"`
	TaxGameScore    int       `json:"tax_game_score"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (Leaderboard) TableName() string {
	return "leaderboard"
}
