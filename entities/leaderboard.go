package Entities

type Leaderboard struct {
	Rank       int  `json:"rank"`
	UserID     uint `json:"user_id"`
	TotalStars int  `json:"total_stars"`
}
type PlayerRank struct {
	Rank       int  `json:"rank"`
	UserID     uint `json:"user_id"`
	TotalStars int  `json:"total_stars"`
}