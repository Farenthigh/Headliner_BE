package Entities

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey" json:"id"`
	Username  string `gorm:"not null" json:"username"`
	Email     string `gorm:"uniqueIndex;not null" json:"email"`
	Password  string `gorm:"not null" json:"password"`
	Character int    `gorm:"not null" json:"character"`
	Chatbotname string `gorm:"not null" json:"chatbotname"`
}
