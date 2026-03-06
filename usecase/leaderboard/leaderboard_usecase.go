package leaderboard_usecase

import (
    Entities "headliner-be/entites"  // ← เอา import adapter ออก
)

// ย้าย Repository interface มาไว้ที่นี่
type LeaderboardRepository interface {
    GetAllLeaderboard() ([]Entities.LeaderboardEntry, error)
}

type LeaderboardUsecase interface {
    GetLeaderboard() ([]Entities.LeaderboardEntry, error)
}

type leaderboardService struct {
    repo LeaderboardRepository  // ← ใช้ interface ที่ประกาศข้างบน
}

func NewLeaderboardService(repo LeaderboardRepository) LeaderboardUsecase {
    return &leaderboardService{repo: repo}
}

func (s *leaderboardService) GetLeaderboard() ([]Entities.LeaderboardEntry, error) {
    return s.repo.GetAllLeaderboard()
}