package Entities

type Leaderboard struct {
	Rank       int  `json:"rank"`
	UserID     uint `json:"user_id"`
	Username   string `json:"username"`
	TotalStars int  `json:"total_stars"`
}

type LeaderboardEntry struct {
    Username        string `json:"username"`
    SavingGameScore int    `json:"saving_game_score"`
    TaxGameScore    int    `json:"tax_game_score"`
}
