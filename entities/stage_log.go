package Entities

type StageLog struct {
	ID uint `gorm:"primaryKey"`

	UserID uint `gorm:"uniqueIndex:idx_user_stage"`
	Stage  int  `gorm:"uniqueIndex:idx_user_stage"`

	Stars int
}
