package Entities

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Username  string `gorm:"uniqueIndex;not null" json:"username"`
	Email     string `gorm:"uniqueIndex;not null" json:"email"`
	Password  string `gorm:"not null" json:"password"`
	Character int    `gorm:"not null" json:"character"`
}
