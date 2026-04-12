package leaderboard_usecase

import (
	Entities "headliner-be/entities" 
)


type LeaderboardRepository interface {
	GetAllLeaderboard() ([]Entities.LeaderboardEntry, error)
	SaveScore(data Entities.Leaderboard) error
}

type LeaderboardUsecase interface {
	GetLeaderboard() ([]Entities.LeaderboardEntry, error)
	SaveScore(data Entities.Leaderboard) error
}

type leaderboardService struct {
	repo LeaderboardRepository 
}

func NewLeaderboardService(repo LeaderboardRepository) LeaderboardUsecase {
	return &leaderboardService{repo: repo}
}

func (s *leaderboardService) GetLeaderboard() ([]Entities.LeaderboardEntry, error) {
	return s.repo.GetAllLeaderboard()
}

func (s *leaderboardService) SaveScore(data Entities.Leaderboard) error {
    return s.repo.SaveScore(data)
}