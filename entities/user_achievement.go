package Entities

import (
	"time"
)

type UserAchievement struct {
	UserID        uint        `gorm:"primaryKey" json:"user_id"`
	AchievementID uint        `gorm:"primaryKey" json:"achievement_id"`
	Achievement   Achievement `gorm:"foreignKey:AchievementID" json:"achievement"`
	UnlockedAt    time.Time   `gorm:"autoCreateTime" json:"unlocked_at"`
}