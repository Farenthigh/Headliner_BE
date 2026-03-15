package Entities

type Leaderboard struct {
	LeaderboardID   int `gorm:"primaryKey;column:leaderboard_id" json:"leaderboard_id"`
	UserID          int `gorm:"column:user_id" json:"user_id"`
	SavingGameScore int `gorm:"column:saving_game_score;default:0" json:"saving_game_score"`
	TaxGameScore    int `gorm:"column:tax_game_score;default:0" json:"tax_game_score"`
}

type LeaderboardEntry struct {
    Username        string `json:"username"`
    SavingGameScore int    `json:"saving_game_score"`
    TaxGameScore    int    `json:"tax_game_score"`
}