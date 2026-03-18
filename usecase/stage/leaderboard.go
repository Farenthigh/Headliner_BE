package stage

import (
	Entities "headliner-be/entities"
)

func (u *StageUsecase) GetLeaderboard(limit int) ([]Entities.Leaderboard, error) {

	var board []Entities.Leaderboard

		err := u.db.Raw(`
		SELECT 
			sl.user_id,
			u.username,
			SUM(sl.stars) as total_stars,
			RANK() OVER (ORDER BY SUM(sl.stars) DESC) as rank
		FROM stage_logs sl
		JOIN users u ON u.id = sl.user_id
		GROUP BY sl.user_id, u.username
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
			sl.user_id,
			u.username,
			SUM(sl.stars) as total_stars,
			RANK() OVER (ORDER BY SUM(sl.stars) DESC) as rank
		FROM stage_logs sl
		JOIN users u ON u.id = sl.user_id
		GROUP BY sl.user_id, u.username
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

