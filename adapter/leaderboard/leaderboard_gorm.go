package leaderboard

import (
	Entities "headliner-be/entities"
	leaderboard_usecase "headliner-be/usecase/leaderboard" 

	"gorm.io/gorm"
)

type LeaderboardGorm struct {
	db *gorm.DB
}

func NewLeaderboardGorm(db *gorm.DB) leaderboard_usecase.LeaderboardRepository {
	return &LeaderboardGorm{db: db}
}

func (r *LeaderboardGorm) GetAllLeaderboard() ([]Entities.LeaderboardEntry, error) {
	var results []Entities.LeaderboardEntry
	err := r.db.Table("leaderboard").
		Select("users.username, leaderboard.saving_game_score, leaderboard.tax_game_score, leaderboard.updated_at").
		Joins("join users on users.id = leaderboard.user_id").
		Scan(&results).Error
	return results, err
}


func (r *LeaderboardGorm) SaveScore(data Entities.Leaderboard) error {
	var existingRecord Entities.Leaderboard

	result := r.db.Where("user_id = ?", data.UserID).First(&existingRecord)

	if result.Error == nil {
		isUpdated := false

		if data.SavingGameScore > existingRecord.SavingGameScore {
			existingRecord.SavingGameScore = data.SavingGameScore
			existingRecord.SavingGameTime = data.SavingGameTime
			isUpdated = true
		}

		if data.TaxGameScore > existingRecord.TaxGameScore {
			existingRecord.TaxGameScore = data.TaxGameScore
			existingRecord.TaxGameTime = data.TaxGameTime
			isUpdated = true
		}

		if isUpdated {
			return r.db.Save(&existingRecord).Error
		}
		
		return nil 
	}

	return r.db.Create(&data).Error
}