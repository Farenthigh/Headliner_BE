package Entities

import (
	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	Subject     string `gorm:"not null" json:"subject"`
	Description string `gorm:"not null" json:"description"`
}