package Entities

type Leaderboard struct {
	Rank       int  `json:"rank"`
	UserID     uint `json:"user_id"`
	Username   string `json:"username"`
	TotalStars int  `json:"total_stars"`
}
