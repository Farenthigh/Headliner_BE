package leaderboard

import (
    Entities "headliner-be/entites"
    leaderboard_usecase "headliner-be/usecase/leaderboard"  // ← import usecase แทน
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
        Select("users.username, leaderboard.saving_game_score, leaderboard.tax_game_score").
        Joins("join users on users.id = leaderboard.user_id").
        Scan(&results).Error
    return results, err
}