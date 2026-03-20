package Entities

type SavingStageLog struct {
	ID uint `gorm:"primaryKey"`

	UserID uint `gorm:"uniqueIndex:idx_user_stage" json:"user_id"`
	Stage  int  `gorm:"uniqueIndex:idx_user_stage" json:"stage"`
	Stars int `json:"stars"`
}
