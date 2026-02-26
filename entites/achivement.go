package Entities

import (
	"gorm.io/gorm"
)

type Achivement struct {
	gorm.Model
	Name        string `gorm:"not null" json:"name"`
	Description string `gorm:"not null" json:"description"`
}