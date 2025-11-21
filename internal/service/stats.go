package service

import (
	"context"

	"Service-for-assigning-reviewers-for-Pull-Requests/internal/repository/postgres"
)

type StatsService struct {
	statsRepo postgres.StatsRepository
}

func NewStatsService(s postgres.StatsRepository) *StatsService {
	return &StatsService{statsRepo: s}
}

func (s *StatsService) GetAssignedCountPerPR(ctx context.Context) (map[string]int, error) {
	return s.statsRepo.GetAssignedReviewersCountPerPR(ctx)
}

func (s *StatsService) GetOpenPRCountPerUser(ctx context.Context) (map[string]int, error) {
	return s.statsRepo.GetOpenPRCountPerUser(ctx)
}
