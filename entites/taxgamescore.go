package Entities

import "gorm.io/gorm"

type TaxGameScore struct {
	gorm.Model
	UserID uint `gorm:"not null" json:"user_id"`
	Score  int  `gorm:"not null" json:"score"`
}