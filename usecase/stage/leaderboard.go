package stage

import (
	Entities "headliner-be/entities"
)

func (u *StageUsecase) GetLeaderboard(limit int) ([]Entities.Leaderboard, error) {

	var board []Entities.Leaderboard

	err := u.db.Raw(`
		SELECT 
			user_id,
			SUM(stars) as total_stars,
			RANK() OVER (ORDER BY SUM(stars) DESC) as rank
		FROM stage_logs
		GROUP BY user_id
		ORDER BY rank
		LIMIT ?
	`, limit).Scan(&board).Error

	return board, err
}

func (u *StageUsecase) GetPlayerRank(userID uint) (*Entities.Leaderboard, error) {

	var result Entities.Leaderboard

	err := u.db.Raw(`
	SELECT *
	FROM (
		SELECT 
			user_id,
			SUM(stars) as total_stars,
			RANK() OVER (ORDER BY SUM(stars) DESC) as rank
		FROM stage_logs
		GROUP BY user_id
	) ranked
	WHERE user_id = ?
	`, userID).Scan(&result).Error

	return &result, err
}

func (u *StageUsecase) GetUserStages(userID uint) ([]Entities.StageLog, error) {

	var stages []Entities.StageLog

	err := u.db.
		Where("user_id = ?", userID).
		Find(&stages).Error

	return stages, err
}

