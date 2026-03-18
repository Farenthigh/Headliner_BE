package stage

import (
	"errors"
	Entities "headliner-be/entities"

	"gorm.io/gorm"
)

type StageUsecase struct {
	db *gorm.DB
}

func NewStageUsecase(db *gorm.DB) *StageUsecase {
	return &StageUsecase{db}
}

func (u *StageUsecase) SaveStage(userID uint, stage int, stars int) error {

	var log Entities.StageLog

	err := u.db.
		Where("user_id = ? AND stage = ?", userID, stage).
		First(&log).Error

	// ถ้าเจอ record อยู่แล้ว
	if err == nil {

		if stars > log.Stars {
			return u.db.Model(&Entities.StageLog{}).
				Where("user_id = ? AND stage = ?", userID, stage).
				Update("stars", stars).Error
		}

		return nil
	}

	// ถ้า error ไม่ใช่ record not found
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// ถ้ายังไม่มี record → create
	newLog := Entities.StageLog{
		UserID: userID,
		Stage:  stage,
		Stars:  stars,
	}

	return u.db.Create(&newLog).Error
}
func (u *StageUsecase) GetUnlockStage(userID uint) (uint, []uint, error) {

    var maxStage uint

    err := u.db.
        Table("stage_logs").
        Select("COALESCE(MAX(stage), 0)").
        Where("user_id = ?", userID).
        Scan(&maxStage).Error

    if err != nil {
        return 0, nil, err
    }

    nextStage := maxStage + 1

    var unlocked []uint
    for i := uint(1); i <= nextStage; i++ {
        unlocked = append(unlocked, i)
    }

    return maxStage, unlocked, nil
}

func (u *StageUsecase) GetStageStars(userID uint) ([]Entities.StageLog, error) {

	var logs []Entities.StageLog

	err := u.db.
		Where("user_id = ?", userID).
		Find(&logs).Error

	return logs, err
}
